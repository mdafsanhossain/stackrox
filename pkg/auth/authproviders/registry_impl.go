package authproviders

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/stackrox/rox/pkg/auth/permissions"
	"github.com/stackrox/rox/pkg/auth/tokens"
	"github.com/stackrox/rox/pkg/logging"
	"github.com/stackrox/rox/pkg/sync"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	log = logging.LoggerForModule()
)

// NewStoreBackedRegistry creates a new auth provider registry that is backed by a store. It also can handle HTTP requests,
// where every incoming HTTP request URL is expected to refer to a path under `urlPathPrefix`. The redirect URL for
// clients upon successful/failed authentication is `clientRedirectURL`.
func NewStoreBackedRegistry(urlPathPrefix string, redirectURL string, store Store, tokenIssuerFactory tokens.IssuerFactory, roleMapperFactory permissions.RoleMapperFactory) (Registry, error) {
	urlPathPrefix = strings.TrimRight(urlPathPrefix, "/") + "/"
	registry := &registryImpl{
		ServeMux:      http.NewServeMux(),
		urlPathPrefix: urlPathPrefix,
		redirectURL:   redirectURL,
		store:         store,
		issuerFactory: tokenIssuerFactory,

		backendFactories: make(map[string]BackendFactory),
		providers:        make(map[string]Provider),

		roleMapperFactory: roleMapperFactory,
	}

	return registry, nil
}

type registryImpl struct {
	*http.ServeMux

	urlPathPrefix string
	redirectURL   string

	store         Store
	issuerFactory tokens.IssuerFactory

	backendFactories map[string]BackendFactory
	providers        map[string]Provider
	mutex            sync.RWMutex

	roleMapperFactory permissions.RoleMapperFactory
}

func (r *registryImpl) Init() error {
	providerDefs, err := r.store.GetAllAuthProviders()
	if err != nil {
		return err
	}

	r.providers = make(map[string]Provider, len(providerDefs))
	for _, storedValue := range providerDefs {
		// Construct the options for the provider, using the stored definition, and the defaults for previously stored objects.
		options := []ProviderOption{
			WithStorageView(storedValue),
		}
		options = append(options, DefaultOptionsForStoredProvider(r.backendFactories, r.issuerFactory, r.roleMapperFactory, r.loginURL)...)

		// Use the options to build the provider.
		provider, err := NewProvider(options...)
		if err != nil {
			panic(err)
		}
		r.addProvider(provider)
	}

	r.initHTTPMux()
	return nil
}

// Accessors that read the registry.
////////////////////////////////////

func (r *registryImpl) GetProvider(id string) Provider {
	return r.getAuthProvider(id)
}

func (r *registryImpl) GetProviders(name, typ *string) []Provider {
	var result []Provider

	r.mutex.RLock()
	defer r.mutex.RUnlock()

	for _, provider := range r.providers {
		if typ != nil && *typ != provider.Type() {
			continue
		}
		if name != nil && provider.Name() != *name {
			continue
		}
		result = append(result, provider)
	}

	return result
}

func (r *registryImpl) getFactory(typ string) BackendFactory {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	return r.backendFactories[typ]
}

func (r *registryImpl) getAuthProvider(id string) Provider {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	return r.providers[id]
}

// Modifiers that update the registry.
//////////////////////////////////////

func (r *registryImpl) RegisterBackendFactory(ctx context.Context, typ string, factoryCreator BackendFactoryCreator) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if r.backendFactories[typ] != nil {
		return fmt.Errorf("backend factory for type %s is already registered", typ)
	}

	pathPrefix := fmt.Sprintf("%s%s/", r.providersURLPrefix(), typ)
	factory := factoryCreator(pathPrefix)
	if factory == nil {
		return errors.New("factory creator returned nil factory")
	}
	r.backendFactories[typ] = factory

	for _, provider := range r.providers {
		if provider.Type() != typ || provider.Backend() != nil {
			continue
		}
		go func(p Provider) {
			if err := p.ApplyOptions(WithBackendFromFactory(ctx, factory)); err != nil {
				log.Errorf("Failed to apply options: %v", err)
			}
		}(provider)
	}

	return nil
}

func (r *registryImpl) CreateProvider(ctx context.Context, options ...ProviderOption) (Provider, error) {
	// Add default options for creation.
	options = append(options, DefaultOptionsForNewProvider(ctx, r.store, r.backendFactories, r.issuerFactory, r.roleMapperFactory, r.loginURL)...)

	// Create provider and add to pool.
	newProvider, err := NewProvider(options...)
	if err != nil {
		return nil, err
	}
	r.addProvider(newProvider)

	return newProvider, nil
}

func (r *registryImpl) UpdateProvider(ctx context.Context, id string, options ...ProviderOption) (Provider, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	provider := r.providers[id]
	if provider == nil {
		return nil, fmt.Errorf("provider with ID %s not found", id)
	}

	// Run the updates with an update to the store added.
	// This will perform name validation since it is a secondary key in the store.
	if err := provider.ApplyOptions(append(options, UpdateStore(ctx, r.store))...); err != nil {
		return nil, err
	}
	r.updatedNoLock(provider)

	return provider, nil
}

func (r *registryImpl) DeleteProvider(ctx context.Context, id string, ignoreActive bool) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	provider := r.providers[id]
	if provider == nil {
		return nil
	}

	if provider.Active() && !ignoreActive {
		return errors.New("cannot update an auth provider once it has been used. Please delete and then re-add to modify")
	}

	if err := provider.ApplyOptions(DeleteFromStore(ctx, r.store), UnregisterSource(r.issuerFactory)); err != nil {
		return err
	}
	delete(r.providers, id)
	r.deletedNoLock(provider)
	return nil
}

func (r *registryImpl) ExchangeToken(ctx context.Context, externalToken, typ, state string) (string, string, error) {
	factory := r.getFactory(typ)
	if factory == nil {
		return "", "", status.Errorf(codes.InvalidArgument, "invalid auth provider type %q", typ)
	}

	providerID, err := factory.ResolveProvider(state)
	if err != nil {
		return "", "", err
	}
	provider := r.getAuthProvider(providerID)
	if provider == nil {
		return "", "", status.Errorf(codes.NotFound, "could not locate auth provider %q", providerID)
	}

	claim, opts, clientState, err := provider.Backend().ExchangeToken(ctx, externalToken, state)
	if err != nil {
		return "", "", err
	}
	token, err := provider.Issuer().Issue(ctx, tokens.RoxClaims{ExternalUser: claim}, opts...)
	if err != nil {
		return "", "", err
	}
	return token.Token, clientState, nil
}

func (r *registryImpl) addProvider(provider Provider) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.providers[provider.ID()] = provider
	r.updatedNoLock(provider)
}

// updatedNoLock fires the callback for the provider backend. Must be called under mutex.
func (r *registryImpl) updatedNoLock(provider Provider) {
	backend := provider.Backend()
	if backend == nil {
		return
	}
	if provider.Enabled() {
		backend.OnEnable(provider)
	} else {
		backend.OnDisable(provider)
	}
}

// deletedNoLock fires the callback for the provider backend. Must be called under mutex.
func (r *registryImpl) deletedNoLock(provider Provider) {
	backend := provider.Backend()
	if backend == nil {
		return
	}
	backend.OnDisable(provider)
}
