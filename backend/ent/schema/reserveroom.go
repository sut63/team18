package schema

import "github.com/facebookincubator/ent"

// ReserveRoom holds the schema definition for the ReserveRoom entity.
type ReserveRoom struct {
	ent.Schema
}

// Fields of the ReserveRoom.
func (ReserveRoom) Fields() []ent.Field {
	return nil //add1
}

// Edges of the ReserveRoom.
func (ReserveRoom) Edges() []ent.Edge {
	return nil
}
