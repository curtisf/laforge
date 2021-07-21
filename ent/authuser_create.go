// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gen0cide/laforge/ent/authuser"
	"github.com/gen0cide/laforge/ent/servertask"
	"github.com/gen0cide/laforge/ent/token"
	"github.com/google/uuid"
)

// AuthUserCreate is the builder for creating a AuthUser entity.
type AuthUserCreate struct {
	config
	mutation *AuthUserMutation
	hooks    []Hook
}

// SetUsername sets the "username" field.
func (auc *AuthUserCreate) SetUsername(s string) *AuthUserCreate {
	auc.mutation.SetUsername(s)
	return auc
}

// SetPassword sets the "password" field.
func (auc *AuthUserCreate) SetPassword(s string) *AuthUserCreate {
	auc.mutation.SetPassword(s)
	return auc
}

// SetFirstName sets the "first_name" field.
func (auc *AuthUserCreate) SetFirstName(s string) *AuthUserCreate {
	auc.mutation.SetFirstName(s)
	return auc
}

// SetNillableFirstName sets the "first_name" field if the given value is not nil.
func (auc *AuthUserCreate) SetNillableFirstName(s *string) *AuthUserCreate {
	if s != nil {
		auc.SetFirstName(*s)
	}
	return auc
}

// SetLastName sets the "last_name" field.
func (auc *AuthUserCreate) SetLastName(s string) *AuthUserCreate {
	auc.mutation.SetLastName(s)
	return auc
}

// SetNillableLastName sets the "last_name" field if the given value is not nil.
func (auc *AuthUserCreate) SetNillableLastName(s *string) *AuthUserCreate {
	if s != nil {
		auc.SetLastName(*s)
	}
	return auc
}

