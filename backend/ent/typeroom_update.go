// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team18/app/ent/dataroom"
	"github.com/team18/app/ent/predicate"
	"github.com/team18/app/ent/typeroom"
)

// TypeRoomUpdate is the builder for updating TypeRoom entities.
type TypeRoomUpdate struct {
	config
	hooks      []Hook
	mutation   *TypeRoomMutation
	predicates []predicate.TypeRoom
}

// Where adds a new predicate for the builder.
func (tru *TypeRoomUpdate) Where(ps ...predicate.TypeRoom) *TypeRoomUpdate {
	tru.predicates = append(tru.predicates, ps...)
	return tru
}

// SetTypeName sets the type_name field.
func (tru *TypeRoomUpdate) SetTypeName(s string) *TypeRoomUpdate {
	tru.mutation.SetTypeName(s)
	return tru
}

// AddDataroomIDs adds the datarooms edge to DataRoom by ids.
func (tru *TypeRoomUpdate) AddDataroomIDs(ids ...int) *TypeRoomUpdate {
	tru.mutation.AddDataroomIDs(ids...)
	return tru
}

// AddDatarooms adds the datarooms edges to DataRoom.
func (tru *TypeRoomUpdate) AddDatarooms(d ...*DataRoom) *TypeRoomUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return tru.AddDataroomIDs(ids...)
}

// Mutation returns the TypeRoomMutation object of the builder.
func (tru *TypeRoomUpdate) Mutation() *TypeRoomMutation {
	return tru.mutation
}

// RemoveDataroomIDs removes the datarooms edge to DataRoom by ids.
func (tru *TypeRoomUpdate) RemoveDataroomIDs(ids ...int) *TypeRoomUpdate {
	tru.mutation.RemoveDataroomIDs(ids...)
	return tru
}

// RemoveDatarooms removes datarooms edges to DataRoom.
func (tru *TypeRoomUpdate) RemoveDatarooms(d ...*DataRoom) *TypeRoomUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return tru.RemoveDataroomIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (tru *TypeRoomUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := tru.mutation.TypeName(); ok {
		if err := typeroom.TypeNameValidator(v); err != nil {
			return 0, &ValidationError{Name: "type_name", err: fmt.Errorf("ent: validator failed for field \"type_name\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(tru.hooks) == 0 {
		affected, err = tru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TypeRoomMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tru.mutation = mutation
			affected, err = tru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tru.hooks) - 1; i >= 0; i-- {
			mut = tru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tru *TypeRoomUpdate) SaveX(ctx context.Context) int {
	affected, err := tru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tru *TypeRoomUpdate) Exec(ctx context.Context) error {
	_, err := tru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tru *TypeRoomUpdate) ExecX(ctx context.Context) {
	if err := tru.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tru *TypeRoomUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   typeroom.Table,
			Columns: typeroom.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: typeroom.FieldID,
			},
		},
	}
	if ps := tru.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tru.mutation.TypeName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: typeroom.FieldTypeName,
		})
	}
	if nodes := tru.mutation.RemovedDataroomsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   typeroom.DataroomsTable,
			Columns: []string{typeroom.DataroomsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: dataroom.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tru.mutation.DataroomsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   typeroom.DataroomsTable,
			Columns: []string{typeroom.DataroomsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: dataroom.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{typeroom.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// TypeRoomUpdateOne is the builder for updating a single TypeRoom entity.
type TypeRoomUpdateOne struct {
	config
	hooks    []Hook
	mutation *TypeRoomMutation
}

// SetTypeName sets the type_name field.
func (truo *TypeRoomUpdateOne) SetTypeName(s string) *TypeRoomUpdateOne {
	truo.mutation.SetTypeName(s)
	return truo
}

// AddDataroomIDs adds the datarooms edge to DataRoom by ids.
func (truo *TypeRoomUpdateOne) AddDataroomIDs(ids ...int) *TypeRoomUpdateOne {
	truo.mutation.AddDataroomIDs(ids...)
	return truo
}

// AddDatarooms adds the datarooms edges to DataRoom.
func (truo *TypeRoomUpdateOne) AddDatarooms(d ...*DataRoom) *TypeRoomUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return truo.AddDataroomIDs(ids...)
}

// Mutation returns the TypeRoomMutation object of the builder.
func (truo *TypeRoomUpdateOne) Mutation() *TypeRoomMutation {
	return truo.mutation
}

// RemoveDataroomIDs removes the datarooms edge to DataRoom by ids.
func (truo *TypeRoomUpdateOne) RemoveDataroomIDs(ids ...int) *TypeRoomUpdateOne {
	truo.mutation.RemoveDataroomIDs(ids...)
	return truo
}

// RemoveDatarooms removes datarooms edges to DataRoom.
func (truo *TypeRoomUpdateOne) RemoveDatarooms(d ...*DataRoom) *TypeRoomUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return truo.RemoveDataroomIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (truo *TypeRoomUpdateOne) Save(ctx context.Context) (*TypeRoom, error) {
	if v, ok := truo.mutation.TypeName(); ok {
		if err := typeroom.TypeNameValidator(v); err != nil {
			return nil, &ValidationError{Name: "type_name", err: fmt.Errorf("ent: validator failed for field \"type_name\": %w", err)}
		}
	}

	var (
		err  error
		node *TypeRoom
	)
	if len(truo.hooks) == 0 {
		node, err = truo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TypeRoomMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			truo.mutation = mutation
			node, err = truo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(truo.hooks) - 1; i >= 0; i-- {
			mut = truo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, truo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (truo *TypeRoomUpdateOne) SaveX(ctx context.Context) *TypeRoom {
	tr, err := truo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return tr
}

// Exec executes the query on the entity.
func (truo *TypeRoomUpdateOne) Exec(ctx context.Context) error {
	_, err := truo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (truo *TypeRoomUpdateOne) ExecX(ctx context.Context) {
	if err := truo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (truo *TypeRoomUpdateOne) sqlSave(ctx context.Context) (tr *TypeRoom, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   typeroom.Table,
			Columns: typeroom.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: typeroom.FieldID,
			},
		},
	}
	id, ok := truo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing TypeRoom.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := truo.mutation.TypeName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: typeroom.FieldTypeName,
		})
	}
	if nodes := truo.mutation.RemovedDataroomsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   typeroom.DataroomsTable,
			Columns: []string{typeroom.DataroomsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: dataroom.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := truo.mutation.DataroomsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   typeroom.DataroomsTable,
			Columns: []string{typeroom.DataroomsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: dataroom.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	tr = &TypeRoom{config: truo.config}
	_spec.Assign = tr.assignValues
	_spec.ScanValues = tr.scanValues()
	if err = sqlgraph.UpdateNode(ctx, truo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{typeroom.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return tr, nil
}