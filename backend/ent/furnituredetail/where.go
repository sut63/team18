// Code generated by entc, DO NOT EDIT.

package furnituredetail

import (
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/team18/app/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.FurnitureDetail {
	return predicate.FurnitureDetail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.FurnitureDetail {
	return predicate.FurnitureDetail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.FurnitureDetail {
	return predicate.FurnitureDetail(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.FurnitureDetail {
	return predicate.FurnitureDetail(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.FurnitureDetail {
	return predicate.FurnitureDetail(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.FurnitureDetail {
	return predicate.FurnitureDetail(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.FurnitureDetail {
	return predicate.FurnitureDetail(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.FurnitureDetail {
	return predicate.FurnitureDetail(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.FurnitureDetail {
	return predicate.FurnitureDetail(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// DateAdd applies equality check predicate on the "date_add" field. It's identical to DateAddEQ.
func DateAdd(v time.Time) predicate.FurnitureDetail {
	return predicate.FurnitureDetail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDateAdd), v))
	})
}

// FurnitureAmount applies equality check predicate on the "furniture_amount" field. It's identical to FurnitureAmountEQ.
func FurnitureAmount(v int) predicate.FurnitureDetail {
	return predicate.FurnitureDetail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFurnitureAmount), v))
	})
}

// FurnitureColour applies equality check predicate on the "furniture_colour" field. It's identical to FurnitureColourEQ.
func FurnitureColour(v string) predicate.FurnitureDetail {
	return predicate.FurnitureDetail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFurnitureColour), v))
	})
}

// Detail applies equality check predicate on the "detail" field. It's identical to DetailEQ.
func Detail(v string) predicate.FurnitureDetail {
	return predicate.FurnitureDetail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDetail), v))
	})
}

// DateAddEQ applies the EQ predicate on the "date_add" field.
func DateAddEQ(v time.Time) predicate.FurnitureDetail {
	return predicate.FurnitureDetail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDateAdd), v))
	})
}

// DateAddNEQ applies the NEQ predicate on the "date_add" field.
func DateAddNEQ(v time.Time) predicate.FurnitureDetail {
	return predicate.FurnitureDetail(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDateAdd), v))
	})
}

// DateAddIn applies the In predicate on the "date_add" field.
func DateAddIn(vs ...time.Time) predicate.FurnitureDetail {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FurnitureDetail(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDateAdd), v...))
	})
}

// DateAddNotIn applies the NotIn predicate on the "date_add" field.
func DateAddNotIn(vs ...time.Time) predicate.FurnitureDetail {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FurnitureDetail(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDateAdd), v...))
	})
}

// DateAddGT applies the GT predicate on the "date_add" field.
func DateAddGT(v time.Time) predicate.FurnitureDetail {
	return predicate.FurnitureDetail(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDateAdd), v))
	})
}

// DateAddGTE applies the GTE predicate on the "date_add" field.
func DateAddGTE(v time.Time) predicate.FurnitureDetail {
	return predicate.FurnitureDetail(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDateAdd), v))
	})
}

// DateAddLT applies the LT predicate on the "date_add" field.
func DateAddLT(v time.Time) predicate.FurnitureDetail {
	return predicate.FurnitureDetail(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDateAdd), v))
	})
}

// DateAddLTE applies the LTE predicate on the "date_add" field.
func DateAddLTE(v time.Time) predicate.FurnitureDetail {
	return predicate.FurnitureDetail(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDateAdd), v))
	})
}

// FurnitureAmountEQ applies the EQ predicate on the "furniture_amount" field.
func FurnitureAmountEQ(v int) predicate.FurnitureDetail {
	return predicate.FurnitureDetail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFurnitureAmount), v))
	})
}

// FurnitureAmountNEQ applies the NEQ predicate on the "furniture_amount" field.
func FurnitureAmountNEQ(v int) predicate.FurnitureDetail {
	return predicate.FurnitureDetail(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFurnitureAmount), v))
	})
}

// FurnitureAmountIn applies the In predicate on the "furniture_amount" field.
func FurnitureAmountIn(vs ...int) predicate.FurnitureDetail {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FurnitureDetail(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldFurnitureAmount), v...))
	})
}

// FurnitureAmountNotIn applies the NotIn predicate on the "furniture_amount" field.
func FurnitureAmountNotIn(vs ...int) predicate.FurnitureDetail {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FurnitureDetail(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldFurnitureAmount), v...))
	})
}

// FurnitureAmountGT applies the GT predicate on the "furniture_amount" field.
func FurnitureAmountGT(v int) predicate.FurnitureDetail {
	return predicate.FurnitureDetail(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFurnitureAmount), v))
	})
}

