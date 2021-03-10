// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gen0cide/laforge/ent/host"
	"github.com/gen0cide/laforge/ent/hostdependency"
	"github.com/gen0cide/laforge/ent/network"
)

// HostDependencyCreate is the builder for creating a HostDependency entity.
type HostDependencyCreate struct {
	config
	mutation *HostDependencyMutation
	hooks    []Hook
}

// SetHostID sets the "host_id" field.
func (hdc *HostDependencyCreate) SetHostID(s string) *HostDependencyCreate {
	hdc.mutation.SetHostID(s)
	return hdc
}

// SetNetworkID sets the "network_id" field.
func (hdc *HostDependencyCreate) SetNetworkID(s string) *HostDependencyCreate {
	hdc.mutation.SetNetworkID(s)
	return hdc
}

// AddHostDependencyToHostIDs adds the "HostDependencyToHost" edge to the Host entity by IDs.
func (hdc *HostDependencyCreate) AddHostDependencyToHostIDs(ids ...int) *HostDependencyCreate {
	hdc.mutation.AddHostDependencyToHostIDs(ids...)
	return hdc
}

// AddHostDependencyToHost adds the "HostDependencyToHost" edges to the Host entity.
func (hdc *HostDependencyCreate) AddHostDependencyToHost(h ...*Host) *HostDependencyCreate {
	ids := make([]int, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return hdc.AddHostDependencyToHostIDs(ids...)
}

// AddHostDependencyToNetworkIDs adds the "HostDependencyToNetwork" edge to the Network entity by IDs.
func (hdc *HostDependencyCreate) AddHostDependencyToNetworkIDs(ids ...int) *HostDependencyCreate {
	hdc.mutation.AddHostDependencyToNetworkIDs(ids...)
	return hdc
}

// AddHostDependencyToNetwork adds the "HostDependencyToNetwork" edges to the Network entity.
func (hdc *HostDependencyCreate) AddHostDependencyToNetwork(n ...*Network) *HostDependencyCreate {
	ids := make([]int, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return hdc.AddHostDependencyToNetworkIDs(ids...)
}

// Mutation returns the HostDependencyMutation object of the builder.
func (hdc *HostDependencyCreate) Mutation() *HostDependencyMutation {
	return hdc.mutation
}

// Save creates the HostDependency in the database.
func (hdc *HostDependencyCreate) Save(ctx context.Context) (*HostDependency, error) {
	var (
		err  error
		node *HostDependency
	)
	if len(hdc.hooks) == 0 {
		if err = hdc.check(); err != nil {
			return nil, err
		}
		node, err = hdc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*HostDependencyMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = hdc.check(); err != nil {
				return nil, err
			}
			hdc.mutation = mutation
			node, err = hdc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(hdc.hooks) - 1; i >= 0; i-- {
			mut = hdc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, hdc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (hdc *HostDependencyCreate) SaveX(ctx context.Context) *HostDependency {
	v, err := hdc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (hdc *HostDependencyCreate) check() error {
	if _, ok := hdc.mutation.HostID(); !ok {
		return &ValidationError{Name: "host_id", err: errors.New("ent: missing required field \"host_id\"")}
	}
	if _, ok := hdc.mutation.NetworkID(); !ok {
		return &ValidationError{Name: "network_id", err: errors.New("ent: missing required field \"network_id\"")}
	}
	return nil
}

func (hdc *HostDependencyCreate) sqlSave(ctx context.Context) (*HostDependency, error) {
	_node, _spec := hdc.createSpec()
	if err := sqlgraph.CreateNode(ctx, hdc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (hdc *HostDependencyCreate) createSpec() (*HostDependency, *sqlgraph.CreateSpec) {
	var (
		_node = &HostDependency{config: hdc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: hostdependency.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: hostdependency.FieldID,
			},
		}
	)
	if value, ok := hdc.mutation.HostID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: hostdependency.FieldHostID,
		})
		_node.HostID = value
	}
	if value, ok := hdc.mutation.NetworkID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: hostdependency.FieldNetworkID,
		})
		_node.NetworkID = value
	}
	if nodes := hdc.mutation.HostDependencyToHostIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   hostdependency.HostDependencyToHostTable,
			Columns: hostdependency.HostDependencyToHostPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: host.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := hdc.mutation.HostDependencyToNetworkIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   hostdependency.HostDependencyToNetworkTable,
			Columns: hostdependency.HostDependencyToNetworkPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: network.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// HostDependencyCreateBulk is the builder for creating many HostDependency entities in bulk.
type HostDependencyCreateBulk struct {
	config
	builders []*HostDependencyCreate
}

// Save creates the HostDependency entities in the database.
func (hdcb *HostDependencyCreateBulk) Save(ctx context.Context) ([]*HostDependency, error) {
	specs := make([]*sqlgraph.CreateSpec, len(hdcb.builders))
	nodes := make([]*HostDependency, len(hdcb.builders))
	mutators := make([]Mutator, len(hdcb.builders))
	for i := range hdcb.builders {
		func(i int, root context.Context) {
			builder := hdcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*HostDependencyMutation)
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
					_, err = mutators[i+1].Mutate(root, hdcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, hdcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
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
		if _, err := mutators[0].Mutate(ctx, hdcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (hdcb *HostDependencyCreateBulk) SaveX(ctx context.Context) []*HostDependency {
	v, err := hdcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
