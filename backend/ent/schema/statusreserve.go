package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// StatusReserve holds the schema definition for the StatusReserve entity.
type StatusReserve struct {
	ent.Schema
}

// Fields of the StatusReserve.
func (StatusReserve) Fields() []ent.Field {
	return []ent.Field{
		field.String("status_name").NotEmpty(),
	}
}

// Edges of the StatusReserve.
func (StatusReserve) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("reserves", ReserveRoom.Type).StorageKey(edge.Column("status_id")),
	}
}
