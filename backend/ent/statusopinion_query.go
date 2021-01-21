// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team18/app/ent/checkout"
	"github.com/team18/app/ent/predicate"
	"github.com/team18/app/ent/statusopinion"
)

// StatusOpinionQuery is the builder for querying StatusOpinion entities.
type StatusOpinionQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.StatusOpinion
	// eager-loading edges.
	withCheckouts *CheckoutQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (soq *StatusOpinionQuery) Where(ps ...predicate.StatusOpinion) *StatusOpinionQuery {
	soq.predicates = append(soq.predicates, ps...)
	return soq
}

// Limit adds a limit step to the query.
func (soq *StatusOpinionQuery) Limit(limit int) *StatusOpinionQuery {
	soq.limit = &limit
	return soq
}

// Offset adds an offset step to the query.
func (soq *StatusOpinionQuery) Offset(offset int) *StatusOpinionQuery {
	soq.offset = &offset
	return soq
}

// Order adds an order step to the query.
func (soq *StatusOpinionQuery) Order(o ...OrderFunc) *StatusOpinionQuery {
	soq.order = append(soq.order, o...)
	return soq
}

// QueryCheckouts chains the current query on the checkouts edge.
func (soq *StatusOpinionQuery) QueryCheckouts() *CheckoutQuery {
	query := &CheckoutQuery{config: soq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := soq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(statusopinion.Table, statusopinion.FieldID, soq.sqlQuery()),
			sqlgraph.To(checkout.Table, checkout.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, statusopinion.CheckoutsTable, statusopinion.CheckoutsColumn),
		)
		fromU = sqlgraph.SetNeighbors(soq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first StatusOpinion entity in the query. Returns *NotFoundError when no statusopinion was found.
func (soq *StatusOpinionQuery) First(ctx context.Context) (*StatusOpinion, error) {
	sos, err := soq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(sos) == 0 {
		return nil, &NotFoundError{statusopinion.Label}
	}
	return sos[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (soq *StatusOpinionQuery) FirstX(ctx context.Context) *StatusOpinion {
	so, err := soq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return so
}

// FirstID returns the first StatusOpinion id in the query. Returns *NotFoundError when no id was found.
func (soq *StatusOpinionQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = soq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{statusopinion.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (soq *StatusOpinionQuery) FirstXID(ctx context.Context) int {
	id, err := soq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only StatusOpinion entity in the query, returns an error if not exactly one entity was returned.
func (soq *StatusOpinionQuery) Only(ctx context.Context) (*StatusOpinion, error) {
	sos, err := soq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(sos) {
	case 1:
		return sos[0], nil
	case 0:
		return nil, &NotFoundError{statusopinion.Label}
	default:
		return nil, &NotSingularError{statusopinion.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (soq *StatusOpinionQuery) OnlyX(ctx context.Context) *StatusOpinion {
	so, err := soq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return so
}

// OnlyID returns the only StatusOpinion id in the query, returns an error if not exactly one id was returned.
func (soq *StatusOpinionQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = soq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{statusopinion.Label}
	default:
		err = &NotSingularError{statusopinion.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (soq *StatusOpinionQuery) OnlyIDX(ctx context.Context) int {
	id, err := soq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of StatusOpinions.
func (soq *StatusOpinionQuery) All(ctx context.Context) ([]*StatusOpinion, error) {
	if err := soq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return soq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (soq *StatusOpinionQuery) AllX(ctx context.Context) []*StatusOpinion {
	sos, err := soq.All(ctx)
	if err != nil {
		panic(err)
	}
	return sos
}

// IDs executes the query and returns a list of StatusOpinion ids.
func (soq *StatusOpinionQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := soq.Select(statusopinion.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (soq *StatusOpinionQuery) IDsX(ctx context.Context) []int {
	ids, err := soq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (soq *StatusOpinionQuery) Count(ctx context.Context) (int, error) {
	if err := soq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return soq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (soq *StatusOpinionQuery) CountX(ctx context.Context) int {
	count, err := soq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (soq *StatusOpinionQuery) Exist(ctx context.Context) (bool, error) {
	if err := soq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return soq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (soq *StatusOpinionQuery) ExistX(ctx context.Context) bool {
	exist, err := soq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (soq *StatusOpinionQuery) Clone() *StatusOpinionQuery {
	return &StatusOpinionQuery{
		config:     soq.config,
		limit:      soq.limit,
		offset:     soq.offset,
		order:      append([]OrderFunc{}, soq.order...),
		unique:     append([]string{}, soq.unique...),
		predicates: append([]predicate.StatusOpinion{}, soq.predicates...),
		// clone intermediate query.
		sql:  soq.sql.Clone(),
		path: soq.path,
	}
}

//  WithCheckouts tells the query-builder to eager-loads the nodes that are connected to
// the "checkouts" edge. The optional arguments used to configure the query builder of the edge.
func (soq *StatusOpinionQuery) WithCheckouts(opts ...func(*CheckoutQuery)) *StatusOpinionQuery {
	query := &CheckoutQuery{config: soq.config}
	for _, opt := range opts {
		opt(query)
	}
	soq.withCheckouts = query
	return soq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Opinion string `json:"opinion,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.StatusOpinion.Query().
//		GroupBy(statusopinion.FieldOpinion).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (soq *StatusOpinionQuery) GroupBy(field string, fields ...string) *StatusOpinionGroupBy {
	group := &StatusOpinionGroupBy{config: soq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := soq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return soq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		Opinion string `json:"opinion,omitempty"`
//	}
//
//	client.StatusOpinion.Query().
//		Select(statusopinion.FieldOpinion).
//		Scan(ctx, &v)
//
func (soq *StatusOpinionQuery) Select(field string, fields ...string) *StatusOpinionSelect {
	selector := &StatusOpinionSelect{config: soq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := soq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return soq.sqlQuery(), nil
	}
	return selector
}

func (soq *StatusOpinionQuery) prepareQuery(ctx context.Context) error {
	if soq.path != nil {
		prev, err := soq.path(ctx)
		if err != nil {
			return err
		}
		soq.sql = prev
	}
	return nil
}

func (soq *StatusOpinionQuery) sqlAll(ctx context.Context) ([]*StatusOpinion, error) {
	var (
		nodes       = []*StatusOpinion{}
		_spec       = soq.querySpec()
		loadedTypes = [1]bool{
			soq.withCheckouts != nil,
		}
	)
	_spec.ScanValues = func() []interface{} {
		node := &StatusOpinion{config: soq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		return values
	}
	_spec.Assign = func(values ...interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(values...)
	}
	if err := sqlgraph.QueryNodes(ctx, soq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := soq.withCheckouts; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*StatusOpinion)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.Checkout(func(s *sql.Selector) {
			s.Where(sql.InValues(statusopinion.CheckoutsColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.status_opinion_checkouts
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "status_opinion_checkouts" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "status_opinion_checkouts" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Checkouts = append(node.Edges.Checkouts, n)
		}
	}

	return nodes, nil
}

func (soq *StatusOpinionQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := soq.querySpec()
	return sqlgraph.CountNodes(ctx, soq.driver, _spec)
}

func (soq *StatusOpinionQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := soq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (soq *StatusOpinionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   statusopinion.Table,
			Columns: statusopinion.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: statusopinion.FieldID,
			},
		},
		From:   soq.sql,
		Unique: true,
	}
	if ps := soq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := soq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := soq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := soq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (soq *StatusOpinionQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(soq.driver.Dialect())
	t1 := builder.Table(statusopinion.Table)
	selector := builder.Select(t1.Columns(statusopinion.Columns...)...).From(t1)
	if soq.sql != nil {
		selector = soq.sql
		selector.Select(selector.Columns(statusopinion.Columns...)...)
	}
	for _, p := range soq.predicates {
		p(selector)
	}
	for _, p := range soq.order {
		p(selector)
	}
	if offset := soq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := soq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// StatusOpinionGroupBy is the builder for group-by StatusOpinion entities.
type StatusOpinionGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sogb *StatusOpinionGroupBy) Aggregate(fns ...AggregateFunc) *StatusOpinionGroupBy {
	sogb.fns = append(sogb.fns, fns...)
	return sogb
}

// Scan applies the group-by query and scan the result into the given value.
func (sogb *StatusOpinionGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := sogb.path(ctx)
	if err != nil {
		return err
	}
	sogb.sql = query
	return sogb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (sogb *StatusOpinionGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := sogb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (sogb *StatusOpinionGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(sogb.fields) > 1 {
		return nil, errors.New("ent: StatusOpinionGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := sogb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (sogb *StatusOpinionGroupBy) StringsX(ctx context.Context) []string {
	v, err := sogb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (sogb *StatusOpinionGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = sogb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{statusopinion.Label}
	default:
		err = fmt.Errorf("ent: StatusOpinionGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (sogb *StatusOpinionGroupBy) StringX(ctx context.Context) string {
	v, err := sogb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (sogb *StatusOpinionGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(sogb.fields) > 1 {
		return nil, errors.New("ent: StatusOpinionGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := sogb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (sogb *StatusOpinionGroupBy) IntsX(ctx context.Context) []int {
	v, err := sogb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (sogb *StatusOpinionGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = sogb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{statusopinion.Label}
	default:
		err = fmt.Errorf("ent: StatusOpinionGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (sogb *StatusOpinionGroupBy) IntX(ctx context.Context) int {
	v, err := sogb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (sogb *StatusOpinionGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(sogb.fields) > 1 {
		return nil, errors.New("ent: StatusOpinionGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := sogb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (sogb *StatusOpinionGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := sogb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (sogb *StatusOpinionGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = sogb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{statusopinion.Label}
	default:
		err = fmt.Errorf("ent: StatusOpinionGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (sogb *StatusOpinionGroupBy) Float64X(ctx context.Context) float64 {
	v, err := sogb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (sogb *StatusOpinionGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(sogb.fields) > 1 {
		return nil, errors.New("ent: StatusOpinionGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := sogb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (sogb *StatusOpinionGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := sogb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (sogb *StatusOpinionGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = sogb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{statusopinion.Label}
	default:
		err = fmt.Errorf("ent: StatusOpinionGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (sogb *StatusOpinionGroupBy) BoolX(ctx context.Context) bool {
	v, err := sogb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (sogb *StatusOpinionGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := sogb.sqlQuery().Query()
	if err := sogb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (sogb *StatusOpinionGroupBy) sqlQuery() *sql.Selector {
	selector := sogb.sql
	columns := make([]string, 0, len(sogb.fields)+len(sogb.fns))
	columns = append(columns, sogb.fields...)
	for _, fn := range sogb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(sogb.fields...)
}

// StatusOpinionSelect is the builder for select fields of StatusOpinion entities.
type StatusOpinionSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (sos *StatusOpinionSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := sos.path(ctx)
	if err != nil {
		return err
	}
	sos.sql = query
	return sos.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (sos *StatusOpinionSelect) ScanX(ctx context.Context, v interface{}) {
	if err := sos.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (sos *StatusOpinionSelect) Strings(ctx context.Context) ([]string, error) {
	if len(sos.fields) > 1 {
		return nil, errors.New("ent: StatusOpinionSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := sos.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (sos *StatusOpinionSelect) StringsX(ctx context.Context) []string {
	v, err := sos.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (sos *StatusOpinionSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = sos.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{statusopinion.Label}
	default:
		err = fmt.Errorf("ent: StatusOpinionSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (sos *StatusOpinionSelect) StringX(ctx context.Context) string {
	v, err := sos.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (sos *StatusOpinionSelect) Ints(ctx context.Context) ([]int, error) {
	if len(sos.fields) > 1 {
		return nil, errors.New("ent: StatusOpinionSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := sos.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (sos *StatusOpinionSelect) IntsX(ctx context.Context) []int {
	v, err := sos.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (sos *StatusOpinionSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = sos.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{statusopinion.Label}
	default:
		err = fmt.Errorf("ent: StatusOpinionSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (sos *StatusOpinionSelect) IntX(ctx context.Context) int {
	v, err := sos.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (sos *StatusOpinionSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(sos.fields) > 1 {
		return nil, errors.New("ent: StatusOpinionSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := sos.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (sos *StatusOpinionSelect) Float64sX(ctx context.Context) []float64 {
	v, err := sos.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (sos *StatusOpinionSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = sos.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{statusopinion.Label}
	default:
		err = fmt.Errorf("ent: StatusOpinionSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (sos *StatusOpinionSelect) Float64X(ctx context.Context) float64 {
	v, err := sos.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (sos *StatusOpinionSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(sos.fields) > 1 {
		return nil, errors.New("ent: StatusOpinionSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := sos.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (sos *StatusOpinionSelect) BoolsX(ctx context.Context) []bool {
	v, err := sos.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (sos *StatusOpinionSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = sos.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{statusopinion.Label}
	default:
		err = fmt.Errorf("ent: StatusOpinionSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (sos *StatusOpinionSelect) BoolX(ctx context.Context) bool {
	v, err := sos.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (sos *StatusOpinionSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := sos.sqlQuery().Query()
	if err := sos.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (sos *StatusOpinionSelect) sqlQuery() sql.Querier {
	selector := sos.sql
	selector.Select(selector.Columns(sos.fields...)...)
	return selector
}
