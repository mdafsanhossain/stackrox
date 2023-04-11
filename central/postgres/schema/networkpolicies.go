// Code generated by pg-bindings generator. DO NOT EDIT.

package schema

import (
	"reflect"

	v1 "github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/postgres"
	schemaPkg "github.com/stackrox/rox/pkg/postgres/schema"
	"github.com/stackrox/rox/pkg/postgres/walker"
	"github.com/stackrox/rox/pkg/search"
	"github.com/stackrox/rox/pkg/search/postgres/mapping"
)

var (
	// CreateTableNetworkpoliciesStmt holds the create statement for table `networkpolicies`.
	CreateTableNetworkpoliciesStmt = &postgres.CreateStmts{
		GormModel: (*Networkpolicies)(nil),
		Children:  []*postgres.CreateStmts{},
	}

	// NetworkpoliciesSchema is the go schema for table `networkpolicies`.
	NetworkpoliciesSchema = func() *walker.Schema {
		schema := schemaPkg.GetSchemaForTable("networkpolicies")
		if schema != nil {
			return schema
		}
		schema = walker.Walk(reflect.TypeOf((*storage.NetworkPolicy)(nil)), "networkpolicies")
		schema.SetOptionsMap(search.Walk(v1.SearchCategory_NETWORK_POLICIES, "networkpolicy", (*storage.NetworkPolicy)(nil)))
		schemaPkg.RegisterTable(schema, CreateTableNetworkpoliciesStmt)
		mapping.RegisterCategoryToTable(v1.SearchCategory_NETWORK_POLICIES, schema)
		return schema
	}()
)

const (
	NetworkpoliciesTableName = "networkpolicies"
)

// Networkpolicies holds the Gorm model for Postgres table `networkpolicies`.
type Networkpolicies struct {
	Id         string `gorm:"column:id;type:varchar;primaryKey"`
	ClusterId  string `gorm:"column:clusterid;type:uuid;index:networkpolicies_sac_filter,type:btree"`
	Namespace  string `gorm:"column:namespace;type:varchar;index:networkpolicies_sac_filter,type:btree"`
	Serialized []byte `gorm:"column:serialized;type:bytea"`
}
