// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/yongliu1992/smartpms/app/property/service/internal/data/ent/predicate"
	"github.com/yongliu1992/smartpms/app/property/service/internal/data/ent/shop"
)

// ShopUpdate is the builder for updating Shop entities.
type ShopUpdate struct {
	config
	hooks    []Hook
	mutation *ShopMutation
}

// Where appends a list predicates to the ShopUpdate builder.
func (su *ShopUpdate) Where(ps ...predicate.Shop) *ShopUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetFloorID sets the "floor_id" field.
func (su *ShopUpdate) SetFloorID(s string) *ShopUpdate {
	su.mutation.SetFloorID(s)
	return su
}

// SetRoomNum sets the "room_num" field.
func (su *ShopUpdate) SetRoomNum(i int) *ShopUpdate {
	su.mutation.ResetRoomNum()
	su.mutation.SetRoomNum(i)
	return su
}

// AddRoomNum adds i to the "room_num" field.
func (su *ShopUpdate) AddRoomNum(i int) *ShopUpdate {
	su.mutation.AddRoomNum(i)
	return su
}

// SetLayer sets the "layer" field.
func (su *ShopUpdate) SetLayer(i int) *ShopUpdate {
	su.mutation.ResetLayer()
	su.mutation.SetLayer(i)
	return su
}

// AddLayer adds i to the "layer" field.
func (su *ShopUpdate) AddLayer(i int) *ShopUpdate {
	su.mutation.AddLayer(i)
	return su
}

// SetBuiltUpArea sets the "built_up_area" field.
func (su *ShopUpdate) SetBuiltUpArea(f float32) *ShopUpdate {
	su.mutation.ResetBuiltUpArea()
	su.mutation.SetBuiltUpArea(f)
	return su
}

// AddBuiltUpArea adds f to the "built_up_area" field.
func (su *ShopUpdate) AddBuiltUpArea(f float32) *ShopUpdate {
	su.mutation.AddBuiltUpArea(f)
	return su
}

// SetCommunityID sets the "community_id" field.
func (su *ShopUpdate) SetCommunityID(i int) *ShopUpdate {
	su.mutation.ResetCommunityID()
	su.mutation.SetCommunityID(i)
	return su
}

// AddCommunityID adds i to the "community_id" field.
func (su *ShopUpdate) AddCommunityID(i int) *ShopUpdate {
	su.mutation.AddCommunityID(i)
	return su
}

// SetFeeRate sets the "fee_rate" field.
func (su *ShopUpdate) SetFeeRate(f float32) *ShopUpdate {
	su.mutation.ResetFeeRate()
	su.mutation.SetFeeRate(f)
	return su
}

// AddFeeRate adds f to the "fee_rate" field.
func (su *ShopUpdate) AddFeeRate(f float32) *ShopUpdate {
	su.mutation.AddFeeRate(f)
	return su
}

// SetRoomArea sets the "room_area" field.
func (su *ShopUpdate) SetRoomArea(f float32) *ShopUpdate {
	su.mutation.ResetRoomArea()
	su.mutation.SetRoomArea(f)
	return su
}

// AddRoomArea adds f to the "room_area" field.
func (su *ShopUpdate) AddRoomArea(f float32) *ShopUpdate {
	su.mutation.AddRoomArea(f)
	return su
}

// SetRent sets the "rent" field.
func (su *ShopUpdate) SetRent(f float32) *ShopUpdate {
	su.mutation.ResetRent()
	su.mutation.SetRent(f)
	return su
}

// AddRent adds f to the "rent" field.
func (su *ShopUpdate) AddRent(f float32) *ShopUpdate {
	su.mutation.AddRent(f)
	return su
}

// SetRemark sets the "remark" field.
func (su *ShopUpdate) SetRemark(s string) *ShopUpdate {
	su.mutation.SetRemark(s)
	return su
}

// SetCreatedAt sets the "created_at" field.
func (su *ShopUpdate) SetCreatedAt(t time.Time) *ShopUpdate {
	su.mutation.SetCreatedAt(t)
	return su
}

// SetUpdatedAt sets the "updated_at" field.
func (su *ShopUpdate) SetUpdatedAt(t time.Time) *ShopUpdate {
	su.mutation.SetUpdatedAt(t)
	return su
}

// SetDeletedAt sets the "deleted_at" field.
func (su *ShopUpdate) SetDeletedAt(t time.Time) *ShopUpdate {
	su.mutation.SetDeletedAt(t)
	return su
}

