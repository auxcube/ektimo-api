// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/auxcube/ektimo-api/ent/predicate"
	"github.com/auxcube/ektimo-api/ent/textquestion"
)

// TextQuestionUpdate is the builder for updating TextQuestion entities.
type TextQuestionUpdate struct {
	config
	hooks    []Hook
	mutation *TextQuestionMutation
}

// Where appends a list predicates to the TextQuestionUpdate builder.
func (tqu *TextQuestionUpdate) Where(ps ...predicate.TextQuestion) *TextQuestionUpdate {
	tqu.mutation.Where(ps...)
	return tqu
}

// SetUpdatedAt sets the "updated_at" field.
func (tqu *TextQuestionUpdate) SetUpdatedAt(t time.Time) *TextQuestionUpdate {
	tqu.mutation.SetUpdatedAt(t)
	return tqu
}

// SetHumanID sets the "human_id" field.
func (tqu *TextQuestionUpdate) SetHumanID(s string) *TextQuestionUpdate {
	tqu.mutation.SetHumanID(s)
	return tqu
}

// SetText sets the "text" field.
func (tqu *TextQuestionUpdate) SetText(s string) *TextQuestionUpdate {
	tqu.mutation.SetText(s)
	return tqu
}

// SetAnswer sets the "answer" field.
func (tqu *TextQuestionUpdate) SetAnswer(s string) *TextQuestionUpdate {
	tqu.mutation.SetAnswer(s)
	return tqu
}

// SetNillableAnswer sets the "answer" field if the given value is not nil.
func (tqu *TextQuestionUpdate) SetNillableAnswer(s *string) *TextQuestionUpdate {
	if s != nil {
		tqu.SetAnswer(*s)
	}
	return tqu
}

// ClearAnswer clears the value of the "answer" field.
func (tqu *TextQuestionUpdate) ClearAnswer() *TextQuestionUpdate {
	tqu.mutation.ClearAnswer()
	return tqu
}

