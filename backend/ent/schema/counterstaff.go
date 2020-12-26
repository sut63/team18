package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
	"github.com/facebook/ent/schema/edge"
)

// CounterStaff holds the schema definition for the CounterStaff entity.
type CounterStaff struct {
	ent.Schema
}

// Fields of the CounterStaff.
func (CounterStaff) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique().NotEmpty(),
		field.String("password").Unique().NotEmpty(),
	}
}

// Edges of the CounterStaff.
func (CounterStaff) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("checkins", CheckIn.Type).StorageKey(edge.Column("staff id")),
		edge.To("checkouts", Checkout.Type).StorageKey(edge.Column("staff id")),		
	}
}
