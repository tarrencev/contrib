// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/contrib/entproto/internal/entprototest/ent/messagewithenum"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MessageWithEnumCreate is the builder for creating a MessageWithEnum entity.
type MessageWithEnumCreate struct {
	config
	mutation *MessageWithEnumMutation
	hooks    []Hook
}

// SetEnumType sets the "enum_type" field.
func (mwec *MessageWithEnumCreate) SetEnumType(mt messagewithenum.EnumType) *MessageWithEnumCreate {
	mwec.mutation.SetEnumType(mt)
	return mwec
}

// SetNillableEnumType sets the "enum_type" field if the given value is not nil.
func (mwec *MessageWithEnumCreate) SetNillableEnumType(mt *messagewithenum.EnumType) *MessageWithEnumCreate {
	if mt != nil {
		mwec.SetEnumType(*mt)
	}
	return mwec
}

// SetEnumWithoutDefault sets the "enum_without_default" field.
func (mwec *MessageWithEnumCreate) SetEnumWithoutDefault(mwd messagewithenum.EnumWithoutDefault) *MessageWithEnumCreate {
	mwec.mutation.SetEnumWithoutDefault(mwd)
	return mwec
}

// Mutation returns the MessageWithEnumMutation object of the builder.
func (mwec *MessageWithEnumCreate) Mutation() *MessageWithEnumMutation {
	return mwec.mutation
}

// Save creates the MessageWithEnum in the database.
func (mwec *MessageWithEnumCreate) Save(ctx context.Context) (*MessageWithEnum, error) {
	var (
		err  error
		node *MessageWithEnum
	)
	mwec.defaults()
	if len(mwec.hooks) == 0 {
		if err = mwec.check(); err != nil {
			return nil, err
		}
		node, err = mwec.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MessageWithEnumMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = mwec.check(); err != nil {
				return nil, err
			}
			mwec.mutation = mutation
			if node, err = mwec.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(mwec.hooks) - 1; i >= 0; i-- {
			mut = mwec.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mwec.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (mwec *MessageWithEnumCreate) SaveX(ctx context.Context) *MessageWithEnum {
	v, err := mwec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (mwec *MessageWithEnumCreate) defaults() {
	if _, ok := mwec.mutation.EnumType(); !ok {
		v := messagewithenum.DefaultEnumType
		mwec.mutation.SetEnumType(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mwec *MessageWithEnumCreate) check() error {
	if _, ok := mwec.mutation.EnumType(); !ok {
		return &ValidationError{Name: "enum_type", err: errors.New("ent: missing required field \"enum_type\"")}
	}
	if v, ok := mwec.mutation.EnumType(); ok {
		if err := messagewithenum.EnumTypeValidator(v); err != nil {
			return &ValidationError{Name: "enum_type", err: fmt.Errorf("ent: validator failed for field \"enum_type\": %w", err)}
		}
	}
	if _, ok := mwec.mutation.EnumWithoutDefault(); !ok {
		return &ValidationError{Name: "enum_without_default", err: errors.New("ent: missing required field \"enum_without_default\"")}
	}
	if v, ok := mwec.mutation.EnumWithoutDefault(); ok {
		if err := messagewithenum.EnumWithoutDefaultValidator(v); err != nil {
			return &ValidationError{Name: "enum_without_default", err: fmt.Errorf("ent: validator failed for field \"enum_without_default\": %w", err)}
		}
	}
	return nil
}

func (mwec *MessageWithEnumCreate) sqlSave(ctx context.Context) (*MessageWithEnum, error) {
	_node, _spec := mwec.createSpec()
	if err := sqlgraph.CreateNode(ctx, mwec.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (mwec *MessageWithEnumCreate) createSpec() (*MessageWithEnum, *sqlgraph.CreateSpec) {
	var (
		_node = &MessageWithEnum{config: mwec.config}
		_spec = &sqlgraph.CreateSpec{
			Table: messagewithenum.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: messagewithenum.FieldID,
			},
		}
	)
	if value, ok := mwec.mutation.EnumType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: messagewithenum.FieldEnumType,
		})
		_node.EnumType = value
	}
	if value, ok := mwec.mutation.EnumWithoutDefault(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: messagewithenum.FieldEnumWithoutDefault,
		})
		_node.EnumWithoutDefault = value
	}
	return _node, _spec
}

// MessageWithEnumCreateBulk is the builder for creating many MessageWithEnum entities in bulk.
type MessageWithEnumCreateBulk struct {
	config
	builders []*MessageWithEnumCreate
}

// Save creates the MessageWithEnum entities in the database.
func (mwecb *MessageWithEnumCreateBulk) Save(ctx context.Context) ([]*MessageWithEnum, error) {
	specs := make([]*sqlgraph.CreateSpec, len(mwecb.builders))
	nodes := make([]*MessageWithEnum, len(mwecb.builders))
	mutators := make([]Mutator, len(mwecb.builders))
	for i := range mwecb.builders {
		func(i int, root context.Context) {
			builder := mwecb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MessageWithEnumMutation)
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
					_, err = mutators[i+1].Mutate(root, mwecb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mwecb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, mwecb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mwecb *MessageWithEnumCreateBulk) SaveX(ctx context.Context) []*MessageWithEnum {
	v, err := mwecb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
