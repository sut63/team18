package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// ReserveRoom holds the schema definition for the ReserveRoom entity.
type ReserveRoom struct {
	ent.Schema
}

// Fields of the ReserveRoom.
func (ReserveRoom) Fields() []ent.Field {
	return []ent.Field{
		field.Time("reserve_date"),
		field.Time("reserve_date"),
		field.Float("net_price").Positive(),
	}
}

// Edges of the ReserveRoom.
func (ReserveRoom) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("customer",Customer.Type).Ref("reserves").Unique(),
		edge.From("promotion",Promotion.Type).Ref("reserves").Unique(),
		edge.From("room",DataRoom.Type).Ref("reserves").Unique(),
		edge.To("checkins", CheckIn.Type).StorageKey(edge.Column("reserves id")),
	}
}
