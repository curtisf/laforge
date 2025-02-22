// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gen0cide/laforge/ent/build"
	"github.com/gen0cide/laforge/ent/network"
	"github.com/gen0cide/laforge/ent/plan"
	"github.com/gen0cide/laforge/ent/predicate"
	"github.com/gen0cide/laforge/ent/provisionedhost"
	"github.com/gen0cide/laforge/ent/provisionednetwork"
	"github.com/gen0cide/laforge/ent/status"
	"github.com/gen0cide/laforge/ent/team"
	"github.com/google/uuid"
)

// ProvisionedNetworkUpdate is the builder for updating ProvisionedNetwork entities.
type ProvisionedNetworkUpdate struct {
	config
	hooks    []Hook
	mutation *ProvisionedNetworkMutation
}

// Where appends a list predicates to the ProvisionedNetworkUpdate builder.
func (pnu *ProvisionedNetworkUpdate) Where(ps ...predicate.ProvisionedNetwork) *ProvisionedNetworkUpdate {
	pnu.mutation.Where(ps...)
	return pnu
}

// SetName sets the "name" field.
func (pnu *ProvisionedNetworkUpdate) SetName(s string) *ProvisionedNetworkUpdate {
	pnu.mutation.SetName(s)
	return pnu
}

// SetCidr sets the "cidr" field.
func (pnu *ProvisionedNetworkUpdate) SetCidr(s string) *ProvisionedNetworkUpdate {
	pnu.mutation.SetCidr(s)
	return pnu
}

// SetProvisionedNetworkToStatusID sets the "ProvisionedNetworkToStatus" edge to the Status entity by ID.
func (pnu *ProvisionedNetworkUpdate) SetProvisionedNetworkToStatusID(id uuid.UUID) *ProvisionedNetworkUpdate {
	pnu.mutation.SetProvisionedNetworkToStatusID(id)
	return pnu
}

// SetNillableProvisionedNetworkToStatusID sets the "ProvisionedNetworkToStatus" edge to the Status entity by ID if the given value is not nil.
func (pnu *ProvisionedNetworkUpdate) SetNillableProvisionedNetworkToStatusID(id *uuid.UUID) *ProvisionedNetworkUpdate {
	if id != nil {
		pnu = pnu.SetProvisionedNetworkToStatusID(*id)
	}
	return pnu
}

// SetProvisionedNetworkToStatus sets the "ProvisionedNetworkToStatus" edge to the Status entity.
func (pnu *ProvisionedNetworkUpdate) SetProvisionedNetworkToStatus(s *Status) *ProvisionedNetworkUpdate {
	return pnu.SetProvisionedNetworkToStatusID(s.ID)
}

// SetProvisionedNetworkToNetworkID sets the "ProvisionedNetworkToNetwork" edge to the Network entity by ID.
func (pnu *ProvisionedNetworkUpdate) SetProvisionedNetworkToNetworkID(id uuid.UUID) *ProvisionedNetworkUpdate {
	pnu.mutation.SetProvisionedNetworkToNetworkID(id)
	return pnu
}

// SetNillableProvisionedNetworkToNetworkID sets the "ProvisionedNetworkToNetwork" edge to the Network entity by ID if the given value is not nil.
func (pnu *ProvisionedNetworkUpdate) SetNillableProvisionedNetworkToNetworkID(id *uuid.UUID) *ProvisionedNetworkUpdate {
	if id != nil {
		pnu = pnu.SetProvisionedNetworkToNetworkID(*id)
	}
	return pnu
}

// SetProvisionedNetworkToNetwork sets the "ProvisionedNetworkToNetwork" edge to the Network entity.
func (pnu *ProvisionedNetworkUpdate) SetProvisionedNetworkToNetwork(n *Network) *ProvisionedNetworkUpdate {
	return pnu.SetProvisionedNetworkToNetworkID(n.ID)
}

