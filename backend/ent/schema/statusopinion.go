package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// StatusOpinion holds the schema definition for the StatusOpinion entity.
type StatusOpinion struct {
	ent.Schema
}

// Fields of the StatusOpinion.
func (StatusOpinion) Fields() []ent.Field {
	return []ent.Field{
		field.String("opinion").NotEmpty().Unique(),
	}
}

// Edges of the StatusOpinion.
func (StatusOpinion) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("checkouts", Checkout.Type),
	}
}
