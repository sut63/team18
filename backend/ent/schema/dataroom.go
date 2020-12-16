package schema

import "github.com/facebookincubator/ent"

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
	return nil
}
