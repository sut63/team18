package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Customer holds the schema definition for the Customer entity.
type Customer struct {
	ent.Schema
}

// Fields of the Customer.
func (Customer) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique().NotEmpty(),
		field.String("email").Unique().NotEmpty(),
		field.String("password").Unique().NotEmpty(),
	}
}

// Edges of the Customer.
func (Customer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("reserves", ReserveRoom.Type).StorageKey(edge.Column("customer_id")),
		edge.To("fixs", FixRoom.Type).StorageKey(edge.Column("customer_id")),
		edge.To("checkins", CheckIn.Type).StorageKey(edge.Column("customer_id")),
	}
}
