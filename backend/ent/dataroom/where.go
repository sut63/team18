// Code generated by entc, DO NOT EDIT.

package dataroom

import (
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/team18/app/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.DataRoom {
	return predicate.DataRoom(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.DataRoom {
	return predicate.DataRoom(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.DataRoom {
	return predicate.DataRoom(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.DataRoom {
	return predicate.DataRoom(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.DataRoom {
	return predicate.DataRoom(func(s *sql.Selector) {
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
func IDGT(id int) predicate.DataRoom {
	return predicate.DataRoom(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.DataRoom {
	return predicate.DataRoom(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.DataRoom {
	return predicate.DataRoom(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.DataRoom {
	return predicate.DataRoom(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Price applies equality check predicate on the "price" field. It's identical to PriceEQ.
func Price(v float64) predicate.DataRoom {
	return predicate.DataRoom(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPrice), v))
	})
}

// Roomnumber applies equality check predicate on the "roomnumber" field. It's identical to RoomnumberEQ.
func Roomnumber(v string) predicate.DataRoom {
	return predicate.DataRoom(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRoomnumber), v))
	})
}

// PriceEQ applies the EQ predicate on the "price" field.
func PriceEQ(v float64) predicate.DataRoom {
	return predicate.DataRoom(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPrice), v))
	})
}

// PriceNEQ applies the NEQ predicate on the "price" field.
func PriceNEQ(v float64) predicate.DataRoom {
	return predicate.DataRoom(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPrice), v))
	})
}

// PriceIn applies the In predicate on the "price" field.
func PriceIn(vs ...float64) predicate.DataRoom {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DataRoom(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPrice), v...))
	})
}

// PriceNotIn applies the NotIn predicate on the "price" field.
func PriceNotIn(vs ...float64) predicate.DataRoom {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DataRoom(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPrice), v...))
	})
}

// PriceGT applies the GT predicate on the "price" field.
func PriceGT(v float64) predicate.DataRoom {
	return predicate.DataRoom(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPrice), v))
	})
}

// PriceGTE applies the GTE predicate on the "price" field.
func PriceGTE(v float64) predicate.DataRoom {
	return predicate.DataRoom(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPrice), v))
	})
}

// PriceLT applies the LT predicate on the "price" field.
func PriceLT(v float64) predicate.DataRoom {
	return predicate.DataRoom(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPrice), v))
	})
}

// PriceLTE applies the LTE predicate on the "price" field.
func PriceLTE(v float64) predicate.DataRoom {
	return predicate.DataRoom(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPrice), v))
	})
}

// RoomnumberEQ applies the EQ predicate on the "roomnumber" field.
func RoomnumberEQ(v string) predicate.DataRoom {
	return predicate.DataRoom(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRoomnumber), v))
	})
}

// RoomnumberNEQ applies the NEQ predicate on the "roomnumber" field.
func RoomnumberNEQ(v string) predicate.DataRoom {
	return predicate.DataRoom(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRoomnumber), v))
	})
}

// RoomnumberIn applies the In predicate on the "roomnumber" field.
func RoomnumberIn(vs ...string) predicate.DataRoom {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DataRoom(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldRoomnumber), v...))
	})
}

// RoomnumberNotIn applies the NotIn predicate on the "roomnumber" field.
func RoomnumberNotIn(vs ...string) predicate.DataRoom {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DataRoom(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRoomnumber), v...))
	})
}

// RoomnumberGT applies the GT predicate on the "roomnumber" field.
func RoomnumberGT(v string) predicate.DataRoom {
	return predicate.DataRoom(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRoomnumber), v))
	})
}

// RoomnumberGTE applies the GTE predicate on the "roomnumber" field.
func RoomnumberGTE(v string) predicate.DataRoom {
	return predicate.DataRoom(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRoomnumber), v))
	})
}

// RoomnumberLT applies the LT predicate on the "roomnumber" field.
func RoomnumberLT(v string) predicate.DataRoom {
	return predicate.DataRoom(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRoomnumber), v))
	})
}

// RoomnumberLTE applies the LTE predicate on the "roomnumber" field.
func RoomnumberLTE(v string) predicate.DataRoom {
	return predicate.DataRoom(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRoomnumber), v))
	})
}

// RoomnumberContains applies the Contains predicate on the "roomnumber" field.
func RoomnumberContains(v string) predicate.DataRoom {
	return predicate.DataRoom(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldRoomnumber), v))
	})
}