// SetProvisionedNetworkToBuildID sets the "ProvisionedNetworkToBuild" edge to the Build entity by ID.
func (pnu *ProvisionedNetworkUpdate) SetProvisionedNetworkToBuildID(id uuid.UUID) *ProvisionedNetworkUpdate {
	pnu.mutation.SetProvisionedNetworkToBuildID(id)
	return pnu
}

// SetNillableProvisionedNetworkToBuildID sets the "ProvisionedNetworkToBuild" edge to the Build entity by ID if the given value is not nil.
func (pnu *ProvisionedNetworkUpdate) SetNillableProvisionedNetworkToBuildID(id *uuid.UUID) *ProvisionedNetworkUpdate {
	if id != nil {
		pnu = pnu.SetProvisionedNetworkToBuildID(*id)
	}
	return pnu
}

// SetProvisionedNetworkToBuild sets the "ProvisionedNetworkToBuild" edge to the Build entity.
func (pnu *ProvisionedNetworkUpdate) SetProvisionedNetworkToBuild(b *Build) *ProvisionedNetworkUpdate {
	return pnu.SetProvisionedNetworkToBuildID(b.ID)
}

// SetProvisionedNetworkToTeamID sets the "ProvisionedNetworkToTeam" edge to the Team entity by ID.
func (pnu *ProvisionedNetworkUpdate) SetProvisionedNetworkToTeamID(id uuid.UUID) *ProvisionedNetworkUpdate {
	pnu.mutation.SetProvisionedNetworkToTeamID(id)
	return pnu
}

// SetNillableProvisionedNetworkToTeamID sets the "ProvisionedNetworkToTeam" edge to the Team entity by ID if the given value is not nil.
func (pnu *ProvisionedNetworkUpdate) SetNillableProvisionedNetworkToTeamID(id *uuid.UUID) *ProvisionedNetworkUpdate {
	if id != nil {
		pnu = pnu.SetProvisionedNetworkToTeamID(*id)
	}
	return pnu
}

// SetProvisionedNetworkToTeam sets the "ProvisionedNetworkToTeam" edge to the Team entity.
func (pnu *ProvisionedNetworkUpdate) SetProvisionedNetworkToTeam(t *Team) *ProvisionedNetworkUpdate {
	return pnu.SetProvisionedNetworkToTeamID(t.ID)
}

// AddProvisionedNetworkToProvisionedHostIDs adds the "ProvisionedNetworkToProvisionedHost" edge to the ProvisionedHost entity by IDs.
func (pnu *ProvisionedNetworkUpdate) AddProvisionedNetworkToProvisionedHostIDs(ids ...uuid.UUID) *ProvisionedNetworkUpdate {
	pnu.mutation.AddProvisionedNetworkToProvisionedHostIDs(ids...)
	return pnu
}

// AddProvisionedNetworkToProvisionedHost adds the "ProvisionedNetworkToProvisionedHost" edges to the ProvisionedHost entity.
func (pnu *ProvisionedNetworkUpdate) AddProvisionedNetworkToProvisionedHost(p ...*ProvisionedHost) *ProvisionedNetworkUpdate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pnu.AddProvisionedNetworkToProvisionedHostIDs(ids...)
}

// SetProvisionedNetworkToPlanID sets the "ProvisionedNetworkToPlan" edge to the Plan entity by ID.
func (pnu *ProvisionedNetworkUpdate) SetProvisionedNetworkToPlanID(id uuid.UUID) *ProvisionedNetworkUpdate {
	pnu.mutation.SetProvisionedNetworkToPlanID(id)
	return pnu
}

// SetNillableProvisionedNetworkToPlanID sets the "ProvisionedNetworkToPlan" edge to the Plan entity by ID if the given value is not nil.
func (pnu *ProvisionedNetworkUpdate) SetNillableProvisionedNetworkToPlanID(id *uuid.UUID) *ProvisionedNetworkUpdate {
	if id != nil {
		pnu = pnu.SetProvisionedNetworkToPlanID(*id)
	}
	return pnu
}

