// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/yongliu1992/smartpms/app/property/service/internal/data/ent/shop"
)

// ShopCreate is the builder for creating a Shop entity.
type ShopCreate struct {
	config
	mutation *ShopMutation
	hooks    []Hook
}

// SetFloorID sets the "floor_id" field.
func (sc *ShopCreate) SetFloorID(s string) *ShopCreate {
	sc.mutation.SetFloorID(s)
	return sc
}

// SetRoomNum sets the "room_num" field.
func (sc *ShopCreate) SetRoomNum(i int) *ShopCreate {
	sc.mutation.SetRoomNum(i)
	return sc
}

// SetLayer sets the "layer" field.
func (sc *ShopCreate) SetLayer(i int) *ShopCreate {
	sc.mutation.SetLayer(i)
	return sc
}

// SetBuiltUpArea sets the "built_up_area" field.
func (sc *ShopCreate) SetBuiltUpArea(f float32) *ShopCreate {
	sc.mutation.SetBuiltUpArea(f)
	return sc
}

// SetCommunityID sets the "community_id" field.
func (sc *ShopCreate) SetCommunityID(i int) *ShopCreate {
	sc.mutation.SetCommunityID(i)
	return sc
}

// SetFeeRate sets the "fee_rate" field.
func (sc *ShopCreate) SetFeeRate(f float32) *ShopCreate {
	sc.mutation.SetFeeRate(f)
	return sc
}

// SetRoomArea sets the "room_area" field.
func (sc *ShopCreate) SetRoomArea(f float32) *ShopCreate {
	sc.mutation.SetRoomArea(f)
	return sc
}

// SetRent sets the "rent" field.
func (sc *ShopCreate) SetRent(f float32) *ShopCreate {
	sc.mutation.SetRent(f)
	return sc
}

// SetRemark sets the "remark" field.
func (sc *ShopCreate) SetRemark(s string) *ShopCreate {
	sc.mutation.SetRemark(s)
	return sc
}

// SetCreatedAt sets the "created_at" field.
func (sc *ShopCreate) SetCreatedAt(t time.Time) *ShopCreate {
	sc.mutation.SetCreatedAt(t)
	return sc
}

// SetUpdatedAt sets the "updated_at" field.
func (sc *ShopCreate) SetUpdatedAt(t time.Time) *ShopCreate {
	sc.mutation.SetUpdatedAt(t)
	return sc
}

// SetDeletedAt sets the "deleted_at" field.
func (sc *ShopCreate) SetDeletedAt(t time.Time) *ShopCreate {
	sc.mutation.SetDeletedAt(t)
	return sc
}

// SetID sets the "id" field.
func (sc *ShopCreate) SetID(i int) *ShopCreate {
	sc.mutation.SetID(i)
	return sc
}

// Mutation returns the ShopMutation object of the builder.
func (sc *ShopCreate) Mutation() *ShopMutation {
	return sc.mutation
}

// Save creates the Shop in the database.
func (sc *ShopCreate) Save(ctx context.Context) (*Shop, error) {
	var (
		err  error
		node *Shop
	)
	if len(sc.hooks) == 0 {
		if err = sc.check(); err != nil {
			return nil, err
		}
		node, err = sc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ShopMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sc.check(); err != nil {
				return nil, err
			}
			sc.mutation = mutation
			if node, err = sc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(sc.hooks) - 1; i >= 0; i-- {
			if sc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (sc *ShopCreate) SaveX(ctx context.Context) *Shop {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *ShopCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *ShopCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *ShopCreate) check() error {
	if _, ok := sc.mutation.FloorID(); !ok {
		return &ValidationError{Name: "floor_id", err: errors.New(`ent: missing required field "Shop.floor_id"`)}
	}
	if _, ok := sc.mutation.RoomNum(); !ok {
		return &ValidationError{Name: "room_num", err: errors.New(`ent: missing required field "Shop.room_num"`)}
	}
	if _, ok := sc.mutation.Layer(); !ok {
		return &ValidationError{Name: "layer", err: errors.New(`ent: missing required field "Shop.layer"`)}
	}
	if _, ok := sc.mutation.BuiltUpArea(); !ok {
		return &ValidationError{Name: "built_up_area", err: errors.New(`ent: missing required field "Shop.built_up_area"`)}
	}
	if _, ok := sc.mutation.CommunityID(); !ok {
		return &ValidationError{Name: "community_id", err: errors.New(`ent: missing required field "Shop.community_id"`)}
	}
	if _, ok := sc.mutation.FeeRate(); !ok {
		return &ValidationError{Name: "fee_rate", err: errors.New(`ent: missing required field "Shop.fee_rate"`)}
	}
	if _, ok := sc.mutation.RoomArea(); !ok {
		return &ValidationError{Name: "room_area", err: errors.New(`ent: missing required field "Shop.room_area"`)}
	}
	if _, ok := sc.mutation.Rent(); !ok {
		return &ValidationError{Name: "rent", err: errors.New(`ent: missing required field "Shop.rent"`)}
	}
	if _, ok := sc.mutation.Remark(); !ok {
		return &ValidationError{Name: "remark", err: errors.New(`ent: missing required field "Shop.remark"`)}
	}
	if _, ok := sc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Shop.created_at"`)}
	}
	if _, ok := sc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Shop.updated_at"`)}
	}
	if _, ok := sc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "Shop.deleted_at"`)}
	}
	return nil
}

func (sc *ShopCreate) sqlSave(ctx context.Context) (*Shop, error) {
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	return _node, nil
}

func (sc *ShopCreate) createSpec() (*Shop, *sqlgraph.CreateSpec) {
	var (
		_node = &Shop{config: sc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: shop.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: shop.FieldID,
			},
		}
	)
	if id, ok := sc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := sc.mutation.FloorID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: shop.FieldFloorID,
		})
		_node.FloorID = value
	}
	if value, ok := sc.mutation.RoomNum(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: shop.FieldRoomNum,
		})
		_node.RoomNum = value
	}
	if value, ok := sc.mutation.Layer(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: shop.FieldLayer,
		})
		_node.Layer = value
	}
	if value, ok := sc.mutation.BuiltUpArea(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: shop.FieldBuiltUpArea,
		})
		_node.BuiltUpArea = value
	}
	if value, ok := sc.mutation.CommunityID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: shop.FieldCommunityID,
		})
		_node.CommunityID = value
	}
	if value, ok := sc.mutation.FeeRate(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: shop.FieldFeeRate,
		})
		_node.FeeRate = value
	}
	if value, ok := sc.mutation.RoomArea(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: shop.FieldRoomArea,
		})
		_node.RoomArea = value
	}
	if value, ok := sc.mutation.Rent(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: shop.FieldRent,
		})
		_node.Rent = value
	}
	if value, ok := sc.mutation.Remark(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: shop.FieldRemark,
		})
		_node.Remark = value
	}
	if value, ok := sc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: shop.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := sc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: shop.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := sc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: shop.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	return _node, _spec
}

// ShopCreateBulk is the builder for creating many Shop entities in bulk.
type ShopCreateBulk struct {
	config
	builders []*ShopCreate
}

// Save creates the Shop entities in the database.
func (scb *ShopCreateBulk) Save(ctx context.Context) ([]*Shop, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Shop, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ShopMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *ShopCreateBulk) SaveX(ctx context.Context) []*Shop {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *ShopCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *ShopCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}
