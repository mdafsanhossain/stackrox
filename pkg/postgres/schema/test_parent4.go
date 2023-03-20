// Code generated by pg-bindings generator. DO NOT EDIT.

package schema

import (
	"fmt"
	"reflect"

	v1 "github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/postgres"
	"github.com/stackrox/rox/pkg/postgres/walker"
	"github.com/stackrox/rox/pkg/sac/resources"
	"github.com/stackrox/rox/pkg/search"
	"github.com/stackrox/rox/pkg/search/postgres/mapping"
)

var (
	// CreateTableTestParent4Stmt holds the create statement for table `test_parent4`.
	CreateTableTestParent4Stmt = &postgres.CreateStmts{
		GormModel: (*TestParent4)(nil),
		Children:  []*postgres.CreateStmts{},
	}

	// TestParent4Schema is the go schema for table `test_parent4`.
	TestParent4Schema = func() *walker.Schema {
		schema := GetSchemaForTable("test_parent4")
		if schema != nil {
			return schema
		}
		schema = walker.Walk(reflect.TypeOf((*storage.TestParent4)(nil)), "test_parent4")
		referencedSchemas := map[string]*walker.Schema{
			"storage.TestGrandparent": TestGrandparentsSchema,
		}

		schema.ResolveReferences(func(messageTypeName string) *walker.Schema {
			return referencedSchemas[fmt.Sprintf("storage.%s", messageTypeName)]
		})
		schema.ScopingResource = &resources.Namespace
		schema.SetOptionsMap(search.Walk(v1.SearchCategory(113), "testparent4", (*storage.TestParent4)(nil)))
		schema.SetSearchScope([]v1.SearchCategory{
			v1.SearchCategory(109),
			v1.SearchCategory(103),
		}...)
		RegisterTable(schema, CreateTableTestParent4Stmt)
		mapping.RegisterCategoryToTable(v1.SearchCategory(113), schema)
		return schema
	}()
)

const (
	TestParent4TableName = "test_parent4"
)

// TestParent4 holds the Gorm model for Postgres table `test_parent4`.
type TestParent4 struct {
	Id                  string           `gorm:"column:id;type:uuid;primaryKey"`
	ParentId            string           `gorm:"column:parentid;type:varchar"`
	Val                 string           `gorm:"column:val;type:varchar"`
	Serialized          []byte           `gorm:"column:serialized;type:bytea"`
	TestGrandparentsRef TestGrandparents `gorm:"foreignKey:parentid;references:id;belongsTo;constraint:OnDelete:CASCADE"`
}