// SetProvisionedNetworkToPlan sets the "ProvisionedNetworkToPlan" edge to the Plan entity.
func (pnu *ProvisionedNetworkUpdate) SetProvisionedNetworkToPlan(p *Plan) *ProvisionedNetworkUpdate {
	return pnu.SetProvisionedNetworkToPlanID(p.ID)
}

// Mutation returns the ProvisionedNetworkMutation object of the builder.
func (pnu *ProvisionedNetworkUpdate) Mutation() *ProvisionedNetworkMutation {
	return pnu.mutation
}

// ClearProvisionedNetworkToStatus clears the "ProvisionedNetworkToStatus" edge to the Status entity.
func (pnu *ProvisionedNetworkUpdate) ClearProvisionedNetworkToStatus() *ProvisionedNetworkUpdate {
	pnu.mutation.ClearProvisionedNetworkToStatus()
	return pnu
}

// ClearProvisionedNetworkToNetwork clears the "ProvisionedNetworkToNetwork" edge to the Network entity.
func (pnu *ProvisionedNetworkUpdate) ClearProvisionedNetworkToNetwork() *ProvisionedNetworkUpdate {
	pnu.mutation.ClearProvisionedNetworkToNetwork()
	return pnu
}

// ClearProvisionedNetworkToBuild clears the "ProvisionedNetworkToBuild" edge to the Build entity.
func (pnu *ProvisionedNetworkUpdate) ClearProvisionedNetworkToBuild() *ProvisionedNetworkUpdate {
	pnu.mutation.ClearProvisionedNetworkToBuild()
	return pnu
}

// ClearProvisionedNetworkToTeam clears the "ProvisionedNetworkToTeam" edge to the Team entity.
func (pnu *ProvisionedNetworkUpdate) ClearProvisionedNetworkToTeam() *ProvisionedNetworkUpdate {
	pnu.mutation.ClearProvisionedNetworkToTeam()
	return pnu
}

// ClearProvisionedNetworkToProvisionedHost clears all "ProvisionedNetworkToProvisionedHost" edges to the ProvisionedHost entity.
func (pnu *ProvisionedNetworkUpdate) ClearProvisionedNetworkToProvisionedHost() *ProvisionedNetworkUpdate {
	pnu.mutation.ClearProvisionedNetworkToProvisionedHost()
	return pnu
}

// RemoveProvisionedNetworkToProvisionedHostIDs removes the "ProvisionedNetworkToProvisionedHost" edge to ProvisionedHost entities by IDs.
func (pnu *ProvisionedNetworkUpdate) RemoveProvisionedNetworkToProvisionedHostIDs(ids ...uuid.UUID) *ProvisionedNetworkUpdate {
	pnu.mutation.RemoveProvisionedNetworkToProvisionedHostIDs(ids...)
	return pnu
}

// RemoveProvisionedNetworkToProvisionedHost removes "ProvisionedNetworkToProvisionedHost" edges to ProvisionedHost entities.
func (pnu *ProvisionedNetworkUpdate) RemoveProvisionedNetworkToProvisionedHost(p ...*ProvisionedHost) *ProvisionedNetworkUpdate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pnu.RemoveProvisionedNetworkToProvisionedHostIDs(ids...)
}

