package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// FurnitureDetail holds the schema definition for the FurnitureDetail entity.
type FurnitureDetail struct {
	ent.Schema
}

// Fields of the FurnitureDetail.
func (FurnitureDetail) Fields() []ent.Field {
	return []ent.Field{
		field.Time("date_add"),
		field.Int("furniture_amount").Max(10).Positive(),
		field.String("furniture_colour").MaxLen(10),
		field.String("detail").Unique().MaxLen(50),
	}
}

// Edges of the FurnitureDetail.
func (FurnitureDetail) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("fixs", FixRoom.Type).StorageKey(edge.Column("object_id")),

		edge.From("furnitures", Furniture.Type).Ref("details").Unique(),
		edge.From("counterstaffs", CounterStaff.Type).Ref("details").Unique(),
		edge.From("types", FurnitureType.Type).Ref("details").Unique(),
		edge.From("rooms", DataRoom.Type).Ref("details").Unique(),
	}
}
