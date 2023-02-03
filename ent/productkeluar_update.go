// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/renaldyhidayatt/inventorygoent/ent/category"
	"github.com/renaldyhidayatt/inventorygoent/ent/predicate"
	"github.com/renaldyhidayatt/inventorygoent/ent/product"
	"github.com/renaldyhidayatt/inventorygoent/ent/productkeluar"
)

// ProductKeluarUpdate is the builder for updating ProductKeluar entities.
type ProductKeluarUpdate struct {
	config
	hooks    []Hook
	mutation *ProductKeluarMutation
}

// Where appends a list predicates to the ProductKeluarUpdate builder.
func (pku *ProductKeluarUpdate) Where(ps ...predicate.ProductKeluar) *ProductKeluarUpdate {
	pku.mutation.Where(ps...)
	return pku
}

// SetQty sets the "qty" field.
func (pku *ProductKeluarUpdate) SetQty(s string) *ProductKeluarUpdate {
	pku.mutation.SetQty(s)
	return pku
}

// SetCreatedAt sets the "created_at" field.
func (pku *ProductKeluarUpdate) SetCreatedAt(t time.Time) *ProductKeluarUpdate {
	pku.mutation.SetCreatedAt(t)
	return pku
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pku *ProductKeluarUpdate) SetNillableCreatedAt(t *time.Time) *ProductKeluarUpdate {
	if t != nil {
		pku.SetCreatedAt(*t)
	}
	return pku
}

// SetUpdatedAt sets the "updated_at" field.
func (pku *ProductKeluarUpdate) SetUpdatedAt(t time.Time) *ProductKeluarUpdate {
	pku.mutation.SetUpdatedAt(t)
	return pku
}

// AddProductIDs adds the "products" edge to the Product entity by IDs.
func (pku *ProductKeluarUpdate) AddProductIDs(ids ...int) *ProductKeluarUpdate {
	pku.mutation.AddProductIDs(ids...)
	return pku
}

// AddProducts adds the "products" edges to the Product entity.
func (pku *ProductKeluarUpdate) AddProducts(p ...*Product) *ProductKeluarUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pku.AddProductIDs(ids...)
}

// AddCategoryIDs adds the "category" edge to the Category entity by IDs.
func (pku *ProductKeluarUpdate) AddCategoryIDs(ids ...int) *ProductKeluarUpdate {
	pku.mutation.AddCategoryIDs(ids...)
	return pku
}

// AddCategory adds the "category" edges to the Category entity.
func (pku *ProductKeluarUpdate) AddCategory(c ...*Category) *ProductKeluarUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return pku.AddCategoryIDs(ids...)
}

// Mutation returns the ProductKeluarMutation object of the builder.
func (pku *ProductKeluarUpdate) Mutation() *ProductKeluarMutation {
	return pku.mutation
}

// ClearProducts clears all "products" edges to the Product entity.
func (pku *ProductKeluarUpdate) ClearProducts() *ProductKeluarUpdate {
	pku.mutation.ClearProducts()
	return pku
}

// RemoveProductIDs removes the "products" edge to Product entities by IDs.
func (pku *ProductKeluarUpdate) RemoveProductIDs(ids ...int) *ProductKeluarUpdate {
	pku.mutation.RemoveProductIDs(ids...)
	return pku
}

// RemoveProducts removes "products" edges to Product entities.
func (pku *ProductKeluarUpdate) RemoveProducts(p ...*Product) *ProductKeluarUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pku.RemoveProductIDs(ids...)
}

// ClearCategory clears all "category" edges to the Category entity.
func (pku *ProductKeluarUpdate) ClearCategory() *ProductKeluarUpdate {
	pku.mutation.ClearCategory()
	return pku
}

// RemoveCategoryIDs removes the "category" edge to Category entities by IDs.
func (pku *ProductKeluarUpdate) RemoveCategoryIDs(ids ...int) *ProductKeluarUpdate {
	pku.mutation.RemoveCategoryIDs(ids...)
	return pku
}

