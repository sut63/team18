package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Furniture holds the schema definition for the Furniture entity.
type Furniture struct {
	ent.Schema
}

// Fields of the Furniture.
func (Furniture) Fields() []ent.Field {
	return []ent.Field{
		field.String("furniture_name").NotEmpty().Unique(),
	}
}

// Edges of the Furniture.
func (Furniture) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("details", FurnitureDetail.Type).StorageKey(edge.Column("furniture_id")),
	}
}
