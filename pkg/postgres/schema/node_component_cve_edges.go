// Code generated by pg-bindings generator. DO NOT EDIT.

package schema

import (
	"fmt"
	"reflect"

	v1 "github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/postgres"
	"github.com/stackrox/rox/pkg/postgres/walker"
	"github.com/stackrox/rox/pkg/search"
)

var (
	// CreateTableNodeComponentCveEdgesStmt holds the create statement for table `node_component_cve_edges`.
	CreateTableNodeComponentCveEdgesStmt = &postgres.CreateStmts{
		Table: `
               create table if not exists node_component_cve_edges (
                   Id varchar,
                   IsFixable bool,
                   FixedBy varchar,
                   ComponentId varchar,
                   CveId varchar,
                   serialized bytea,
                   PRIMARY KEY(Id),
                   CONSTRAINT fk_parent_table_0 FOREIGN KEY (ComponentId) REFERENCES node_components(Id) ON DELETE CASCADE
               )
               `,
		GormModel: (*NodeComponentCveEdges)(nil),
		Indexes:   []string{},
		Children:  []*postgres.CreateStmts{},
	}

	// NodeComponentCveEdgesSchema is the go schema for table `node_component_cve_edges`.
	NodeComponentCveEdgesSchema = func() *walker.Schema {
		schema := GetSchemaForTable("node_component_cve_edges")
		if schema != nil {
			return schema
		}
		schema = walker.Walk(reflect.TypeOf((*storage.NodeComponentCVEEdge)(nil)), "node_component_cve_edges")
		referencedSchemas := map[string]*walker.Schema{
			"storage.ImageComponent": NodeComponentsSchema,
			"storage.CVE":            NodeCvesSchema,
		}

		schema.ResolveReferences(func(messageTypeName string) *walker.Schema {
			return referencedSchemas[fmt.Sprintf("storage.%s", messageTypeName)]
		})
		schema.SetOptionsMap(search.Walk(v1.SearchCategory_NODE_COMPONENT_CVE_EDGE, "nodecomponentcveedge", (*storage.NodeComponentCVEEdge)(nil)))
		RegisterTable(schema, CreateTableNodeComponentCveEdgesStmt)
		return schema
	}()
)

const (
	NodeComponentCveEdgesTableName = "node_component_cve_edges"
)

// NodeComponentCveEdges holds the Gorm model for Postgres table `node_component_cve_edges`.
type NodeComponentCveEdges struct {
	Id                string         `gorm:"column:id;type:varchar;primaryKey"`
	IsFixable         bool           `gorm:"column:isfixable;type:bool"`
	FixedBy           string         `gorm:"column:fixedby;type:varchar"`
	ComponentId       string         `gorm:"column:componentid;type:varchar"`
	CveId             string         `gorm:"column:cveid;type:varchar"`
	Serialized        []byte         `gorm:"column:serialized;type:bytea"`
	NodeComponentsRef NodeComponents `gorm:"foreignKey:componentid;references:id;belongsTo;constraint:OnDelete:CASCADE"`
}