// Mutation returns the ShopMutation object of the builder.
func (su *ShopUpdate) Mutation() *ShopMutation {
	return su.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *ShopUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(su.hooks) == 0 {
		affected, err = su.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ShopMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			su.mutation = mutation
			affected, err = su.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(su.hooks) - 1; i >= 0; i-- {
			if su.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = su.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, su.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (su *ShopUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *ShopUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *ShopUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

func (su *ShopUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   shop.Table,
			Columns: shop.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: shop.FieldID,
			},
		},
	}
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.FloorID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: shop.FieldFloorID,
		})
	}
	if value, ok := su.mutation.RoomNum(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: shop.FieldRoomNum,
		})
	}
	if value, ok := su.mutation.AddedRoomNum(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: shop.FieldRoomNum,
		})
	}
	if value, ok := su.mutation.Layer(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: shop.FieldLayer,
		})
	}
	if value, ok := su.mutation.AddedLayer(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: shop.FieldLayer,
		})
	}
	if value, ok := su.mutation.BuiltUpArea(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: shop.FieldBuiltUpArea,
		})
	}
	if value, ok := su.mutation.AddedBuiltUpArea(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: shop.FieldBuiltUpArea,
		})
	}
	if value, ok := su.mutation.CommunityID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: shop.FieldCommunityID,
		})
	}
	if value, ok := su.mutation.AddedCommunityID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: shop.FieldCommunityID,
		})
	}
	if value, ok := su.mutation.FeeRate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: shop.FieldFeeRate,
		})
	}
	if value, ok := su.mutation.AddedFeeRate(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: shop.FieldFeeRate,
		})
	}
	if value, ok := su.mutation.RoomArea(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: shop.FieldRoomArea,
		})
	}
	if value, ok := su.mutation.AddedRoomArea(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: shop.FieldRoomArea,
		})
	}
	if value, ok := su.mutation.Rent(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: shop.FieldRent,
		})
	}
	if value, ok := su.mutation.AddedRent(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: shop.FieldRent,
		})
	}
	if value, ok := su.mutation.Remark(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: shop.FieldRemark,
		})
	}
	if value, ok := su.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: shop.FieldCreatedAt,
		})
	}
	if value, ok := su.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: shop.FieldUpdatedAt,
		})
	}
	if value, ok := su.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: shop.FieldDeletedAt,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{shop.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// ShopUpdateOne is the builder for updating a single Shop entity.
type ShopUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ShopMutation
}

// SetFloorID sets the "floor_id" field.
func (suo *ShopUpdateOne) SetFloorID(s string) *ShopUpdateOne {
	suo.mutation.SetFloorID(s)
	return suo
}

// SetRoomNum sets the "room_num" field.
func (suo *ShopUpdateOne) SetRoomNum(i int) *ShopUpdateOne {
	suo.mutation.ResetRoomNum()
	suo.mutation.SetRoomNum(i)
	return suo
}

// AddRoomNum adds i to the "room_num" field.
func (suo *ShopUpdateOne) AddRoomNum(i int) *ShopUpdateOne {
	suo.mutation.AddRoomNum(i)
	return suo
}

// SetLayer sets the "layer" field.
func (suo *ShopUpdateOne) SetLayer(i int) *ShopUpdateOne {
	suo.mutation.ResetLayer()
	suo.mutation.SetLayer(i)
	return suo
}

// AddLayer adds i to the "layer" field.
func (suo *ShopUpdateOne) AddLayer(i int) *ShopUpdateOne {
	suo.mutation.AddLayer(i)
	return suo
}

// SetBuiltUpArea sets the "built_up_area" field.
func (suo *ShopUpdateOne) SetBuiltUpArea(f float32) *ShopUpdateOne {
	suo.mutation.ResetBuiltUpArea()
	suo.mutation.SetBuiltUpArea(f)
	return suo
}

// AddBuiltUpArea adds f to the "built_up_area" field.
func (suo *ShopUpdateOne) AddBuiltUpArea(f float32) *ShopUpdateOne {
	suo.mutation.AddBuiltUpArea(f)
	return suo
}

// SetCommunityID sets the "community_id" field.
func (suo *ShopUpdateOne) SetCommunityID(i int) *ShopUpdateOne {
	suo.mutation.ResetCommunityID()
	suo.mutation.SetCommunityID(i)
	return suo
}

// AddCommunityID adds i to the "community_id" field.
func (suo *ShopUpdateOne) AddCommunityID(i int) *ShopUpdateOne {
	suo.mutation.AddCommunityID(i)
	return suo
}

// SetFeeRate sets the "fee_rate" field.
func (suo *ShopUpdateOne) SetFeeRate(f float32) *ShopUpdateOne {
	suo.mutation.ResetFeeRate()
	suo.mutation.SetFeeRate(f)
	return suo
}

// AddFeeRate adds f to the "fee_rate" field.
func (suo *ShopUpdateOne) AddFeeRate(f float32) *ShopUpdateOne {
	suo.mutation.AddFeeRate(f)
	return suo
}

