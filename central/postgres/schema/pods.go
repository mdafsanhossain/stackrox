// Code generated by pg-bindings generator. DO NOT EDIT.

package schema

import (
	"fmt"
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
	// CreateTablePodsStmt holds the create statement for table `pods`.
	CreateTablePodsStmt = &postgres.CreateStmts{
		GormModel: (*Pods)(nil),
		Children: []*postgres.CreateStmts{
			&postgres.CreateStmts{
				GormModel: (*PodsLiveInstances)(nil),
				Children:  []*postgres.CreateStmts{},
			},
		},
	}

	// PodsSchema is the go schema for table `pods`.
	PodsSchema = func() *walker.Schema {
		schema := schemaPkg.GetSchemaForTable("pods")
		if schema != nil {
			return schema
		}
		schema = walker.Walk(reflect.TypeOf((*storage.Pod)(nil)), "pods")
		referencedSchemas := map[string]*walker.Schema{
			"storage.Deployment": DeploymentsSchema,
		}

		schema.ResolveReferences(func(messageTypeName string) *walker.Schema {
			return referencedSchemas[fmt.Sprintf("storage.%s", messageTypeName)]
		})
		schema.SetOptionsMap(search.Walk(v1.SearchCategory_PODS, "pod", (*storage.Pod)(nil)))
		schemaPkg.RegisterTable(schema, CreateTablePodsStmt)
		mapping.RegisterCategoryToTable(v1.SearchCategory_PODS, schema)
		return schema
	}()
)

const (
	PodsTableName              = "pods"
	PodsLiveInstancesTableName = "pods_live_instances"
)

// Pods holds the Gorm model for Postgres table `pods`.
type Pods struct {
	Id           string `gorm:"column:id;type:uuid;primaryKey"`
	Name         string `gorm:"column:name;type:varchar"`
	DeploymentId string `gorm:"column:deploymentid;type:uuid"`
	Namespace    string `gorm:"column:namespace;type:varchar;index:pods_sac_filter,type:btree"`
	ClusterId    string `gorm:"column:clusterid;type:uuid;index:pods_sac_filter,type:btree"`
	Serialized   []byte `gorm:"column:serialized;type:bytea"`
}

// PodsLiveInstances holds the Gorm model for Postgres table `pods_live_instances`.
type PodsLiveInstances struct {
	PodsId      string `gorm:"column:pods_id;type:uuid;primaryKey"`
	Idx         int    `gorm:"column:idx;type:integer;primaryKey;index:podsliveinstances_idx,type:btree"`
	ImageDigest string `gorm:"column:imagedigest;type:varchar"`
	PodsRef     Pods   `gorm:"foreignKey:pods_id;references:id;belongsTo;constraint:OnDelete:CASCADE"`
}
