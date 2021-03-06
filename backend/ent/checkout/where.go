// Code generated by entc, DO NOT EDIT.

package checkout

import (
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/team18/app/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CheckoutDate applies equality check predicate on the "checkout_date" field. It's identical to CheckoutDateEQ.
func CheckoutDate(v time.Time) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCheckoutDate), v))
	})
}

// IdentityCard applies equality check predicate on the "identity_card" field. It's identical to IdentityCardEQ.
func IdentityCard(v string) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIdentityCard), v))
	})
}

// Price applies equality check predicate on the "price" field. It's identical to PriceEQ.
func Price(v float64) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPrice), v))
	})
}

// Comment applies equality check predicate on the "comment" field. It's identical to CommentEQ.
func Comment(v string) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldComment), v))
	})
}

// CheckoutDateEQ applies the EQ predicate on the "checkout_date" field.
func CheckoutDateEQ(v time.Time) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCheckoutDate), v))
	})
}

// CheckoutDateNEQ applies the NEQ predicate on the "checkout_date" field.
func CheckoutDateNEQ(v time.Time) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCheckoutDate), v))
	})
}

// CheckoutDateIn applies the In predicate on the "checkout_date" field.
func CheckoutDateIn(vs ...time.Time) predicate.Checkout {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Checkout(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCheckoutDate), v...))
	})
}

// CheckoutDateNotIn applies the NotIn predicate on the "checkout_date" field.
func CheckoutDateNotIn(vs ...time.Time) predicate.Checkout {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Checkout(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCheckoutDate), v...))
	})
}

// CheckoutDateGT applies the GT predicate on the "checkout_date" field.
func CheckoutDateGT(v time.Time) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCheckoutDate), v))
	})
}

// CheckoutDateGTE applies the GTE predicate on the "checkout_date" field.
func CheckoutDateGTE(v time.Time) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCheckoutDate), v))
	})
}

// CheckoutDateLT applies the LT predicate on the "checkout_date" field.
func CheckoutDateLT(v time.Time) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCheckoutDate), v))
	})
}

// CheckoutDateLTE applies the LTE predicate on the "checkout_date" field.
func CheckoutDateLTE(v time.Time) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCheckoutDate), v))
	})
}

// IdentityCardEQ applies the EQ predicate on the "identity_card" field.
func IdentityCardEQ(v string) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIdentityCard), v))
	})
}

// IdentityCardNEQ applies the NEQ predicate on the "identity_card" field.
func IdentityCardNEQ(v string) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIdentityCard), v))
	})
}

// IdentityCardIn applies the In predicate on the "identity_card" field.
func IdentityCardIn(vs ...string) predicate.Checkout {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Checkout(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldIdentityCard), v...))
	})
}

// IdentityCardNotIn applies the NotIn predicate on the "identity_card" field.
func IdentityCardNotIn(vs ...string) predicate.Checkout {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Checkout(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldIdentityCard), v...))
	})
}

// IdentityCardGT applies the GT predicate on the "identity_card" field.
func IdentityCardGT(v string) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldIdentityCard), v))
	})
}

// IdentityCardGTE applies the GTE predicate on the "identity_card" field.
func IdentityCardGTE(v string) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldIdentityCard), v))
	})
}

// IdentityCardLT applies the LT predicate on the "identity_card" field.
func IdentityCardLT(v string) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldIdentityCard), v))
	})
}

// IdentityCardLTE applies the LTE predicate on the "identity_card" field.
func IdentityCardLTE(v string) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldIdentityCard), v))
	})
}

// IdentityCardContains applies the Contains predicate on the "identity_card" field.
func IdentityCardContains(v string) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldIdentityCard), v))
	})
}

// IdentityCardHasPrefix applies the HasPrefix predicate on the "identity_card" field.
func IdentityCardHasPrefix(v string) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldIdentityCard), v))
	})
}

// IdentityCardHasSuffix applies the HasSuffix predicate on the "identity_card" field.
func IdentityCardHasSuffix(v string) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldIdentityCard), v))
	})
}

// IdentityCardEqualFold applies the EqualFold predicate on the "identity_card" field.
func IdentityCardEqualFold(v string) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldIdentityCard), v))
	})
}

// IdentityCardContainsFold applies the ContainsFold predicate on the "identity_card" field.
func IdentityCardContainsFold(v string) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldIdentityCard), v))
	})
}

