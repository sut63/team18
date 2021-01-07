// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team18/app/ent/furnituredetail"
	"github.com/team18/app/ent/furnituretype"
)

// FurnitureTypeCreate is the builder for creating a FurnitureType entity.
type FurnitureTypeCreate struct {
	config
	mutation *FurnitureTypeMutation
	hooks    []Hook
}

// SetFurnitureType sets the furniture_type field.
func (ftc *FurnitureTypeCreate) SetFurnitureType(s string) *FurnitureTypeCreate {
	ftc.mutation.SetFurnitureType(s)
	return ftc
}

// AddDetailIDs adds the details edge to FurnitureDetail by ids.
func (ftc *FurnitureTypeCreate) AddDetailIDs(ids ...int) *FurnitureTypeCreate {
	ftc.mutation.AddDetailIDs(ids...)
	return ftc
}

// AddDetails adds the details edges to FurnitureDetail.
func (ftc *FurnitureTypeCreate) AddDetails(f ...*FurnitureDetail) *FurnitureTypeCreate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return ftc.AddDetailIDs(ids...)
}

// Mutation returns the FurnitureTypeMutation object of the builder.
func (ftc *FurnitureTypeCreate) Mutation() *FurnitureTypeMutation {
	return ftc.mutation
}

// Save creates the FurnitureType in the database.
func (ftc *FurnitureTypeCreate) Save(ctx context.Context) (*FurnitureType, error) {
	if _, ok := ftc.mutation.FurnitureType(); !ok {
		return nil, &ValidationError{Name: "furniture_type", err: errors.New("ent: missing required field \"furniture_type\"")}
	}
	var (
		err  error
		node *FurnitureType
	)
	if len(ftc.hooks) == 0 {
		node, err = ftc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FurnitureTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ftc.mutation = mutation
			node, err = ftc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ftc.hooks) - 1; i >= 0; i-- {
			mut = ftc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ftc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ftc *FurnitureTypeCreate) SaveX(ctx context.Context) *FurnitureType {
	v, err := ftc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ftc *FurnitureTypeCreate) sqlSave(ctx context.Context) (*FurnitureType, error) {
	ft, _spec := ftc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ftc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	ft.ID = int(id)
	return ft, nil
}

func (ftc *FurnitureTypeCreate) createSpec() (*FurnitureType, *sqlgraph.CreateSpec) {
	var (
		ft    = &FurnitureType{config: ftc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: furnituretype.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: furnituretype.FieldID,
			},
		}
	)
	if value, ok := ftc.mutation.FurnitureType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: furnituretype.FieldFurnitureType,
		})
		ft.FurnitureType = value
	}
	if nodes := ftc.mutation.DetailsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   furnituretype.DetailsTable,
			Columns: []string{furnituretype.DetailsColumn},
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
	return ft, _spec
}