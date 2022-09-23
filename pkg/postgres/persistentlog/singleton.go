package persistentlog

import (
	"github.com/stackrox/rox/pkg/features"
	"github.com/stackrox/rox/pkg/postgres/persistentlog/store"
	"github.com/stackrox/rox/pkg/postgres/pgadmin"
	"github.com/stackrox/rox/pkg/postgres/pgconfig"
	"github.com/stackrox/rox/pkg/sync"
)

var (
	once sync.Once

	logStore store.Store
)

func initialize() {
	_, config, err := pgconfig.GetPostgresConfig()
	if err != nil {
		return
	}
	logStore = store.New(pgadmin.GetAdminPool(config))
}

// Singleton provides the interface for non-service external interaction.
func Singleton() store.Store {
	if !features.PostgresDatastore.Enabled() {
		return nil
	}
	once.Do(initialize)
	return logStore
}
