// Code generated by pg-bindings generator. DO NOT EDIT.

package schema

import (
	"reflect"

	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/postgres"
	schemaPkg "github.com/stackrox/rox/pkg/postgres/schema"
	"github.com/stackrox/rox/pkg/postgres/walker"
)

var (
	// CreateTableNotificationSchedulesStmt holds the create statement for table `notification_schedules`.
	CreateTableNotificationSchedulesStmt = &postgres.CreateStmts{
		GormModel: (*NotificationSchedules)(nil),
		Children:  []*postgres.CreateStmts{},
	}

	// NotificationSchedulesSchema is the go schema for table `notification_schedules`.
	NotificationSchedulesSchema = func() *walker.Schema {
		schema := schemaPkg.GetSchemaForTable("notification_schedules")
		if schema != nil {
			return schema
		}
		schema = walker.Walk(reflect.TypeOf((*storage.NotificationSchedule)(nil)), "notification_schedules")
		schemaPkg.RegisterTable(schema, CreateTableNotificationSchedulesStmt)
		return schema
	}()
)

const (
	NotificationSchedulesTableName = "notification_schedules"
)

// NotificationSchedules holds the Gorm model for Postgres table `notification_schedules`.
type NotificationSchedules struct {
	Serialized []byte `gorm:"column:serialized;type:bytea"`
}
