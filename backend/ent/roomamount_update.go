// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/Sujitnapa29/app/ent/books"
	"github.com/Sujitnapa29/app/ent/predicate"
	"github.com/Sujitnapa29/app/ent/roomamount"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// RoomamountUpdate is the builder for updating Roomamount entities.
type RoomamountUpdate struct {
	config
	hooks      []Hook
	mutation   *RoomamountMutation
	predicates []predicate.Roomamount
}

// Where adds a new predicate for the builder.
func (ru *RoomamountUpdate) Where(ps ...predicate.Roomamount) *RoomamountUpdate {
	ru.predicates = append(ru.predicates, ps...)
	return ru
}

// SetAmount sets the Amount field.
func (ru *RoomamountUpdate) SetAmount(i int) *RoomamountUpdate {
	ru.mutation.ResetAmount()
	ru.mutation.SetAmount(i)
	return ru
}

// AddAmount adds i to Amount.
func (ru *RoomamountUpdate) AddAmount(i int) *RoomamountUpdate {
	ru.mutation.AddAmount(i)
	return ru
}

// AddBookIDs adds the books edge to Books by ids.
func (ru *RoomamountUpdate) AddBookIDs(ids ...int) *RoomamountUpdate {
	ru.mutation.AddBookIDs(ids...)
	return ru
}

// AddBooks adds the books edges to Books.
func (ru *RoomamountUpdate) AddBooks(b ...*Books) *RoomamountUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return ru.AddBookIDs(ids...)
}

// Mutation returns the RoomamountMutation object of the builder.
func (ru *RoomamountUpdate) Mutation() *RoomamountMutation {
	return ru.mutation
}

// RemoveBookIDs removes the books edge to Books by ids.
func (ru *RoomamountUpdate) RemoveBookIDs(ids ...int) *RoomamountUpdate {
	ru.mutation.RemoveBookIDs(ids...)
	return ru
}

// RemoveBooks removes books edges to Books.
func (ru *RoomamountUpdate) RemoveBooks(b ...*Books) *RoomamountUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return ru.RemoveBookIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (ru *RoomamountUpdate) Save(ctx context.Context) (int, error) {

	var (
		err      error
		affected int
	)
	if len(ru.hooks) == 0 {
		affected, err = ru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RoomamountMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ru.mutation = mutation
			affected, err = ru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ru.hooks) - 1; i >= 0; i-- {
			mut = ru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ru *RoomamountUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *RoomamountUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *RoomamountUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ru *RoomamountUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   roomamount.Table,
			Columns: roomamount.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: roomamount.FieldID,
			},
		},
	}
	if ps := ru.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.Amount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: roomamount.FieldAmount,
		})
	}
	if value, ok := ru.mutation.AddedAmount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: roomamount.FieldAmount,
		})
	}
	if nodes := ru.mutation.RemovedBooksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   roomamount.BooksTable,
			Columns: []string{roomamount.BooksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: books.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.BooksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   roomamount.BooksTable,
			Columns: []string{roomamount.BooksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: books.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{roomamount.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// RoomamountUpdateOne is the builder for updating a single Roomamount entity.
type RoomamountUpdateOne struct {
	config
	hooks    []Hook
	mutation *RoomamountMutation
}

// SetAmount sets the Amount field.
func (ruo *RoomamountUpdateOne) SetAmount(i int) *RoomamountUpdateOne {
	ruo.mutation.ResetAmount()
	ruo.mutation.SetAmount(i)
	return ruo
}

// AddAmount adds i to Amount.
func (ruo *RoomamountUpdateOne) AddAmount(i int) *RoomamountUpdateOne {
	ruo.mutation.AddAmount(i)
	return ruo
}

// AddBookIDs adds the books edge to Books by ids.
func (ruo *RoomamountUpdateOne) AddBookIDs(ids ...int) *RoomamountUpdateOne {
	ruo.mutation.AddBookIDs(ids...)
	return ruo
}

// AddBooks adds the books edges to Books.
func (ruo *RoomamountUpdateOne) AddBooks(b ...*Books) *RoomamountUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return ruo.AddBookIDs(ids...)
}

// Mutation returns the RoomamountMutation object of the builder.
func (ruo *RoomamountUpdateOne) Mutation() *RoomamountMutation {
	return ruo.mutation
}

// RemoveBookIDs removes the books edge to Books by ids.
func (ruo *RoomamountUpdateOne) RemoveBookIDs(ids ...int) *RoomamountUpdateOne {
	ruo.mutation.RemoveBookIDs(ids...)
	return ruo
}

// RemoveBooks removes books edges to Books.
func (ruo *RoomamountUpdateOne) RemoveBooks(b ...*Books) *RoomamountUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return ruo.RemoveBookIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (ruo *RoomamountUpdateOne) Save(ctx context.Context) (*Roomamount, error) {

	var (
		err  error
		node *Roomamount
	)
	if len(ruo.hooks) == 0 {
		node, err = ruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RoomamountMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ruo.mutation = mutation
			node, err = ruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ruo.hooks) - 1; i >= 0; i-- {
			mut = ruo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ruo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *RoomamountUpdateOne) SaveX(ctx context.Context) *Roomamount {
	r, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return r
}

// Exec executes the query on the entity.
func (ruo *RoomamountUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *RoomamountUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ruo *RoomamountUpdateOne) sqlSave(ctx context.Context) (r *Roomamount, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   roomamount.Table,
			Columns: roomamount.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: roomamount.FieldID,
			},
		},
	}
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Roomamount.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := ruo.mutation.Amount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: roomamount.FieldAmount,
		})
	}
	if value, ok := ruo.mutation.AddedAmount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: roomamount.FieldAmount,
		})
	}
	if nodes := ruo.mutation.RemovedBooksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   roomamount.BooksTable,
			Columns: []string{roomamount.BooksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: books.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.BooksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   roomamount.BooksTable,
			Columns: []string{roomamount.BooksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: books.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	r = &Roomamount{config: ruo.config}
	_spec.Assign = r.assignValues
	_spec.ScanValues = r.scanValues()
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{roomamount.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return r, nil
}
