package schema

import (
	"time"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// CheckIn holds the schema definition for the CheckIn entity.
type CheckIn struct {
	ent.Schema
}

// Fields of the CheckIn.
func (CheckIn) Fields() []ent.Field {
	return []ent.Field{
		field.Time("checkin_date").Default(time.Now().Local),
	}
}

// Edges of the CheckIn.
func (CheckIn) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("customer",Customer.Type).Ref("checkins").Unique(),
		edge.From("counter",CounterStaff.Type).Ref("checkins").Unique(),
		edge.From("reserveroom",ReserveRoom.Type).Ref("checkins").Unique(),
		edge.To("checkouts", Checkout.Type).Unique(),
	}
}
