package schema

import (
	"regexp"

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
		field.String("mobile_key").MinLen(10).MaxLen(10),
		field.String("phone_number").Match(regexp.MustCompile("[0]\\d{9}")).MinLen(10).MaxLen(10),
		field.String("person_number").Match(regexp.MustCompile("[0-9]\\d{12}")).MinLen(13).MaxLen(13),
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
