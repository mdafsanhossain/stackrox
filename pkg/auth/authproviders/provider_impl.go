package authproviders

import (
	"context"

	"github.com/gogo/protobuf/proto"
	"github.com/pkg/errors"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/auth/permissions"
	"github.com/stackrox/rox/pkg/auth/tokens"
	"github.com/stackrox/rox/pkg/sync"
)

var (
	errProviderDisabled = errors.New("provider has been deleted or disabled")
)

// If you add new data fields to this class, make sure you make commensurate modifications
// to the cloneWithoutMutex and copyWithoutMutex functions below.
type providerImpl struct {
	mutex sync.RWMutex

	storedInfo storage.AuthProvider
	backend    Backend
	roleMapper permissions.RoleMapper
	issuer     tokens.Issuer

	doNotStore bool

	validateCallback func() error
}

// Accessor functions.
//////////////////////

func (p *providerImpl) ID() string {
	p.mutex.RLock()
	defer p.mutex.RUnlock()

	return p.storedInfo.GetId()
}

func (p *providerImpl) Type() string {
	p.mutex.RLock()
	defer p.mutex.RUnlock()

	return p.storedInfo.Type
}

func (p *providerImpl) Name() string {
	p.mutex.RLock()
	defer p.mutex.RUnlock()

	return p.storedInfo.Name
}

func (p *providerImpl) Enabled() bool {
	p.mutex.RLock()
	defer p.mutex.RUnlock()

	return p.backend != nil && p.storedInfo.Enabled
}

func (p *providerImpl) Active() bool {
	p.mutex.RLock()
	defer p.mutex.RUnlock()

	return p.storedInfo.GetActive()
}

func (p *providerImpl) StorageView() *storage.AuthProvider {
	p.mutex.RLock()
	defer p.mutex.RUnlock()

	result := p.storedInfo
	if p.backend == nil {
		result.Enabled = false
	}
	return &result
}

func (p *providerImpl) Backend() Backend {
	p.mutex.RLock()
	defer p.mutex.RUnlock()

	return p.backend
}

func (p *providerImpl) RoleMapper() permissions.RoleMapper {
	p.mutex.RLock()
	defer p.mutex.RUnlock()

	return p.roleMapper
}

func (p *providerImpl) Issuer() tokens.Issuer {
	p.mutex.RLock()
	defer p.mutex.RUnlock()

	return p.issuer
}

// Modifier functions.
//////////////////////

func (p *providerImpl) Validate(ctx context.Context, claims *tokens.Claims) error {
	enabled := p.Enabled()
	if !enabled {
		return errProviderDisabled
	}
	backend := p.Backend()
	return backend.Validate(ctx, claims)
}

// We must lock the provider when applying options to it.
func (p *providerImpl) ApplyOptions(options ...ProviderOption) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	// Try updates on a copy of the provider
	modifiedProvider := cloneWithoutMutex(p)
	if err := applyOptions(modifiedProvider, options...); err != nil {
		return err
	}

	// If updates succeed, apply them.
	copyWithoutMutex(p, modifiedProvider)
	return nil
}

func (p *providerImpl) MarkAsActive() error {
	if p.Active() || p.validateCallback == nil {
		return nil
	}
	return p.validateCallback()
}

// Does a deep copy of the proto field 'storedInfo' so that it can support nested message fields.
func cloneWithoutMutex(pr *providerImpl) *providerImpl {
	return &providerImpl{
		storedInfo: *proto.Clone(&pr.storedInfo).(*storage.AuthProvider),
		backend:    pr.backend,
		roleMapper: pr.roleMapper,
		issuer:     pr.issuer,
	}
}

// No need to do a deep copy of the 'storedInfo' field here since the 'from' input was created with a deep copy.
func copyWithoutMutex(to *providerImpl, from *providerImpl) {
	to.storedInfo = from.storedInfo
	to.backend = from.backend
	to.roleMapper = from.roleMapper
	to.issuer = from.issuer
}
