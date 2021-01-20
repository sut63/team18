// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team18/app/ent/customer"
	"github.com/team18/app/ent/dataroom"
	"github.com/team18/app/ent/fixroom"
	"github.com/team18/app/ent/furnituredetail"
	"github.com/team18/app/ent/predicate"
)

// FixRoomUpdate is the builder for updating FixRoom entities.
type FixRoomUpdate struct {
	config
	hooks      []Hook
	mutation   *FixRoomMutation
	predicates []predicate.FixRoom
}

// Where adds a new predicate for the builder.
func (fru *FixRoomUpdate) Where(ps ...predicate.FixRoom) *FixRoomUpdate {
	fru.predicates = append(fru.predicates, ps...)
	return fru
}

// SetFixDetail sets the fix_detail field.
func (fru *FixRoomUpdate) SetFixDetail(s string) *FixRoomUpdate {
	fru.mutation.SetFixDetail(s)
	return fru
}

// SetPhoneNumber sets the phone_number field.
func (fru *FixRoomUpdate) SetPhoneNumber(s string) *FixRoomUpdate {
	fru.mutation.SetPhoneNumber(s)
	return fru
}

// SetFacebook sets the facebook field.
func (fru *FixRoomUpdate) SetFacebook(s string) *FixRoomUpdate {
	fru.mutation.SetFacebook(s)
	return fru
}

// SetCustomerID sets the customer edge to Customer by id.
func (fru *FixRoomUpdate) SetCustomerID(id int) *FixRoomUpdate {
	fru.mutation.SetCustomerID(id)
	return fru
}

// SetNillableCustomerID sets the customer edge to Customer by id if the given value is not nil.
func (fru *FixRoomUpdate) SetNillableCustomerID(id *int) *FixRoomUpdate {
	if id != nil {
		fru = fru.SetCustomerID(*id)
	}
	return fru
}

// SetCustomer sets the customer edge to Customer.
func (fru *FixRoomUpdate) SetCustomer(c *Customer) *FixRoomUpdate {
	return fru.SetCustomerID(c.ID)
}

// SetFurnitureDetailID sets the furnitureDetail edge to FurnitureDetail by id.
func (fru *FixRoomUpdate) SetFurnitureDetailID(id int) *FixRoomUpdate {
	fru.mutation.SetFurnitureDetailID(id)
	return fru
}

// SetNillableFurnitureDetailID sets the furnitureDetail edge to FurnitureDetail by id if the given value is not nil.
func (fru *FixRoomUpdate) SetNillableFurnitureDetailID(id *int) *FixRoomUpdate {
	if id != nil {
		fru = fru.SetFurnitureDetailID(*id)
	}
	return fru
}

// SetFurnitureDetail sets the furnitureDetail edge to FurnitureDetail.
func (fru *FixRoomUpdate) SetFurnitureDetail(f *FurnitureDetail) *FixRoomUpdate {
	return fru.SetFurnitureDetailID(f.ID)
}

// SetRoomID sets the room edge to DataRoom by id.
func (fru *FixRoomUpdate) SetRoomID(id int) *FixRoomUpdate {
	fru.mutation.SetRoomID(id)
	return fru
}

// SetNillableRoomID sets the room edge to DataRoom by id if the given value is not nil.
func (fru *FixRoomUpdate) SetNillableRoomID(id *int) *FixRoomUpdate {
	if id != nil {
		fru = fru.SetRoomID(*id)
	}
	return fru
}

// SetRoom sets the room edge to DataRoom.
func (fru *FixRoomUpdate) SetRoom(d *DataRoom) *FixRoomUpdate {
	return fru.SetRoomID(d.ID)
}

// Mutation returns the FixRoomMutation object of the builder.
func (fru *FixRoomUpdate) Mutation() *FixRoomMutation {
	return fru.mutation
}

// ClearCustomer clears the customer edge to Customer.
func (fru *FixRoomUpdate) ClearCustomer() *FixRoomUpdate {
	fru.mutation.ClearCustomer()
	return fru
}

// ClearFurnitureDetail clears the furnitureDetail edge to FurnitureDetail.
func (fru *FixRoomUpdate) ClearFurnitureDetail() *FixRoomUpdate {
	fru.mutation.ClearFurnitureDetail()
	return fru
}