// RoomnumberHasPrefix applies the HasPrefix predicate on the "roomnumber" field.
func RoomnumberHasPrefix(v string) predicate.DataRoom {
	return predicate.DataRoom(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldRoomnumber), v))
	})
}

// RoomnumberHasSuffix applies the HasSuffix predicate on the "roomnumber" field.
func RoomnumberHasSuffix(v string) predicate.DataRoom {
	return predicate.DataRoom(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldRoomnumber), v))
	})
}

// RoomnumberEqualFold applies the EqualFold predicate on the "roomnumber" field.
func RoomnumberEqualFold(v string) predicate.DataRoom {
	return predicate.DataRoom(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldRoomnumber), v))
	})
}

// RoomnumberContainsFold applies the ContainsFold predicate on the "roomnumber" field.
func RoomnumberContainsFold(v string) predicate.DataRoom {
	return predicate.DataRoom(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldRoomnumber), v))
	})
}

// HasReserves applies the HasEdge predicate on the "reserves" edge.
func HasReserves() predicate.DataRoom {
	return predicate.DataRoom(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ReservesTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ReservesTable, ReservesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasReservesWith applies the HasEdge predicate on the "reserves" edge with a given conditions (other predicates).
func HasReservesWith(preds ...predicate.ReserveRoom) predicate.DataRoom {
	return predicate.DataRoom(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ReservesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ReservesTable, ReservesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasFixs applies the HasEdge predicate on the "fixs" edge.
func HasFixs() predicate.DataRoom {
	return predicate.DataRoom(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(FixsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, FixsTable, FixsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFixsWith applies the HasEdge predicate on the "fixs" edge with a given conditions (other predicates).
func HasFixsWith(preds ...predicate.FixRoom) predicate.DataRoom {
	return predicate.DataRoom(func(s *sql.Selector) {
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

// HasDetails applies the HasEdge predicate on the "details" edge.
func HasDetails() predicate.DataRoom {
	return predicate.DataRoom(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DetailsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, DetailsTable, DetailsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDetailsWith applies the HasEdge predicate on the "details" edge with a given conditions (other predicates).
func HasDetailsWith(preds ...predicate.FurnitureDetail) predicate.DataRoom {
	return predicate.DataRoom(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DetailsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, DetailsTable, DetailsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCheckins applies the HasEdge predicate on the "checkins" edge.
func HasCheckins() predicate.DataRoom {
	return predicate.DataRoom(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CheckinsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, CheckinsTable, CheckinsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCheckinsWith applies the HasEdge predicate on the "checkins" edge with a given conditions (other predicates).
func HasCheckinsWith(preds ...predicate.CheckIn) predicate.DataRoom {
	return predicate.DataRoom(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CheckinsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, CheckinsTable, CheckinsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPromotion applies the HasEdge predicate on the "promotion" edge.
func HasPromotion() predicate.DataRoom {
	return predicate.DataRoom(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PromotionTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PromotionTable, PromotionColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPromotionWith applies the HasEdge predicate on the "promotion" edge with a given conditions (other predicates).
func HasPromotionWith(preds ...predicate.Promotion) predicate.DataRoom {
	return predicate.DataRoom(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PromotionInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PromotionTable, PromotionColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasStatusroom applies the HasEdge predicate on the "statusroom" edge.
func HasStatusroom() predicate.DataRoom {
	return predicate.DataRoom(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(StatusroomTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, StatusroomTable, StatusroomColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasStatusroomWith applies the HasEdge predicate on the "statusroom" edge with a given conditions (other predicates).
func HasStatusroomWith(preds ...predicate.StatusRoom) predicate.DataRoom {
	return predicate.DataRoom(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(StatusroomInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, StatusroomTable, StatusroomColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTyperoom applies the HasEdge predicate on the "typeroom" edge.
func HasTyperoom() predicate.DataRoom {
	return predicate.DataRoom(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TyperoomTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, TyperoomTable, TyperoomColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTyperoomWith applies the HasEdge predicate on the "typeroom" edge with a given conditions (other predicates).
func HasTyperoomWith(preds ...predicate.TypeRoom) predicate.DataRoom {
	return predicate.DataRoom(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TyperoomInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, TyperoomTable, TyperoomColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.DataRoom) predicate.DataRoom {
	return predicate.DataRoom(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.DataRoom) predicate.DataRoom {
	return predicate.DataRoom(func(s *sql.Selector) {
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
func Not(p predicate.DataRoom) predicate.DataRoom {
	return predicate.DataRoom(func(s *sql.Selector) {
		p(s.Not())
	})
}
