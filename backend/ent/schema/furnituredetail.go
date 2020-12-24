package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
)

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
	return []ent.Edge{
		edge.To("fixs", ReserveRoom.Type).StorageKey(edge.Column("object id")),
	}
}