// ClearProvisionedNetworkToPlan clears the "ProvisionedNetworkToPlan" edge to the Plan entity.
func (pnu *ProvisionedNetworkUpdate) ClearProvisionedNetworkToPlan() *ProvisionedNetworkUpdate {
	pnu.mutation.ClearProvisionedNetworkToPlan()
	return pnu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pnu *ProvisionedNetworkUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(pnu.hooks) == 0 {
		affected, err = pnu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProvisionedNetworkMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pnu.mutation = mutation
			affected, err = pnu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pnu.hooks) - 1; i >= 0; i-- {
			if pnu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pnu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pnu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pnu *ProvisionedNetworkUpdate) SaveX(ctx context.Context) int {
	affected, err := pnu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pnu *ProvisionedNetworkUpdate) Exec(ctx context.Context) error {
	_, err := pnu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pnu *ProvisionedNetworkUpdate) ExecX(ctx context.Context) {
	if err := pnu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pnu *ProvisionedNetworkUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   provisionednetwork.Table,
			Columns: provisionednetwork.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: provisionednetwork.FieldID,
			},
		},
	}
	if ps := pnu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pnu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: provisionednetwork.FieldName,
		})
	}
	if value, ok := pnu.mutation.Cidr(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: provisionednetwork.FieldCidr,
		})
	}
	if pnu.mutation.ProvisionedNetworkToStatusCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   provisionednetwork.ProvisionedNetworkToStatusTable,
			Columns: []string{provisionednetwork.ProvisionedNetworkToStatusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: status.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pnu.mutation.ProvisionedNetworkToStatusIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   provisionednetwork.ProvisionedNetworkToStatusTable,
			Columns: []string{provisionednetwork.ProvisionedNetworkToStatusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: status.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pnu.mutation.ProvisionedNetworkToNetworkCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   provisionednetwork.ProvisionedNetworkToNetworkTable,
			Columns: []string{provisionednetwork.ProvisionedNetworkToNetworkColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: network.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pnu.mutation.ProvisionedNetworkToNetworkIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   provisionednetwork.ProvisionedNetworkToNetworkTable,
			Columns: []string{provisionednetwork.ProvisionedNetworkToNetworkColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: network.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pnu.mutation.ProvisionedNetworkToBuildCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   provisionednetwork.ProvisionedNetworkToBuildTable,
			Columns: []string{provisionednetwork.ProvisionedNetworkToBuildColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: build.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pnu.mutation.ProvisionedNetworkToBuildIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   provisionednetwork.ProvisionedNetworkToBuildTable,
			Columns: []string{provisionednetwork.ProvisionedNetworkToBuildColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: build.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pnu.mutation.ProvisionedNetworkToTeamCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   provisionednetwork.ProvisionedNetworkToTeamTable,
			Columns: []string{provisionednetwork.ProvisionedNetworkToTeamColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: team.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pnu.mutation.ProvisionedNetworkToTeamIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   provisionednetwork.ProvisionedNetworkToTeamTable,
			Columns: []string{provisionednetwork.ProvisionedNetworkToTeamColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: team.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pnu.mutation.ProvisionedNetworkToProvisionedHostCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   provisionednetwork.ProvisionedNetworkToProvisionedHostTable,
			Columns: []string{provisionednetwork.ProvisionedNetworkToProvisionedHostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: provisionedhost.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pnu.mutation.RemovedProvisionedNetworkToProvisionedHostIDs(); len(nodes) > 0 && !pnu.mutation.ProvisionedNetworkToProvisionedHostCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   provisionednetwork.ProvisionedNetworkToProvisionedHostTable,
			Columns: []string{provisionednetwork.ProvisionedNetworkToProvisionedHostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: provisionedhost.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pnu.mutation.ProvisionedNetworkToProvisionedHostIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   provisionednetwork.ProvisionedNetworkToProvisionedHostTable,
			Columns: []string{provisionednetwork.ProvisionedNetworkToProvisionedHostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: provisionedhost.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pnu.mutation.ProvisionedNetworkToPlanCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   provisionednetwork.ProvisionedNetworkToPlanTable,
			Columns: []string{provisionednetwork.ProvisionedNetworkToPlanColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: plan.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pnu.mutation.ProvisionedNetworkToPlanIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   provisionednetwork.ProvisionedNetworkToPlanTable,
			Columns: []string{provisionednetwork.ProvisionedNetworkToPlanColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: plan.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pnu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{provisionednetwork.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// ProvisionedNetworkUpdateOne is the builder for updating a single ProvisionedNetwork entity.
type ProvisionedNetworkUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ProvisionedNetworkMutation
}

// SetName sets the "name" field.
func (pnuo *ProvisionedNetworkUpdateOne) SetName(s string) *ProvisionedNetworkUpdateOne {
	pnuo.mutation.SetName(s)
	return pnuo
}

// SetCidr sets the "cidr" field.
func (pnuo *ProvisionedNetworkUpdateOne) SetCidr(s string) *ProvisionedNetworkUpdateOne {
	pnuo.mutation.SetCidr(s)
	return pnuo
}

// SetProvisionedNetworkToStatusID sets the "ProvisionedNetworkToStatus" edge to the Status entity by ID.
func (pnuo *ProvisionedNetworkUpdateOne) SetProvisionedNetworkToStatusID(id uuid.UUID) *ProvisionedNetworkUpdateOne {
	pnuo.mutation.SetProvisionedNetworkToStatusID(id)
	return pnuo
}

// SetNillableProvisionedNetworkToStatusID sets the "ProvisionedNetworkToStatus" edge to the Status entity by ID if the given value is not nil.
func (pnuo *ProvisionedNetworkUpdateOne) SetNillableProvisionedNetworkToStatusID(id *uuid.UUID) *ProvisionedNetworkUpdateOne {
	if id != nil {
		pnuo = pnuo.SetProvisionedNetworkToStatusID(*id)
	}
	return pnuo
}

// SetProvisionedNetworkToStatus sets the "ProvisionedNetworkToStatus" edge to the Status entity.
func (pnuo *ProvisionedNetworkUpdateOne) SetProvisionedNetworkToStatus(s *Status) *ProvisionedNetworkUpdateOne {
	return pnuo.SetProvisionedNetworkToStatusID(s.ID)
}

// SetProvisionedNetworkToNetworkID sets the "ProvisionedNetworkToNetwork" edge to the Network entity by ID.
func (pnuo *ProvisionedNetworkUpdateOne) SetProvisionedNetworkToNetworkID(id uuid.UUID) *ProvisionedNetworkUpdateOne {
	pnuo.mutation.SetProvisionedNetworkToNetworkID(id)
	return pnuo
}

// SetNillableProvisionedNetworkToNetworkID sets the "ProvisionedNetworkToNetwork" edge to the Network entity by ID if the given value is not nil.
func (pnuo *ProvisionedNetworkUpdateOne) SetNillableProvisionedNetworkToNetworkID(id *uuid.UUID) *ProvisionedNetworkUpdateOne {
	if id != nil {
		pnuo = pnuo.SetProvisionedNetworkToNetworkID(*id)
	}
	return pnuo
}

// SetProvisionedNetworkToNetwork sets the "ProvisionedNetworkToNetwork" edge to the Network entity.
func (pnuo *ProvisionedNetworkUpdateOne) SetProvisionedNetworkToNetwork(n *Network) *ProvisionedNetworkUpdateOne {
	return pnuo.SetProvisionedNetworkToNetworkID(n.ID)
}

// SetProvisionedNetworkToBuildID sets the "ProvisionedNetworkToBuild" edge to the Build entity by ID.
func (pnuo *ProvisionedNetworkUpdateOne) SetProvisionedNetworkToBuildID(id uuid.UUID) *ProvisionedNetworkUpdateOne {
	pnuo.mutation.SetProvisionedNetworkToBuildID(id)
	return pnuo
}

// SetNillableProvisionedNetworkToBuildID sets the "ProvisionedNetworkToBuild" edge to the Build entity by ID if the given value is not nil.
func (pnuo *ProvisionedNetworkUpdateOne) SetNillableProvisionedNetworkToBuildID(id *uuid.UUID) *ProvisionedNetworkUpdateOne {
	if id != nil {
		pnuo = pnuo.SetProvisionedNetworkToBuildID(*id)
	}
	return pnuo
}

// SetProvisionedNetworkToBuild sets the "ProvisionedNetworkToBuild" edge to the Build entity.
func (pnuo *ProvisionedNetworkUpdateOne) SetProvisionedNetworkToBuild(b *Build) *ProvisionedNetworkUpdateOne {
	return pnuo.SetProvisionedNetworkToBuildID(b.ID)
}

// SetProvisionedNetworkToTeamID sets the "ProvisionedNetworkToTeam" edge to the Team entity by ID.
func (pnuo *ProvisionedNetworkUpdateOne) SetProvisionedNetworkToTeamID(id uuid.UUID) *ProvisionedNetworkUpdateOne {
	pnuo.mutation.SetProvisionedNetworkToTeamID(id)
	return pnuo
}

// SetNillableProvisionedNetworkToTeamID sets the "ProvisionedNetworkToTeam" edge to the Team entity by ID if the given value is not nil.
func (pnuo *ProvisionedNetworkUpdateOne) SetNillableProvisionedNetworkToTeamID(id *uuid.UUID) *ProvisionedNetworkUpdateOne {
	if id != nil {
		pnuo = pnuo.SetProvisionedNetworkToTeamID(*id)
	}
	return pnuo
}

// SetProvisionedNetworkToTeam sets the "ProvisionedNetworkToTeam" edge to the Team entity.
func (pnuo *ProvisionedNetworkUpdateOne) SetProvisionedNetworkToTeam(t *Team) *ProvisionedNetworkUpdateOne {
	return pnuo.SetProvisionedNetworkToTeamID(t.ID)
}

// AddProvisionedNetworkToProvisionedHostIDs adds the "ProvisionedNetworkToProvisionedHost" edge to the ProvisionedHost entity by IDs.
func (pnuo *ProvisionedNetworkUpdateOne) AddProvisionedNetworkToProvisionedHostIDs(ids ...uuid.UUID) *ProvisionedNetworkUpdateOne {
	pnuo.mutation.AddProvisionedNetworkToProvisionedHostIDs(ids...)
	return pnuo
}

// AddProvisionedNetworkToProvisionedHost adds the "ProvisionedNetworkToProvisionedHost" edges to the ProvisionedHost entity.
func (pnuo *ProvisionedNetworkUpdateOne) AddProvisionedNetworkToProvisionedHost(p ...*ProvisionedHost) *ProvisionedNetworkUpdateOne {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pnuo.AddProvisionedNetworkToProvisionedHostIDs(ids...)
}

// SetProvisionedNetworkToPlanID sets the "ProvisionedNetworkToPlan" edge to the Plan entity by ID.
func (pnuo *ProvisionedNetworkUpdateOne) SetProvisionedNetworkToPlanID(id uuid.UUID) *ProvisionedNetworkUpdateOne {
	pnuo.mutation.SetProvisionedNetworkToPlanID(id)
	return pnuo
}

// SetNillableProvisionedNetworkToPlanID sets the "ProvisionedNetworkToPlan" edge to the Plan entity by ID if the given value is not nil.
func (pnuo *ProvisionedNetworkUpdateOne) SetNillableProvisionedNetworkToPlanID(id *uuid.UUID) *ProvisionedNetworkUpdateOne {
	if id != nil {
		pnuo = pnuo.SetProvisionedNetworkToPlanID(*id)
	}
	return pnuo
}

// SetProvisionedNetworkToPlan sets the "ProvisionedNetworkToPlan" edge to the Plan entity.
func (pnuo *ProvisionedNetworkUpdateOne) SetProvisionedNetworkToPlan(p *Plan) *ProvisionedNetworkUpdateOne {
	return pnuo.SetProvisionedNetworkToPlanID(p.ID)
}

// Mutation returns the ProvisionedNetworkMutation object of the builder.
func (pnuo *ProvisionedNetworkUpdateOne) Mutation() *ProvisionedNetworkMutation {
	return pnuo.mutation
}

// ClearProvisionedNetworkToStatus clears the "ProvisionedNetworkToStatus" edge to the Status entity.
func (pnuo *ProvisionedNetworkUpdateOne) ClearProvisionedNetworkToStatus() *ProvisionedNetworkUpdateOne {
	pnuo.mutation.ClearProvisionedNetworkToStatus()
	return pnuo
}

// ClearProvisionedNetworkToNetwork clears the "ProvisionedNetworkToNetwork" edge to the Network entity.
func (pnuo *ProvisionedNetworkUpdateOne) ClearProvisionedNetworkToNetwork() *ProvisionedNetworkUpdateOne {
	pnuo.mutation.ClearProvisionedNetworkToNetwork()
	return pnuo
}

// ClearProvisionedNetworkToBuild clears the "ProvisionedNetworkToBuild" edge to the Build entity.
func (pnuo *ProvisionedNetworkUpdateOne) ClearProvisionedNetworkToBuild() *ProvisionedNetworkUpdateOne {
	pnuo.mutation.ClearProvisionedNetworkToBuild()
	return pnuo
}

// ClearProvisionedNetworkToTeam clears the "ProvisionedNetworkToTeam" edge to the Team entity.
func (pnuo *ProvisionedNetworkUpdateOne) ClearProvisionedNetworkToTeam() *ProvisionedNetworkUpdateOne {
	pnuo.mutation.ClearProvisionedNetworkToTeam()
	return pnuo
}

// ClearProvisionedNetworkToProvisionedHost clears all "ProvisionedNetworkToProvisionedHost" edges to the ProvisionedHost entity.
func (pnuo *ProvisionedNetworkUpdateOne) ClearProvisionedNetworkToProvisionedHost() *ProvisionedNetworkUpdateOne {
	pnuo.mutation.ClearProvisionedNetworkToProvisionedHost()
	return pnuo
}

// RemoveProvisionedNetworkToProvisionedHostIDs removes the "ProvisionedNetworkToProvisionedHost" edge to ProvisionedHost entities by IDs.
func (pnuo *ProvisionedNetworkUpdateOne) RemoveProvisionedNetworkToProvisionedHostIDs(ids ...uuid.UUID) *ProvisionedNetworkUpdateOne {
	pnuo.mutation.RemoveProvisionedNetworkToProvisionedHostIDs(ids...)
	return pnuo
}

// RemoveProvisionedNetworkToProvisionedHost removes "ProvisionedNetworkToProvisionedHost" edges to ProvisionedHost entities.
func (pnuo *ProvisionedNetworkUpdateOne) RemoveProvisionedNetworkToProvisionedHost(p ...*ProvisionedHost) *ProvisionedNetworkUpdateOne {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pnuo.RemoveProvisionedNetworkToProvisionedHostIDs(ids...)
}

// ClearProvisionedNetworkToPlan clears the "ProvisionedNetworkToPlan" edge to the Plan entity.
func (pnuo *ProvisionedNetworkUpdateOne) ClearProvisionedNetworkToPlan() *ProvisionedNetworkUpdateOne {
	pnuo.mutation.ClearProvisionedNetworkToPlan()
	return pnuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (pnuo *ProvisionedNetworkUpdateOne) Select(field string, fields ...string) *ProvisionedNetworkUpdateOne {
	pnuo.fields = append([]string{field}, fields...)
	return pnuo
}

// Save executes the query and returns the updated ProvisionedNetwork entity.
func (pnuo *ProvisionedNetworkUpdateOne) Save(ctx context.Context) (*ProvisionedNetwork, error) {
	var (
		err  error
		node *ProvisionedNetwork
	)
	if len(pnuo.hooks) == 0 {
		node, err = pnuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProvisionedNetworkMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pnuo.mutation = mutation
			node, err = pnuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(pnuo.hooks) - 1; i >= 0; i-- {
			if pnuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pnuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pnuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (pnuo *ProvisionedNetworkUpdateOne) SaveX(ctx context.Context) *ProvisionedNetwork {
	node, err := pnuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (pnuo *ProvisionedNetworkUpdateOne) Exec(ctx context.Context) error {
	_, err := pnuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pnuo *ProvisionedNetworkUpdateOne) ExecX(ctx context.Context) {
	if err := pnuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pnuo *ProvisionedNetworkUpdateOne) sqlSave(ctx context.Context) (_node *ProvisionedNetwork, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   provisionednetwork.Table,
			Columns: provisionednetwork.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: provisionednetwork.FieldID,
			},
		},
	}
	id, ok := pnuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing ProvisionedNetwork.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := pnuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, provisionednetwork.FieldID)
		for _, f := range fields {
			if !provisionednetwork.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != provisionednetwork.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := pnuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pnuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: provisionednetwork.FieldName,
		})
	}
	if value, ok := pnuo.mutation.Cidr(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: provisionednetwork.FieldCidr,
		})
	}
	if pnuo.mutation.ProvisionedNetworkToStatusCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   provisionednetwork.ProvisionedNetworkToStatusTable,
			Columns: []string{provisionednetwork.ProvisionedNetworkToStatusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: status.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pnuo.mutation.ProvisionedNetworkToStatusIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   provisionednetwork.ProvisionedNetworkToStatusTable,
			Columns: []string{provisionednetwork.ProvisionedNetworkToStatusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: status.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pnuo.mutation.ProvisionedNetworkToNetworkCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   provisionednetwork.ProvisionedNetworkToNetworkTable,
			Columns: []string{provisionednetwork.ProvisionedNetworkToNetworkColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: network.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pnuo.mutation.ProvisionedNetworkToNetworkIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   provisionednetwork.ProvisionedNetworkToNetworkTable,
			Columns: []string{provisionednetwork.ProvisionedNetworkToNetworkColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: network.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pnuo.mutation.ProvisionedNetworkToBuildCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   provisionednetwork.ProvisionedNetworkToBuildTable,
			Columns: []string{provisionednetwork.ProvisionedNetworkToBuildColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: build.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pnuo.mutation.ProvisionedNetworkToBuildIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   provisionednetwork.ProvisionedNetworkToBuildTable,
			Columns: []string{provisionednetwork.ProvisionedNetworkToBuildColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: build.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pnuo.mutation.ProvisionedNetworkToTeamCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   provisionednetwork.ProvisionedNetworkToTeamTable,
			Columns: []string{provisionednetwork.ProvisionedNetworkToTeamColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: team.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pnuo.mutation.ProvisionedNetworkToTeamIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   provisionednetwork.ProvisionedNetworkToTeamTable,
			Columns: []string{provisionednetwork.ProvisionedNetworkToTeamColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: team.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pnuo.mutation.ProvisionedNetworkToProvisionedHostCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   provisionednetwork.ProvisionedNetworkToProvisionedHostTable,
			Columns: []string{provisionednetwork.ProvisionedNetworkToProvisionedHostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: provisionedhost.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pnuo.mutation.RemovedProvisionedNetworkToProvisionedHostIDs(); len(nodes) > 0 && !pnuo.mutation.ProvisionedNetworkToProvisionedHostCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   provisionednetwork.ProvisionedNetworkToProvisionedHostTable,
			Columns: []string{provisionednetwork.ProvisionedNetworkToProvisionedHostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: provisionedhost.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pnuo.mutation.ProvisionedNetworkToProvisionedHostIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   provisionednetwork.ProvisionedNetworkToProvisionedHostTable,
			Columns: []string{provisionednetwork.ProvisionedNetworkToProvisionedHostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: provisionedhost.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pnuo.mutation.ProvisionedNetworkToPlanCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   provisionednetwork.ProvisionedNetworkToPlanTable,
			Columns: []string{provisionednetwork.ProvisionedNetworkToPlanColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: plan.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pnuo.mutation.ProvisionedNetworkToPlanIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   provisionednetwork.ProvisionedNetworkToPlanTable,
			Columns: []string{provisionednetwork.ProvisionedNetworkToPlanColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: plan.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &ProvisionedNetwork{config: pnuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, pnuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{provisionednetwork.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
