package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// StatusCheckIn holds the schema definition for the StatusCheckIn entity.
type StatusCheckIn struct {
	ent.Schema
}

// Fields of the StatusCheckIn.
func (StatusCheckIn) Fields() []ent.Field {
	return []ent.Field{
		field.String("status_name").NotEmpty().Unique(),
	}
}

// Edges of the StatusCheckIn.
func (StatusCheckIn) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("checkins", CheckIn.Type).StorageKey(edge.Column("status_id")),
	}
}
