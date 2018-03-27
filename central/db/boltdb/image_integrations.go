package boltdb

import (
	"fmt"

	"bitbucket.org/stack-rox/apollo/central/db"
	"bitbucket.org/stack-rox/apollo/generated/api/v1"
	"bitbucket.org/stack-rox/apollo/pkg/uuid"
	"github.com/boltdb/bolt"
	"github.com/golang/protobuf/proto"
)

const imageIntegrationBucket = "imageintegrations"

func (b *BoltDB) getImageIntegration(id string, bucket *bolt.Bucket) (integration *v1.ImageIntegration, exists bool, err error) {
	integration = new(v1.ImageIntegration)
	val := bucket.Get([]byte(id))
	if val == nil {
		return
	}
	exists = true
	err = proto.Unmarshal(val, integration)
	return
}

// GetImageIntegration returns integration with given id.
func (b *BoltDB) GetImageIntegration(id string) (integration *v1.ImageIntegration, exists bool, err error) {
	err = b.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(imageIntegrationBucket))
		integration, exists, err = b.getImageIntegration(id, bucket)
		return err
	})
	return
}

// GetImageIntegrations retrieves integrations from bolt
func (b *BoltDB) GetImageIntegrations(request *v1.GetImageIntegrationsRequest) ([]*v1.ImageIntegration, error) {
	var integrations []*v1.ImageIntegration
	err := b.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(imageIntegrationBucket))
		return b.ForEach(func(k, v []byte) error {
			var integration v1.ImageIntegration
			if err := proto.Unmarshal(v, &integration); err != nil {
				return err
			}
			integrations = append(integrations, &integration)
			return nil
		})
	})
	return integrations, err
}

// AddImageIntegration adds a integration into bolt
func (b *BoltDB) AddImageIntegration(integration *v1.ImageIntegration) (string, error) {
	integration.Id = uuid.NewV4().String()
	err := b.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(imageIntegrationBucket))
		_, exists, err := b.getImageIntegration(integration.GetId(), bucket)
		if err != nil {
			return err
		}
		if exists {
			return fmt.Errorf("Image integration %s (%s) cannot be added because it already exists", integration.GetId(), integration.GetName())
		}
		if err := checkUniqueKeyExistsAndInsert(tx, imageIntegrationBucket, integration.GetId(), integration.GetName()); err != nil {
			return fmt.Errorf("Could not add image integration due to name validation: %s", err)
		}
		bytes, err := proto.Marshal(integration)
		if err != nil {
			return err
		}
		return bucket.Put([]byte(integration.GetId()), bytes)
	})
	return integration.Id, err
}

// UpdateImageIntegration upserts a integration into bolt
func (b *BoltDB) UpdateImageIntegration(integration *v1.ImageIntegration) error {
	return b.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(imageIntegrationBucket))
		// If the update is changing the name, check if the name has already been taken
		if getCurrentUniqueKey(tx, imageIntegrationBucket, integration.GetId()) != integration.GetName() {
			if err := checkUniqueKeyExistsAndInsert(tx, imageIntegrationBucket, integration.GetId(), integration.GetName()); err != nil {
				return fmt.Errorf("Could not update integration due to name validation: %s", err)
			}
		}
		bytes, err := proto.Marshal(integration)
		if err != nil {
			return err
		}
		return b.Put([]byte(integration.GetId()), bytes)
	})
}

// RemoveImageIntegration removes a integration from bolt
func (b *BoltDB) RemoveImageIntegration(id string) error {
	return b.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(imageIntegrationBucket))
		key := []byte(id)
		if exists := b.Get(key) != nil; !exists {
			return db.ErrNotFound{Type: "ImageIntegration", ID: string(key)}
		}
		if err := removeUniqueKey(tx, imageIntegrationBucket, id); err != nil {
			return err
		}
		return b.Delete(key)
	})
}