// RemoveCategory removes "category" edges to Category entities.
func (pku *ProductKeluarUpdate) RemoveCategory(c ...*Category) *ProductKeluarUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return pku.RemoveCategoryIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pku *ProductKeluarUpdate) Save(ctx context.Context) (int, error) {
	pku.defaults()
	return withHooks[int, ProductKeluarMutation](ctx, pku.sqlSave, pku.mutation, pku.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pku *ProductKeluarUpdate) SaveX(ctx context.Context) int {
	affected, err := pku.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pku *ProductKeluarUpdate) Exec(ctx context.Context) error {
	_, err := pku.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pku *ProductKeluarUpdate) ExecX(ctx context.Context) {
	if err := pku.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pku *ProductKeluarUpdate) defaults() {
	if _, ok := pku.mutation.UpdatedAt(); !ok {
		v := productkeluar.UpdateDefaultUpdatedAt()
		pku.mutation.SetUpdatedAt(v)
	}
}

func (pku *ProductKeluarUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   productkeluar.Table,
			Columns: productkeluar.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: productkeluar.FieldID,
			},
		},
	}
	if ps := pku.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pku.mutation.Qty(); ok {
		_spec.SetField(productkeluar.FieldQty, field.TypeString, value)
	}
	if value, ok := pku.mutation.CreatedAt(); ok {
		_spec.SetField(productkeluar.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := pku.mutation.UpdatedAt(); ok {
		_spec.SetField(productkeluar.FieldUpdatedAt, field.TypeTime, value)
	}
	if pku.mutation.ProductsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   productkeluar.ProductsTable,
			Columns: productkeluar.ProductsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: product.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pku.mutation.RemovedProductsIDs(); len(nodes) > 0 && !pku.mutation.ProductsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   productkeluar.ProductsTable,
			Columns: productkeluar.ProductsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: product.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pku.mutation.ProductsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   productkeluar.ProductsTable,
			Columns: productkeluar.ProductsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: product.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pku.mutation.CategoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   productkeluar.CategoryTable,
			Columns: productkeluar.CategoryPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: category.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pku.mutation.RemovedCategoryIDs(); len(nodes) > 0 && !pku.mutation.CategoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   productkeluar.CategoryTable,
			Columns: productkeluar.CategoryPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: category.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pku.mutation.CategoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   productkeluar.CategoryTable,
			Columns: productkeluar.CategoryPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: category.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pku.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{productkeluar.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pku.mutation.done = true
	return n, nil
}

// ProductKeluarUpdateOne is the builder for updating a single ProductKeluar entity.
type ProductKeluarUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ProductKeluarMutation
}

// SetQty sets the "qty" field.
func (pkuo *ProductKeluarUpdateOne) SetQty(s string) *ProductKeluarUpdateOne {
	pkuo.mutation.SetQty(s)
	return pkuo
}

// SetCreatedAt sets the "created_at" field.
func (pkuo *ProductKeluarUpdateOne) SetCreatedAt(t time.Time) *ProductKeluarUpdateOne {
	pkuo.mutation.SetCreatedAt(t)
	return pkuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pkuo *ProductKeluarUpdateOne) SetNillableCreatedAt(t *time.Time) *ProductKeluarUpdateOne {
	if t != nil {
		pkuo.SetCreatedAt(*t)
	}
	return pkuo
}

// SetUpdatedAt sets the "updated_at" field.
func (pkuo *ProductKeluarUpdateOne) SetUpdatedAt(t time.Time) *ProductKeluarUpdateOne {
	pkuo.mutation.SetUpdatedAt(t)
	return pkuo
}

// AddProductIDs adds the "products" edge to the Product entity by IDs.
func (pkuo *ProductKeluarUpdateOne) AddProductIDs(ids ...int) *ProductKeluarUpdateOne {
	pkuo.mutation.AddProductIDs(ids...)
	return pkuo
}

// AddProducts adds the "products" edges to the Product entity.
func (pkuo *ProductKeluarUpdateOne) AddProducts(p ...*Product) *ProductKeluarUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pkuo.AddProductIDs(ids...)
}

// AddCategoryIDs adds the "category" edge to the Category entity by IDs.
func (pkuo *ProductKeluarUpdateOne) AddCategoryIDs(ids ...int) *ProductKeluarUpdateOne {
	pkuo.mutation.AddCategoryIDs(ids...)
	return pkuo
}

// AddCategory adds the "category" edges to the Category entity.
func (pkuo *ProductKeluarUpdateOne) AddCategory(c ...*Category) *ProductKeluarUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return pkuo.AddCategoryIDs(ids...)
}

// Mutation returns the ProductKeluarMutation object of the builder.
func (pkuo *ProductKeluarUpdateOne) Mutation() *ProductKeluarMutation {
	return pkuo.mutation
}

