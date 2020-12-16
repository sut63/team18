package schema

import "github.com/facebookincubator/ent"

// CheckIn holds the schema definition for the CheckIn entity.
type CheckIn struct {
	ent.Schema
}

// Fields of the CheckIn.
func (CheckIn) Fields() []ent.Field {
	return nil
}

// Edges of the CheckIn.
func (CheckIn) Edges() []ent.Edge {
	return nil
}