// SetEmail sets the "email" field.
func (auc *AuthUserCreate) SetEmail(s string) *AuthUserCreate {
	auc.mutation.SetEmail(s)
	return auc
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (auc *AuthUserCreate) SetNillableEmail(s *string) *AuthUserCreate {
	if s != nil {
		auc.SetEmail(*s)
	}
	return auc
}

// SetPhone sets the "phone" field.
func (auc *AuthUserCreate) SetPhone(s string) *AuthUserCreate {
	auc.mutation.SetPhone(s)
	return auc
}

// SetNillablePhone sets the "phone" field if the given value is not nil.
func (auc *AuthUserCreate) SetNillablePhone(s *string) *AuthUserCreate {
	if s != nil {
		auc.SetPhone(*s)
	}
	return auc
}

// SetCompany sets the "company" field.
func (auc *AuthUserCreate) SetCompany(s string) *AuthUserCreate {
	auc.mutation.SetCompany(s)
	return auc
}

// SetNillableCompany sets the "company" field if the given value is not nil.
func (auc *AuthUserCreate) SetNillableCompany(s *string) *AuthUserCreate {
	if s != nil {
		auc.SetCompany(*s)
	}
	return auc
}

// SetOccupation sets the "occupation" field.
func (auc *AuthUserCreate) SetOccupation(s string) *AuthUserCreate {
	auc.mutation.SetOccupation(s)
	return auc
}

// SetNillableOccupation sets the "occupation" field if the given value is not nil.
func (auc *AuthUserCreate) SetNillableOccupation(s *string) *AuthUserCreate {
	if s != nil {
		auc.SetOccupation(*s)
	}
	return auc
}

// SetPrivateKeyPath sets the "private_key_path" field.
func (auc *AuthUserCreate) SetPrivateKeyPath(s string) *AuthUserCreate {
	auc.mutation.SetPrivateKeyPath(s)
	return auc
}

// SetNillablePrivateKeyPath sets the "private_key_path" field if the given value is not nil.
func (auc *AuthUserCreate) SetNillablePrivateKeyPath(s *string) *AuthUserCreate {
	if s != nil {
		auc.SetPrivateKeyPath(*s)
	}
	return auc
}

// SetRole sets the "role" field.
func (auc *AuthUserCreate) SetRole(a authuser.Role) *AuthUserCreate {
	auc.mutation.SetRole(a)
	return auc
}

// SetProvider sets the "provider" field.
func (auc *AuthUserCreate) SetProvider(a authuser.Provider) *AuthUserCreate {
	auc.mutation.SetProvider(a)
	return auc
}

// SetID sets the "id" field.
func (auc *AuthUserCreate) SetID(u uuid.UUID) *AuthUserCreate {
	auc.mutation.SetID(u)
	return auc
}

// AddAuthUserToTokenIDs adds the "AuthUserToToken" edge to the Token entity by IDs.
func (auc *AuthUserCreate) AddAuthUserToTokenIDs(ids ...uuid.UUID) *AuthUserCreate {
	auc.mutation.AddAuthUserToTokenIDs(ids...)
	return auc
}

// AddAuthUserToToken adds the "AuthUserToToken" edges to the Token entity.
func (auc *AuthUserCreate) AddAuthUserToToken(t ...*Token) *AuthUserCreate {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return auc.AddAuthUserToTokenIDs(ids...)
}

// AddAuthUserToServerTaskIDs adds the "AuthUserToServerTasks" edge to the ServerTask entity by IDs.
func (auc *AuthUserCreate) AddAuthUserToServerTaskIDs(ids ...uuid.UUID) *AuthUserCreate {
	auc.mutation.AddAuthUserToServerTaskIDs(ids...)
	return auc
}

// AddAuthUserToServerTasks adds the "AuthUserToServerTasks" edges to the ServerTask entity.
func (auc *AuthUserCreate) AddAuthUserToServerTasks(s ...*ServerTask) *AuthUserCreate {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return auc.AddAuthUserToServerTaskIDs(ids...)
}

// Mutation returns the AuthUserMutation object of the builder.
func (auc *AuthUserCreate) Mutation() *AuthUserMutation {
	return auc.mutation
}

// Save creates the AuthUser in the database.
func (auc *AuthUserCreate) Save(ctx context.Context) (*AuthUser, error) {
	var (
		err  error
		node *AuthUser
	)
	auc.defaults()
	if len(auc.hooks) == 0 {
		if err = auc.check(); err != nil {
			return nil, err
		}
		node, err = auc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AuthUserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = auc.check(); err != nil {
				return nil, err
			}
			auc.mutation = mutation
			node, err = auc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(auc.hooks) - 1; i >= 0; i-- {
			mut = auc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, auc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (auc *AuthUserCreate) SaveX(ctx context.Context) *AuthUser {
	v, err := auc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (auc *AuthUserCreate) defaults() {
	if _, ok := auc.mutation.FirstName(); !ok {
		v := authuser.DefaultFirstName
		auc.mutation.SetFirstName(v)
	}
	if _, ok := auc.mutation.LastName(); !ok {
		v := authuser.DefaultLastName
		auc.mutation.SetLastName(v)
	}
	if _, ok := auc.mutation.Email(); !ok {
		v := authuser.DefaultEmail
		auc.mutation.SetEmail(v)
	}
	if _, ok := auc.mutation.Phone(); !ok {
		v := authuser.DefaultPhone
		auc.mutation.SetPhone(v)
	}
	if _, ok := auc.mutation.Company(); !ok {
		v := authuser.DefaultCompany
		auc.mutation.SetCompany(v)
	}
	if _, ok := auc.mutation.Occupation(); !ok {
		v := authuser.DefaultOccupation
		auc.mutation.SetOccupation(v)
	}
	if _, ok := auc.mutation.PrivateKeyPath(); !ok {
		v := authuser.DefaultPrivateKeyPath
		auc.mutation.SetPrivateKeyPath(v)
	}
	if _, ok := auc.mutation.ID(); !ok {
		v := authuser.DefaultID()
		auc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (auc *AuthUserCreate) check() error {
	if _, ok := auc.mutation.Username(); !ok {
		return &ValidationError{Name: "username", err: errors.New("ent: missing required field \"username\"")}
	}
	if _, ok := auc.mutation.Password(); !ok {
		return &ValidationError{Name: "password", err: errors.New("ent: missing required field \"password\"")}
	}
	if _, ok := auc.mutation.FirstName(); !ok {
		return &ValidationError{Name: "first_name", err: errors.New("ent: missing required field \"first_name\"")}
	}
	if _, ok := auc.mutation.LastName(); !ok {
		return &ValidationError{Name: "last_name", err: errors.New("ent: missing required field \"last_name\"")}
	}
	if _, ok := auc.mutation.Email(); !ok {
		return &ValidationError{Name: "email", err: errors.New("ent: missing required field \"email\"")}
	}
	if _, ok := auc.mutation.Phone(); !ok {
		return &ValidationError{Name: "phone", err: errors.New("ent: missing required field \"phone\"")}
	}
	if _, ok := auc.mutation.Company(); !ok {
		return &ValidationError{Name: "company", err: errors.New("ent: missing required field \"company\"")}
	}
	if _, ok := auc.mutation.Occupation(); !ok {
		return &ValidationError{Name: "occupation", err: errors.New("ent: missing required field \"occupation\"")}
	}
	if _, ok := auc.mutation.PrivateKeyPath(); !ok {
		return &ValidationError{Name: "private_key_path", err: errors.New("ent: missing required field \"private_key_path\"")}
	}
	if _, ok := auc.mutation.Role(); !ok {
		return &ValidationError{Name: "role", err: errors.New("ent: missing required field \"role\"")}
	}
	if v, ok := auc.mutation.Role(); ok {
		if err := authuser.RoleValidator(v); err != nil {
			return &ValidationError{Name: "role", err: fmt.Errorf("ent: validator failed for field \"role\": %w", err)}
		}
	}
	if _, ok := auc.mutation.Provider(); !ok {
		return &ValidationError{Name: "provider", err: errors.New("ent: missing required field \"provider\"")}
	}
	if v, ok := auc.mutation.Provider(); ok {
		if err := authuser.ProviderValidator(v); err != nil {
			return &ValidationError{Name: "provider", err: fmt.Errorf("ent: validator failed for field \"provider\": %w", err)}
		}
	}
	return nil
}

func (auc *AuthUserCreate) sqlSave(ctx context.Context) (*AuthUser, error) {
	_node, _spec := auc.createSpec()
	if err := sqlgraph.CreateNode(ctx, auc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}

func (auc *AuthUserCreate) createSpec() (*AuthUser, *sqlgraph.CreateSpec) {
	var (
		_node = &AuthUser{config: auc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: authuser.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: authuser.FieldID,
			},
		}
	)
	if id, ok := auc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := auc.mutation.Username(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: authuser.FieldUsername,
		})
		_node.Username = value
	}
	if value, ok := auc.mutation.Password(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: authuser.FieldPassword,
		})
		_node.Password = value
	}
	if value, ok := auc.mutation.FirstName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: authuser.FieldFirstName,
		})
		_node.FirstName = value
	}
	if value, ok := auc.mutation.LastName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: authuser.FieldLastName,
		})
		_node.LastName = value
	}
	if value, ok := auc.mutation.Email(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: authuser.FieldEmail,
		})
		_node.Email = value
	}
	if value, ok := auc.mutation.Phone(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: authuser.FieldPhone,
		})
		_node.Phone = value
	}
	if value, ok := auc.mutation.Company(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: authuser.FieldCompany,
		})
		_node.Company = value
	}
	if value, ok := auc.mutation.Occupation(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: authuser.FieldOccupation,
		})
		_node.Occupation = value
	}
	if value, ok := auc.mutation.PrivateKeyPath(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: authuser.FieldPrivateKeyPath,
		})
		_node.PrivateKeyPath = value
	}
	if value, ok := auc.mutation.Role(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: authuser.FieldRole,
		})
		_node.Role = value
	}
	if value, ok := auc.mutation.Provider(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: authuser.FieldProvider,
		})
		_node.Provider = value
	}
	if nodes := auc.mutation.AuthUserToTokenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   authuser.AuthUserToTokenTable,
			Columns: []string{authuser.AuthUserToTokenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: token.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := auc.mutation.AuthUserToServerTasksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   authuser.AuthUserToServerTasksTable,
			Columns: []string{authuser.AuthUserToServerTasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: servertask.FieldID,
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

// AuthUserCreateBulk is the builder for creating many AuthUser entities in bulk.
type AuthUserCreateBulk struct {
	config
	builders []*AuthUserCreate
}

// Save creates the AuthUser entities in the database.
func (aucb *AuthUserCreateBulk) Save(ctx context.Context) ([]*AuthUser, error) {
	specs := make([]*sqlgraph.CreateSpec, len(aucb.builders))
	nodes := make([]*AuthUser, len(aucb.builders))
	mutators := make([]Mutator, len(aucb.builders))
	for i := range aucb.builders {
		func(i int, root context.Context) {
			builder := aucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AuthUserMutation)
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
					_, err = mutators[i+1].Mutate(root, aucb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, aucb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
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
		if _, err := mutators[0].Mutate(ctx, aucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (aucb *AuthUserCreateBulk) SaveX(ctx context.Context) []*AuthUser {
	v, err := aucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
