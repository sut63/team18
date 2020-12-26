package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// DataRoom holds the schema definition for the DataRoom entity.
type DataRoom struct {
	ent.Schema
}

// Fields of the DataRoom.
func (DataRoom) Fields() []ent.Field {
	return []ent.Field{
		field.Float("price").Positive(),
		field.String("roomnumber").NotEmpty(),
	}
}

// Edges of the DataRoom.
func (DataRoom) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("reserves", ReserveRoom.Type).StorageKey(edge.Column("room id")),
		edge.To("fixs", ReserveRoom.Type).StorageKey(edge.Column("room id")),
		edge.From("promotion", Promotion.Type).Ref("datarooms").Unique(),
		edge.From("statusroom", StatusRoom.Type).Ref("datarooms").Unique(),
		edge.From("typeroom", TypeRoom.Type).Ref("datarooms").Unique(),
	}
}
