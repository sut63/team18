package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// CheckIn holds the schema definition for the CheckIn entity.
type CheckIn struct {
	ent.Schema
}

// Fields of the CheckIn.
func (CheckIn) Fields() []ent.Field {
	return []ent.Field{
		field.Time("checkin_date"),
	}
}

// Edges of the CheckIn.
func (CheckIn) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("customer", Customer.Type).Ref("checkins").Unique(),
		edge.From("counter", CounterStaff.Type).Ref("checkins").Unique(),
		edge.From("reserveroom", ReserveRoom.Type).Ref("checkins").Unique(),
		edge.From("dataroom", DataRoom.Type).Ref("checkins").Unique(),
		edge.From("status", StatusCheckIn.Type).Ref("checkins").Unique(),

		edge.To("checkouts", Checkout.Type).Unique(),
	}
}
