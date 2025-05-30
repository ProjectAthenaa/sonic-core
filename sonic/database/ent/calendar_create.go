// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ProjectAthenaa/sonic-core/sonic/database/ent/calendar"
	"github.com/ProjectAthenaa/sonic-core/sonic/database/ent/product"
	"github.com/google/uuid"
)

// CalendarCreate is the builder for creating a Calendar entity.
type CalendarCreate struct {
	config
	mutation *CalendarMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (cc *CalendarCreate) SetCreatedAt(t time.Time) *CalendarCreate {
	cc.mutation.SetCreatedAt(t)
	return cc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cc *CalendarCreate) SetNillableCreatedAt(t *time.Time) *CalendarCreate {
	if t != nil {
		cc.SetCreatedAt(*t)
	}
	return cc
}

// SetUpdatedAt sets the "updated_at" field.
func (cc *CalendarCreate) SetUpdatedAt(t time.Time) *CalendarCreate {
	cc.mutation.SetUpdatedAt(t)
	return cc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cc *CalendarCreate) SetNillableUpdatedAt(t *time.Time) *CalendarCreate {
	if t != nil {
		cc.SetUpdatedAt(*t)
	}
	return cc
}

// SetReleaseDate sets the "ReleaseDate" field.
func (cc *CalendarCreate) SetReleaseDate(t time.Time) *CalendarCreate {
	cc.mutation.SetReleaseDate(t)
	return cc
}

// SetProductImage sets the "ProductImage" field.
func (cc *CalendarCreate) SetProductImage(s string) *CalendarCreate {
	cc.mutation.SetProductImage(s)
	return cc
}

// SetNillableProductImage sets the "ProductImage" field if the given value is not nil.
func (cc *CalendarCreate) SetNillableProductImage(s *string) *CalendarCreate {
	if s != nil {
		cc.SetProductImage(*s)
	}
	return cc
}

// SetProductName sets the "ProductName" field.
func (cc *CalendarCreate) SetProductName(s string) *CalendarCreate {
	cc.mutation.SetProductName(s)
	return cc
}

// SetHypedRelease sets the "HypedRelease" field.
func (cc *CalendarCreate) SetHypedRelease(b bool) *CalendarCreate {
	cc.mutation.SetHypedRelease(b)
	return cc
}

// SetUsersRunning sets the "UsersRunning" field.
func (cc *CalendarCreate) SetUsersRunning(i int) *CalendarCreate {
	cc.mutation.SetUsersRunning(i)
	return cc
}

// SetID sets the "id" field.
func (cc *CalendarCreate) SetID(u uuid.UUID) *CalendarCreate {
	cc.mutation.SetID(u)
	return cc
}

// SetQuickTaskID sets the "QuickTask" edge to the Product entity by ID.
func (cc *CalendarCreate) SetQuickTaskID(id uuid.UUID) *CalendarCreate {
	cc.mutation.SetQuickTaskID(id)
	return cc
}

// SetNillableQuickTaskID sets the "QuickTask" edge to the Product entity by ID if the given value is not nil.
func (cc *CalendarCreate) SetNillableQuickTaskID(id *uuid.UUID) *CalendarCreate {
	if id != nil {
		cc = cc.SetQuickTaskID(*id)
	}
	return cc
}

// SetQuickTask sets the "QuickTask" edge to the Product entity.
func (cc *CalendarCreate) SetQuickTask(p *Product) *CalendarCreate {
	return cc.SetQuickTaskID(p.ID)
}

// Mutation returns the CalendarMutation object of the builder.
func (cc *CalendarCreate) Mutation() *CalendarMutation {
	return cc.mutation
}

// Save creates the Calendar in the database.
func (cc *CalendarCreate) Save(ctx context.Context) (*Calendar, error) {
	var (
		err  error
		node *Calendar
	)
	cc.defaults()
	if len(cc.hooks) == 0 {
		if err = cc.check(); err != nil {
			return nil, err
		}
		node, err = cc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CalendarMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cc.check(); err != nil {
				return nil, err
			}
			cc.mutation = mutation
			if node, err = cc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(cc.hooks) - 1; i >= 0; i-- {
			if cc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CalendarCreate) SaveX(ctx context.Context) *Calendar {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CalendarCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CalendarCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *CalendarCreate) defaults() {
	if _, ok := cc.mutation.CreatedAt(); !ok {
		v := calendar.DefaultCreatedAt()
		cc.mutation.SetCreatedAt(v)
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		v := calendar.DefaultUpdatedAt()
		cc.mutation.SetUpdatedAt(v)
	}
	if _, ok := cc.mutation.ProductImage(); !ok {
		v := calendar.DefaultProductImage
		cc.mutation.SetProductImage(v)
	}
	if _, ok := cc.mutation.ID(); !ok {
		v := calendar.DefaultID()
		cc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CalendarCreate) check() error {
	if _, ok := cc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "updated_at"`)}
	}
	if _, ok := cc.mutation.ReleaseDate(); !ok {
		return &ValidationError{Name: "ReleaseDate", err: errors.New(`ent: missing required field "ReleaseDate"`)}
	}
	if _, ok := cc.mutation.ProductImage(); !ok {
		return &ValidationError{Name: "ProductImage", err: errors.New(`ent: missing required field "ProductImage"`)}
	}
	if _, ok := cc.mutation.ProductName(); !ok {
		return &ValidationError{Name: "ProductName", err: errors.New(`ent: missing required field "ProductName"`)}
	}
	if _, ok := cc.mutation.HypedRelease(); !ok {
		return &ValidationError{Name: "HypedRelease", err: errors.New(`ent: missing required field "HypedRelease"`)}
	}
	if _, ok := cc.mutation.UsersRunning(); !ok {
		return &ValidationError{Name: "UsersRunning", err: errors.New(`ent: missing required field "UsersRunning"`)}
	}
	return nil
}

func (cc *CalendarCreate) sqlSave(ctx context.Context) (*Calendar, error) {
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
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

func (cc *CalendarCreate) createSpec() (*Calendar, *sqlgraph.CreateSpec) {
	var (
		_node = &Calendar{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: calendar.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: calendar.FieldID,
			},
		}
	)
	if id, ok := cc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := cc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: calendar.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := cc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: calendar.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := cc.mutation.ReleaseDate(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: calendar.FieldReleaseDate,
		})
		_node.ReleaseDate = value
	}
	if value, ok := cc.mutation.ProductImage(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: calendar.FieldProductImage,
		})
		_node.ProductImage = value
	}
	if value, ok := cc.mutation.ProductName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: calendar.FieldProductName,
		})
		_node.ProductName = value
	}
	if value, ok := cc.mutation.HypedRelease(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: calendar.FieldHypedRelease,
		})
		_node.HypedRelease = value
	}
	if value, ok := cc.mutation.UsersRunning(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: calendar.FieldUsersRunning,
		})
		_node.UsersRunning = value
	}
	if nodes := cc.mutation.QuickTaskIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   calendar.QuickTaskTable,
			Columns: []string{calendar.QuickTaskColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: product.FieldID,
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

// CalendarCreateBulk is the builder for creating many Calendar entities in bulk.
type CalendarCreateBulk struct {
	config
	builders []*CalendarCreate
}

// Save creates the Calendar entities in the database.
func (ccb *CalendarCreateBulk) Save(ctx context.Context) ([]*Calendar, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Calendar, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CalendarMutation)
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
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CalendarCreateBulk) SaveX(ctx context.Context) []*Calendar {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CalendarCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CalendarCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
