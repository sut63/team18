// Code generated by entc, DO NOT EDIT.

package hook

import (
	"context"
	"fmt"

	"github.com/team18/app/ent"
)

// The CheckInFunc type is an adapter to allow the use of ordinary
// function as CheckIn mutator.
type CheckInFunc func(context.Context, *ent.CheckInMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f CheckInFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.CheckInMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.CheckInMutation", m)
	}
	return f(ctx, mv)
}

// The CheckoutFunc type is an adapter to allow the use of ordinary
// function as Checkout mutator.
type CheckoutFunc func(context.Context, *ent.CheckoutMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f CheckoutFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.CheckoutMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.CheckoutMutation", m)
	}
	return f(ctx, mv)
}

// The CounterStaffFunc type is an adapter to allow the use of ordinary
// function as CounterStaff mutator.
type CounterStaffFunc func(context.Context, *ent.CounterStaffMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f CounterStaffFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.CounterStaffMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.CounterStaffMutation", m)
	}
	return f(ctx, mv)
}

// The CustomerFunc type is an adapter to allow the use of ordinary
// function as Customer mutator.
type CustomerFunc func(context.Context, *ent.CustomerMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f CustomerFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.CustomerMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.CustomerMutation", m)
	}
	return f(ctx, mv)
}

// The DataRoomFunc type is an adapter to allow the use of ordinary
// function as DataRoom mutator.
type DataRoomFunc func(context.Context, *ent.DataRoomMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f DataRoomFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.DataRoomMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.DataRoomMutation", m)
	}
	return f(ctx, mv)
}

// The FixRoomFunc type is an adapter to allow the use of ordinary
// function as FixRoom mutator.
type FixRoomFunc func(context.Context, *ent.FixRoomMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f FixRoomFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.FixRoomMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.FixRoomMutation", m)
	}
	return f(ctx, mv)
}

// The FurnitureFunc type is an adapter to allow the use of ordinary
// function as Furniture mutator.
type FurnitureFunc func(context.Context, *ent.FurnitureMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f FurnitureFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.FurnitureMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.FurnitureMutation", m)
	}
	return f(ctx, mv)
}

// The FurnitureDetailFunc type is an adapter to allow the use of ordinary
// function as FurnitureDetail mutator.
type FurnitureDetailFunc func(context.Context, *ent.FurnitureDetailMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f FurnitureDetailFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.FurnitureDetailMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.FurnitureDetailMutation", m)
	}
	return f(ctx, mv)
}

// The FurnitureTypeFunc type is an adapter to allow the use of ordinary
// function as FurnitureType mutator.
type FurnitureTypeFunc func(context.Context, *ent.FurnitureTypeMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f FurnitureTypeFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.FurnitureTypeMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.FurnitureTypeMutation", m)
	}
	return f(ctx, mv)
}

// The PromotionFunc type is an adapter to allow the use of ordinary
// function as Promotion mutator.
type PromotionFunc func(context.Context, *ent.PromotionMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f PromotionFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.PromotionMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.PromotionMutation", m)
	}
	return f(ctx, mv)
}

// The ReserveRoomFunc type is an adapter to allow the use of ordinary
// function as ReserveRoom mutator.
type ReserveRoomFunc func(context.Context, *ent.ReserveRoomMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ReserveRoomFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.ReserveRoomMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ReserveRoomMutation", m)
	}
	return f(ctx, mv)
}

// The StatusFunc type is an adapter to allow the use of ordinary
// function as Status mutator.
type StatusFunc func(context.Context, *ent.StatusMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f StatusFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.StatusMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.StatusMutation", m)
	}
	return f(ctx, mv)
}

// The StatusReserveFunc type is an adapter to allow the use of ordinary
// function as StatusReserve mutator.
type StatusReserveFunc func(context.Context, *ent.StatusReserveMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f StatusReserveFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.StatusReserveMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.StatusReserveMutation", m)
	}
	return f(ctx, mv)
}

