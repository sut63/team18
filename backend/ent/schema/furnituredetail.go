package schema

import "github.com/facebookincubator/ent"

// FurnitureDetail holds the schema definition for the FurnitureDetail entity.
type FurnitureDetail struct {
	ent.Schema
}

// Fields of the FurnitureDetail.
func (FurnitureDetail) Fields() []ent.Field {
	return nil
}

// Edges of the FurnitureDetail.
func (FurnitureDetail) Edges() []ent.Edge {
	return nil
}
