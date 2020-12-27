package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// FurnitureType holds the schema definition for the FurnitureType entity.
type FurnitureType struct {
	ent.Schema
}

// Fields of the FurnitureType.
func (FurnitureType) Fields() []ent.Field {
	return []ent.Field{
		field.String("furniture_type"),
	}
}

// Edges of the FurnitureType.
func (FurnitureType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("details", FurnitureDetail.Type).StorageKey(edge.Column("type_id")),
	}
}
