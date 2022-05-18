package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Community holds the schema definition for the Community entity.
type Community struct {
	ent.Schema
}

// Fields of the Community.
func (Community) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.String("name"),
		field.Int("province_id"),
		field.Int("city_id"),
		field.Int("area_id"),
		field.Int("area_num"),
		field.Int("admin_id"),
		field.Int("comm_number"),
		field.Int("state"),
		field.Time("start_time"),
		field.Time("end_time"),
		field.Time("created_at"),
		field.Time("updated_at"),
		field.Time("deleted_at"),
	}
}

// Edges of the Community.
func (Community) Edges() []ent.Edge {
	return nil
}
