package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Checkout holds the schema definition for the Checkout entity.
type Checkout struct {
	ent.Schema
}

// Fields of the Checkout.
func (Checkout) Fields() []ent.Field {
	return []ent.Field{
		field.Time("checkout_date"),
		field.String("identity_card").MinLen(13).MaxLen(13),
		field.Float("price").Min(0).Positive(),
		field.String("comment").MaxLen(70),
	}
}

// Edges of the Checkout.
func (Checkout) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("statuss", Status.Type).Ref("checkouts").Unique(),
		edge.From("statusopinion", StatusOpinion.Type).Ref("checkouts").Unique(),
		edge.From("counterstaffs", CounterStaff.Type).Ref("checkouts").Unique(),
		edge.From("checkins", CheckIn.Type).Ref("checkouts").Unique().Required(),
	}
}
