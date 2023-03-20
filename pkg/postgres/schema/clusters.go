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
	"github.com/stackrox/rox/pkg/search/postgres/mapping"
)

var (
	// CreateTableClustersStmt holds the create statement for table `clusters`.
	CreateTableClustersStmt = &postgres.CreateStmts{
		GormModel: (*Clusters)(nil),
		Children:  []*postgres.CreateStmts{},
	}

	// ClustersSchema is the go schema for table `clusters`.
	ClustersSchema = func() *walker.Schema {
		schema := GetSchemaForTable("clusters")
		if schema != nil {
			return schema
		}
		schema = walker.Walk(reflect.TypeOf((*storage.Cluster)(nil)), "clusters")
		schema.ScopingResource = &resources.Cluster
		schema.SetOptionsMap(search.Walk(v1.SearchCategory_CLUSTERS, "cluster", (*storage.Cluster)(nil)))
		RegisterTable(schema, CreateTableClustersStmt)
		mapping.RegisterCategoryToTable(v1.SearchCategory_CLUSTERS, schema)
		return schema
	}()
)

const (
	ClustersTableName = "clusters"
)

// Clusters holds the Gorm model for Postgres table `clusters`.
type Clusters struct {
	Id         string            `gorm:"column:id;type:uuid;primaryKey"`
	Name       string            `gorm:"column:name;type:varchar;unique"`
	Labels     map[string]string `gorm:"column:labels;type:jsonb"`
	Serialized []byte            `gorm:"column:serialized;type:bytea"`
}
