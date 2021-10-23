// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ktakenaka/gosample/ent/sample"
)

// SampleCreate is the builder for creating a Sample entity.
type SampleCreate struct {
	config
	mutation *SampleMutation
	hooks    []Hook
}

// SetOfficeID sets the "office_id" field.
func (sc *SampleCreate) SetOfficeID(i int) *SampleCreate {
	sc.mutation.SetOfficeID(i)
	return sc
}

// SetCode sets the "code" field.
func (sc *SampleCreate) SetCode(s string) *SampleCreate {
	sc.mutation.SetCode(s)
	return sc
}

// SetSize sets the "size" field.
func (sc *SampleCreate) SetSize(s sample.Size) *SampleCreate {
	sc.mutation.SetSize(s)
	return sc
}

// SetAmount sets the "amount" field.
func (sc *SampleCreate) SetAmount(f float64) *SampleCreate {
	sc.mutation.SetAmount(f)
	return sc
}

// SetMemo sets the "memo" field.
func (sc *SampleCreate) SetMemo(s string) *SampleCreate {
	sc.mutation.SetMemo(s)
	return sc
}

// SetNillableMemo sets the "memo" field if the given value is not nil.
func (sc *SampleCreate) SetNillableMemo(s *string) *SampleCreate {
	if s != nil {
		sc.SetMemo(*s)
	}
	return sc
}

// SetURL sets the "url" field.
func (sc *SampleCreate) SetURL(u *url.URL) *SampleCreate {
	sc.mutation.SetURL(u)
	return sc
}

// SetActive sets the "active" field.
func (sc *SampleCreate) SetActive(b bool) *SampleCreate {
	sc.mutation.SetActive(b)
	return sc
}

// SetNillableActive sets the "active" field if the given value is not nil.
func (sc *SampleCreate) SetNillableActive(b *bool) *SampleCreate {
	if b != nil {
		sc.SetActive(*b)
	}
	return sc
}

// SetCreatedAt sets the "created_at" field.
func (sc *SampleCreate) SetCreatedAt(t time.Time) *SampleCreate {
	sc.mutation.SetCreatedAt(t)
	return sc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (sc *SampleCreate) SetNillableCreatedAt(t *time.Time) *SampleCreate {
	if t != nil {
		sc.SetCreatedAt(*t)
	}
	return sc
}

// SetUpdatedAt sets the "updated_at" field.
func (sc *SampleCreate) SetUpdatedAt(t time.Time) *SampleCreate {
	sc.mutation.SetUpdatedAt(t)
	return sc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (sc *SampleCreate) SetNillableUpdatedAt(t *time.Time) *SampleCreate {
	if t != nil {
		sc.SetUpdatedAt(*t)
	}
	return sc
}

// SetID sets the "id" field.
func (sc *SampleCreate) SetID(s string) *SampleCreate {
	sc.mutation.SetID(s)
	return sc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (sc *SampleCreate) SetNillableID(s *string) *SampleCreate {
	if s != nil {
		sc.SetID(*s)
	}
	return sc
}

// Mutation returns the SampleMutation object of the builder.
func (sc *SampleCreate) Mutation() *SampleMutation {
	return sc.mutation
}

// Save creates the Sample in the database.
func (sc *SampleCreate) Save(ctx context.Context) (*Sample, error) {
	var (
		err  error
		node *Sample
	)
	sc.defaults()
	if len(sc.hooks) == 0 {
		if err = sc.check(); err != nil {
			return nil, err
		}
		node, err = sc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SampleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sc.check(); err != nil {
				return nil, err
			}
			sc.mutation = mutation
			if node, err = sc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(sc.hooks) - 1; i >= 0; i-- {
			if sc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (sc *SampleCreate) SaveX(ctx context.Context) *Sample {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *SampleCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *SampleCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *SampleCreate) defaults() {
	if _, ok := sc.mutation.Active(); !ok {
		v := sample.DefaultActive
		sc.mutation.SetActive(v)
	}
	if _, ok := sc.mutation.CreatedAt(); !ok {
		v := sample.DefaultCreatedAt()
		sc.mutation.SetCreatedAt(v)
	}
	if _, ok := sc.mutation.UpdatedAt(); !ok {
		v := sample.DefaultUpdatedAt()
		sc.mutation.SetUpdatedAt(v)
	}
	if _, ok := sc.mutation.ID(); !ok {
		v := sample.DefaultID()
		sc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *SampleCreate) check() error {
	if _, ok := sc.mutation.OfficeID(); !ok {
		return &ValidationError{Name: "office_id", err: errors.New(`ent: missing required field "office_id"`)}
	}
	if _, ok := sc.mutation.Code(); !ok {
		return &ValidationError{Name: "code", err: errors.New(`ent: missing required field "code"`)}
	}
	if v, ok := sc.mutation.Code(); ok {
		if err := sample.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`ent: validator failed for field "code": %w`, err)}
		}
	}
	if _, ok := sc.mutation.Size(); !ok {
		return &ValidationError{Name: "size", err: errors.New(`ent: missing required field "size"`)}
	}
	if v, ok := sc.mutation.Size(); ok {
		if err := sample.SizeValidator(v); err != nil {
			return &ValidationError{Name: "size", err: fmt.Errorf(`ent: validator failed for field "size": %w`, err)}
		}
	}
	if _, ok := sc.mutation.Amount(); !ok {
		return &ValidationError{Name: "amount", err: errors.New(`ent: missing required field "amount"`)}
	}
	if _, ok := sc.mutation.Active(); !ok {
		return &ValidationError{Name: "active", err: errors.New(`ent: missing required field "active"`)}
	}
	if _, ok := sc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	if _, ok := sc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "updated_at"`)}
	}
	return nil
}

func (sc *SampleCreate) sqlSave(ctx context.Context) (*Sample, error) {
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		_node.ID = _spec.ID.Value.(string)
	}
	return _node, nil
}

func (sc *SampleCreate) createSpec() (*Sample, *sqlgraph.CreateSpec) {
	var (
		_node = &Sample{config: sc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: sample.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: sample.FieldID,
			},
		}
	)
	if id, ok := sc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := sc.mutation.OfficeID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: sample.FieldOfficeID,
		})
		_node.OfficeID = value
	}
	if value, ok := sc.mutation.Code(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sample.FieldCode,
		})
		_node.Code = value
	}
	if value, ok := sc.mutation.Size(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: sample.FieldSize,
		})
		_node.Size = value
	}
	if value, ok := sc.mutation.Amount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: sample.FieldAmount,
		})
		_node.Amount = value
	}
	if value, ok := sc.mutation.Memo(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sample.FieldMemo,
		})
		_node.Memo = &value
	}
	if value, ok := sc.mutation.URL(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: sample.FieldURL,
		})
		_node.URL = value
	}
	if value, ok := sc.mutation.Active(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: sample.FieldActive,
		})
		_node.Active = value
	}
	if value, ok := sc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sample.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := sc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sample.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	return _node, _spec
}

// SampleCreateBulk is the builder for creating many Sample entities in bulk.
type SampleCreateBulk struct {
	config
	builders []*SampleCreate
}

// Save creates the Sample entities in the database.
func (scb *SampleCreateBulk) Save(ctx context.Context) ([]*Sample, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Sample, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SampleMutation)
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
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *SampleCreateBulk) SaveX(ctx context.Context) []*Sample {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *SampleCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *SampleCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}
