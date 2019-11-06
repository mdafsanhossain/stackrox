package bolt

import (
	"fmt"
	"time"

	bolt "github.com/etcd-io/bbolt"
	"github.com/gogo/protobuf/proto"
	"github.com/stackrox/rox/central/alert/convert"
	"github.com/stackrox/rox/central/alert/datastore/internal/store"
	"github.com/stackrox/rox/central/metrics"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/bolthelper"
	"github.com/stackrox/rox/pkg/logging"
	ops "github.com/stackrox/rox/pkg/metrics"
)

var (
	log = logging.LoggerForModule()

	alertBucket     = []byte("alerts")
	alertListBucket = []byte("alerts_list")
)

type storeImpl struct {
	*bolthelper.BoltWrapper
}

// New returns a new Store instance using the provided bolt DB instance.
func New(db *bolt.DB) store.Store {
	bolthelper.RegisterBucketOrPanic(db, alertBucket)
	bolthelper.RegisterBucketOrPanic(db, alertListBucket)

	wrapper, err := bolthelper.NewBoltWrapper(db, alertBucket)
	if err != nil {
		panic(err)
	}

	return &storeImpl{
		BoltWrapper: wrapper,
	}
}

// GetAlert returns an alert with given id.
func (b *storeImpl) ListAlert(id string) (alert *storage.ListAlert, exists bool, err error) {
	defer metrics.SetBoltOperationDurationTime(time.Now(), ops.Get, "ListAlert")
	err = b.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(alertListBucket)
		alert = new(storage.ListAlert)
		val := bucket.Get([]byte(id))
		if val == nil {
			return nil
		}
		exists = true
		err = proto.Unmarshal(val, alert)
		return nil
	})

	return
}

// GetAlertStates returns a minimal message in order to determine the state of the alerts
func (b *storeImpl) GetAlertIDs() ([]string, error) {
	defer metrics.SetBoltOperationDurationTime(time.Now(), ops.GetAll, "AlertIDs")

	var ids []string
	err := b.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(alertListBucket)
		return b.ForEach(func(k, v []byte) error {
			ids = append(ids, string(k))
			return nil
		})
	})
	return ids, err
}

// ListAlerts returns a minimal form of the Alert struct for faster marshalling
func (b *storeImpl) ListAlerts() ([]*storage.ListAlert, error) {
	defer metrics.SetBoltOperationDurationTime(time.Now(), ops.GetMany, "ListAlert")

	var alerts []*storage.ListAlert
	err := b.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(alertListBucket)
		return b.ForEach(func(k, v []byte) error {
			var alert storage.ListAlert
			if err := proto.Unmarshal(v, &alert); err != nil {
				return err
			}
			alerts = append(alerts, &alert)
			return nil
		})
	})
	return alerts, err
}

// GetAlert returns an alert with given id.
func (b *storeImpl) GetAlert(id string) (alert *storage.Alert, exists bool, err error) {
	defer metrics.SetBoltOperationDurationTime(time.Now(), ops.Get, "Alert")

	err = b.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(alertBucket)
		alert, exists, err = getAlert(id, bucket)
		return err
	})

	return
}

func (b *storeImpl) GetListAlerts(ids []string) ([]*storage.ListAlert, []int, error) {
	defer metrics.SetBoltOperationDurationTime(time.Now(), ops.GetMany, "ListAlert")

	alerts := make([]*storage.ListAlert, 0, len(ids))
	var missingIndices []int
	err := b.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(alertListBucket)
		for i, id := range ids {
			v := b.Get([]byte(id))
			if v == nil {
				missingIndices = append(missingIndices, i)
				continue
			}
			var alert storage.ListAlert
			if err := proto.Unmarshal(v, &alert); err != nil {
				return err
			}
			alerts = append(alerts, &alert)
		}
		return nil
	})

	return alerts, missingIndices, err
}

