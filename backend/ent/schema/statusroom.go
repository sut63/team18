package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// StatusRoom holds the schema definition for the StatusRoom entity.
type StatusRoom struct {
	ent.Schema
}

// Fields of the StatusRoom.
func (StatusRoom) Fields() []ent.Field {
	return []ent.Field{
		field.String("statusname").NotEmpty(),
	}
}

// Edges of the StatusRoom.
func (StatusRoom) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("datarooms", DataRoom.Type).StorageKey(edge.Column("statusroom id")),
	}
}
