// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gen0cide/laforge/ent/command"
	"github.com/gen0cide/laforge/ent/tag"
	"github.com/gen0cide/laforge/ent/user"
)

// CommandCreate is the builder for creating a Command entity.
type CommandCreate struct {
	config
	mutation *CommandMutation
	hooks    []Hook
}

// SetHclID sets the "hcl_id" field.
func (cc *CommandCreate) SetHclID(s string) *CommandCreate {
	cc.mutation.SetHclID(s)
	return cc
}

// SetName sets the "name" field.
func (cc *CommandCreate) SetName(s string) *CommandCreate {
	cc.mutation.SetName(s)
	return cc
}

// SetDescription sets the "description" field.
func (cc *CommandCreate) SetDescription(s string) *CommandCreate {
	cc.mutation.SetDescription(s)
	return cc
}

// SetProgram sets the "program" field.
func (cc *CommandCreate) SetProgram(s string) *CommandCreate {
	cc.mutation.SetProgram(s)
	return cc
}

// SetArgs sets the "args" field.
func (cc *CommandCreate) SetArgs(s []string) *CommandCreate {
	cc.mutation.SetArgs(s)
	return cc
}

// SetIgnoreErrors sets the "ignore_errors" field.
func (cc *CommandCreate) SetIgnoreErrors(b bool) *CommandCreate {
	cc.mutation.SetIgnoreErrors(b)
	return cc
}

// SetDisabled sets the "disabled" field.
func (cc *CommandCreate) SetDisabled(b bool) *CommandCreate {
	cc.mutation.SetDisabled(b)
	return cc
}

// SetCooldown sets the "cooldown" field.
func (cc *CommandCreate) SetCooldown(i int) *CommandCreate {
	cc.mutation.SetCooldown(i)
	return cc
}

// SetTimeout sets the "timeout" field.
func (cc *CommandCreate) SetTimeout(i int) *CommandCreate {
	cc.mutation.SetTimeout(i)
	return cc
}

// SetVars sets the "vars" field.
func (cc *CommandCreate) SetVars(m map[string]string) *CommandCreate {
	cc.mutation.SetVars(m)
	return cc
}

// SetTags sets the "tags" field.
func (cc *CommandCreate) SetTags(m map[string]string) *CommandCreate {
	cc.mutation.SetTags(m)
	return cc
}

// AddCommandToUserIDs adds the "CommandToUser" edge to the User entity by IDs.
func (cc *CommandCreate) AddCommandToUserIDs(ids ...int) *CommandCreate {
	cc.mutation.AddCommandToUserIDs(ids...)
	return cc
}

// AddCommandToUser adds the "CommandToUser" edges to the User entity.
func (cc *CommandCreate) AddCommandToUser(u ...*User) *CommandCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return cc.AddCommandToUserIDs(ids...)
}

// AddCommandToTagIDs adds the "CommandToTag" edge to the Tag entity by IDs.
func (cc *CommandCreate) AddCommandToTagIDs(ids ...int) *CommandCreate {
	cc.mutation.AddCommandToTagIDs(ids...)
	return cc
}

// AddCommandToTag adds the "CommandToTag" edges to the Tag entity.
func (cc *CommandCreate) AddCommandToTag(t ...*Tag) *CommandCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return cc.AddCommandToTagIDs(ids...)
}

// Mutation returns the CommandMutation object of the builder.
func (cc *CommandCreate) Mutation() *CommandMutation {
	return cc.mutation
}

