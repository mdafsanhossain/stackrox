package version

import (
	"fmt"

	"github.com/dgraph-io/badger"
	bolt "github.com/etcd-io/bbolt"
	"github.com/gogo/protobuf/proto"
	"github.com/stackrox/rox/central/version/store"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/logging"
	"github.com/stackrox/rox/pkg/migrations"
)

var (
	log = logging.LoggerForModule()
)

// Ensure is an opaque command that ensures that the DB is in a good state by the time it returns.
// It will simply return if the DB version is already what Central expects.
// If the DB is empty, then it will write the current DB version to the DB.
// It will returns an error if the DB is of an old version.
// If Ensure returns an error, the state of the DB is undefined, and it is not safe for Central to try to
// function normally.
func Ensure(boltDB *bolt.DB, badgerDB *badger.DB) error {
	versionStore := store.New(boltDB, badgerDB)
	version, err := versionStore.GetVersion()
	if err != nil {
		return fmt.Errorf("failed to read version from DB: %v", err)
	}

	// No version in the DB. This means that we're starting from scratch, with a blank DB, so we can just
	// write the current version in and move on.
	if version == nil {
		if err := versionStore.UpdateVersion(&storage.Version{SeqNum: migrations.CurrentDBVersionSeqNum}); err != nil {
			return fmt.Errorf("failed to write version to the DB: %vs", err)
		}
		log.Info("No version found in the DB. Assuming that this is a fresh install...")
		return nil
	}

	if version.GetSeqNum() != migrations.CurrentDBVersionSeqNum {
		return fmt.Errorf("invalid DB version found: %s", proto.MarshalTextString(version))
	}

	// TYPICAL CASE: DB is of the same version. This happens if Central does a regular restart, and was not patched.

	log.Info("Version found in the DB was current. We're good to go!")
	return nil
}
