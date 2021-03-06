// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/Sujitnapa29/app/ent/adult"
	"github.com/Sujitnapa29/app/ent/books"
	"github.com/Sujitnapa29/app/ent/predicate"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// AdultUpdate is the builder for updating Adult entities.
type AdultUpdate struct {
	config
	hooks      []Hook
	mutation   *AdultMutation
	predicates []predicate.Adult
}

// Where adds a new predicate for the builder.
func (au *AdultUpdate) Where(ps ...predicate.Adult) *AdultUpdate {
	au.predicates = append(au.predicates, ps...)
	return au
}

// SetAmount sets the Amount field.
func (au *AdultUpdate) SetAmount(i int) *AdultUpdate {
	au.mutation.ResetAmount()
	au.mutation.SetAmount(i)
	return au
}

// AddAmount adds i to Amount.
func (au *AdultUpdate) AddAmount(i int) *AdultUpdate {
	au.mutation.AddAmount(i)
	return au
}

// AddBookIDs adds the books edge to Books by ids.
func (au *AdultUpdate) AddBookIDs(ids ...int) *AdultUpdate {
	au.mutation.AddBookIDs(ids...)
	return au
}

// AddBooks adds the books edges to Books.
func (au *AdultUpdate) AddBooks(b ...*Books) *AdultUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return au.AddBookIDs(ids...)
}

// Mutation returns the AdultMutation object of the builder.
func (au *AdultUpdate) Mutation() *AdultMutation {
	return au.mutation
}

// RemoveBookIDs removes the books edge to Books by ids.
func (au *AdultUpdate) RemoveBookIDs(ids ...int) *AdultUpdate {
	au.mutation.RemoveBookIDs(ids...)
	return au
}

// RemoveBooks removes books edges to Books.
func (au *AdultUpdate) RemoveBooks(b ...*Books) *AdultUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return au.RemoveBookIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (au *AdultUpdate) Save(ctx context.Context) (int, error) {

	var (
		err      error
		affected int
	)
	if len(au.hooks) == 0 {
		affected, err = au.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AdultMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			au.mutation = mutation
			affected, err = au.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(au.hooks) - 1; i >= 0; i-- {
			mut = au.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, au.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (au *AdultUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AdultUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AdultUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

func (au *AdultUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   adult.Table,
			Columns: adult.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: adult.FieldID,
			},
		},
	}
	if ps := au.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.Amount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: adult.FieldAmount,
		})
	}
	if value, ok := au.mutation.AddedAmount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: adult.FieldAmount,
		})
	}
	if nodes := au.mutation.RemovedBooksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   adult.BooksTable,
			Columns: []string{adult.BooksColumn},
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
	if nodes := au.mutation.BooksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   adult.BooksTable,
			Columns: []string{adult.BooksColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{adult.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// AdultUpdateOne is the builder for updating a single Adult entity.
type AdultUpdateOne struct {
	config
	hooks    []Hook
	mutation *AdultMutation
}

// SetAmount sets the Amount field.
func (auo *AdultUpdateOne) SetAmount(i int) *AdultUpdateOne {
	auo.mutation.ResetAmount()
	auo.mutation.SetAmount(i)
	return auo
}

// AddAmount adds i to Amount.
func (auo *AdultUpdateOne) AddAmount(i int) *AdultUpdateOne {
	auo.mutation.AddAmount(i)
	return auo
}

// AddBookIDs adds the books edge to Books by ids.
func (auo *AdultUpdateOne) AddBookIDs(ids ...int) *AdultUpdateOne {
	auo.mutation.AddBookIDs(ids...)
	return auo
}

// AddBooks adds the books edges to Books.
func (auo *AdultUpdateOne) AddBooks(b ...*Books) *AdultUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return auo.AddBookIDs(ids...)
}

// Mutation returns the AdultMutation object of the builder.
func (auo *AdultUpdateOne) Mutation() *AdultMutation {
	return auo.mutation
}

// RemoveBookIDs removes the books edge to Books by ids.
func (auo *AdultUpdateOne) RemoveBookIDs(ids ...int) *AdultUpdateOne {
	auo.mutation.RemoveBookIDs(ids...)
	return auo
}

// RemoveBooks removes books edges to Books.
func (auo *AdultUpdateOne) RemoveBooks(b ...*Books) *AdultUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return auo.RemoveBookIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (auo *AdultUpdateOne) Save(ctx context.Context) (*Adult, error) {

	var (
		err  error
		node *Adult
	)
	if len(auo.hooks) == 0 {
		node, err = auo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AdultMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			auo.mutation = mutation
			node, err = auo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(auo.hooks) - 1; i >= 0; i-- {
			mut = auo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, auo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AdultUpdateOne) SaveX(ctx context.Context) *Adult {
	a, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return a
}

// Exec executes the query on the entity.
func (auo *AdultUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AdultUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (auo *AdultUpdateOne) sqlSave(ctx context.Context) (a *Adult, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   adult.Table,
			Columns: adult.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: adult.FieldID,
			},
		},
	}
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Adult.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := auo.mutation.Amount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: adult.FieldAmount,
		})
	}
	if value, ok := auo.mutation.AddedAmount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: adult.FieldAmount,
		})
	}
	if nodes := auo.mutation.RemovedBooksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   adult.BooksTable,
			Columns: []string{adult.BooksColumn},
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
	if nodes := auo.mutation.BooksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   adult.BooksTable,
			Columns: []string{adult.BooksColumn},
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
	a = &Adult{config: auo.config}
	_spec.Assign = a.assignValues
	_spec.ScanValues = a.scanValues()
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{adult.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return a, nil
}
