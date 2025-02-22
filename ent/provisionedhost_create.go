// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gen0cide/laforge/ent/agentstatus"
	"github.com/gen0cide/laforge/ent/agenttask"
	"github.com/gen0cide/laforge/ent/build"
	"github.com/gen0cide/laforge/ent/ginfilemiddleware"
	"github.com/gen0cide/laforge/ent/host"
	"github.com/gen0cide/laforge/ent/plan"
	"github.com/gen0cide/laforge/ent/provisionedhost"
	"github.com/gen0cide/laforge/ent/provisionednetwork"
	"github.com/gen0cide/laforge/ent/provisioningstep"
	"github.com/gen0cide/laforge/ent/status"
	"github.com/google/uuid"
)

// ProvisionedHostCreate is the builder for creating a ProvisionedHost entity.
type ProvisionedHostCreate struct {
	config
	mutation *ProvisionedHostMutation
	hooks    []Hook
}

// SetSubnetIP sets the "subnet_ip" field.
func (phc *ProvisionedHostCreate) SetSubnetIP(s string) *ProvisionedHostCreate {
	phc.mutation.SetSubnetIP(s)
	return phc
}

// SetAddonType sets the "addon_type" field.
func (phc *ProvisionedHostCreate) SetAddonType(pt provisionedhost.AddonType) *ProvisionedHostCreate {
	phc.mutation.SetAddonType(pt)
	return phc
}

// SetNillableAddonType sets the "addon_type" field if the given value is not nil.
func (phc *ProvisionedHostCreate) SetNillableAddonType(pt *provisionedhost.AddonType) *ProvisionedHostCreate {
	if pt != nil {
		phc.SetAddonType(*pt)
	}
	return phc
}

// SetID sets the "id" field.
func (phc *ProvisionedHostCreate) SetID(u uuid.UUID) *ProvisionedHostCreate {
	phc.mutation.SetID(u)
	return phc
}

// SetProvisionedHostToStatusID sets the "ProvisionedHostToStatus" edge to the Status entity by ID.
func (phc *ProvisionedHostCreate) SetProvisionedHostToStatusID(id uuid.UUID) *ProvisionedHostCreate {
	phc.mutation.SetProvisionedHostToStatusID(id)
	return phc
}

// SetProvisionedHostToStatus sets the "ProvisionedHostToStatus" edge to the Status entity.
func (phc *ProvisionedHostCreate) SetProvisionedHostToStatus(s *Status) *ProvisionedHostCreate {
	return phc.SetProvisionedHostToStatusID(s.ID)
}

// SetProvisionedHostToProvisionedNetworkID sets the "ProvisionedHostToProvisionedNetwork" edge to the ProvisionedNetwork entity by ID.
func (phc *ProvisionedHostCreate) SetProvisionedHostToProvisionedNetworkID(id uuid.UUID) *ProvisionedHostCreate {
	phc.mutation.SetProvisionedHostToProvisionedNetworkID(id)
	return phc
}

// SetProvisionedHostToProvisionedNetwork sets the "ProvisionedHostToProvisionedNetwork" edge to the ProvisionedNetwork entity.
func (phc *ProvisionedHostCreate) SetProvisionedHostToProvisionedNetwork(p *ProvisionedNetwork) *ProvisionedHostCreate {
	return phc.SetProvisionedHostToProvisionedNetworkID(p.ID)
}

// SetProvisionedHostToHostID sets the "ProvisionedHostToHost" edge to the Host entity by ID.
func (phc *ProvisionedHostCreate) SetProvisionedHostToHostID(id uuid.UUID) *ProvisionedHostCreate {
	phc.mutation.SetProvisionedHostToHostID(id)
	return phc
}

// SetProvisionedHostToHost sets the "ProvisionedHostToHost" edge to the Host entity.
func (phc *ProvisionedHostCreate) SetProvisionedHostToHost(h *Host) *ProvisionedHostCreate {
	return phc.SetProvisionedHostToHostID(h.ID)
}

// SetProvisionedHostToEndStepPlanID sets the "ProvisionedHostToEndStepPlan" edge to the Plan entity by ID.
func (phc *ProvisionedHostCreate) SetProvisionedHostToEndStepPlanID(id uuid.UUID) *ProvisionedHostCreate {
	phc.mutation.SetProvisionedHostToEndStepPlanID(id)
	return phc
}

