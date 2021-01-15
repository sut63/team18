// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team18/app/ent/checkin"
	"github.com/team18/app/ent/statuscheckin"
)

// StatusCheckInCreate is the builder for creating a StatusCheckIn entity.
type StatusCheckInCreate struct {
	config
	mutation *StatusCheckInMutation
	hooks    []Hook
}

// SetStatusName sets the status_name field.
func (scic *StatusCheckInCreate) SetStatusName(s string) *StatusCheckInCreate {
	scic.mutation.SetStatusName(s)
	return scic
}

// AddCheckinIDs adds the checkins edge to CheckIn by ids.
func (scic *StatusCheckInCreate) AddCheckinIDs(ids ...int) *StatusCheckInCreate {
	scic.mutation.AddCheckinIDs(ids...)
	return scic
}

// AddCheckins adds the checkins edges to CheckIn.
func (scic *StatusCheckInCreate) AddCheckins(c ...*CheckIn) *StatusCheckInCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return scic.AddCheckinIDs(ids...)
}

// Mutation returns the StatusCheckInMutation object of the builder.
func (scic *StatusCheckInCreate) Mutation() *StatusCheckInMutation {
	return scic.mutation
}

// Save creates the StatusCheckIn in the database.
func (scic *StatusCheckInCreate) Save(ctx context.Context) (*StatusCheckIn, error) {
	if _, ok := scic.mutation.StatusName(); !ok {
		return nil, &ValidationError{Name: "status_name", err: errors.New("ent: missing required field \"status_name\"")}
	}
	if v, ok := scic.mutation.StatusName(); ok {
		if err := statuscheckin.StatusNameValidator(v); err != nil {
			return nil, &ValidationError{Name: "status_name", err: fmt.Errorf("ent: validator failed for field \"status_name\": %w", err)}
		}
	}
	var (
		err  error
		node *StatusCheckIn
	)
	if len(scic.hooks) == 0 {
		node, err = scic.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*StatusCheckInMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			scic.mutation = mutation
			node, err = scic.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(scic.hooks) - 1; i >= 0; i-- {
			mut = scic.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, scic.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (scic *StatusCheckInCreate) SaveX(ctx context.Context) *StatusCheckIn {
	v, err := scic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (scic *StatusCheckInCreate) sqlSave(ctx context.Context) (*StatusCheckIn, error) {
	sci, _spec := scic.createSpec()
	if err := sqlgraph.CreateNode(ctx, scic.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	sci.ID = int(id)
	return sci, nil
}

func (scic *StatusCheckInCreate) createSpec() (*StatusCheckIn, *sqlgraph.CreateSpec) {
	var (
		sci   = &StatusCheckIn{config: scic.config}
		_spec = &sqlgraph.CreateSpec{
			Table: statuscheckin.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: statuscheckin.FieldID,
			},
		}
	)
	if value, ok := scic.mutation.StatusName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: statuscheckin.FieldStatusName,
		})
		sci.StatusName = value
	}
	if nodes := scic.mutation.CheckinsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   statuscheckin.CheckinsTable,
			Columns: []string{statuscheckin.CheckinsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: checkin.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return sci, _spec
}