// ClearProducts clears all "products" edges to the Product entity.
func (pkuo *ProductKeluarUpdateOne) ClearProducts() *ProductKeluarUpdateOne {
	pkuo.mutation.ClearProducts()
	return pkuo
}

// RemoveProductIDs removes the "products" edge to Product entities by IDs.
func (pkuo *ProductKeluarUpdateOne) RemoveProductIDs(ids ...int) *ProductKeluarUpdateOne {
	pkuo.mutation.RemoveProductIDs(ids...)
	return pkuo
}

// RemoveProducts removes "products" edges to Product entities.
func (pkuo *ProductKeluarUpdateOne) RemoveProducts(p ...*Product) *ProductKeluarUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pkuo.RemoveProductIDs(ids...)
}

// ClearCategory clears all "category" edges to the Category entity.
func (pkuo *ProductKeluarUpdateOne) ClearCategory() *ProductKeluarUpdateOne {
	pkuo.mutation.ClearCategory()
	return pkuo
}

// RemoveCategoryIDs removes the "category" edge to Category entities by IDs.
func (pkuo *ProductKeluarUpdateOne) RemoveCategoryIDs(ids ...int) *ProductKeluarUpdateOne {
	pkuo.mutation.RemoveCategoryIDs(ids...)
	return pkuo
}

// RemoveCategory removes "category" edges to Category entities.
func (pkuo *ProductKeluarUpdateOne) RemoveCategory(c ...*Category) *ProductKeluarUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return pkuo.RemoveCategoryIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (pkuo *ProductKeluarUpdateOne) Select(field string, fields ...string) *ProductKeluarUpdateOne {
	pkuo.fields = append([]string{field}, fields...)
	return pkuo
}

// Save executes the query and returns the updated ProductKeluar entity.
func (pkuo *ProductKeluarUpdateOne) Save(ctx context.Context) (*ProductKeluar, error) {
	pkuo.defaults()
	return withHooks[*ProductKeluar, ProductKeluarMutation](ctx, pkuo.sqlSave, pkuo.mutation, pkuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pkuo *ProductKeluarUpdateOne) SaveX(ctx context.Context) *ProductKeluar {
	node, err := pkuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (pkuo *ProductKeluarUpdateOne) Exec(ctx context.Context) error {
	_, err := pkuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pkuo *ProductKeluarUpdateOne) ExecX(ctx context.Context) {
	if err := pkuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pkuo *ProductKeluarUpdateOne) defaults() {
	if _, ok := pkuo.mutation.UpdatedAt(); !ok {
		v := productkeluar.UpdateDefaultUpdatedAt()
		pkuo.mutation.SetUpdatedAt(v)
	}
}

func (pkuo *ProductKeluarUpdateOne) sqlSave(ctx context.Context) (_node *ProductKeluar, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   productkeluar.Table,
			Columns: productkeluar.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: productkeluar.FieldID,
			},
		},
	}
	id, ok := pkuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ProductKeluar.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := pkuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, productkeluar.FieldID)
		for _, f := range fields {
			if !productkeluar.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != productkeluar.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := pkuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pkuo.mutation.Qty(); ok {
		_spec.SetField(productkeluar.FieldQty, field.TypeString, value)
	}
	if value, ok := pkuo.mutation.CreatedAt(); ok {
		_spec.SetField(productkeluar.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := pkuo.mutation.UpdatedAt(); ok {
		_spec.SetField(productkeluar.FieldUpdatedAt, field.TypeTime, value)
	}
	if pkuo.mutation.ProductsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   productkeluar.ProductsTable,
			Columns: productkeluar.ProductsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: product.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pkuo.mutation.RemovedProductsIDs(); len(nodes) > 0 && !pkuo.mutation.ProductsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   productkeluar.ProductsTable,
			Columns: productkeluar.ProductsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: product.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pkuo.mutation.ProductsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   productkeluar.ProductsTable,
			Columns: productkeluar.ProductsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: product.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pkuo.mutation.CategoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   productkeluar.CategoryTable,
			Columns: productkeluar.CategoryPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: category.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pkuo.mutation.RemovedCategoryIDs(); len(nodes) > 0 && !pkuo.mutation.CategoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   productkeluar.CategoryTable,
			Columns: productkeluar.CategoryPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: category.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pkuo.mutation.CategoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   productkeluar.CategoryTable,
			Columns: productkeluar.CategoryPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: category.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &ProductKeluar{config: pkuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, pkuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{productkeluar.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	pkuo.mutation.done = true
	return _node, nil
}
