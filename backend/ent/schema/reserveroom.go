package schema

import (
	"regexp"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// ReserveRoom holds the schema definition for the ReserveRoom entity.
type ReserveRoom struct {
	ent.Schema
}

// Fields of the ReserveRoom.
func (ReserveRoom) Fields() []ent.Field {
	return []ent.Field{
		field.Time("reserve_date"),
		field.String("request").MaxLen(50),
		field.Int("amount").Min(1).Max(5).Positive(),
		field.String("phone_number").Match(regexp.MustCompile("[0]\\d{9}")).MinLen(10).MaxLen(10),
		field.Float("net_price").Positive(),
	}
}

// Edges of the ReserveRoom.
func (ReserveRoom) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("customer", Customer.Type).Ref("reserves").Unique(),
		edge.From("promotion", Promotion.Type).Ref("reserves").Unique(),
		edge.From("room", DataRoom.Type).Ref("reserves").Unique(),
		edge.From("status", StatusReserve.Type).Ref("reserves").Unique(),
		edge.To("checkins", CheckIn.Type).StorageKey(edge.Column("reserves_id")),
	}
}
