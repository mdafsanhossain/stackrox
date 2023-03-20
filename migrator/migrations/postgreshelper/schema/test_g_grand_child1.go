// Code generated by pg-bindings generator. DO NOT EDIT.

package schema

import (
	"reflect"

	v1 "github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/postgres"
	"github.com/stackrox/rox/pkg/postgres/walker"
	"github.com/stackrox/rox/pkg/sac/resources"
	"github.com/stackrox/rox/pkg/search"
)

var (
	// CreateTableTestGGrandChild1Stmt holds the create statement for table `test_g_grand_child1`.
	CreateTableTestGGrandChild1Stmt = &postgres.CreateStmts{
		GormModel: (*TestGGrandChild1)(nil),
		Children:  []*postgres.CreateStmts{},
	}

	// TestGGrandChild1Schema is the go schema for table `test_g_grand_child1`.
	TestGGrandChild1Schema = func() *walker.Schema {
		schema := walker.Walk(reflect.TypeOf((*storage.TestGGrandChild1)(nil)), "test_g_grand_child1")
		schema.ScopingResource = &resources.Namespace
		schema.SetOptionsMap(search.Walk(v1.SearchCategory(65), "testggrandchild1", (*storage.TestGGrandChild1)(nil)))
		return schema
	}()
)

const (
	TestGGrandChild1TableName = "test_g_grand_child1"
)

// TestGGrandChild1 holds the Gorm model for Postgres table `test_g_grand_child1`.
type TestGGrandChild1 struct {
	Id         string `gorm:"column:id;type:varchar;primaryKey"`
	Val        string `gorm:"column:val;type:varchar"`
	Serialized []byte `gorm:"column:serialized;type:bytea"`
}