func (b *storeImpl) GetAlerts(ids []string) ([]*storage.Alert, []int, error) {
	defer metrics.SetBoltOperationDurationTime(time.Now(), ops.GetMany, "Alert")

	alerts := make([]*storage.Alert, 0, len(ids))
	var missingIndices []int
	err := b.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(alertBucket)
		for i, id := range ids {
			v := b.Get([]byte(id))
			if v == nil {
				missingIndices = append(missingIndices, i)
				continue
			}
			var alert storage.Alert
			if err := proto.Unmarshal(v, &alert); err != nil {
				return err
			}
			alerts = append(alerts, &alert)
		}
		return nil
	})

	return alerts, missingIndices, err
}

// AddAlert adds an alert into Bolt
func (b *storeImpl) AddAlert(alert *storage.Alert) error {
	defer metrics.SetBoltOperationDurationTime(time.Now(), ops.Add, "Alert")

	bytes, err := proto.Marshal(alert)
	if err != nil {
		return err
	}

	listBytes, err := marshalAsListAlert(alert)
	if err != nil {
		return err
	}

	return b.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(alertBucket)
		if bucket.Get([]byte(alert.Id)) != nil {
			return fmt.Errorf("Alert %v cannot be added because it already exists", alert.GetId())
		}
		if err := bucket.Put([]byte(alert.Id), bytes); err != nil {
			return err
		}
		bucket = tx.Bucket(alertListBucket)
		return bucket.Put([]byte(alert.Id), listBytes)
	})
}

// UpdateAlert upserts an alert into Bolt
func (b *storeImpl) UpdateAlert(alert *storage.Alert) error {
	defer metrics.SetBoltOperationDurationTime(time.Now(), ops.Update, "Alert")

	bytes, err := proto.Marshal(alert)
	if err != nil {
		return err
	}

	listBytes, err := marshalAsListAlert(alert)
	if err != nil {
		return err
	}

	return b.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(alertBucket)
		if err := bucket.Put([]byte(alert.Id), bytes); err != nil {
			return err
		}
		bucket = tx.Bucket(alertListBucket)
		return bucket.Put([]byte(alert.Id), listBytes)
	})
}

func (b *storeImpl) GetTxnCount() (txNum uint64, err error) {
	err = b.View(func(tx *bolt.Tx) error {
		txNum = b.BoltWrapper.GetTxnCount(tx)
		return nil
	})
	return
}

func (b *storeImpl) IncTxnCount() error {
	return b.Update(func(tx *bolt.Tx) error {
		// The b.Update increments the txn count automatically
		return nil
	})
}

func (b *storeImpl) deleteAlert(id string, tx *bolt.Tx) error {
	alertBucket := tx.Bucket(alertBucket)
	if err := alertBucket.Delete([]byte(id)); err != nil {
		return err
	}
	listAlertBucket := tx.Bucket(alertListBucket)
	if err := listAlertBucket.Delete([]byte(id)); err != nil {
		return err
	}
	return nil
}

// DeleteAlert removes an alert
func (b *storeImpl) DeleteAlert(id string) error {
	defer metrics.SetBoltOperationDurationTime(time.Now(), ops.Remove, "Alert")

	return b.Update(func(tx *bolt.Tx) error {
		return b.deleteAlert(id, tx)
	})
}

// DeleteAlerts removes multiple alerts
func (b *storeImpl) DeleteAlerts(ids ...string) error {
	defer metrics.SetBoltOperationDurationTime(time.Now(), ops.RemoveMany, "Alert")

	return b.Update(func(tx *bolt.Tx) error {
		for _, id := range ids {
			if err := b.deleteAlert(id, tx); err != nil {
				return err
			}
		}
		return nil
	})
}

func (b *storeImpl) WalkAll(fn func(*storage.ListAlert) error) error {
	return b.DB.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(alertListBucket)
		return bucket.ForEach(func(k, v []byte) error {
			var alert storage.ListAlert
			if err := proto.Unmarshal(v, &alert); err != nil {
				return err
			}
			return fn(&alert)
		})
	})
}

func getAlert(id string, bucket *bolt.Bucket) (alert *storage.Alert, exists bool, err error) {
	alert = new(storage.Alert)
	val := bucket.Get([]byte(id))
	if val == nil {
		return
	}
	exists = true
	err = proto.Unmarshal(val, alert)
	return
}

func marshalAsListAlert(alert *storage.Alert) ([]byte, error) {
	listAlert := convert.AlertToListAlert(alert)
	return proto.Marshal(listAlert)
}
