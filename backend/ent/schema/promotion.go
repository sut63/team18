package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
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
		edge.To("datarooms", DataRoom.Type).StorageKey(edge.Column("promotion id")),
	}
}