// ClearRoom clears the room edge to DataRoom.
func (fru *FixRoomUpdate) ClearRoom() *FixRoomUpdate {
	fru.mutation.ClearRoom()
	return fru
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (fru *FixRoomUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := fru.mutation.FixDetail(); ok {
		if err := fixroom.FixDetailValidator(v); err != nil {
			return 0, &ValidationError{Name: "fix_detail", err: fmt.Errorf("ent: validator failed for field \"fix_detail\": %w", err)}
		}
	}
	if v, ok := fru.mutation.PhoneNumber(); ok {
		if err := fixroom.PhoneNumberValidator(v); err != nil {
			return 0, &ValidationError{Name: "phone_number", err: fmt.Errorf("ent: validator failed for field \"phone_number\": %w", err)}
		}
	}
	if v, ok := fru.mutation.Facebook(); ok {
		if err := fixroom.FacebookValidator(v); err != nil {
			return 0, &ValidationError{Name: "facebook", err: fmt.Errorf("ent: validator failed for field \"facebook\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(fru.hooks) == 0 {
		affected, err = fru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FixRoomMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			fru.mutation = mutation
			affected, err = fru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(fru.hooks) - 1; i >= 0; i-- {
			mut = fru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (fru *FixRoomUpdate) SaveX(ctx context.Context) int {
	affected, err := fru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (fru *FixRoomUpdate) Exec(ctx context.Context) error {
	_, err := fru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fru *FixRoomUpdate) ExecX(ctx context.Context) {
	if err := fru.Exec(ctx); err != nil {
		panic(err)
	}
}

func (fru *FixRoomUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   fixroom.Table,
			Columns: fixroom.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: fixroom.FieldID,
			},
		},
	}
	if ps := fru.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fru.mutation.FixDetail(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: fixroom.FieldFixDetail,
		})
	}
	if value, ok := fru.mutation.PhoneNumber(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: fixroom.FieldPhoneNumber,
		})
	}
	if value, ok := fru.mutation.Facebook(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: fixroom.FieldFacebook,
		})
	}
	if fru.mutation.CustomerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   fixroom.CustomerTable,
			Columns: []string{fixroom.CustomerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: customer.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fru.mutation.CustomerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   fixroom.CustomerTable,
			Columns: []string{fixroom.CustomerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: customer.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if fru.mutation.FurnitureDetailCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   fixroom.FurnitureDetailTable,
			Columns: []string{fixroom.FurnitureDetailColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: furnituredetail.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fru.mutation.FurnitureDetailIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   fixroom.FurnitureDetailTable,
			Columns: []string{fixroom.FurnitureDetailColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if fru.mutation.RoomCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   fixroom.RoomTable,
			Columns: []string{fixroom.RoomColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: dataroom.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fru.mutation.RoomIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   fixroom.RoomTable,
			Columns: []string{fixroom.RoomColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, fru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{fixroom.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// FixRoomUpdateOne is the builder for updating a single FixRoom entity.
type FixRoomUpdateOne struct {
	config
	hooks    []Hook
	mutation *FixRoomMutation
}

// SetFixDetail sets the fix_detail field.
func (fruo *FixRoomUpdateOne) SetFixDetail(s string) *FixRoomUpdateOne {
	fruo.mutation.SetFixDetail(s)
	return fruo
}

// SetPhoneNumber sets the phone_number field.
func (fruo *FixRoomUpdateOne) SetPhoneNumber(s string) *FixRoomUpdateOne {
	fruo.mutation.SetPhoneNumber(s)
	return fruo
}

// SetFacebook sets the facebook field.
func (fruo *FixRoomUpdateOne) SetFacebook(s string) *FixRoomUpdateOne {
	fruo.mutation.SetFacebook(s)
	return fruo
}

// SetCustomerID sets the customer edge to Customer by id.
func (fruo *FixRoomUpdateOne) SetCustomerID(id int) *FixRoomUpdateOne {
	fruo.mutation.SetCustomerID(id)
	return fruo
}

// SetNillableCustomerID sets the customer edge to Customer by id if the given value is not nil.
func (fruo *FixRoomUpdateOne) SetNillableCustomerID(id *int) *FixRoomUpdateOne {
	if id != nil {
		fruo = fruo.SetCustomerID(*id)
	}
	return fruo
}

// SetCustomer sets the customer edge to Customer.
func (fruo *FixRoomUpdateOne) SetCustomer(c *Customer) *FixRoomUpdateOne {
	return fruo.SetCustomerID(c.ID)
}

// SetFurnitureDetailID sets the furnitureDetail edge to FurnitureDetail by id.
func (fruo *FixRoomUpdateOne) SetFurnitureDetailID(id int) *FixRoomUpdateOne {
	fruo.mutation.SetFurnitureDetailID(id)
	return fruo
}

// SetNillableFurnitureDetailID sets the furnitureDetail edge to FurnitureDetail by id if the given value is not nil.
func (fruo *FixRoomUpdateOne) SetNillableFurnitureDetailID(id *int) *FixRoomUpdateOne {
	if id != nil {
		fruo = fruo.SetFurnitureDetailID(*id)
	}
	return fruo
}

// SetFurnitureDetail sets the furnitureDetail edge to FurnitureDetail.
func (fruo *FixRoomUpdateOne) SetFurnitureDetail(f *FurnitureDetail) *FixRoomUpdateOne {
	return fruo.SetFurnitureDetailID(f.ID)
}

// SetRoomID sets the room edge to DataRoom by id.
func (fruo *FixRoomUpdateOne) SetRoomID(id int) *FixRoomUpdateOne {
	fruo.mutation.SetRoomID(id)
	return fruo
}

// SetNillableRoomID sets the room edge to DataRoom by id if the given value is not nil.
func (fruo *FixRoomUpdateOne) SetNillableRoomID(id *int) *FixRoomUpdateOne {
	if id != nil {
		fruo = fruo.SetRoomID(*id)
	}
	return fruo
}

// SetRoom sets the room edge to DataRoom.
func (fruo *FixRoomUpdateOne) SetRoom(d *DataRoom) *FixRoomUpdateOne {
	return fruo.SetRoomID(d.ID)
}

// Mutation returns the FixRoomMutation object of the builder.
func (fruo *FixRoomUpdateOne) Mutation() *FixRoomMutation {
	return fruo.mutation
}

// ClearCustomer clears the customer edge to Customer.
func (fruo *FixRoomUpdateOne) ClearCustomer() *FixRoomUpdateOne {
	fruo.mutation.ClearCustomer()
	return fruo
}

// ClearFurnitureDetail clears the furnitureDetail edge to FurnitureDetail.
func (fruo *FixRoomUpdateOne) ClearFurnitureDetail() *FixRoomUpdateOne {
	fruo.mutation.ClearFurnitureDetail()
	return fruo
}

// ClearRoom clears the room edge to DataRoom.
func (fruo *FixRoomUpdateOne) ClearRoom() *FixRoomUpdateOne {
	fruo.mutation.ClearRoom()
	return fruo
}

// Save executes the query and returns the updated entity.
func (fruo *FixRoomUpdateOne) Save(ctx context.Context) (*FixRoom, error) {
	if v, ok := fruo.mutation.FixDetail(); ok {
		if err := fixroom.FixDetailValidator(v); err != nil {
			return nil, &ValidationError{Name: "fix_detail", err: fmt.Errorf("ent: validator failed for field \"fix_detail\": %w", err)}
		}
	}
	if v, ok := fruo.mutation.PhoneNumber(); ok {
		if err := fixroom.PhoneNumberValidator(v); err != nil {
			return nil, &ValidationError{Name: "phone_number", err: fmt.Errorf("ent: validator failed for field \"phone_number\": %w", err)}
		}
	}
	if v, ok := fruo.mutation.Facebook(); ok {
		if err := fixroom.FacebookValidator(v); err != nil {
			return nil, &ValidationError{Name: "facebook", err: fmt.Errorf("ent: validator failed for field \"facebook\": %w", err)}
		}
	}

	var (
		err  error
		node *FixRoom
	)
	if len(fruo.hooks) == 0 {
		node, err = fruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FixRoomMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			fruo.mutation = mutation
			node, err = fruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(fruo.hooks) - 1; i >= 0; i-- {
			mut = fruo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fruo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (fruo *FixRoomUpdateOne) SaveX(ctx context.Context) *FixRoom {
	fr, err := fruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return fr
}

// Exec executes the query on the entity.
func (fruo *FixRoomUpdateOne) Exec(ctx context.Context) error {
	_, err := fruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fruo *FixRoomUpdateOne) ExecX(ctx context.Context) {
	if err := fruo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (fruo *FixRoomUpdateOne) sqlSave(ctx context.Context) (fr *FixRoom, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   fixroom.Table,
			Columns: fixroom.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: fixroom.FieldID,
			},
		},
	}
	id, ok := fruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing FixRoom.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := fruo.mutation.FixDetail(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: fixroom.FieldFixDetail,
		})
	}
	if value, ok := fruo.mutation.PhoneNumber(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: fixroom.FieldPhoneNumber,
		})
	}
	if value, ok := fruo.mutation.Facebook(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: fixroom.FieldFacebook,
		})
	}
	if fruo.mutation.CustomerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   fixroom.CustomerTable,
			Columns: []string{fixroom.CustomerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: customer.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fruo.mutation.CustomerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   fixroom.CustomerTable,
			Columns: []string{fixroom.CustomerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: customer.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if fruo.mutation.FurnitureDetailCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   fixroom.FurnitureDetailTable,
			Columns: []string{fixroom.FurnitureDetailColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: furnituredetail.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fruo.mutation.FurnitureDetailIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   fixroom.FurnitureDetailTable,
			Columns: []string{fixroom.FurnitureDetailColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if fruo.mutation.RoomCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   fixroom.RoomTable,
			Columns: []string{fixroom.RoomColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: dataroom.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fruo.mutation.RoomIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   fixroom.RoomTable,
			Columns: []string{fixroom.RoomColumn},
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
	fr = &FixRoom{config: fruo.config}
	_spec.Assign = fr.assignValues
	_spec.ScanValues = fr.scanValues()
	if err = sqlgraph.UpdateNode(ctx, fruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{fixroom.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return fr, nil
}
