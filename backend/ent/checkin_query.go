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
	"github.com/team18/app/ent/checkin"
	"github.com/team18/app/ent/checkout"
	"github.com/team18/app/ent/counterstaff"
	"github.com/team18/app/ent/customer"
	"github.com/team18/app/ent/dataroom"
	"github.com/team18/app/ent/predicate"
	"github.com/team18/app/ent/reserveroom"
	"github.com/team18/app/ent/statuscheckin"
)

// CheckInQuery is the builder for querying CheckIn entities.
type CheckInQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.CheckIn
	// eager-loading edges.
	withCustomer    *CustomerQuery
	withCounter     *CounterStaffQuery
	withReserveroom *ReserveRoomQuery
	withDataroom    *DataRoomQuery
	withStatus      *StatusCheckInQuery
	withCheckouts   *CheckoutQuery
	withFKs         bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (ciq *CheckInQuery) Where(ps ...predicate.CheckIn) *CheckInQuery {
	ciq.predicates = append(ciq.predicates, ps...)
	return ciq
}

// Limit adds a limit step to the query.
func (ciq *CheckInQuery) Limit(limit int) *CheckInQuery {
	ciq.limit = &limit
	return ciq
}

// Offset adds an offset step to the query.
func (ciq *CheckInQuery) Offset(offset int) *CheckInQuery {
	ciq.offset = &offset
	return ciq
}

// Order adds an order step to the query.
func (ciq *CheckInQuery) Order(o ...OrderFunc) *CheckInQuery {
	ciq.order = append(ciq.order, o...)
	return ciq
}

