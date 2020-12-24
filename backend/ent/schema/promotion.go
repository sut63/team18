package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Promotion holds the schema definition for the Promotion entity.
type Promotion struct {
	ent.Schema
}

// Fields of the Promotion.
func (Promotion) Fields() []ent.Field {
	return []ent.Field{
		field.String("promotion_name").Unique().NotEmpty(),
		field.Float("discount").Positive(),
	}
}

// Edges of the Promotion.
func (Promotion) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("reserves", ReserveRoom.Type).StorageKey(edge.Column("promotion id")),
	}
}
