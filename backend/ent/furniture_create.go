// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team18/app/ent/furniture"
	"github.com/team18/app/ent/furnituredetail"
)

// FurnitureCreate is the builder for creating a Furniture entity.
type FurnitureCreate struct {
	config
	mutation *FurnitureMutation
	hooks    []Hook
}

// SetFurnitureName sets the furniture_name field.
func (fc *FurnitureCreate) SetFurnitureName(s string) *FurnitureCreate {
	fc.mutation.SetFurnitureName(s)
	return fc
}

// AddDetailIDs adds the details edge to FurnitureDetail by ids.
func (fc *FurnitureCreate) AddDetailIDs(ids ...int) *FurnitureCreate {
	fc.mutation.AddDetailIDs(ids...)
	return fc
}

// AddDetails adds the details edges to FurnitureDetail.
func (fc *FurnitureCreate) AddDetails(f ...*FurnitureDetail) *FurnitureCreate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fc.AddDetailIDs(ids...)
}

// Mutation returns the FurnitureMutation object of the builder.
func (fc *FurnitureCreate) Mutation() *FurnitureMutation {
	return fc.mutation
}

// Save creates the Furniture in the database.
func (fc *FurnitureCreate) Save(ctx context.Context) (*Furniture, error) {
	if _, ok := fc.mutation.FurnitureName(); !ok {
		return nil, &ValidationError{Name: "furniture_name", err: errors.New("ent: missing required field \"furniture_name\"")}
	}
	if v, ok := fc.mutation.FurnitureName(); ok {
		if err := furniture.FurnitureNameValidator(v); err != nil {
			return nil, &ValidationError{Name: "furniture_name", err: fmt.Errorf("ent: validator failed for field \"furniture_name\": %w", err)}
		}
	}
	var (
		err  error
		node *Furniture
	)
	if len(fc.hooks) == 0 {
		node, err = fc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FurnitureMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			fc.mutation = mutation
			node, err = fc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(fc.hooks) - 1; i >= 0; i-- {
			mut = fc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (fc *FurnitureCreate) SaveX(ctx context.Context) *Furniture {
	v, err := fc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (fc *FurnitureCreate) sqlSave(ctx context.Context) (*Furniture, error) {
	f, _spec := fc.createSpec()
	if err := sqlgraph.CreateNode(ctx, fc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	f.ID = int(id)
	return f, nil
}

func (fc *FurnitureCreate) createSpec() (*Furniture, *sqlgraph.CreateSpec) {
	var (
		f     = &Furniture{config: fc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: furniture.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: furniture.FieldID,
			},
		}
	)
	if value, ok := fc.mutation.FurnitureName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: furniture.FieldFurnitureName,
		})
		f.FurnitureName = value
	}
	if nodes := fc.mutation.DetailsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   furniture.DetailsTable,
			Columns: []string{furniture.DetailsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: furnituredetail.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return f, _spec
}