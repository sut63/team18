package schema

import (
	"errors"
	"regexp"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// DataRoom holds the schema definition for the DataRoom entity.
type DataRoom struct {
	ent.Schema
}

// Fields of the DataRoom.
func (DataRoom) Fields() []ent.Field {
	return []ent.Field{
		field.Float("price").Min(0).Positive(),
		field.String("roomnumber").Validate(func(s string) error {
			match, _ := regexp.MatchString("[A-Z]\\d3", s)
			if !match {
				return errors.New("รูปแบบหมายเลขห้องพักไม่ถูกต้อง")
			}
			return nil
		}),
		field.String("roomdetail").MaxLen(70),
	}
}

// Edges of the DataRoom.
func (DataRoom) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("reserves", ReserveRoom.Type).StorageKey(edge.Column("room_id")),
		edge.To("fixs", FixRoom.Type).StorageKey(edge.Column("room_id")),
		edge.To("details", FurnitureDetail.Type).StorageKey(edge.Column("room_id")),
		edge.To("checkins", CheckIn.Type).StorageKey(edge.Column("room_id")),

		edge.From("promotion", Promotion.Type).Ref("datarooms").Unique(),
		edge.From("statusroom", StatusRoom.Type).Ref("datarooms").Unique(),
		edge.From("typeroom", TypeRoom.Type).Ref("datarooms").Unique(),
	}
}
