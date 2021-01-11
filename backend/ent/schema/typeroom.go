package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// TypeRoom holds the schema definition for the TypeRoom entity.
type TypeRoom struct {
	ent.Schema
}

// Fields of the TypeRoom.
func (TypeRoom) Fields() []ent.Field {
	return []ent.Field{
		field.String("type_name").NotEmpty().Unique(),
	}
}

// Edges of the TypeRoom.
func (TypeRoom) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("datarooms", DataRoom.Type).StorageKey(edge.Column("typeroom_id")),
	}
}
