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
	// CreateTablePolicyCategoryEdgesStmt holds the create statement for table `policy_category_edges`.
	CreateTablePolicyCategoryEdgesStmt = &postgres.CreateStmts{
		GormModel: (*PolicyCategoryEdges)(nil),
		Children:  []*postgres.CreateStmts{},
	}

	// PolicyCategoryEdgesSchema is the go schema for table `policy_category_edges`.
	PolicyCategoryEdgesSchema = func() *walker.Schema {
		schema := GetSchemaForTable("policy_category_edges")
		if schema != nil {
			return schema
		}
		schema = walker.Walk(reflect.TypeOf((*storage.PolicyCategoryEdge)(nil)), "policy_category_edges")
		referencedSchemas := map[string]*walker.Schema{
			"storage.Policy":         PoliciesSchema,
			"storage.PolicyCategory": PolicyCategoriesSchema,
		}

		schema.ResolveReferences(func(messageTypeName string) *walker.Schema {
			return referencedSchemas[fmt.Sprintf("storage.%s", messageTypeName)]
		})
		schema.ScopingResource = &resources.Policy
		schema.SetOptionsMap(search.Walk(v1.SearchCategory_POLICY_CATEGORY_EDGE, "policycategoryedge", (*storage.PolicyCategoryEdge)(nil)))
		schema.SetSearchScope([]v1.SearchCategory{
			v1.SearchCategory_POLICY_CATEGORY_EDGE,
			v1.SearchCategory_POLICY_CATEGORIES,
		}...)
		RegisterTable(schema, CreateTablePolicyCategoryEdgesStmt)
		mapping.RegisterCategoryToTable(v1.SearchCategory_POLICY_CATEGORY_EDGE, schema)
		return schema
	}()
)

const (
	PolicyCategoryEdgesTableName = "policy_category_edges"
)

// PolicyCategoryEdges holds the Gorm model for Postgres table `policy_category_edges`.
type PolicyCategoryEdges struct {
	Id                  string           `gorm:"column:id;type:varchar;primaryKey"`
	PolicyId            string           `gorm:"column:policyid;type:varchar"`
	CategoryId          string           `gorm:"column:categoryid;type:varchar"`
	Serialized          []byte           `gorm:"column:serialized;type:bytea"`
	PoliciesRef         Policies         `gorm:"foreignKey:policyid;references:id;belongsTo;constraint:OnDelete:CASCADE"`
	PolicyCategoriesRef PolicyCategories `gorm:"foreignKey:categoryid;references:id;belongsTo;constraint:OnDelete:CASCADE"`
}
