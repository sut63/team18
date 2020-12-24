package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
)

// DataRoom holds the schema definition for the DataRoom entity.
type DataRoom struct {
	ent.Schema
}

// Fields of the DataRoom.
func (DataRoom) Fields() []ent.Field {
	return nil
}

// Edges of the DataRoom a.
func (DataRoom) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("reserves", ReserveRoom.Type).StorageKey(edge.Column("room id")),
	}
}
