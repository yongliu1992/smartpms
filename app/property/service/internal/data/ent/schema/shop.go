package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Shop holds the schema definition for the Shop entity.
type Shop struct {
	ent.Schema
}

// Fields of the Shop.
func (Shop) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.String("floor_id"),
		field.Int("room_num"),
		field.Int("layer"),
		field.Float32("built_up_area"),
		field.Int("community_id"),
		field.Float32("fee_rate"),
		field.Float32("room_area"),
		field.Float32("rent"),
		field.String("remark"),
		field.Time("created_at"),
		field.Time("updated_at"),
		field.Time("deleted_at"),
	}
}

// Edges of the Shop.
func (Shop) Edges() []ent.Edge {
	return nil
}
