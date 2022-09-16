// Code generated by pg-bindings generator. DO NOT EDIT.
package schema

import (
	"testing"

	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/postgres/schema"
	"github.com/stackrox/rox/pkg/testutils"
	"github.com/stretchr/testify/assert"
)

// ConvertTestParent3FromProto converts a `*storage.TestParent3` to Gorm model
func ConvertTestParent3FromProto(obj *storage.TestParent3) (*schema.TestParent3, error) {
	serialized, err := obj.Marshal()
	if err != nil {
		return nil, err
	}
	model := &schema.TestParent3{
		Id:         obj.GetId(),
		ParentId:   obj.GetParentId(),
		Val:        obj.GetVal(),
		Serialized: serialized,
	}
	return model, nil
}

// ConvertTestParent3ToProto converts Gorm model `TestParent3` to its protobuf type object
func ConvertTestParent3ToProto(m *schema.TestParent3) (*storage.TestParent3, error) {
	var msg storage.TestParent3
	if err := msg.Unmarshal(m.Serialized); err != nil {
		return nil, err
	}
	return &msg, nil
}

func TestTestParent3Serialization(t *testing.T) {
	obj := &storage.TestParent3{}
	assert.NoError(t, testutils.FullInit(obj, testutils.UniqueInitializer(), testutils.JSONFieldsFilter))
	m, err := ConvertTestParent3FromProto(obj)
	assert.NoError(t, err)
	conv, err := ConvertTestParent3ToProto(m)
	assert.NoError(t, err)
	assert.Equal(t, obj, conv)
}