// Mutation returns the TextQuestionMutation object of the builder.
func (tqu *TextQuestionUpdate) Mutation() *TextQuestionMutation {
	return tqu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tqu *TextQuestionUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	tqu.defaults()
	if len(tqu.hooks) == 0 {
		if err = tqu.check(); err != nil {
			return 0, err
		}
		affected, err = tqu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TextQuestionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tqu.check(); err != nil {
				return 0, err
			}
			tqu.mutation = mutation
			affected, err = tqu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tqu.hooks) - 1; i >= 0; i-- {
			if tqu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tqu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tqu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tqu *TextQuestionUpdate) SaveX(ctx context.Context) int {
	affected, err := tqu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tqu *TextQuestionUpdate) Exec(ctx context.Context) error {
	_, err := tqu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tqu *TextQuestionUpdate) ExecX(ctx context.Context) {
	if err := tqu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tqu *TextQuestionUpdate) defaults() {
	if _, ok := tqu.mutation.UpdatedAt(); !ok {
		v := textquestion.UpdateDefaultUpdatedAt()
		tqu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tqu *TextQuestionUpdate) check() error {
	if v, ok := tqu.mutation.HumanID(); ok {
		if err := textquestion.HumanIDValidator(v); err != nil {
			return &ValidationError{Name: "human_id", err: fmt.Errorf(`ent: validator failed for field "TextQuestion.human_id": %w`, err)}
		}
	}
	if v, ok := tqu.mutation.Text(); ok {
		if err := textquestion.TextValidator(v); err != nil {
			return &ValidationError{Name: "text", err: fmt.Errorf(`ent: validator failed for field "TextQuestion.text": %w`, err)}
		}
	}
	return nil
}

func (tqu *TextQuestionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   textquestion.Table,
			Columns: textquestion.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: textquestion.FieldID,
			},
		},
	}
	if ps := tqu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tqu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: textquestion.FieldUpdatedAt,
		})
	}
	if value, ok := tqu.mutation.HumanID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: textquestion.FieldHumanID,
		})
	}
	if value, ok := tqu.mutation.Text(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: textquestion.FieldText,
		})
	}
	if value, ok := tqu.mutation.Answer(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: textquestion.FieldAnswer,
		})
	}
	if tqu.mutation.AnswerCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: textquestion.FieldAnswer,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tqu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{textquestion.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// TextQuestionUpdateOne is the builder for updating a single TextQuestion entity.
type TextQuestionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TextQuestionMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (tquo *TextQuestionUpdateOne) SetUpdatedAt(t time.Time) *TextQuestionUpdateOne {
	tquo.mutation.SetUpdatedAt(t)
	return tquo
}

// SetHumanID sets the "human_id" field.
func (tquo *TextQuestionUpdateOne) SetHumanID(s string) *TextQuestionUpdateOne {
	tquo.mutation.SetHumanID(s)
	return tquo
}

// SetText sets the "text" field.
func (tquo *TextQuestionUpdateOne) SetText(s string) *TextQuestionUpdateOne {
	tquo.mutation.SetText(s)
	return tquo
}

// SetAnswer sets the "answer" field.
func (tquo *TextQuestionUpdateOne) SetAnswer(s string) *TextQuestionUpdateOne {
	tquo.mutation.SetAnswer(s)
	return tquo
}

// SetNillableAnswer sets the "answer" field if the given value is not nil.
func (tquo *TextQuestionUpdateOne) SetNillableAnswer(s *string) *TextQuestionUpdateOne {
	if s != nil {
		tquo.SetAnswer(*s)
	}
	return tquo
}

// ClearAnswer clears the value of the "answer" field.
func (tquo *TextQuestionUpdateOne) ClearAnswer() *TextQuestionUpdateOne {
	tquo.mutation.ClearAnswer()
	return tquo
}

// Mutation returns the TextQuestionMutation object of the builder.
func (tquo *TextQuestionUpdateOne) Mutation() *TextQuestionMutation {
	return tquo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tquo *TextQuestionUpdateOne) Select(field string, fields ...string) *TextQuestionUpdateOne {
	tquo.fields = append([]string{field}, fields...)
	return tquo
}

// Save executes the query and returns the updated TextQuestion entity.
func (tquo *TextQuestionUpdateOne) Save(ctx context.Context) (*TextQuestion, error) {
	var (
		err  error
		node *TextQuestion
	)
	tquo.defaults()
	if len(tquo.hooks) == 0 {
		if err = tquo.check(); err != nil {
			return nil, err
		}
		node, err = tquo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TextQuestionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tquo.check(); err != nil {
				return nil, err
			}
			tquo.mutation = mutation
			node, err = tquo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tquo.hooks) - 1; i >= 0; i-- {
			if tquo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tquo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tquo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tquo *TextQuestionUpdateOne) SaveX(ctx context.Context) *TextQuestion {
	node, err := tquo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tquo *TextQuestionUpdateOne) Exec(ctx context.Context) error {
	_, err := tquo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tquo *TextQuestionUpdateOne) ExecX(ctx context.Context) {
	if err := tquo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tquo *TextQuestionUpdateOne) defaults() {
	if _, ok := tquo.mutation.UpdatedAt(); !ok {
		v := textquestion.UpdateDefaultUpdatedAt()
		tquo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tquo *TextQuestionUpdateOne) check() error {
	if v, ok := tquo.mutation.HumanID(); ok {
		if err := textquestion.HumanIDValidator(v); err != nil {
			return &ValidationError{Name: "human_id", err: fmt.Errorf(`ent: validator failed for field "TextQuestion.human_id": %w`, err)}
		}
	}
	if v, ok := tquo.mutation.Text(); ok {
		if err := textquestion.TextValidator(v); err != nil {
			return &ValidationError{Name: "text", err: fmt.Errorf(`ent: validator failed for field "TextQuestion.text": %w`, err)}
		}
	}
	return nil
}

func (tquo *TextQuestionUpdateOne) sqlSave(ctx context.Context) (_node *TextQuestion, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   textquestion.Table,
			Columns: textquestion.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: textquestion.FieldID,
			},
		},
	}
	id, ok := tquo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "TextQuestion.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tquo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, textquestion.FieldID)
		for _, f := range fields {
			if !textquestion.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != textquestion.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tquo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tquo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: textquestion.FieldUpdatedAt,
		})
	}
	if value, ok := tquo.mutation.HumanID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: textquestion.FieldHumanID,
		})
	}
	if value, ok := tquo.mutation.Text(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: textquestion.FieldText,
		})
	}
	if value, ok := tquo.mutation.Answer(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: textquestion.FieldAnswer,
		})
	}
	if tquo.mutation.AnswerCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: textquestion.FieldAnswer,
		})
	}
	_node = &TextQuestion{config: tquo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tquo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{textquestion.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}