// FurnitureAmountGTE applies the GTE predicate on the "furniture_amount" field.
func FurnitureAmountGTE(v int) predicate.FurnitureDetail {
	return predicate.FurnitureDetail(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFurnitureAmount), v))
	})
}

// FurnitureAmountLT applies the LT predicate on the "furniture_amount" field.
func FurnitureAmountLT(v int) predicate.FurnitureDetail {
	return predicate.FurnitureDetail(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFurnitureAmount), v))
	})
}

// FurnitureAmountLTE applies the LTE predicate on the "furniture_amount" field.
func FurnitureAmountLTE(v int) predicate.FurnitureDetail {
	return predicate.FurnitureDetail(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFurnitureAmount), v))
	})
}

// FurnitureColourEQ applies the EQ predicate on the "furniture_colour" field.
func FurnitureColourEQ(v string) predicate.FurnitureDetail {
	return predicate.FurnitureDetail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFurnitureColour), v))
	})
}

// FurnitureColourNEQ applies the NEQ predicate on the "furniture_colour" field.
func FurnitureColourNEQ(v string) predicate.FurnitureDetail {
	return predicate.FurnitureDetail(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFurnitureColour), v))
	})
}

// FurnitureColourIn applies the In predicate on the "furniture_colour" field.
func FurnitureColourIn(vs ...string) predicate.FurnitureDetail {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FurnitureDetail(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldFurnitureColour), v...))
	})
}

// FurnitureColourNotIn applies the NotIn predicate on the "furniture_colour" field.
func FurnitureColourNotIn(vs ...string) predicate.FurnitureDetail {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FurnitureDetail(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldFurnitureColour), v...))
	})
}

// FurnitureColourGT applies the GT predicate on the "furniture_colour" field.
func FurnitureColourGT(v string) predicate.FurnitureDetail {
	return predicate.FurnitureDetail(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFurnitureColour), v))
	})
}

// FurnitureColourGTE applies the GTE predicate on the "furniture_colour" field.
func FurnitureColourGTE(v string) predicate.FurnitureDetail {
	return predicate.FurnitureDetail(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFurnitureColour), v))
	})
}

// FurnitureColourLT applies the LT predicate on the "furniture_colour" field.
func FurnitureColourLT(v string) predicate.FurnitureDetail {
	return predicate.FurnitureDetail(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFurnitureColour), v))
	})
}

// FurnitureColourLTE applies the LTE predicate on the "furniture_colour" field.
func FurnitureColourLTE(v string) predicate.FurnitureDetail {
	return predicate.FurnitureDetail(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFurnitureColour), v))
	})
}

// FurnitureColourContains applies the Contains predicate on the "furniture_colour" field.
func FurnitureColourContains(v string) predicate.FurnitureDetail {
	return predicate.FurnitureDetail(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldFurnitureColour), v))
	})
}

// FurnitureColourHasPrefix applies the HasPrefix predicate on the "furniture_colour" field.
func FurnitureColourHasPrefix(v string) predicate.FurnitureDetail {
	return predicate.FurnitureDetail(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldFurnitureColour), v))
	})
}

// FurnitureColourHasSuffix applies the HasSuffix predicate on the "furniture_colour" field.
func FurnitureColourHasSuffix(v string) predicate.FurnitureDetail {
	return predicate.FurnitureDetail(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldFurnitureColour), v))
	})
}

// FurnitureColourEqualFold applies the EqualFold predicate on the "furniture_colour" field.
func FurnitureColourEqualFold(v string) predicate.FurnitureDetail {
	return predicate.FurnitureDetail(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldFurnitureColour), v))
	})
}

// FurnitureColourContainsFold applies the ContainsFold predicate on the "furniture_colour" field.
func FurnitureColourContainsFold(v string) predicate.FurnitureDetail {
	return predicate.FurnitureDetail(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldFurnitureColour), v))
	})
}

// DetailEQ applies the EQ predicate on the "detail" field.
func DetailEQ(v string) predicate.FurnitureDetail {
	return predicate.FurnitureDetail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDetail), v))
	})
}

// DetailNEQ applies the NEQ predicate on the "detail" field.
func DetailNEQ(v string) predicate.FurnitureDetail {
	return predicate.FurnitureDetail(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDetail), v))
	})
}

// DetailIn applies the In predicate on the "detail" field.
func DetailIn(vs ...string) predicate.FurnitureDetail {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FurnitureDetail(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDetail), v...))
	})
}

// DetailNotIn applies the NotIn predicate on the "detail" field.
func DetailNotIn(vs ...string) predicate.FurnitureDetail {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FurnitureDetail(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDetail), v...))
	})
}