// SetNillableProvisionedHostToEndStepPlanID sets the "ProvisionedHostToEndStepPlan" edge to the Plan entity by ID if the given value is not nil.
func (phc *ProvisionedHostCreate) SetNillableProvisionedHostToEndStepPlanID(id *uuid.UUID) *ProvisionedHostCreate {
	if id != nil {
		phc = phc.SetProvisionedHostToEndStepPlanID(*id)
	}
	return phc
}

// SetProvisionedHostToEndStepPlan sets the "ProvisionedHostToEndStepPlan" edge to the Plan entity.
func (phc *ProvisionedHostCreate) SetProvisionedHostToEndStepPlan(p *Plan) *ProvisionedHostCreate {
	return phc.SetProvisionedHostToEndStepPlanID(p.ID)
}

// SetProvisionedHostToBuildID sets the "ProvisionedHostToBuild" edge to the Build entity by ID.
func (phc *ProvisionedHostCreate) SetProvisionedHostToBuildID(id uuid.UUID) *ProvisionedHostCreate {
	phc.mutation.SetProvisionedHostToBuildID(id)
	return phc
}

// SetProvisionedHostToBuild sets the "ProvisionedHostToBuild" edge to the Build entity.
func (phc *ProvisionedHostCreate) SetProvisionedHostToBuild(b *Build) *ProvisionedHostCreate {
	return phc.SetProvisionedHostToBuildID(b.ID)
}

// AddProvisionedHostToProvisioningStepIDs adds the "ProvisionedHostToProvisioningStep" edge to the ProvisioningStep entity by IDs.
func (phc *ProvisionedHostCreate) AddProvisionedHostToProvisioningStepIDs(ids ...uuid.UUID) *ProvisionedHostCreate {
	phc.mutation.AddProvisionedHostToProvisioningStepIDs(ids...)
	return phc
}

// AddProvisionedHostToProvisioningStep adds the "ProvisionedHostToProvisioningStep" edges to the ProvisioningStep entity.
func (phc *ProvisionedHostCreate) AddProvisionedHostToProvisioningStep(p ...*ProvisioningStep) *ProvisionedHostCreate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return phc.AddProvisionedHostToProvisioningStepIDs(ids...)
}

// AddProvisionedHostToAgentStatuIDs adds the "ProvisionedHostToAgentStatus" edge to the AgentStatus entity by IDs.
func (phc *ProvisionedHostCreate) AddProvisionedHostToAgentStatuIDs(ids ...uuid.UUID) *ProvisionedHostCreate {
	phc.mutation.AddProvisionedHostToAgentStatuIDs(ids...)
	return phc
}

// AddProvisionedHostToAgentStatus adds the "ProvisionedHostToAgentStatus" edges to the AgentStatus entity.
func (phc *ProvisionedHostCreate) AddProvisionedHostToAgentStatus(a ...*AgentStatus) *ProvisionedHostCreate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return phc.AddProvisionedHostToAgentStatuIDs(ids...)
}

// AddProvisionedHostToAgentTaskIDs adds the "ProvisionedHostToAgentTask" edge to the AgentTask entity by IDs.
func (phc *ProvisionedHostCreate) AddProvisionedHostToAgentTaskIDs(ids ...uuid.UUID) *ProvisionedHostCreate {
	phc.mutation.AddProvisionedHostToAgentTaskIDs(ids...)
	return phc
}

// AddProvisionedHostToAgentTask adds the "ProvisionedHostToAgentTask" edges to the AgentTask entity.
func (phc *ProvisionedHostCreate) AddProvisionedHostToAgentTask(a ...*AgentTask) *ProvisionedHostCreate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return phc.AddProvisionedHostToAgentTaskIDs(ids...)
}

// SetProvisionedHostToPlanID sets the "ProvisionedHostToPlan" edge to the Plan entity by ID.
func (phc *ProvisionedHostCreate) SetProvisionedHostToPlanID(id uuid.UUID) *ProvisionedHostCreate {
	phc.mutation.SetProvisionedHostToPlanID(id)
	return phc
}

// SetNillableProvisionedHostToPlanID sets the "ProvisionedHostToPlan" edge to the Plan entity by ID if the given value is not nil.
func (phc *ProvisionedHostCreate) SetNillableProvisionedHostToPlanID(id *uuid.UUID) *ProvisionedHostCreate {
	if id != nil {
		phc = phc.SetProvisionedHostToPlanID(*id)
	}
	return phc
}

// SetProvisionedHostToPlan sets the "ProvisionedHostToPlan" edge to the Plan entity.
func (phc *ProvisionedHostCreate) SetProvisionedHostToPlan(p *Plan) *ProvisionedHostCreate {
	return phc.SetProvisionedHostToPlanID(p.ID)
}

