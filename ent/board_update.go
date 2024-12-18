// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"github.com/jaehyeonkim2358/gin-api-server/ent/board"
	"github.com/jaehyeonkim2358/gin-api-server/ent/predicate"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BoardUpdate is the builder for updating Board entities.
type BoardUpdate struct {
	config
	hooks    []Hook
	mutation *BoardMutation
}

// Where appends a list predicates to the BoardUpdate builder.
func (bu *BoardUpdate) Where(ps ...predicate.Board) *BoardUpdate {
	bu.mutation.Where(ps...)
	return bu
}

// SetTitle sets the "title" field.
func (bu *BoardUpdate) SetTitle(s string) *BoardUpdate {
	bu.mutation.SetTitle(s)
	return bu
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (bu *BoardUpdate) SetNillableTitle(s *string) *BoardUpdate {
	if s != nil {
		bu.SetTitle(*s)
	}
	return bu
}

// SetContent sets the "content" field.
func (bu *BoardUpdate) SetContent(s string) *BoardUpdate {
	bu.mutation.SetContent(s)
	return bu
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (bu *BoardUpdate) SetNillableContent(s *string) *BoardUpdate {
	if s != nil {
		bu.SetContent(*s)
	}
	return bu
}

// Mutation returns the BoardMutation object of the builder.
func (bu *BoardUpdate) Mutation() *BoardMutation {
	return bu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bu *BoardUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, bu.sqlSave, bu.mutation, bu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (bu *BoardUpdate) SaveX(ctx context.Context) int {
	affected, err := bu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bu *BoardUpdate) Exec(ctx context.Context) error {
	_, err := bu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bu *BoardUpdate) ExecX(ctx context.Context) {
	if err := bu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (bu *BoardUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(board.Table, board.Columns, sqlgraph.NewFieldSpec(board.FieldID, field.TypeInt))
	if ps := bu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bu.mutation.Title(); ok {
		_spec.SetField(board.FieldTitle, field.TypeString, value)
	}
	if value, ok := bu.mutation.Content(); ok {
		_spec.SetField(board.FieldContent, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{board.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	bu.mutation.done = true
	return n, nil
}

// BoardUpdateOne is the builder for updating a single Board entity.
type BoardUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BoardMutation
}

// SetTitle sets the "title" field.
func (buo *BoardUpdateOne) SetTitle(s string) *BoardUpdateOne {
	buo.mutation.SetTitle(s)
	return buo
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (buo *BoardUpdateOne) SetNillableTitle(s *string) *BoardUpdateOne {
	if s != nil {
		buo.SetTitle(*s)
	}
	return buo
}

// SetContent sets the "content" field.
func (buo *BoardUpdateOne) SetContent(s string) *BoardUpdateOne {
	buo.mutation.SetContent(s)
	return buo
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (buo *BoardUpdateOne) SetNillableContent(s *string) *BoardUpdateOne {
	if s != nil {
		buo.SetContent(*s)
	}
	return buo
}

// Mutation returns the BoardMutation object of the builder.
func (buo *BoardUpdateOne) Mutation() *BoardMutation {
	return buo.mutation
}

// Where appends a list predicates to the BoardUpdate builder.
func (buo *BoardUpdateOne) Where(ps ...predicate.Board) *BoardUpdateOne {
	buo.mutation.Where(ps...)
	return buo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (buo *BoardUpdateOne) Select(field string, fields ...string) *BoardUpdateOne {
	buo.fields = append([]string{field}, fields...)
	return buo
}

// Save executes the query and returns the updated Board entity.
func (buo *BoardUpdateOne) Save(ctx context.Context) (*Board, error) {
	return withHooks(ctx, buo.sqlSave, buo.mutation, buo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (buo *BoardUpdateOne) SaveX(ctx context.Context) *Board {
	node, err := buo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (buo *BoardUpdateOne) Exec(ctx context.Context) error {
	_, err := buo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (buo *BoardUpdateOne) ExecX(ctx context.Context) {
	if err := buo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (buo *BoardUpdateOne) sqlSave(ctx context.Context) (_node *Board, err error) {
	_spec := sqlgraph.NewUpdateSpec(board.Table, board.Columns, sqlgraph.NewFieldSpec(board.FieldID, field.TypeInt))
	id, ok := buo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Board.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := buo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, board.FieldID)
		for _, f := range fields {
			if !board.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != board.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := buo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := buo.mutation.Title(); ok {
		_spec.SetField(board.FieldTitle, field.TypeString, value)
	}
	if value, ok := buo.mutation.Content(); ok {
		_spec.SetField(board.FieldContent, field.TypeString, value)
	}
	_node = &Board{config: buo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, buo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{board.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	buo.mutation.done = true
	return _node, nil
}
