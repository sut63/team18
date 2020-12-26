package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// TypeRoom holds the schema definition for the TypeRoom entity.
type TypeRoom struct {
	ent.Schema
}

// Fields of the TypeRoom.
func (TypeRoom) Fields() []ent.Field {
	return []ent.Field{
		field.String("type_name").NotEmpty(),
	}
}

// Edges of the TypeRoom.
func (TypeRoom) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("datarooms", DataRoom.Type).StorageKey(edge.Column("typeroom id")),
	}
}
