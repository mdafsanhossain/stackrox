package registries

import (
	"sort"

	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/registries/types"
	"github.com/stackrox/rox/pkg/sync"
)

type setImpl struct {
	lock sync.RWMutex

	factory      Factory
	integrations map[string]types.ImageRegistry
}

func sortIntegrations(integrations []types.ImageRegistry) {
	// This just ensures that the registries that have username/passwords are processed first
	sort.SliceStable(integrations, func(i, j int) bool {
		conf1 := integrations[i].Config()
		conf2 := integrations[j].Config()
		// Sort by username existing first
		if conf1.Username != "" && conf2.Username == "" {
			return true
		}
		if conf1.Username == "" && conf2.Username != "" {
			return false
		}
		// Then sort by if autogenerated or not, this should prefer non-autogenerated registries (user input over autogenerated)
		return !conf1.Autogenerated && conf2.Autogenerated
	})
}

func (e *setImpl) getSortedRegistriesNoLock() []types.ImageRegistry {
	integrations := make([]types.ImageRegistry, 0, len(e.integrations))
	for _, i := range e.integrations {
		integrations = append(integrations, i)
	}
	sortIntegrations(integrations)
	return integrations
}

// GetAll returns the set of integrations that are active.
func (e *setImpl) GetAll() []types.ImageRegistry {
	e.lock.RLock()
	defer e.lock.RUnlock()
	return e.getSortedRegistriesNoLock()
}

// GetRegistryMetadataByImage returns the config for a registry that contains the input image.
func (e *setImpl) GetRegistryMetadataByImage(image *storage.Image) *types.Config {
	e.lock.RLock()
	defer e.lock.RUnlock()

	reg := e.getRegistryByImageNoLock(image)
	if reg != nil {
		return reg.Config()
	}

	return nil
}

// GetRegistryByImage returns the registry that contains the input image.
func (e *setImpl) GetRegistryByImage(image *storage.Image) types.Registry {
	e.lock.RLock()
	defer e.lock.RUnlock()

	return e.getRegistryByImageNoLock(image)
}

func (e *setImpl) getRegistryByImageNoLock(image *storage.Image) types.Registry {
	if sourceID := image.GetMetadata().GetDataSource().GetId(); sourceID != "" {
		reg, ok := e.integrations[sourceID]
		if ok {
			return reg
		}
	}

	integrations := e.getSortedRegistriesNoLock()
	for _, i := range integrations {
		if i.Match(image.GetName()) {
			return i
		}
	}

	return nil
}

// Match returns whether a registry in the set has the given image.
func (e *setImpl) Match(image *storage.ImageName) bool {
	e.lock.RLock()
	defer e.lock.RUnlock()

	integrations := e.getSortedRegistriesNoLock()
	for _, i := range integrations {
		if i.Match(image) {
			return true
		}
	}
	return false
}

// IsEmpty returns whether the set is empty.
func (e *setImpl) IsEmpty() bool {
	e.lock.RLock()
	defer e.lock.RUnlock()

	return len(e.integrations) == 0
}

// Clear removes all present integrations.
func (e *setImpl) Clear() {
	e.lock.Lock()
	defer e.lock.Unlock()

	e.integrations = make(map[string]types.ImageRegistry)
}

// UpdateImageIntegration updates the integration with the matching id to a new configuration.
// This does not update a pre-existing registry, instead it replaces it with a new one.
func (e *setImpl) UpdateImageIntegration(integration *storage.ImageIntegration) error {
	i, err := e.factory.CreateRegistry(integration)
	if err != nil {
		return err
	}

	e.lock.Lock()
	defer e.lock.Unlock()

	e.integrations[integration.GetId()] = i
	return nil
}

// RemoveImageIntegration removes the integration with a matching id if one exists.
func (e *setImpl) RemoveImageIntegration(id string) error {
	e.lock.Lock()
	defer e.lock.Unlock()

	delete(e.integrations, id)
	return nil
}
