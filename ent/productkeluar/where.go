// Code generated by ent, DO NOT EDIT.

package productkeluar

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/renaldyhidayatt/inventorygoent/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.ProductKeluar {
	return predicate.ProductKeluar(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.ProductKeluar {
	return predicate.ProductKeluar(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.ProductKeluar {
	return predicate.ProductKeluar(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.ProductKeluar {
	return predicate.ProductKeluar(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.ProductKeluar {
	return predicate.ProductKeluar(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.ProductKeluar {
	return predicate.ProductKeluar(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.ProductKeluar {
	return predicate.ProductKeluar(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.ProductKeluar {
	return predicate.ProductKeluar(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.ProductKeluar {
	return predicate.ProductKeluar(sql.FieldLTE(FieldID, id))
}

// Qty applies equality check predicate on the "qty" field. It's identical to QtyEQ.
func Qty(v string) predicate.ProductKeluar {
	return predicate.ProductKeluar(sql.FieldEQ(FieldQty, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.ProductKeluar {
	return predicate.ProductKeluar(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.ProductKeluar {
	return predicate.ProductKeluar(sql.FieldEQ(FieldUpdatedAt, v))
}

// QtyEQ applies the EQ predicate on the "qty" field.
func QtyEQ(v string) predicate.ProductKeluar {
	return predicate.ProductKeluar(sql.FieldEQ(FieldQty, v))
}

// QtyNEQ applies the NEQ predicate on the "qty" field.
func QtyNEQ(v string) predicate.ProductKeluar {
	return predicate.ProductKeluar(sql.FieldNEQ(FieldQty, v))
}

// QtyIn applies the In predicate on the "qty" field.
func QtyIn(vs ...string) predicate.ProductKeluar {
	return predicate.ProductKeluar(sql.FieldIn(FieldQty, vs...))
}

// QtyNotIn applies the NotIn predicate on the "qty" field.
func QtyNotIn(vs ...string) predicate.ProductKeluar {
	return predicate.ProductKeluar(sql.FieldNotIn(FieldQty, vs...))
}

// QtyGT applies the GT predicate on the "qty" field.
func QtyGT(v string) predicate.ProductKeluar {
	return predicate.ProductKeluar(sql.FieldGT(FieldQty, v))
}

// QtyGTE applies the GTE predicate on the "qty" field.
func QtyGTE(v string) predicate.ProductKeluar {
	return predicate.ProductKeluar(sql.FieldGTE(FieldQty, v))
}

// QtyLT applies the LT predicate on the "qty" field.
func QtyLT(v string) predicate.ProductKeluar {
	return predicate.ProductKeluar(sql.FieldLT(FieldQty, v))
}

// QtyLTE applies the LTE predicate on the "qty" field.
func QtyLTE(v string) predicate.ProductKeluar {
	return predicate.ProductKeluar(sql.FieldLTE(FieldQty, v))
}

// QtyContains applies the Contains predicate on the "qty" field.
func QtyContains(v string) predicate.ProductKeluar {
	return predicate.ProductKeluar(sql.FieldContains(FieldQty, v))
}

// QtyHasPrefix applies the HasPrefix predicate on the "qty" field.
func QtyHasPrefix(v string) predicate.ProductKeluar {
	return predicate.ProductKeluar(sql.FieldHasPrefix(FieldQty, v))
}

// QtyHasSuffix applies the HasSuffix predicate on the "qty" field.
func QtyHasSuffix(v string) predicate.ProductKeluar {
	return predicate.ProductKeluar(sql.FieldHasSuffix(FieldQty, v))
}

// QtyEqualFold applies the EqualFold predicate on the "qty" field.
func QtyEqualFold(v string) predicate.ProductKeluar {
	return predicate.ProductKeluar(sql.FieldEqualFold(FieldQty, v))
}

// QtyContainsFold applies the ContainsFold predicate on the "qty" field.
func QtyContainsFold(v string) predicate.ProductKeluar {
	return predicate.ProductKeluar(sql.FieldContainsFold(FieldQty, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.ProductKeluar {
	return predicate.ProductKeluar(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.ProductKeluar {
	return predicate.ProductKeluar(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.ProductKeluar {
	return predicate.ProductKeluar(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.ProductKeluar {
	return predicate.ProductKeluar(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.ProductKeluar {
	return predicate.ProductKeluar(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.ProductKeluar {
	return predicate.ProductKeluar(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.ProductKeluar {
	return predicate.ProductKeluar(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.ProductKeluar {
	return predicate.ProductKeluar(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.ProductKeluar {
	return predicate.ProductKeluar(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.ProductKeluar {
	return predicate.ProductKeluar(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.ProductKeluar {
	return predicate.ProductKeluar(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.ProductKeluar {
	return predicate.ProductKeluar(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.ProductKeluar {
	return predicate.ProductKeluar(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.ProductKeluar {
	return predicate.ProductKeluar(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.ProductKeluar {
	return predicate.ProductKeluar(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.ProductKeluar {
	return predicate.ProductKeluar(sql.FieldLTE(FieldUpdatedAt, v))
}

// HasProducts applies the HasEdge predicate on the "products" edge.
func HasProducts() predicate.ProductKeluar {
	return predicate.ProductKeluar(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, ProductsTable, ProductsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProductsWith applies the HasEdge predicate on the "products" edge with a given conditions (other predicates).
func HasProductsWith(preds ...predicate.Product) predicate.ProductKeluar {
	return predicate.ProductKeluar(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProductsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, ProductsTable, ProductsPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCategory applies the HasEdge predicate on the "category" edge.
func HasCategory() predicate.ProductKeluar {
	return predicate.ProductKeluar(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, CategoryTable, CategoryPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCategoryWith applies the HasEdge predicate on the "category" edge with a given conditions (other predicates).
func HasCategoryWith(preds ...predicate.Category) predicate.ProductKeluar {
	return predicate.ProductKeluar(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CategoryInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, CategoryTable, CategoryPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ProductKeluar) predicate.ProductKeluar {
	return predicate.ProductKeluar(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ProductKeluar) predicate.ProductKeluar {
	return predicate.ProductKeluar(func(s *sql.Selector) {
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
func Not(p predicate.ProductKeluar) predicate.ProductKeluar {
	return predicate.ProductKeluar(func(s *sql.Selector) {
		p(s.Not())
	})
}