// SetProvisionedHostToGinFileMiddlewareID sets the "ProvisionedHostToGinFileMiddleware" edge to the GinFileMiddleware entity by ID.
func (phc *ProvisionedHostCreate) SetProvisionedHostToGinFileMiddlewareID(id uuid.UUID) *ProvisionedHostCreate {
	phc.mutation.SetProvisionedHostToGinFileMiddlewareID(id)
	return phc
}

// SetNillableProvisionedHostToGinFileMiddlewareID sets the "ProvisionedHostToGinFileMiddleware" edge to the GinFileMiddleware entity by ID if the given value is not nil.
func (phc *ProvisionedHostCreate) SetNillableProvisionedHostToGinFileMiddlewareID(id *uuid.UUID) *ProvisionedHostCreate {
	if id != nil {
		phc = phc.SetProvisionedHostToGinFileMiddlewareID(*id)
	}
	return phc
}

// SetProvisionedHostToGinFileMiddleware sets the "ProvisionedHostToGinFileMiddleware" edge to the GinFileMiddleware entity.
func (phc *ProvisionedHostCreate) SetProvisionedHostToGinFileMiddleware(g *GinFileMiddleware) *ProvisionedHostCreate {
	return phc.SetProvisionedHostToGinFileMiddlewareID(g.ID)
}

// Mutation returns the ProvisionedHostMutation object of the builder.
func (phc *ProvisionedHostCreate) Mutation() *ProvisionedHostMutation {
	return phc.mutation
}