// PriceEQ applies the EQ predicate on the "price" field.
func PriceEQ(v float64) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPrice), v))
	})
}

// PriceNEQ applies the NEQ predicate on the "price" field.
func PriceNEQ(v float64) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPrice), v))
	})
}

// PriceIn applies the In predicate on the "price" field.
func PriceIn(vs ...float64) predicate.Checkout {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Checkout(func(s *sql.Selector) {
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
func PriceNotIn(vs ...float64) predicate.Checkout {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Checkout(func(s *sql.Selector) {
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
func PriceGT(v float64) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPrice), v))
	})
}

// PriceGTE applies the GTE predicate on the "price" field.
func PriceGTE(v float64) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPrice), v))
	})
}

// PriceLT applies the LT predicate on the "price" field.
func PriceLT(v float64) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPrice), v))
	})
}

// PriceLTE applies the LTE predicate on the "price" field.
func PriceLTE(v float64) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPrice), v))
	})
}

// CommentEQ applies the EQ predicate on the "comment" field.
func CommentEQ(v string) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldComment), v))
	})
}

// CommentNEQ applies the NEQ predicate on the "comment" field.
func CommentNEQ(v string) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldComment), v))
	})
}

// CommentIn applies the In predicate on the "comment" field.
func CommentIn(vs ...string) predicate.Checkout {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Checkout(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldComment), v...))
	})
}

// CommentNotIn applies the NotIn predicate on the "comment" field.
func CommentNotIn(vs ...string) predicate.Checkout {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Checkout(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldComment), v...))
	})
}

// CommentGT applies the GT predicate on the "comment" field.
func CommentGT(v string) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldComment), v))
	})
}

// CommentGTE applies the GTE predicate on the "comment" field.
func CommentGTE(v string) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldComment), v))
	})
}

// CommentLT applies the LT predicate on the "comment" field.
func CommentLT(v string) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldComment), v))
	})
}

// CommentLTE applies the LTE predicate on the "comment" field.
func CommentLTE(v string) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldComment), v))
	})
}

// CommentContains applies the Contains predicate on the "comment" field.
func CommentContains(v string) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldComment), v))
	})
}

// CommentHasPrefix applies the HasPrefix predicate on the "comment" field.
func CommentHasPrefix(v string) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldComment), v))
	})
}

// CommentHasSuffix applies the HasSuffix predicate on the "comment" field.
func CommentHasSuffix(v string) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldComment), v))
	})
}

// CommentEqualFold applies the EqualFold predicate on the "comment" field.
func CommentEqualFold(v string) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldComment), v))
	})
}

// CommentContainsFold applies the ContainsFold predicate on the "comment" field.
func CommentContainsFold(v string) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldComment), v))
	})
}

// HasStatuss applies the HasEdge predicate on the "statuss" edge.
func HasStatuss() predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(StatussTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, StatussTable, StatussColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasStatussWith applies the HasEdge predicate on the "statuss" edge with a given conditions (other predicates).
func HasStatussWith(preds ...predicate.Status) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(StatussInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, StatussTable, StatussColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasStatusopinion applies the HasEdge predicate on the "statusopinion" edge.
func HasStatusopinion() predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(StatusopinionTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, StatusopinionTable, StatusopinionColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasStatusopinionWith applies the HasEdge predicate on the "statusopinion" edge with a given conditions (other predicates).
func HasStatusopinionWith(preds ...predicate.StatusOpinion) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(StatusopinionInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, StatusopinionTable, StatusopinionColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCounterstaffs applies the HasEdge predicate on the "counterstaffs" edge.
func HasCounterstaffs() predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CounterstaffsTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CounterstaffsTable, CounterstaffsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCounterstaffsWith applies the HasEdge predicate on the "counterstaffs" edge with a given conditions (other predicates).
func HasCounterstaffsWith(preds ...predicate.CounterStaff) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
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

// HasCheckins applies the HasEdge predicate on the "checkins" edge.
func HasCheckins() predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CheckinsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, CheckinsTable, CheckinsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCheckinsWith applies the HasEdge predicate on the "checkins" edge with a given conditions (other predicates).
func HasCheckinsWith(preds ...predicate.CheckIn) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CheckinsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, CheckinsTable, CheckinsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Checkout) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Checkout) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
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
func Not(p predicate.Checkout) predicate.Checkout {
	return predicate.Checkout(func(s *sql.Selector) {
		p(s.Not())
	})
}
