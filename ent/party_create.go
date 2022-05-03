// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/KawashiroNitori/lootbot/ent/party"
)

// PartyCreate is the builder for creating a Party entity.
type PartyCreate struct {
	config
	mutation *PartyMutation
	hooks    []Hook
}

// SetChannelID sets the "channel_id" field.
func (pc *PartyCreate) SetChannelID(s string) *PartyCreate {
	pc.mutation.SetChannelID(s)
	return pc
}

// SetID sets the "id" field.
func (pc *PartyCreate) SetID(s string) *PartyCreate {
	pc.mutation.SetID(s)
	return pc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (pc *PartyCreate) SetNillableID(s *string) *PartyCreate {
	if s != nil {
		pc.SetID(*s)
	}
	return pc
}

// Mutation returns the PartyMutation object of the builder.
func (pc *PartyCreate) Mutation() *PartyMutation {
	return pc.mutation
}

// Save creates the Party in the database.
func (pc *PartyCreate) Save(ctx context.Context) (*Party, error) {
	var (
		err  error
		node *Party
	)
	pc.defaults()
	if len(pc.hooks) == 0 {
		if err = pc.check(); err != nil {
			return nil, err
		}
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PartyMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pc.check(); err != nil {
				return nil, err
			}
			pc.mutation = mutation
			if node, err = pc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(pc.hooks) - 1; i >= 0; i-- {
			if pc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PartyCreate) SaveX(ctx context.Context) *Party {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PartyCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PartyCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *PartyCreate) defaults() {
	if _, ok := pc.mutation.ID(); !ok {
		v := party.DefaultID()
		pc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PartyCreate) check() error {
	if _, ok := pc.mutation.ChannelID(); !ok {
		return &ValidationError{Name: "channel_id", err: errors.New(`ent: missing required field "Party.channel_id"`)}
	}
	return nil
}

func (pc *PartyCreate) sqlSave(ctx context.Context) (*Party, error) {
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Party.ID type: %T", _spec.ID.Value)
		}
	}
	return _node, nil
}

func (pc *PartyCreate) createSpec() (*Party, *sqlgraph.CreateSpec) {
	var (
		_node = &Party{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: party.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: party.FieldID,
			},
		}
	)
	if id, ok := pc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := pc.mutation.ChannelID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: party.FieldChannelID,
		})
		_node.ChannelID = value
	}
	return _node, _spec
}

// PartyCreateBulk is the builder for creating many Party entities in bulk.
type PartyCreateBulk struct {
	config
	builders []*PartyCreate
}

// Save creates the Party entities in the database.
func (pcb *PartyCreateBulk) Save(ctx context.Context) ([]*Party, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Party, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PartyMutation)
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
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
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
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PartyCreateBulk) SaveX(ctx context.Context) []*Party {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PartyCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PartyCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
