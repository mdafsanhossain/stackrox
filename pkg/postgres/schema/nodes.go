// Code generated by pg-bindings generator. DO NOT EDIT.

package schema

import (
	"fmt"
	"reflect"
	"time"

	v1 "github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/postgres"
	"github.com/stackrox/rox/pkg/postgres/walker"
	"github.com/stackrox/rox/pkg/search"
)

var (
	// CreateTableNodesStmt holds the create statement for table `nodes`.
	CreateTableNodesStmt = &postgres.CreateStmts{
		Table: `
               create table if not exists nodes (
                   Id varchar,
                   Name varchar,
                   ClusterId varchar,
                   ClusterName varchar,
                   Labels jsonb,
                   Annotations jsonb,
                   JoinedAt timestamp,
                   ContainerRuntime_Version varchar,
                   OsImage varchar,
                   LastUpdated timestamp,
                   Scan_ScanTime timestamp,
                   Components integer,
                   Cves integer,
                   FixableCves integer,
                   RiskScore numeric,
                   TopCvss numeric,
                   serialized bytea,
                   PRIMARY KEY(Id)
               )
               `,
		GormModel: (*Nodes)(nil),
		Indexes:   []string{},
		Children: []*postgres.CreateStmts{
			&postgres.CreateStmts{
				Table: `
               create table if not exists nodes_taints (
                   nodes_Id varchar,
                   idx integer,
                   Key varchar,
                   Value varchar,
                   TaintEffect integer,
                   PRIMARY KEY(nodes_Id, idx),
                   CONSTRAINT fk_parent_table_0 FOREIGN KEY (nodes_Id) REFERENCES nodes(Id) ON DELETE CASCADE
               )
               `,
				GormModel: (*NodesTaints)(nil),
				Indexes: []string{
					"create index if not exists nodesTaints_idx on nodes_taints using btree(idx)",
				},
				Children: []*postgres.CreateStmts{},
			},
		},
	}

	// NodesSchema is the go schema for table `nodes`.
	NodesSchema = func() *walker.Schema {
		schema := GetSchemaForTable("nodes")
		if schema != nil {
			return schema
		}
		schema = walker.Walk(reflect.TypeOf((*storage.Node)(nil)), "nodes")
		referencedSchemas := map[string]*walker.Schema{
			"storage.Cluster": ClustersSchema,
		}

		schema.ResolveReferences(func(messageTypeName string) *walker.Schema {
			return referencedSchemas[fmt.Sprintf("storage.%s", messageTypeName)]
		})
		schema.SetOptionsMap(search.Walk(v1.SearchCategory_NODES, "node", (*storage.Node)(nil)))
		RegisterTable(schema, CreateTableNodesStmt)
		return schema
	}()
)

const (
	NodesTableName       = "nodes"
	NodesTaintsTableName = "nodes_taints"
)

// Nodes holds the Gorm model for Postgres table `nodes`.
type Nodes struct {
	Id                      string            `gorm:"column:id;type:varchar;primaryKey"`
	Name                    string            `gorm:"column:name;type:varchar"`
	ClusterId               string            `gorm:"column:clusterid;type:varchar"`
	ClusterName             string            `gorm:"column:clustername;type:varchar"`
	Labels                  map[string]string `gorm:"column:labels;type:jsonb"`
	Annotations             map[string]string `gorm:"column:annotations;type:jsonb"`
	JoinedAt                *time.Time        `gorm:"column:joinedat;type:timestamp"`
	ContainerRuntimeVersion string            `gorm:"column:containerruntime_version;type:varchar"`
	OsImage                 string            `gorm:"column:osimage;type:varchar"`
	LastUpdated             *time.Time        `gorm:"column:lastupdated;type:timestamp"`
	ScanScanTime            *time.Time        `gorm:"column:scan_scantime;type:timestamp"`
	Components              int32             `gorm:"column:components;type:integer"`
	Cves                    int32             `gorm:"column:cves;type:integer"`
	FixableCves             int32             `gorm:"column:fixablecves;type:integer"`
	RiskScore               float32           `gorm:"column:riskscore;type:numeric"`
	TopCvss                 float32           `gorm:"column:topcvss;type:numeric"`
	Serialized              []byte            `gorm:"column:serialized;type:bytea"`
}

// NodesTaints holds the Gorm model for Postgres table `nodes_taints`.
type NodesTaints struct {
	NodesId     string              `gorm:"column:nodes_id;type:varchar;primaryKey"`
	Idx         int                 `gorm:"column:idx;type:integer;primaryKey;index:nodestaints_idx,type:btree"`
	Key         string              `gorm:"column:key;type:varchar"`
	Value       string              `gorm:"column:value;type:varchar"`
	TaintEffect storage.TaintEffect `gorm:"column:tainteffect;type:integer"`
	NodesRef    Nodes               `gorm:"foreignKey:nodes_id;references:id;belongsTo;constraint:OnDelete:CASCADE"`
}