// QueryCustomer chains the current query on the customer edge.
func (ciq *CheckInQuery) QueryCustomer() *CustomerQuery {
	query := &CustomerQuery{config: ciq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ciq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(checkin.Table, checkin.FieldID, ciq.sqlQuery()),
			sqlgraph.To(customer.Table, customer.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, checkin.CustomerTable, checkin.CustomerColumn),
		)
		fromU = sqlgraph.SetNeighbors(ciq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCounter chains the current query on the counter edge.
func (ciq *CheckInQuery) QueryCounter() *CounterStaffQuery {
	query := &CounterStaffQuery{config: ciq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ciq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(checkin.Table, checkin.FieldID, ciq.sqlQuery()),
			sqlgraph.To(counterstaff.Table, counterstaff.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, checkin.CounterTable, checkin.CounterColumn),
		)
		fromU = sqlgraph.SetNeighbors(ciq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryReserveroom chains the current query on the reserveroom edge.
func (ciq *CheckInQuery) QueryReserveroom() *ReserveRoomQuery {
	query := &ReserveRoomQuery{config: ciq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ciq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(checkin.Table, checkin.FieldID, ciq.sqlQuery()),
			sqlgraph.To(reserveroom.Table, reserveroom.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, checkin.ReserveroomTable, checkin.ReserveroomColumn),
		)
		fromU = sqlgraph.SetNeighbors(ciq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryDataroom chains the current query on the dataroom edge.
func (ciq *CheckInQuery) QueryDataroom() *DataRoomQuery {
	query := &DataRoomQuery{config: ciq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ciq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(checkin.Table, checkin.FieldID, ciq.sqlQuery()),
			sqlgraph.To(dataroom.Table, dataroom.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, checkin.DataroomTable, checkin.DataroomColumn),
		)
		fromU = sqlgraph.SetNeighbors(ciq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryStatus chains the current query on the status edge.
func (ciq *CheckInQuery) QueryStatus() *StatusCheckInQuery {
	query := &StatusCheckInQuery{config: ciq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ciq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(checkin.Table, checkin.FieldID, ciq.sqlQuery()),
			sqlgraph.To(statuscheckin.Table, statuscheckin.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, checkin.StatusTable, checkin.StatusColumn),
		)
		fromU = sqlgraph.SetNeighbors(ciq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCheckouts chains the current query on the checkouts edge.
func (ciq *CheckInQuery) QueryCheckouts() *CheckoutQuery {
	query := &CheckoutQuery{config: ciq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ciq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(checkin.Table, checkin.FieldID, ciq.sqlQuery()),
			sqlgraph.To(checkout.Table, checkout.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, checkin.CheckoutsTable, checkin.CheckoutsColumn),
		)
		fromU = sqlgraph.SetNeighbors(ciq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first CheckIn entity in the query. Returns *NotFoundError when no checkin was found.
func (ciq *CheckInQuery) First(ctx context.Context) (*CheckIn, error) {
	cis, err := ciq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(cis) == 0 {
		return nil, &NotFoundError{checkin.Label}
	}
	return cis[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ciq *CheckInQuery) FirstX(ctx context.Context) *CheckIn {
	ci, err := ciq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return ci
}

// FirstID returns the first CheckIn id in the query. Returns *NotFoundError when no id was found.
func (ciq *CheckInQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ciq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{checkin.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (ciq *CheckInQuery) FirstXID(ctx context.Context) int {
	id, err := ciq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only CheckIn entity in the query, returns an error if not exactly one entity was returned.
func (ciq *CheckInQuery) Only(ctx context.Context) (*CheckIn, error) {
	cis, err := ciq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(cis) {
	case 1:
		return cis[0], nil
	case 0:
		return nil, &NotFoundError{checkin.Label}
	default:
		return nil, &NotSingularError{checkin.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ciq *CheckInQuery) OnlyX(ctx context.Context) *CheckIn {
	ci, err := ciq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return ci
}

// OnlyID returns the only CheckIn id in the query, returns an error if not exactly one id was returned.
func (ciq *CheckInQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ciq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{checkin.Label}
	default:
		err = &NotSingularError{checkin.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ciq *CheckInQuery) OnlyIDX(ctx context.Context) int {
	id, err := ciq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of CheckIns.
func (ciq *CheckInQuery) All(ctx context.Context) ([]*CheckIn, error) {
	if err := ciq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return ciq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (ciq *CheckInQuery) AllX(ctx context.Context) []*CheckIn {
	cis, err := ciq.All(ctx)
	if err != nil {
		panic(err)
	}
	return cis
}

// IDs executes the query and returns a list of CheckIn ids.
func (ciq *CheckInQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := ciq.Select(checkin.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ciq *CheckInQuery) IDsX(ctx context.Context) []int {
	ids, err := ciq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ciq *CheckInQuery) Count(ctx context.Context) (int, error) {
	if err := ciq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return ciq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (ciq *CheckInQuery) CountX(ctx context.Context) int {
	count, err := ciq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ciq *CheckInQuery) Exist(ctx context.Context) (bool, error) {
	if err := ciq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return ciq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (ciq *CheckInQuery) ExistX(ctx context.Context) bool {
	exist, err := ciq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ciq *CheckInQuery) Clone() *CheckInQuery {
	return &CheckInQuery{
		config:     ciq.config,
		limit:      ciq.limit,
		offset:     ciq.offset,
		order:      append([]OrderFunc{}, ciq.order...),
		unique:     append([]string{}, ciq.unique...),
		predicates: append([]predicate.CheckIn{}, ciq.predicates...),
		// clone intermediate query.
		sql:  ciq.sql.Clone(),
		path: ciq.path,
	}
}

//  WithCustomer tells the query-builder to eager-loads the nodes that are connected to
// the "customer" edge. The optional arguments used to configure the query builder of the edge.
func (ciq *CheckInQuery) WithCustomer(opts ...func(*CustomerQuery)) *CheckInQuery {
	query := &CustomerQuery{config: ciq.config}
	for _, opt := range opts {
		opt(query)
	}
	ciq.withCustomer = query
	return ciq
}

//  WithCounter tells the query-builder to eager-loads the nodes that are connected to
// the "counter" edge. The optional arguments used to configure the query builder of the edge.
func (ciq *CheckInQuery) WithCounter(opts ...func(*CounterStaffQuery)) *CheckInQuery {
	query := &CounterStaffQuery{config: ciq.config}
	for _, opt := range opts {
		opt(query)
	}
	ciq.withCounter = query
	return ciq
}

//  WithReserveroom tells the query-builder to eager-loads the nodes that are connected to
// the "reserveroom" edge. The optional arguments used to configure the query builder of the edge.
func (ciq *CheckInQuery) WithReserveroom(opts ...func(*ReserveRoomQuery)) *CheckInQuery {
	query := &ReserveRoomQuery{config: ciq.config}
	for _, opt := range opts {
		opt(query)
	}
	ciq.withReserveroom = query
	return ciq
}

//  WithDataroom tells the query-builder to eager-loads the nodes that are connected to
// the "dataroom" edge. The optional arguments used to configure the query builder of the edge.
func (ciq *CheckInQuery) WithDataroom(opts ...func(*DataRoomQuery)) *CheckInQuery {
	query := &DataRoomQuery{config: ciq.config}
	for _, opt := range opts {
		opt(query)
	}
	ciq.withDataroom = query
	return ciq
}

//  WithStatus tells the query-builder to eager-loads the nodes that are connected to
// the "status" edge. The optional arguments used to configure the query builder of the edge.
func (ciq *CheckInQuery) WithStatus(opts ...func(*StatusCheckInQuery)) *CheckInQuery {
	query := &StatusCheckInQuery{config: ciq.config}
	for _, opt := range opts {
		opt(query)
	}
	ciq.withStatus = query
	return ciq
}

//  WithCheckouts tells the query-builder to eager-loads the nodes that are connected to
// the "checkouts" edge. The optional arguments used to configure the query builder of the edge.
func (ciq *CheckInQuery) WithCheckouts(opts ...func(*CheckoutQuery)) *CheckInQuery {
	query := &CheckoutQuery{config: ciq.config}
	for _, opt := range opts {
		opt(query)
	}
	ciq.withCheckouts = query
	return ciq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CheckinDate time.Time `json:"checkin_date,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.CheckIn.Query().
//		GroupBy(checkin.FieldCheckinDate).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (ciq *CheckInQuery) GroupBy(field string, fields ...string) *CheckInGroupBy {
	group := &CheckInGroupBy{config: ciq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ciq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ciq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		CheckinDate time.Time `json:"checkin_date,omitempty"`
//	}
//
//	client.CheckIn.Query().
//		Select(checkin.FieldCheckinDate).
//		Scan(ctx, &v)
//
func (ciq *CheckInQuery) Select(field string, fields ...string) *CheckInSelect {
	selector := &CheckInSelect{config: ciq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ciq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ciq.sqlQuery(), nil
	}
	return selector
}

func (ciq *CheckInQuery) prepareQuery(ctx context.Context) error {
	if ciq.path != nil {
		prev, err := ciq.path(ctx)
		if err != nil {
			return err
		}
		ciq.sql = prev
	}
	return nil
}

func (ciq *CheckInQuery) sqlAll(ctx context.Context) ([]*CheckIn, error) {
	var (
		nodes       = []*CheckIn{}
		withFKs     = ciq.withFKs
		_spec       = ciq.querySpec()
		loadedTypes = [6]bool{
			ciq.withCustomer != nil,
			ciq.withCounter != nil,
			ciq.withReserveroom != nil,
			ciq.withDataroom != nil,
			ciq.withStatus != nil,
			ciq.withCheckouts != nil,
		}
	)
	if ciq.withCustomer != nil || ciq.withCounter != nil || ciq.withReserveroom != nil || ciq.withDataroom != nil || ciq.withStatus != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, checkin.ForeignKeys...)
	}
	_spec.ScanValues = func() []interface{} {
		node := &CheckIn{config: ciq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		if withFKs {
			values = append(values, node.fkValues()...)
		}
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
	if err := sqlgraph.QueryNodes(ctx, ciq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := ciq.withCustomer; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*CheckIn)
		for i := range nodes {
			if fk := nodes[i].customer_id; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(customer.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "customer_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Customer = n
			}
		}
	}

	if query := ciq.withCounter; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*CheckIn)
		for i := range nodes {
			if fk := nodes[i].staff_id; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(counterstaff.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "staff_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Counter = n
			}
		}
	}

	if query := ciq.withReserveroom; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*CheckIn)
		for i := range nodes {
			if fk := nodes[i].reserves_id; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(reserveroom.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "reserves_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Reserveroom = n
			}
		}
	}

	if query := ciq.withDataroom; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*CheckIn)
		for i := range nodes {
			if fk := nodes[i].room_id; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(dataroom.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "room_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Dataroom = n
			}
		}
	}

	if query := ciq.withStatus; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*CheckIn)
		for i := range nodes {
			if fk := nodes[i].status_id; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(statuscheckin.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "status_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Status = n
			}
		}
	}

	if query := ciq.withCheckouts; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*CheckIn)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.Checkout(func(s *sql.Selector) {
			s.Where(sql.InValues(checkin.CheckoutsColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.check_in_checkouts
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "check_in_checkouts" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "check_in_checkouts" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Checkouts = n
		}
	}

	return nodes, nil
}

func (ciq *CheckInQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ciq.querySpec()
	return sqlgraph.CountNodes(ctx, ciq.driver, _spec)
}

func (ciq *CheckInQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := ciq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (ciq *CheckInQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   checkin.Table,
			Columns: checkin.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: checkin.FieldID,
			},
		},
		From:   ciq.sql,
		Unique: true,
	}
	if ps := ciq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ciq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ciq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ciq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ciq *CheckInQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(ciq.driver.Dialect())
	t1 := builder.Table(checkin.Table)
	selector := builder.Select(t1.Columns(checkin.Columns...)...).From(t1)
	if ciq.sql != nil {
		selector = ciq.sql
		selector.Select(selector.Columns(checkin.Columns...)...)
	}
	for _, p := range ciq.predicates {
		p(selector)
	}
	for _, p := range ciq.order {
		p(selector)
	}
	if offset := ciq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ciq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// CheckInGroupBy is the builder for group-by CheckIn entities.
type CheckInGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cigb *CheckInGroupBy) Aggregate(fns ...AggregateFunc) *CheckInGroupBy {
	cigb.fns = append(cigb.fns, fns...)
	return cigb
}

// Scan applies the group-by query and scan the result into the given value.
func (cigb *CheckInGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := cigb.path(ctx)
	if err != nil {
		return err
	}
	cigb.sql = query
	return cigb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (cigb *CheckInGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := cigb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (cigb *CheckInGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(cigb.fields) > 1 {
		return nil, errors.New("ent: CheckInGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := cigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (cigb *CheckInGroupBy) StringsX(ctx context.Context) []string {
	v, err := cigb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (cigb *CheckInGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = cigb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{checkin.Label}
	default:
		err = fmt.Errorf("ent: CheckInGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (cigb *CheckInGroupBy) StringX(ctx context.Context) string {
	v, err := cigb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (cigb *CheckInGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(cigb.fields) > 1 {
		return nil, errors.New("ent: CheckInGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := cigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (cigb *CheckInGroupBy) IntsX(ctx context.Context) []int {
	v, err := cigb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (cigb *CheckInGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = cigb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{checkin.Label}
	default:
		err = fmt.Errorf("ent: CheckInGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (cigb *CheckInGroupBy) IntX(ctx context.Context) int {
	v, err := cigb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (cigb *CheckInGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(cigb.fields) > 1 {
		return nil, errors.New("ent: CheckInGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := cigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (cigb *CheckInGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := cigb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (cigb *CheckInGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = cigb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{checkin.Label}
	default:
		err = fmt.Errorf("ent: CheckInGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (cigb *CheckInGroupBy) Float64X(ctx context.Context) float64 {
	v, err := cigb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (cigb *CheckInGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(cigb.fields) > 1 {
		return nil, errors.New("ent: CheckInGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := cigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (cigb *CheckInGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := cigb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (cigb *CheckInGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = cigb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{checkin.Label}
	default:
		err = fmt.Errorf("ent: CheckInGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (cigb *CheckInGroupBy) BoolX(ctx context.Context) bool {
	v, err := cigb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (cigb *CheckInGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := cigb.sqlQuery().Query()
	if err := cigb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (cigb *CheckInGroupBy) sqlQuery() *sql.Selector {
	selector := cigb.sql
	columns := make([]string, 0, len(cigb.fields)+len(cigb.fns))
	columns = append(columns, cigb.fields...)
	for _, fn := range cigb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(cigb.fields...)
}

// CheckInSelect is the builder for select fields of CheckIn entities.
type CheckInSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (cis *CheckInSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := cis.path(ctx)
	if err != nil {
		return err
	}
	cis.sql = query
	return cis.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (cis *CheckInSelect) ScanX(ctx context.Context, v interface{}) {
	if err := cis.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (cis *CheckInSelect) Strings(ctx context.Context) ([]string, error) {
	if len(cis.fields) > 1 {
		return nil, errors.New("ent: CheckInSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := cis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (cis *CheckInSelect) StringsX(ctx context.Context) []string {
	v, err := cis.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (cis *CheckInSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = cis.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{checkin.Label}
	default:
		err = fmt.Errorf("ent: CheckInSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (cis *CheckInSelect) StringX(ctx context.Context) string {
	v, err := cis.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (cis *CheckInSelect) Ints(ctx context.Context) ([]int, error) {
	if len(cis.fields) > 1 {
		return nil, errors.New("ent: CheckInSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := cis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (cis *CheckInSelect) IntsX(ctx context.Context) []int {
	v, err := cis.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (cis *CheckInSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = cis.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{checkin.Label}
	default:
		err = fmt.Errorf("ent: CheckInSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (cis *CheckInSelect) IntX(ctx context.Context) int {
	v, err := cis.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (cis *CheckInSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(cis.fields) > 1 {
		return nil, errors.New("ent: CheckInSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := cis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (cis *CheckInSelect) Float64sX(ctx context.Context) []float64 {
	v, err := cis.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (cis *CheckInSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = cis.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{checkin.Label}
	default:
		err = fmt.Errorf("ent: CheckInSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (cis *CheckInSelect) Float64X(ctx context.Context) float64 {
	v, err := cis.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (cis *CheckInSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(cis.fields) > 1 {
		return nil, errors.New("ent: CheckInSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := cis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (cis *CheckInSelect) BoolsX(ctx context.Context) []bool {
	v, err := cis.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (cis *CheckInSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = cis.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{checkin.Label}
	default:
		err = fmt.Errorf("ent: CheckInSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (cis *CheckInSelect) BoolX(ctx context.Context) bool {
	v, err := cis.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (cis *CheckInSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := cis.sqlQuery().Query()
	if err := cis.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (cis *CheckInSelect) sqlQuery() sql.Querier {
	selector := cis.sql
	selector.Select(selector.Columns(cis.fields...)...)
	return selector
}
