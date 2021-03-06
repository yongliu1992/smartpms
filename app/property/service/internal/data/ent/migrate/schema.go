// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CommunitiesColumns holds the columns for the "communities" table.
	CommunitiesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "province_id", Type: field.TypeInt},
		{Name: "city_id", Type: field.TypeInt},
		{Name: "area_id", Type: field.TypeInt},
		{Name: "area_num", Type: field.TypeInt},
		{Name: "admin_id", Type: field.TypeInt},
		{Name: "comm_number", Type: field.TypeInt},
		{Name: "state", Type: field.TypeInt},
		{Name: "start_time", Type: field.TypeTime},
		{Name: "end_time", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime},
	}
	// CommunitiesTable holds the schema information for the "communities" table.
	CommunitiesTable = &schema.Table{
		Name:       "communities",
		Columns:    CommunitiesColumns,
		PrimaryKey: []*schema.Column{CommunitiesColumns[0]},
	}
	// ShopsColumns holds the columns for the "shops" table.
	ShopsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "floor_id", Type: field.TypeString},
		{Name: "room_num", Type: field.TypeInt},
		{Name: "layer", Type: field.TypeInt},
		{Name: "built_up_area", Type: field.TypeFloat32},
		{Name: "community_id", Type: field.TypeInt},
		{Name: "fee_rate", Type: field.TypeFloat32},
		{Name: "room_area", Type: field.TypeFloat32},
		{Name: "rent", Type: field.TypeFloat32},
		{Name: "remark", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime},
	}
	// ShopsTable holds the schema information for the "shops" table.
	ShopsTable = &schema.Table{
		Name:       "shops",
		Columns:    ShopsColumns,
		PrimaryKey: []*schema.Column{ShopsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CommunitiesTable,
		ShopsTable,
	}
)

func init() {
}
