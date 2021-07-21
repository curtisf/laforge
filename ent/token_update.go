// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gen0cide/laforge/ent/authuser"
	"github.com/gen0cide/laforge/ent/predicate"
	"github.com/gen0cide/laforge/ent/token"
	"github.com/google/uuid"
)

// TokenUpdate is the builder for updating Token entities.
type TokenUpdate struct {
	config
	hooks    []Hook
	mutation *TokenMutation
}

// Where adds a new predicate for the TokenUpdate builder.
func (tu *TokenUpdate) Where(ps ...predicate.Token) *TokenUpdate {
	tu.mutation.predicates = append(tu.mutation.predicates, ps...)
	return tu
}

// SetToken sets the "token" field.
func (tu *TokenUpdate) SetToken(s string) *TokenUpdate {
	tu.mutation.SetToken(s)
	return tu
}

// SetExpireAt sets the "expire_at" field.
func (tu *TokenUpdate) SetExpireAt(i int64) *TokenUpdate {
	tu.mutation.ResetExpireAt()
	tu.mutation.SetExpireAt(i)
	return tu
}

// AddExpireAt adds i to the "expire_at" field.
func (tu *TokenUpdate) AddExpireAt(i int64) *TokenUpdate {
	tu.mutation.AddExpireAt(i)
	return tu
}

// SetTokenToAuthUserID sets the "TokenToAuthUser" edge to the AuthUser entity by ID.
func (tu *TokenUpdate) SetTokenToAuthUserID(id uuid.UUID) *TokenUpdate {
	tu.mutation.SetTokenToAuthUserID(id)
	return tu
}

// SetTokenToAuthUser sets the "TokenToAuthUser" edge to the AuthUser entity.
func (tu *TokenUpdate) SetTokenToAuthUser(a *AuthUser) *TokenUpdate {
	return tu.SetTokenToAuthUserID(a.ID)
}

// Mutation returns the TokenMutation object of the builder.
func (tu *TokenUpdate) Mutation() *TokenMutation {
	return tu.mutation
}

// ClearTokenToAuthUser clears the "TokenToAuthUser" edge to the AuthUser entity.
func (tu *TokenUpdate) ClearTokenToAuthUser() *TokenUpdate {
	tu.mutation.ClearTokenToAuthUser()
	return tu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TokenUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(tu.hooks) == 0 {
		if err = tu.check(); err != nil {
			return 0, err
		}
		affected, err = tu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TokenMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tu.check(); err != nil {
				return 0, err
			}
			tu.mutation = mutation
			affected, err = tu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tu.hooks) - 1; i >= 0; i-- {
			mut = tu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TokenUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TokenUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TokenUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tu *TokenUpdate) check() error {
	if _, ok := tu.mutation.TokenToAuthUserID(); tu.mutation.TokenToAuthUserCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"TokenToAuthUser\"")
	}
	return nil
}

func (tu *TokenUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   token.Table,
			Columns: token.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: token.FieldID,
			},
		},
	}
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.Token(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: token.FieldToken,
		})
	}
	if value, ok := tu.mutation.ExpireAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: token.FieldExpireAt,
		})
	}
	if value, ok := tu.mutation.AddedExpireAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: token.FieldExpireAt,
		})
	}
	if tu.mutation.TokenToAuthUserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   token.TokenToAuthUserTable,
			Columns: []string{token.TokenToAuthUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: authuser.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.TokenToAuthUserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   token.TokenToAuthUserTable,
			Columns: []string{token.TokenToAuthUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: authuser.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{token.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// TokenUpdateOne is the builder for updating a single Token entity.
type TokenUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TokenMutation
}

// SetToken sets the "token" field.
func (tuo *TokenUpdateOne) SetToken(s string) *TokenUpdateOne {
	tuo.mutation.SetToken(s)
	return tuo
}

// SetExpireAt sets the "expire_at" field.
func (tuo *TokenUpdateOne) SetExpireAt(i int64) *TokenUpdateOne {
	tuo.mutation.ResetExpireAt()
	tuo.mutation.SetExpireAt(i)
	return tuo
}

// AddExpireAt adds i to the "expire_at" field.
func (tuo *TokenUpdateOne) AddExpireAt(i int64) *TokenUpdateOne {
	tuo.mutation.AddExpireAt(i)
	return tuo
}

// SetTokenToAuthUserID sets the "TokenToAuthUser" edge to the AuthUser entity by ID.
func (tuo *TokenUpdateOne) SetTokenToAuthUserID(id uuid.UUID) *TokenUpdateOne {
	tuo.mutation.SetTokenToAuthUserID(id)
	return tuo
}

// SetTokenToAuthUser sets the "TokenToAuthUser" edge to the AuthUser entity.
func (tuo *TokenUpdateOne) SetTokenToAuthUser(a *AuthUser) *TokenUpdateOne {
	return tuo.SetTokenToAuthUserID(a.ID)
}

// Mutation returns the TokenMutation object of the builder.
func (tuo *TokenUpdateOne) Mutation() *TokenMutation {
	return tuo.mutation
}

// ClearTokenToAuthUser clears the "TokenToAuthUser" edge to the AuthUser entity.
func (tuo *TokenUpdateOne) ClearTokenToAuthUser() *TokenUpdateOne {
	tuo.mutation.ClearTokenToAuthUser()
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TokenUpdateOne) Select(field string, fields ...string) *TokenUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Token entity.
func (tuo *TokenUpdateOne) Save(ctx context.Context) (*Token, error) {
	var (
		err  error
		node *Token
	)
	if len(tuo.hooks) == 0 {
		if err = tuo.check(); err != nil {
			return nil, err
		}
		node, err = tuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TokenMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tuo.check(); err != nil {
				return nil, err
			}
			tuo.mutation = mutation
			node, err = tuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tuo.hooks) - 1; i >= 0; i-- {
			mut = tuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TokenUpdateOne) SaveX(ctx context.Context) *Token {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TokenUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TokenUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tuo *TokenUpdateOne) check() error {
	if _, ok := tuo.mutation.TokenToAuthUserID(); tuo.mutation.TokenToAuthUserCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"TokenToAuthUser\"")
	}
	return nil
}

func (tuo *TokenUpdateOne) sqlSave(ctx context.Context) (_node *Token, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   token.Table,
			Columns: token.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: token.FieldID,
			},
		},
	}
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Token.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, token.FieldID)
		for _, f := range fields {
			if !token.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != token.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.Token(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: token.FieldToken,
		})
	}
	if value, ok := tuo.mutation.ExpireAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: token.FieldExpireAt,
		})
	}
	if value, ok := tuo.mutation.AddedExpireAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: token.FieldExpireAt,
		})
	}
	if tuo.mutation.TokenToAuthUserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   token.TokenToAuthUserTable,
			Columns: []string{token.TokenToAuthUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: authuser.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.TokenToAuthUserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   token.TokenToAuthUserTable,
			Columns: []string{token.TokenToAuthUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: authuser.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Token{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{token.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
