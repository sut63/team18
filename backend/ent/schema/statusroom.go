package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// StatusRoom holds the schema definition for the StatusRoom entity.
type StatusRoom struct {
	ent.Schema
}

// Fields of the StatusRoom.
func (StatusRoom) Fields() []ent.Field {
	return []ent.Field{
		field.String("status_name").NotEmpty().Unique(),
	}
}

// Edges of the StatusRoom.
func (StatusRoom) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("datarooms", DataRoom.Type).StorageKey(edge.Column("statusroom_id")),
	}
}