// Save creates the Command in the database.
func (cc *CommandCreate) Save(ctx context.Context) (*Command, error) {
	var (
		err  error
		node *Command
	)
	if len(cc.hooks) == 0 {
		if err = cc.check(); err != nil {
			return nil, err
		}
		node, err = cc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CommandMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cc.check(); err != nil {
				return nil, err
			}
			cc.mutation = mutation
			node, err = cc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cc.hooks) - 1; i >= 0; i-- {
			mut = cc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CommandCreate) SaveX(ctx context.Context) *Command {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (cc *CommandCreate) check() error {
	if _, ok := cc.mutation.HclID(); !ok {
		return &ValidationError{Name: "hcl_id", err: errors.New("ent: missing required field \"hcl_id\"")}
	}
	if _, ok := cc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	if _, ok := cc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New("ent: missing required field \"description\"")}
	}
	if _, ok := cc.mutation.Program(); !ok {
		return &ValidationError{Name: "program", err: errors.New("ent: missing required field \"program\"")}
	}
	if _, ok := cc.mutation.Args(); !ok {
		return &ValidationError{Name: "args", err: errors.New("ent: missing required field \"args\"")}
	}
	if _, ok := cc.mutation.IgnoreErrors(); !ok {
		return &ValidationError{Name: "ignore_errors", err: errors.New("ent: missing required field \"ignore_errors\"")}
	}
	if _, ok := cc.mutation.Disabled(); !ok {
		return &ValidationError{Name: "disabled", err: errors.New("ent: missing required field \"disabled\"")}
	}
	if _, ok := cc.mutation.Cooldown(); !ok {
		return &ValidationError{Name: "cooldown", err: errors.New("ent: missing required field \"cooldown\"")}
	}
	if v, ok := cc.mutation.Cooldown(); ok {
		if err := command.CooldownValidator(v); err != nil {
			return &ValidationError{Name: "cooldown", err: fmt.Errorf("ent: validator failed for field \"cooldown\": %w", err)}
		}
	}
	if _, ok := cc.mutation.Timeout(); !ok {
		return &ValidationError{Name: "timeout", err: errors.New("ent: missing required field \"timeout\"")}
	}
	if v, ok := cc.mutation.Timeout(); ok {
		if err := command.TimeoutValidator(v); err != nil {
			return &ValidationError{Name: "timeout", err: fmt.Errorf("ent: validator failed for field \"timeout\": %w", err)}
		}
	}
	if _, ok := cc.mutation.Vars(); !ok {
		return &ValidationError{Name: "vars", err: errors.New("ent: missing required field \"vars\"")}
	}
	if _, ok := cc.mutation.Tags(); !ok {
		return &ValidationError{Name: "tags", err: errors.New("ent: missing required field \"tags\"")}
	}
	return nil
}

func (cc *CommandCreate) sqlSave(ctx context.Context) (*Command, error) {
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (cc *CommandCreate) createSpec() (*Command, *sqlgraph.CreateSpec) {
	var (
		_node = &Command{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: command.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: command.FieldID,
			},
		}
	)
	if value, ok := cc.mutation.HclID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: command.FieldHclID,
		})
		_node.HclID = value
	}
	if value, ok := cc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: command.FieldName,
		})
		_node.Name = value
	}
	if value, ok := cc.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: command.FieldDescription,
		})
		_node.Description = value
	}
	if value, ok := cc.mutation.Program(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: command.FieldProgram,
		})
		_node.Program = value
	}
	if value, ok := cc.mutation.Args(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: command.FieldArgs,
		})
		_node.Args = value
	}
	if value, ok := cc.mutation.IgnoreErrors(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: command.FieldIgnoreErrors,
		})
		_node.IgnoreErrors = value
	}
	if value, ok := cc.mutation.Disabled(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: command.FieldDisabled,
		})
		_node.Disabled = value
	}
	if value, ok := cc.mutation.Cooldown(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: command.FieldCooldown,
		})
		_node.Cooldown = value
	}
	if value, ok := cc.mutation.Timeout(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: command.FieldTimeout,
		})
		_node.Timeout = value
	}
	if value, ok := cc.mutation.Vars(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: command.FieldVars,
		})
		_node.Vars = value
	}
	if value, ok := cc.mutation.Tags(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: command.FieldTags,
		})
		_node.Tags = value
	}
	if nodes := cc.mutation.CommandToUserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   command.CommandToUserTable,
			Columns: []string{command.CommandToUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.CommandToTagIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   command.CommandToTagTable,
			Columns: []string{command.CommandToTagColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tag.FieldID,
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

// CommandCreateBulk is the builder for creating many Command entities in bulk.
type CommandCreateBulk struct {
	config
	builders []*CommandCreate
}

// Save creates the Command entities in the database.
func (ccb *CommandCreateBulk) Save(ctx context.Context) ([]*Command, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Command, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CommandMutation)
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
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CommandCreateBulk) SaveX(ctx context.Context) []*Command {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