// The StatusRoomFunc type is an adapter to allow the use of ordinary
// function as StatusRoom mutator.
type StatusRoomFunc func(context.Context, *ent.StatusRoomMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f StatusRoomFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.StatusRoomMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.StatusRoomMutation", m)
	}
	return f(ctx, mv)
}

// The TypeRoomFunc type is an adapter to allow the use of ordinary
// function as TypeRoom mutator.
type TypeRoomFunc func(context.Context, *ent.TypeRoomMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f TypeRoomFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.TypeRoomMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.TypeRoomMutation", m)
	}
	return f(ctx, mv)
}

// Condition is a hook condition function.
type Condition func(context.Context, ent.Mutation) bool

// And groups conditions with the AND operator.
func And(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if !first(ctx, m) || !second(ctx, m) {
			return false
		}
		for _, cond := range rest {
			if !cond(ctx, m) {
				return false
			}
		}
		return true
	}
}

// Or groups conditions with the OR operator.
func Or(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if first(ctx, m) || second(ctx, m) {
			return true
		}
		for _, cond := range rest {
			if cond(ctx, m) {
				return true
			}
		}
		return false
	}
}

// Not negates a given condition.
func Not(cond Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		return !cond(ctx, m)
	}
}

// HasOp is a condition testing mutation operation.
func HasOp(op ent.Op) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		return m.Op().Is(op)
	}
}

// HasAddedFields is a condition validating `.AddedField` on fields.
func HasAddedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.AddedField(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.AddedField(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasClearedFields is a condition validating `.FieldCleared` on fields.
func HasClearedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if exists := m.FieldCleared(field); !exists {
			return false
		}
		for _, field := range fields {
			if exists := m.FieldCleared(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasFields is a condition validating `.Field` on fields.
func HasFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.Field(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.Field(field); !exists {
				return false
			}
		}
		return true
	}
}

// If executes the given hook under condition.
//
//	Hook.If(ComputeAverage, And(HasFields(...), HasAddedFields(...)))
//
func If(hk ent.Hook, cond Condition) ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if cond(ctx, m) {
				return hk(next).Mutate(ctx, m)
			}
			return next.Mutate(ctx, m)
		})
	}
}

// On executes the given hook only for the given operation.
//
//	hook.On(Log, ent.Delete|ent.Create)
//
func On(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, HasOp(op))
}

// Unless skips the given hook only for the given operation.
//
//	hook.Unless(Log, ent.Update|ent.UpdateOne)
//
func Unless(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, Not(HasOp(op)))
}

// Reject returns a hook that rejects all operations that match op.
//
//	func (T) Hooks() []ent.Hook {
//		return []ent.Hook{
//			Reject(ent.Delete|ent.Update),
//		}
//	}
//
func Reject(op ent.Op) ent.Hook {
	hk := func(ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(_ context.Context, m ent.Mutation) (ent.Value, error) {
			return nil, fmt.Errorf("%s operation is not allowed", m.Op())
		})
	}
	return On(hk, op)
}

// Chain acts as a list of hooks and is effectively immutable.
// Once created, it will always hold the same set of hooks in the same order.
type Chain struct {
	hooks []ent.Hook
}

// NewChain creates a new chain of hooks.
func NewChain(hooks ...ent.Hook) Chain {
	return Chain{append([]ent.Hook(nil), hooks...)}
}

// Hook chains the list of hooks and returns the final hook.
func (c Chain) Hook() ent.Hook {
	return func(mutator ent.Mutator) ent.Mutator {
		for i := len(c.hooks) - 1; i >= 0; i-- {
			mutator = c.hooks[i](mutator)
		}
		return mutator
	}
}

// Append extends a chain, adding the specified hook
// as the last ones in the mutation flow.
func (c Chain) Append(hooks ...ent.Hook) Chain {
	newHooks := make([]ent.Hook, 0, len(c.hooks)+len(hooks))
	newHooks = append(newHooks, c.hooks...)
	newHooks = append(newHooks, hooks...)
	return Chain{newHooks}
}

// Extend extends a chain, adding the specified chain
// as the last ones in the mutation flow.
func (c Chain) Extend(chain Chain) Chain {
	return c.Append(chain.hooks...)
}
