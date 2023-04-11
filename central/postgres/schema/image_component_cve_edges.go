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
	// CreateTableImageComponentCveEdgesStmt holds the create statement for table `image_component_cve_edges`.
	CreateTableImageComponentCveEdgesStmt = &postgres.CreateStmts{
		GormModel: (*ImageComponentCveEdges)(nil),
		Children:  []*postgres.CreateStmts{},
	}

	// ImageComponentCveEdgesSchema is the go schema for table `image_component_cve_edges`.
	ImageComponentCveEdgesSchema = func() *walker.Schema {
		schema := schemaPkg.GetSchemaForTable("image_component_cve_edges")
		if schema != nil {
			return schema
		}
		schema = walker.Walk(reflect.TypeOf((*storage.ComponentCVEEdge)(nil)), "image_component_cve_edges")
		referencedSchemas := map[string]*walker.Schema{
			"storage.ImageComponent": ImageComponentsSchema,
			"storage.ImageCVE":       ImageCvesSchema,
		}

		schema.ResolveReferences(func(messageTypeName string) *walker.Schema {
			return referencedSchemas[fmt.Sprintf("storage.%s", messageTypeName)]
		})
		schema.SetOptionsMap(search.Walk(v1.SearchCategory_COMPONENT_VULN_EDGE, "componentcveedge", (*storage.ComponentCVEEdge)(nil)))
		schema.SetSearchScope([]v1.SearchCategory{
			v1.SearchCategory_IMAGE_VULNERABILITIES,
			v1.SearchCategory_COMPONENT_VULN_EDGE,
			v1.SearchCategory_IMAGE_COMPONENTS,
			v1.SearchCategory_IMAGE_COMPONENT_EDGE,
			v1.SearchCategory_IMAGE_VULN_EDGE,
			v1.SearchCategory_IMAGES,
			v1.SearchCategory_DEPLOYMENTS,
			v1.SearchCategory_NAMESPACES,
			v1.SearchCategory_CLUSTERS,
		}...)
		schemaPkg.RegisterTable(schema, CreateTableImageComponentCveEdgesStmt)
		mapping.RegisterCategoryToTable(v1.SearchCategory_COMPONENT_VULN_EDGE, schema)
		return schema
	}()
)

const (
	ImageComponentCveEdgesTableName = "image_component_cve_edges"
)

// ImageComponentCveEdges holds the Gorm model for Postgres table `image_component_cve_edges`.
type ImageComponentCveEdges struct {
	Id                 string          `gorm:"column:id;type:varchar;primaryKey"`
	IsFixable          bool            `gorm:"column:isfixable;type:bool"`
	FixedBy            string          `gorm:"column:fixedby;type:varchar"`
	ImageComponentId   string          `gorm:"column:imagecomponentid;type:varchar;index:imagecomponentcveedges_imagecomponentid,type:hash"`
	ImageCveId         string          `gorm:"column:imagecveid;type:varchar;index:imagecomponentcveedges_imagecveid,type:hash"`
	Serialized         []byte          `gorm:"column:serialized;type:bytea"`
	ImageComponentsRef ImageComponents `gorm:"foreignKey:imagecomponentid;references:id;belongsTo;constraint:OnDelete:CASCADE"`
}
