package version

import (
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/postgres/schema"
)

// ConvertVersionFromProto converts a `*storage.Version` to Gorm model
func ConvertVersionFromProto(obj *storage.Version) (*schema.Versions, error) {
	serialized, err := obj.Marshal()
	if err != nil {
		return nil, err
	}
	model := &schema.Versions{
		Serialized: serialized,
	}
	return model, nil
}

// ConvertVersionToProto converts Gorm model `Versions` to its protobuf type object
func ConvertVersionToProto(m *schema.Versions) (*storage.Version, error) {
	var msg storage.Version
	if err := msg.Unmarshal(m.Serialized); err != nil {
		return nil, err
	}
	return &msg, nil
}
