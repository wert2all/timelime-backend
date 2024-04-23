// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"timeline/backend/ent/predicate"
	"timeline/backend/ent/timeline"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TimelineDelete is the builder for deleting a Timeline entity.
type TimelineDelete struct {
	config
	hooks    []Hook
	mutation *TimelineMutation
}

// Where appends a list predicates to the TimelineDelete builder.
func (td *TimelineDelete) Where(ps ...predicate.Timeline) *TimelineDelete {
	td.mutation.Where(ps...)
	return td
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (td *TimelineDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, td.sqlExec, td.mutation, td.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (td *TimelineDelete) ExecX(ctx context.Context) int {
	n, err := td.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (td *TimelineDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(timeline.Table, sqlgraph.NewFieldSpec(timeline.FieldID, field.TypeInt))
	if ps := td.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, td.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	td.mutation.done = true
	return affected, err
}

// TimelineDeleteOne is the builder for deleting a single Timeline entity.
type TimelineDeleteOne struct {
	td *TimelineDelete
}

// Where appends a list predicates to the TimelineDelete builder.
func (tdo *TimelineDeleteOne) Where(ps ...predicate.Timeline) *TimelineDeleteOne {
	tdo.td.mutation.Where(ps...)
	return tdo
}

// Exec executes the deletion query.
func (tdo *TimelineDeleteOne) Exec(ctx context.Context) error {
	n, err := tdo.td.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{timeline.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (tdo *TimelineDeleteOne) ExecX(ctx context.Context) {
	if err := tdo.Exec(ctx); err != nil {
		panic(err)
	}
}
