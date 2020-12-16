package schema

import "github.com/facebookincubator/ent"

// FixRoom holds the schema definition for the FixRoom entity.
//V2
type FixRoom struct {
	ent.Schema
}

// Fields of the FixRoom.
func (FixRoom) Fields() []ent.Field {
	return nil
}

// Edges of the FixRoom.
func (FixRoom) Edges() []ent.Edge {
	return nil
}
