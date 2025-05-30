// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ProjectAthenaa/sonic-core/sonic/database/ent/accountgroup"
	"github.com/ProjectAthenaa/sonic-core/sonic/database/ent/predicate"
)

// AccountGroupDelete is the builder for deleting a AccountGroup entity.
type AccountGroupDelete struct {
	config
	hooks    []Hook
	mutation *AccountGroupMutation
}

// Where appends a list predicates to the AccountGroupDelete builder.
func (agd *AccountGroupDelete) Where(ps ...predicate.AccountGroup) *AccountGroupDelete {
	agd.mutation.Where(ps...)
	return agd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (agd *AccountGroupDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(agd.hooks) == 0 {
		affected, err = agd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AccountGroupMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			agd.mutation = mutation
			affected, err = agd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(agd.hooks) - 1; i >= 0; i-- {
			if agd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = agd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, agd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (agd *AccountGroupDelete) ExecX(ctx context.Context) int {
	n, err := agd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (agd *AccountGroupDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: accountgroup.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: accountgroup.FieldID,
			},
		},
	}
	if ps := agd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, agd.driver, _spec)
}

// AccountGroupDeleteOne is the builder for deleting a single AccountGroup entity.
type AccountGroupDeleteOne struct {
	agd *AccountGroupDelete
}

// Exec executes the deletion query.
func (agdo *AccountGroupDeleteOne) Exec(ctx context.Context) error {
	n, err := agdo.agd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{accountgroup.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (agdo *AccountGroupDeleteOne) ExecX(ctx context.Context) {
	agdo.agd.ExecX(ctx)
}
