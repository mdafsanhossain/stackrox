package m18to19

import (
	"github.com/dgraph-io/badger"
	bolt "github.com/etcd-io/bbolt"
	"github.com/pkg/errors"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/migrator/log"
	"github.com/stackrox/rox/migrator/migrations"
	"github.com/stackrox/rox/migrator/types"
	"github.com/stackrox/rox/pkg/features"
)

var migration = types.Migration{
	StartingSeqNum: 18,
	VersionAfter:   storage.Version{SeqNum: 19},
	Run: func(db *bolt.DB, badgerDB *badger.DB) error {
		if !features.BadgerDB.Enabled() {
			return nil
		}
		return rewriteData(db, badgerDB)
	},
}

const (
	batchSize = 2000
)

var (
	deploymentBucket     = []byte("deployments")
	listDeploymentBucket = []byte("deployments_list")

	alertBucket         = []byte("alerts")
	listAlertBucketName = []byte("alerts_list")

	imageBucket     = []byte("imageBucket")
	listImageBucket = []byte("images_list")

	processIndicatorBucket = []byte("process_indicators")
	uniqueProcessesBucket  = []byte("process_indicators_unique")
)

func init() {
	migrations.MustRegisterMigration(migration)
}

func rewriteData(db *bolt.DB, badgerDB *badger.DB) error {
	// Alert
	if err := rewrite(db, badgerDB, alertBucket); err != nil {
		return err
	}
	if err := rewrite(db, badgerDB, listAlertBucketName); err != nil {
		return err
	}

	// Deployment
	if err := rewrite(db, badgerDB, deploymentBucket); err != nil {
		return err
	}
	if err := rewrite(db, badgerDB, listDeploymentBucket); err != nil {
		return err
	}

	// Image
	if err := rewrite(db, badgerDB, imageBucket); err != nil {
		return err
	}
	if err := rewrite(db, badgerDB, listImageBucket); err != nil {
		return err
	}

	// Process Indicators
	if err := rewrite(db, badgerDB, processIndicatorBucket); err != nil {
		return err
	}
	if err := rewrite(db, badgerDB, uniqueProcessesBucket); err != nil {
		return err
	}
	return nil
}

func rewrite(db *bolt.DB, badgerDB *badger.DB, bucketName []byte) error {
	log.WriteToStderr("Rewriting Bucket %q", string(bucketName))
	return db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(bucketName)
		if bucket == nil {
			return nil
		}
		totalKeys := bucket.Stats().KeyN
		log.WriteToStderr("Total keys in bucket: %d", totalKeys)

		keysWritten := 0
		batch := badgerDB.NewWriteBatch()
		err := bucket.ForEach(func(k, v []byte) error {
			if batch.Error() != nil {
				return batch.Error()
			}
			key := make([]byte, 0, len(bucketName)+len(k)+1)
			key = append(key, bucketName...)
			// The separator is a null char
			key = append(key, []byte("\x00")...)
			key = append(key, k...)

			if err := batch.Set(key, v); err != nil {
				return errors.Wrapf(err, "error setting key/value in Badger for bucket %q", string(bucketName))
			}

			keysWritten++
			if keysWritten%batchSize == 0 {
				if err := batch.Flush(); err != nil {
					return err
				}
				log.WriteToStderr("Written %d/%d keys for bucket %q", keysWritten, totalKeys, string(bucketName))
				batch = badgerDB.NewWriteBatch()
			}
			return nil
		})
		defer batch.Cancel()
		if err != nil {
			return err
		}
		log.WriteToStderr("Running final flush for %s into BadgerDB", string(bucketName))
		if err := batch.Flush(); err != nil {
			return errors.Wrapf(err, "error flushing BadgerDB for bucket %q", string(bucketName))
		}
		log.WriteToStderr("Wrote %s into BadgerDB. Deleting Bucket from Bolt", string(bucketName))
		if err := tx.DeleteBucket(bucketName); err != nil {
			return errors.Wrapf(err, "error deleting bucket %q from Bolt", string(bucketName))
		}
		log.WriteToStderr("Successfully deleted bucket %q", string(bucketName))
		return nil
	})
}
