package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// FixRoom holds the schema definition for the FixRoom entity.
type FixRoom struct {
	ent.Schema
}

// Fields of the FixRoom.
func (FixRoom) Fields() []ent.Field {
	return []ent.Field{
		field.String("fix_detail"),
	}
}

// Edges of the FixRoom.
func (FixRoom) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("customer", Customer.Type).Ref("fixs").Unique(),
		edge.From("furnitureDetail", FurnitureDetail.Type).Ref("fixs").Unique(),
		edge.From("room", DataRoom.Type).Ref("fixs").Unique(),
	}
}