// DetailGT applies the GT predicate on the "detail" field.
func DetailGT(v string) predicate.FurnitureDetail {
	return predicate.FurnitureDetail(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDetail), v))
	})
}

// DetailGTE applies the GTE predicate on the "detail" field.
func DetailGTE(v string) predicate.FurnitureDetail {
	return predicate.FurnitureDetail(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDetail), v))
	})
}

// DetailLT applies the LT predicate on the "detail" field.
func DetailLT(v string) predicate.FurnitureDetail {
	return predicate.FurnitureDetail(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDetail), v))
	})
}

// DetailLTE applies the LTE predicate on the "detail" field.
func DetailLTE(v string) predicate.FurnitureDetail {
	return predicate.FurnitureDetail(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDetail), v))
	})
}

// DetailContains applies the Contains predicate on the "detail" field.
func DetailContains(v string) predicate.FurnitureDetail {
	return predicate.FurnitureDetail(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDetail), v))
	})
}

// DetailHasPrefix applies the HasPrefix predicate on the "detail" field.
func DetailHasPrefix(v string) predicate.FurnitureDetail {
	return predicate.FurnitureDetail(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDetail), v))
	})
}

// DetailHasSuffix applies the HasSuffix predicate on the "detail" field.
func DetailHasSuffix(v string) predicate.FurnitureDetail {
	return predicate.FurnitureDetail(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDetail), v))
	})
}

// DetailEqualFold applies the EqualFold predicate on the "detail" field.
func DetailEqualFold(v string) predicate.FurnitureDetail {
	return predicate.FurnitureDetail(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDetail), v))
	})
}

// DetailContainsFold applies the ContainsFold predicate on the "detail" field.
func DetailContainsFold(v string) predicate.FurnitureDetail {
	return predicate.FurnitureDetail(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDetail), v))
	})
}

// HasFixs applies the HasEdge predicate on the "fixs" edge.
func HasFixs() predicate.FurnitureDetail {
	return predicate.FurnitureDetail(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(FixsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, FixsTable, FixsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFixsWith applies the HasEdge predicate on the "fixs" edge with a given conditions (other predicates).
func HasFixsWith(preds ...predicate.FixRoom) predicate.FurnitureDetail {
	return predicate.FurnitureDetail(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(FixsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, FixsTable, FixsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasFurnitures applies the HasEdge predicate on the "furnitures" edge.
func HasFurnitures() predicate.FurnitureDetail {
	return predicate.FurnitureDetail(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(FurnituresTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, FurnituresTable, FurnituresColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFurnituresWith applies the HasEdge predicate on the "furnitures" edge with a given conditions (other predicates).
func HasFurnituresWith(preds ...predicate.Furniture) predicate.FurnitureDetail {
	return predicate.FurnitureDetail(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(FurnituresInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, FurnituresTable, FurnituresColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCounterstaffs applies the HasEdge predicate on the "counterstaffs" edge.
func HasCounterstaffs() predicate.FurnitureDetail {
	return predicate.FurnitureDetail(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CounterstaffsTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CounterstaffsTable, CounterstaffsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCounterstaffsWith applies the HasEdge predicate on the "counterstaffs" edge with a given conditions (other predicates).
func HasCounterstaffsWith(preds ...predicate.CounterStaff) predicate.FurnitureDetail {
	return predicate.FurnitureDetail(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CounterstaffsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CounterstaffsTable, CounterstaffsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTypes applies the HasEdge predicate on the "types" edge.
func HasTypes() predicate.FurnitureDetail {
	return predicate.FurnitureDetail(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TypesTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, TypesTable, TypesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTypesWith applies the HasEdge predicate on the "types" edge with a given conditions (other predicates).
func HasTypesWith(preds ...predicate.FurnitureType) predicate.FurnitureDetail {
	return predicate.FurnitureDetail(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TypesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, TypesTable, TypesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasRooms applies the HasEdge predicate on the "rooms" edge.
func HasRooms() predicate.FurnitureDetail {
	return predicate.FurnitureDetail(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RoomsTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, RoomsTable, RoomsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRoomsWith applies the HasEdge predicate on the "rooms" edge with a given conditions (other predicates).
func HasRoomsWith(preds ...predicate.DataRoom) predicate.FurnitureDetail {
	return predicate.FurnitureDetail(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RoomsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, RoomsTable, RoomsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.FurnitureDetail) predicate.FurnitureDetail {
	return predicate.FurnitureDetail(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.FurnitureDetail) predicate.FurnitureDetail {
	return predicate.FurnitureDetail(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.FurnitureDetail) predicate.FurnitureDetail {
	return predicate.FurnitureDetail(func(s *sql.Selector) {
		p(s.Not())
	})
}