// SetRoomArea sets the "room_area" field.
func (suo *ShopUpdateOne) SetRoomArea(f float32) *ShopUpdateOne {
	suo.mutation.ResetRoomArea()
	suo.mutation.SetRoomArea(f)
	return suo
}

// AddRoomArea adds f to the "room_area" field.
func (suo *ShopUpdateOne) AddRoomArea(f float32) *ShopUpdateOne {
	suo.mutation.AddRoomArea(f)
	return suo
}

// SetRent sets the "rent" field.
func (suo *ShopUpdateOne) SetRent(f float32) *ShopUpdateOne {
	suo.mutation.ResetRent()
	suo.mutation.SetRent(f)
	return suo
}

// AddRent adds f to the "rent" field.
func (suo *ShopUpdateOne) AddRent(f float32) *ShopUpdateOne {
	suo.mutation.AddRent(f)
	return suo
}

// SetRemark sets the "remark" field.
func (suo *ShopUpdateOne) SetRemark(s string) *ShopUpdateOne {
	suo.mutation.SetRemark(s)
	return suo
}

// SetCreatedAt sets the "created_at" field.
func (suo *ShopUpdateOne) SetCreatedAt(t time.Time) *ShopUpdateOne {
	suo.mutation.SetCreatedAt(t)
	return suo
}

// SetUpdatedAt sets the "updated_at" field.
func (suo *ShopUpdateOne) SetUpdatedAt(t time.Time) *ShopUpdateOne {
	suo.mutation.SetUpdatedAt(t)
	return suo
}

// SetDeletedAt sets the "deleted_at" field.
func (suo *ShopUpdateOne) SetDeletedAt(t time.Time) *ShopUpdateOne {
	suo.mutation.SetDeletedAt(t)
	return suo
}

// Mutation returns the ShopMutation object of the builder.
func (suo *ShopUpdateOne) Mutation() *ShopMutation {
	return suo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *ShopUpdateOne) Select(field string, fields ...string) *ShopUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Shop entity.
func (suo *ShopUpdateOne) Save(ctx context.Context) (*Shop, error) {
	var (
		err  error
		node *Shop
	)
	if len(suo.hooks) == 0 {
		node, err = suo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ShopMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			suo.mutation = mutation
			node, err = suo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(suo.hooks) - 1; i >= 0; i-- {
			if suo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = suo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, suo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (suo *ShopUpdateOne) SaveX(ctx context.Context) *Shop {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *ShopUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *ShopUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (suo *ShopUpdateOne) sqlSave(ctx context.Context) (_node *Shop, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   shop.Table,
			Columns: shop.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: shop.FieldID,
			},
		},
	}
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Shop.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, shop.FieldID)
		for _, f := range fields {
			if !shop.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != shop.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.FloorID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: shop.FieldFloorID,
		})
	}
	if value, ok := suo.mutation.RoomNum(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: shop.FieldRoomNum,
		})
	}
	if value, ok := suo.mutation.AddedRoomNum(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: shop.FieldRoomNum,
		})
	}
	if value, ok := suo.mutation.Layer(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: shop.FieldLayer,
		})
	}
	if value, ok := suo.mutation.AddedLayer(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: shop.FieldLayer,
		})
	}
	if value, ok := suo.mutation.BuiltUpArea(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: shop.FieldBuiltUpArea,
		})
	}
	if value, ok := suo.mutation.AddedBuiltUpArea(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: shop.FieldBuiltUpArea,
		})
	}
	if value, ok := suo.mutation.CommunityID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: shop.FieldCommunityID,
		})
	}
	if value, ok := suo.mutation.AddedCommunityID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: shop.FieldCommunityID,
		})
	}
	if value, ok := suo.mutation.FeeRate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: shop.FieldFeeRate,
		})
	}
	if value, ok := suo.mutation.AddedFeeRate(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: shop.FieldFeeRate,
		})
	}
	if value, ok := suo.mutation.RoomArea(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: shop.FieldRoomArea,
		})
	}
	if value, ok := suo.mutation.AddedRoomArea(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: shop.FieldRoomArea,
		})
	}
	if value, ok := suo.mutation.Rent(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: shop.FieldRent,
		})
	}
	if value, ok := suo.mutation.AddedRent(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: shop.FieldRent,
		})
	}
	if value, ok := suo.mutation.Remark(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: shop.FieldRemark,
		})
	}
	if value, ok := suo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: shop.FieldCreatedAt,
		})
	}
	if value, ok := suo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: shop.FieldUpdatedAt,
		})
	}
	if value, ok := suo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: shop.FieldDeletedAt,
		})
	}
	_node = &Shop{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{shop.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