// Save creates the ProvisionedHost in the database.
func (phc *ProvisionedHostCreate) Save(ctx context.Context) (*ProvisionedHost, error) {
	var (
		err  error
		node *ProvisionedHost
	)
	phc.defaults()
	if len(phc.hooks) == 0 {
		if err = phc.check(); err != nil {
			return nil, err
		}
		node, err = phc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProvisionedHostMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = phc.check(); err != nil {
				return nil, err
			}
			phc.mutation = mutation
			if node, err = phc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(phc.hooks) - 1; i >= 0; i-- {
			if phc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = phc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, phc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (phc *ProvisionedHostCreate) SaveX(ctx context.Context) *ProvisionedHost {
	v, err := phc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (phc *ProvisionedHostCreate) Exec(ctx context.Context) error {
	_, err := phc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (phc *ProvisionedHostCreate) ExecX(ctx context.Context) {
	if err := phc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (phc *ProvisionedHostCreate) defaults() {
	if _, ok := phc.mutation.ID(); !ok {
		v := provisionedhost.DefaultID()
		phc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (phc *ProvisionedHostCreate) check() error {
	if _, ok := phc.mutation.SubnetIP(); !ok {
		return &ValidationError{Name: "subnet_ip", err: errors.New(`ent: missing required field "subnet_ip"`)}
	}
	if v, ok := phc.mutation.AddonType(); ok {
		if err := provisionedhost.AddonTypeValidator(v); err != nil {
			return &ValidationError{Name: "addon_type", err: fmt.Errorf(`ent: validator failed for field "addon_type": %w`, err)}
		}
	}
	if _, ok := phc.mutation.ProvisionedHostToStatusID(); !ok {
		return &ValidationError{Name: "ProvisionedHostToStatus", err: errors.New("ent: missing required edge \"ProvisionedHostToStatus\"")}
	}
	if _, ok := phc.mutation.ProvisionedHostToProvisionedNetworkID(); !ok {
		return &ValidationError{Name: "ProvisionedHostToProvisionedNetwork", err: errors.New("ent: missing required edge \"ProvisionedHostToProvisionedNetwork\"")}
	}
	if _, ok := phc.mutation.ProvisionedHostToHostID(); !ok {
		return &ValidationError{Name: "ProvisionedHostToHost", err: errors.New("ent: missing required edge \"ProvisionedHostToHost\"")}
	}
	if _, ok := phc.mutation.ProvisionedHostToBuildID(); !ok {
		return &ValidationError{Name: "ProvisionedHostToBuild", err: errors.New("ent: missing required edge \"ProvisionedHostToBuild\"")}
	}
	return nil
}

func (phc *ProvisionedHostCreate) sqlSave(ctx context.Context) (*ProvisionedHost, error) {
	_node, _spec := phc.createSpec()
	if err := sqlgraph.CreateNode(ctx, phc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		_node.ID = _spec.ID.Value.(uuid.UUID)
	}
	return _node, nil
}

func (phc *ProvisionedHostCreate) createSpec() (*ProvisionedHost, *sqlgraph.CreateSpec) {
	var (
		_node = &ProvisionedHost{config: phc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: provisionedhost.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: provisionedhost.FieldID,
			},
		}
	)
	if id, ok := phc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := phc.mutation.SubnetIP(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: provisionedhost.FieldSubnetIP,
		})
		_node.SubnetIP = value
	}
	if value, ok := phc.mutation.AddonType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: provisionedhost.FieldAddonType,
		})
		_node.AddonType = &value
	}
	if nodes := phc.mutation.ProvisionedHostToStatusIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   provisionedhost.ProvisionedHostToStatusTable,
			Columns: []string{provisionedhost.ProvisionedHostToStatusColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := phc.mutation.ProvisionedHostToProvisionedNetworkIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   provisionedhost.ProvisionedHostToProvisionedNetworkTable,
			Columns: []string{provisionedhost.ProvisionedHostToProvisionedNetworkColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: provisionednetwork.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.provisioned_host_provisioned_host_to_provisioned_network = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := phc.mutation.ProvisionedHostToHostIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   provisionedhost.ProvisionedHostToHostTable,
			Columns: []string{provisionedhost.ProvisionedHostToHostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: host.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.provisioned_host_provisioned_host_to_host = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := phc.mutation.ProvisionedHostToEndStepPlanIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   provisionedhost.ProvisionedHostToEndStepPlanTable,
			Columns: []string{provisionedhost.ProvisionedHostToEndStepPlanColumn},
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
		_node.provisioned_host_provisioned_host_to_end_step_plan = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := phc.mutation.ProvisionedHostToBuildIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   provisionedhost.ProvisionedHostToBuildTable,
			Columns: []string{provisionedhost.ProvisionedHostToBuildColumn},
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
		_node.provisioned_host_provisioned_host_to_build = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := phc.mutation.ProvisionedHostToProvisioningStepIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   provisionedhost.ProvisionedHostToProvisioningStepTable,
			Columns: []string{provisionedhost.ProvisionedHostToProvisioningStepColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: provisioningstep.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := phc.mutation.ProvisionedHostToAgentStatusIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   provisionedhost.ProvisionedHostToAgentStatusTable,
			Columns: []string{provisionedhost.ProvisionedHostToAgentStatusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: agentstatus.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := phc.mutation.ProvisionedHostToAgentTaskIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   provisionedhost.ProvisionedHostToAgentTaskTable,
			Columns: []string{provisionedhost.ProvisionedHostToAgentTaskColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: agenttask.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := phc.mutation.ProvisionedHostToPlanIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   provisionedhost.ProvisionedHostToPlanTable,
			Columns: []string{provisionedhost.ProvisionedHostToPlanColumn},
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
		_node.plan_plan_to_provisioned_host = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := phc.mutation.ProvisionedHostToGinFileMiddlewareIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   provisionedhost.ProvisionedHostToGinFileMiddlewareTable,
			Columns: []string{provisionedhost.ProvisionedHostToGinFileMiddlewareColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: ginfilemiddleware.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.gin_file_middleware_gin_file_middleware_to_provisioned_host = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ProvisionedHostCreateBulk is the builder for creating many ProvisionedHost entities in bulk.
type ProvisionedHostCreateBulk struct {
	config
	builders []*ProvisionedHostCreate
}

// Save creates the ProvisionedHost entities in the database.
func (phcb *ProvisionedHostCreateBulk) Save(ctx context.Context) ([]*ProvisionedHost, error) {
	specs := make([]*sqlgraph.CreateSpec, len(phcb.builders))
	nodes := make([]*ProvisionedHost, len(phcb.builders))
	mutators := make([]Mutator, len(phcb.builders))
	for i := range phcb.builders {
		func(i int, root context.Context) {
			builder := phcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProvisionedHostMutation)
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
					_, err = mutators[i+1].Mutate(root, phcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, phcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, phcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (phcb *ProvisionedHostCreateBulk) SaveX(ctx context.Context) []*ProvisionedHost {
	v, err := phcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (phcb *ProvisionedHostCreateBulk) Exec(ctx context.Context) error {
	_, err := phcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (phcb *ProvisionedHostCreateBulk) ExecX(ctx context.Context) {
	if err := phcb.Exec(ctx); err != nil {
		panic(err)
	}
}
