// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/renaldyhidayatt/inventorygoent/ent/predicate"
	"github.com/renaldyhidayatt/inventorygoent/ent/product"
	"github.com/renaldyhidayatt/inventorygoent/ent/productmasuk"
	"github.com/renaldyhidayatt/inventorygoent/ent/supplier"
)

// ProductMasukQuery is the builder for querying ProductMasuk entities.
type ProductMasukQuery struct {
	config
	ctx          *QueryContext
	order        []OrderFunc
	inters       []Interceptor
	predicates   []predicate.ProductMasuk
	withProduct  *ProductQuery
	withSupplier *SupplierQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ProductMasukQuery builder.
func (pmq *ProductMasukQuery) Where(ps ...predicate.ProductMasuk) *ProductMasukQuery {
	pmq.predicates = append(pmq.predicates, ps...)
	return pmq
}

// Limit the number of records to be returned by this query.
func (pmq *ProductMasukQuery) Limit(limit int) *ProductMasukQuery {
	pmq.ctx.Limit = &limit
	return pmq
}

// Offset to start from.
func (pmq *ProductMasukQuery) Offset(offset int) *ProductMasukQuery {
	pmq.ctx.Offset = &offset
	return pmq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pmq *ProductMasukQuery) Unique(unique bool) *ProductMasukQuery {
	pmq.ctx.Unique = &unique
	return pmq
}

// Order specifies how the records should be ordered.
func (pmq *ProductMasukQuery) Order(o ...OrderFunc) *ProductMasukQuery {
	pmq.order = append(pmq.order, o...)
	return pmq
}

// QueryProduct chains the current query on the "product" edge.
func (pmq *ProductMasukQuery) QueryProduct() *ProductQuery {
	query := (&ProductClient{config: pmq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pmq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(productmasuk.Table, productmasuk.FieldID, selector),
			sqlgraph.To(product.Table, product.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, productmasuk.ProductTable, productmasuk.ProductPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(pmq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QuerySupplier chains the current query on the "supplier" edge.
func (pmq *ProductMasukQuery) QuerySupplier() *SupplierQuery {
	query := (&SupplierClient{config: pmq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pmq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(productmasuk.Table, productmasuk.FieldID, selector),
			sqlgraph.To(supplier.Table, supplier.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, productmasuk.SupplierTable, productmasuk.SupplierPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(pmq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ProductMasuk entity from the query.
// Returns a *NotFoundError when no ProductMasuk was found.
func (pmq *ProductMasukQuery) First(ctx context.Context) (*ProductMasuk, error) {
	nodes, err := pmq.Limit(1).All(setContextOp(ctx, pmq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{productmasuk.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pmq *ProductMasukQuery) FirstX(ctx context.Context) *ProductMasuk {
	node, err := pmq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ProductMasuk ID from the query.
// Returns a *NotFoundError when no ProductMasuk ID was found.
func (pmq *ProductMasukQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pmq.Limit(1).IDs(setContextOp(ctx, pmq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{productmasuk.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pmq *ProductMasukQuery) FirstIDX(ctx context.Context) int {
	id, err := pmq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ProductMasuk entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ProductMasuk entity is found.
// Returns a *NotFoundError when no ProductMasuk entities are found.
func (pmq *ProductMasukQuery) Only(ctx context.Context) (*ProductMasuk, error) {
	nodes, err := pmq.Limit(2).All(setContextOp(ctx, pmq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{productmasuk.Label}
	default:
		return nil, &NotSingularError{productmasuk.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pmq *ProductMasukQuery) OnlyX(ctx context.Context) *ProductMasuk {
	node, err := pmq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ProductMasuk ID in the query.
// Returns a *NotSingularError when more than one ProductMasuk ID is found.
// Returns a *NotFoundError when no entities are found.
func (pmq *ProductMasukQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pmq.Limit(2).IDs(setContextOp(ctx, pmq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{productmasuk.Label}
	default:
		err = &NotSingularError{productmasuk.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pmq *ProductMasukQuery) OnlyIDX(ctx context.Context) int {
	id, err := pmq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ProductMasuks.
func (pmq *ProductMasukQuery) All(ctx context.Context) ([]*ProductMasuk, error) {
	ctx = setContextOp(ctx, pmq.ctx, "All")
	if err := pmq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ProductMasuk, *ProductMasukQuery]()
	return withInterceptors[[]*ProductMasuk](ctx, pmq, qr, pmq.inters)
}

// AllX is like All, but panics if an error occurs.
func (pmq *ProductMasukQuery) AllX(ctx context.Context) []*ProductMasuk {
	nodes, err := pmq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ProductMasuk IDs.
func (pmq *ProductMasukQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	ctx = setContextOp(ctx, pmq.ctx, "IDs")
	if err := pmq.Select(productmasuk.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pmq *ProductMasukQuery) IDsX(ctx context.Context) []int {
	ids, err := pmq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pmq *ProductMasukQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, pmq.ctx, "Count")
	if err := pmq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, pmq, querierCount[*ProductMasukQuery](), pmq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (pmq *ProductMasukQuery) CountX(ctx context.Context) int {
	count, err := pmq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pmq *ProductMasukQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, pmq.ctx, "Exist")
	switch _, err := pmq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (pmq *ProductMasukQuery) ExistX(ctx context.Context) bool {
	exist, err := pmq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ProductMasukQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pmq *ProductMasukQuery) Clone() *ProductMasukQuery {
	if pmq == nil {
		return nil
	}
	return &ProductMasukQuery{
		config:       pmq.config,
		ctx:          pmq.ctx.Clone(),
		order:        append([]OrderFunc{}, pmq.order...),
		inters:       append([]Interceptor{}, pmq.inters...),
		predicates:   append([]predicate.ProductMasuk{}, pmq.predicates...),
		withProduct:  pmq.withProduct.Clone(),
		withSupplier: pmq.withSupplier.Clone(),
		// clone intermediate query.
		sql:  pmq.sql.Clone(),
		path: pmq.path,
	}
}

// WithProduct tells the query-builder to eager-load the nodes that are connected to
// the "product" edge. The optional arguments are used to configure the query builder of the edge.
func (pmq *ProductMasukQuery) WithProduct(opts ...func(*ProductQuery)) *ProductMasukQuery {
	query := (&ProductClient{config: pmq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pmq.withProduct = query
	return pmq
}

// WithSupplier tells the query-builder to eager-load the nodes that are connected to
// the "supplier" edge. The optional arguments are used to configure the query builder of the edge.
func (pmq *ProductMasukQuery) WithSupplier(opts ...func(*SupplierQuery)) *ProductMasukQuery {
	query := (&SupplierClient{config: pmq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pmq.withSupplier = query
	return pmq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ProductMasuk.Query().
//		GroupBy(productmasuk.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (pmq *ProductMasukQuery) GroupBy(field string, fields ...string) *ProductMasukGroupBy {
	pmq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ProductMasukGroupBy{build: pmq}
	grbuild.flds = &pmq.ctx.Fields
	grbuild.label = productmasuk.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.ProductMasuk.Query().
//		Select(productmasuk.FieldName).
//		Scan(ctx, &v)
func (pmq *ProductMasukQuery) Select(fields ...string) *ProductMasukSelect {
	pmq.ctx.Fields = append(pmq.ctx.Fields, fields...)
	sbuild := &ProductMasukSelect{ProductMasukQuery: pmq}
	sbuild.label = productmasuk.Label
	sbuild.flds, sbuild.scan = &pmq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ProductMasukSelect configured with the given aggregations.
func (pmq *ProductMasukQuery) Aggregate(fns ...AggregateFunc) *ProductMasukSelect {
	return pmq.Select().Aggregate(fns...)
}

func (pmq *ProductMasukQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range pmq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, pmq); err != nil {
				return err
			}
		}
	}
	for _, f := range pmq.ctx.Fields {
		if !productmasuk.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if pmq.path != nil {
		prev, err := pmq.path(ctx)
		if err != nil {
			return err
		}
		pmq.sql = prev
	}
	return nil
}

func (pmq *ProductMasukQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ProductMasuk, error) {
	var (
		nodes       = []*ProductMasuk{}
		_spec       = pmq.querySpec()
		loadedTypes = [2]bool{
			pmq.withProduct != nil,
			pmq.withSupplier != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ProductMasuk).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ProductMasuk{config: pmq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, pmq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := pmq.withProduct; query != nil {
		if err := pmq.loadProduct(ctx, query, nodes,
			func(n *ProductMasuk) { n.Edges.Product = []*Product{} },
			func(n *ProductMasuk, e *Product) { n.Edges.Product = append(n.Edges.Product, e) }); err != nil {
			return nil, err
		}
	}
	if query := pmq.withSupplier; query != nil {
		if err := pmq.loadSupplier(ctx, query, nodes,
			func(n *ProductMasuk) { n.Edges.Supplier = []*Supplier{} },
			func(n *ProductMasuk, e *Supplier) { n.Edges.Supplier = append(n.Edges.Supplier, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (pmq *ProductMasukQuery) loadProduct(ctx context.Context, query *ProductQuery, nodes []*ProductMasuk, init func(*ProductMasuk), assign func(*ProductMasuk, *Product)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*ProductMasuk)
	nids := make(map[int]map[*ProductMasuk]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(productmasuk.ProductTable)
		s.Join(joinT).On(s.C(product.FieldID), joinT.C(productmasuk.ProductPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(productmasuk.ProductPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(productmasuk.ProductPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := int(values[0].(*sql.NullInt64).Int64)
				inValue := int(values[1].(*sql.NullInt64).Int64)
				if nids[inValue] == nil {
					nids[inValue] = map[*ProductMasuk]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Product](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "product" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (pmq *ProductMasukQuery) loadSupplier(ctx context.Context, query *SupplierQuery, nodes []*ProductMasuk, init func(*ProductMasuk), assign func(*ProductMasuk, *Supplier)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*ProductMasuk)
	nids := make(map[int]map[*ProductMasuk]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(productmasuk.SupplierTable)
		s.Join(joinT).On(s.C(supplier.FieldID), joinT.C(productmasuk.SupplierPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(productmasuk.SupplierPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(productmasuk.SupplierPrimaryKey[0]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := int(values[0].(*sql.NullInt64).Int64)
				inValue := int(values[1].(*sql.NullInt64).Int64)
				if nids[inValue] == nil {
					nids[inValue] = map[*ProductMasuk]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Supplier](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "supplier" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (pmq *ProductMasukQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pmq.querySpec()
	_spec.Node.Columns = pmq.ctx.Fields
	if len(pmq.ctx.Fields) > 0 {
		_spec.Unique = pmq.ctx.Unique != nil && *pmq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, pmq.driver, _spec)
}

func (pmq *ProductMasukQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   productmasuk.Table,
			Columns: productmasuk.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: productmasuk.FieldID,
			},
		},
		From:   pmq.sql,
		Unique: true,
	}
	if unique := pmq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := pmq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, productmasuk.FieldID)
		for i := range fields {
			if fields[i] != productmasuk.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := pmq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pmq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pmq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pmq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pmq *ProductMasukQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pmq.driver.Dialect())
	t1 := builder.Table(productmasuk.Table)
	columns := pmq.ctx.Fields
	if len(columns) == 0 {
		columns = productmasuk.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pmq.sql != nil {
		selector = pmq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if pmq.ctx.Unique != nil && *pmq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range pmq.predicates {
		p(selector)
	}
	for _, p := range pmq.order {
		p(selector)
	}
	if offset := pmq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pmq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ProductMasukGroupBy is the group-by builder for ProductMasuk entities.
type ProductMasukGroupBy struct {
	selector
	build *ProductMasukQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pmgb *ProductMasukGroupBy) Aggregate(fns ...AggregateFunc) *ProductMasukGroupBy {
	pmgb.fns = append(pmgb.fns, fns...)
	return pmgb
}

// Scan applies the selector query and scans the result into the given value.
func (pmgb *ProductMasukGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pmgb.build.ctx, "GroupBy")
	if err := pmgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ProductMasukQuery, *ProductMasukGroupBy](ctx, pmgb.build, pmgb, pmgb.build.inters, v)
}

func (pmgb *ProductMasukGroupBy) sqlScan(ctx context.Context, root *ProductMasukQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(pmgb.fns))
	for _, fn := range pmgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*pmgb.flds)+len(pmgb.fns))
		for _, f := range *pmgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*pmgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pmgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ProductMasukSelect is the builder for selecting fields of ProductMasuk entities.
type ProductMasukSelect struct {
	*ProductMasukQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (pms *ProductMasukSelect) Aggregate(fns ...AggregateFunc) *ProductMasukSelect {
	pms.fns = append(pms.fns, fns...)
	return pms
}

// Scan applies the selector query and scans the result into the given value.
func (pms *ProductMasukSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pms.ctx, "Select")
	if err := pms.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ProductMasukQuery, *ProductMasukSelect](ctx, pms.ProductMasukQuery, pms, pms.inters, v)
}

func (pms *ProductMasukSelect) sqlScan(ctx context.Context, root *ProductMasukQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(pms.fns))
	for _, fn := range pms.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*pms.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
