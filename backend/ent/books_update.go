// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"github.com/Sujitnapa29/app/ent/adult"
	"github.com/Sujitnapa29/app/ent/books"
	"github.com/Sujitnapa29/app/ent/customer"
	"github.com/Sujitnapa29/app/ent/kid"
	"github.com/Sujitnapa29/app/ent/predicate"
	"github.com/Sujitnapa29/app/ent/room"
	"github.com/Sujitnapa29/app/ent/roomamount"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// BooksUpdate is the builder for updating Books entities.
type BooksUpdate struct {
	config
	hooks      []Hook
	mutation   *BooksMutation
	predicates []predicate.Books
}

// Where adds a new predicate for the builder.
func (bu *BooksUpdate) Where(ps ...predicate.Books) *BooksUpdate {
	bu.predicates = append(bu.predicates, ps...)
	return bu
}

// SetCheckin sets the Checkin field.
func (bu *BooksUpdate) SetCheckin(t time.Time) *BooksUpdate {
	bu.mutation.SetCheckin(t)
	return bu
}

// SetCheckout sets the Checkout field.
func (bu *BooksUpdate) SetCheckout(t time.Time) *BooksUpdate {
	bu.mutation.SetCheckout(t)
	return bu
}

// SetAdultID sets the adult edge to Adult by id.
func (bu *BooksUpdate) SetAdultID(id int) *BooksUpdate {
	bu.mutation.SetAdultID(id)
	return bu
}

// SetNillableAdultID sets the adult edge to Adult by id if the given value is not nil.
func (bu *BooksUpdate) SetNillableAdultID(id *int) *BooksUpdate {
	if id != nil {
		bu = bu.SetAdultID(*id)
	}
	return bu
}

// SetAdult sets the adult edge to Adult.
func (bu *BooksUpdate) SetAdult(a *Adult) *BooksUpdate {
	return bu.SetAdultID(a.ID)
}

// SetKidID sets the kid edge to Kid by id.
func (bu *BooksUpdate) SetKidID(id int) *BooksUpdate {
	bu.mutation.SetKidID(id)
	return bu
}

// SetNillableKidID sets the kid edge to Kid by id if the given value is not nil.
func (bu *BooksUpdate) SetNillableKidID(id *int) *BooksUpdate {
	if id != nil {
		bu = bu.SetKidID(*id)
	}
	return bu
}

// SetKid sets the kid edge to Kid.
func (bu *BooksUpdate) SetKid(k *Kid) *BooksUpdate {
	return bu.SetKidID(k.ID)
}

// SetRoomamountID sets the roomamount edge to Roomamount by id.
func (bu *BooksUpdate) SetRoomamountID(id int) *BooksUpdate {
	bu.mutation.SetRoomamountID(id)
	return bu
}

// SetNillableRoomamountID sets the roomamount edge to Roomamount by id if the given value is not nil.
func (bu *BooksUpdate) SetNillableRoomamountID(id *int) *BooksUpdate {
	if id != nil {
		bu = bu.SetRoomamountID(*id)
	}
	return bu
}

// SetRoomamount sets the roomamount edge to Roomamount.
func (bu *BooksUpdate) SetRoomamount(r *Roomamount) *BooksUpdate {
	return bu.SetRoomamountID(r.ID)
}

// SetRoomID sets the room edge to Room by id.
func (bu *BooksUpdate) SetRoomID(id int) *BooksUpdate {
	bu.mutation.SetRoomID(id)
	return bu
}

// SetNillableRoomID sets the room edge to Room by id if the given value is not nil.
func (bu *BooksUpdate) SetNillableRoomID(id *int) *BooksUpdate {
	if id != nil {
		bu = bu.SetRoomID(*id)
	}
	return bu
}

// SetRoom sets the room edge to Room.
func (bu *BooksUpdate) SetRoom(r *Room) *BooksUpdate {
	return bu.SetRoomID(r.ID)
}

// SetCustomerID sets the customer edge to Customer by id.
func (bu *BooksUpdate) SetCustomerID(id int) *BooksUpdate {
	bu.mutation.SetCustomerID(id)
	return bu
}

// SetNillableCustomerID sets the customer edge to Customer by id if the given value is not nil.
func (bu *BooksUpdate) SetNillableCustomerID(id *int) *BooksUpdate {
	if id != nil {
		bu = bu.SetCustomerID(*id)
	}
	return bu
}

// SetCustomer sets the customer edge to Customer.
func (bu *BooksUpdate) SetCustomer(c *Customer) *BooksUpdate {
	return bu.SetCustomerID(c.ID)
}

// Mutation returns the BooksMutation object of the builder.
func (bu *BooksUpdate) Mutation() *BooksMutation {
	return bu.mutation
}

// ClearAdult clears the adult edge to Adult.
func (bu *BooksUpdate) ClearAdult() *BooksUpdate {
	bu.mutation.ClearAdult()
	return bu
}

// ClearKid clears the kid edge to Kid.
func (bu *BooksUpdate) ClearKid() *BooksUpdate {
	bu.mutation.ClearKid()
	return bu
}

// ClearRoomamount clears the roomamount edge to Roomamount.
func (bu *BooksUpdate) ClearRoomamount() *BooksUpdate {
	bu.mutation.ClearRoomamount()
	return bu
}

// ClearRoom clears the room edge to Room.
func (bu *BooksUpdate) ClearRoom() *BooksUpdate {
	bu.mutation.ClearRoom()
	return bu
}

// ClearCustomer clears the customer edge to Customer.
func (bu *BooksUpdate) ClearCustomer() *BooksUpdate {
	bu.mutation.ClearCustomer()
	return bu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (bu *BooksUpdate) Save(ctx context.Context) (int, error) {

	var (
		err      error
		affected int
	)
	if len(bu.hooks) == 0 {
		affected, err = bu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BooksMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			bu.mutation = mutation
			affected, err = bu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(bu.hooks) - 1; i >= 0; i-- {
			mut = bu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (bu *BooksUpdate) SaveX(ctx context.Context) int {
	affected, err := bu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bu *BooksUpdate) Exec(ctx context.Context) error {
	_, err := bu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bu *BooksUpdate) ExecX(ctx context.Context) {
	if err := bu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (bu *BooksUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   books.Table,
			Columns: books.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: books.FieldID,
			},
		},
	}
	if ps := bu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bu.mutation.Checkin(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: books.FieldCheckin,
		})
	}
	if value, ok := bu.mutation.Checkout(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: books.FieldCheckout,
		})
	}
	if bu.mutation.AdultCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   books.AdultTable,
			Columns: []string{books.AdultColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: adult.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.AdultIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   books.AdultTable,
			Columns: []string{books.AdultColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: adult.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if bu.mutation.KidCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   books.KidTable,
			Columns: []string{books.KidColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: kid.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.KidIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   books.KidTable,
			Columns: []string{books.KidColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: kid.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if bu.mutation.RoomamountCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   books.RoomamountTable,
			Columns: []string{books.RoomamountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: roomamount.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.RoomamountIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   books.RoomamountTable,
			Columns: []string{books.RoomamountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: roomamount.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if bu.mutation.RoomCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   books.RoomTable,
			Columns: []string{books.RoomColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: room.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.RoomIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   books.RoomTable,
			Columns: []string{books.RoomColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: room.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if bu.mutation.CustomerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   books.CustomerTable,
			Columns: []string{books.CustomerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: customer.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.CustomerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   books.CustomerTable,
			Columns: []string{books.CustomerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: customer.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{books.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// BooksUpdateOne is the builder for updating a single Books entity.
type BooksUpdateOne struct {
	config
	hooks    []Hook
	mutation *BooksMutation
}

// SetCheckin sets the Checkin field.
func (buo *BooksUpdateOne) SetCheckin(t time.Time) *BooksUpdateOne {
	buo.mutation.SetCheckin(t)
	return buo
}

// SetCheckout sets the Checkout field.
func (buo *BooksUpdateOne) SetCheckout(t time.Time) *BooksUpdateOne {
	buo.mutation.SetCheckout(t)
	return buo
}

// SetAdultID sets the adult edge to Adult by id.
func (buo *BooksUpdateOne) SetAdultID(id int) *BooksUpdateOne {
	buo.mutation.SetAdultID(id)
	return buo
}

// SetNillableAdultID sets the adult edge to Adult by id if the given value is not nil.
func (buo *BooksUpdateOne) SetNillableAdultID(id *int) *BooksUpdateOne {
	if id != nil {
		buo = buo.SetAdultID(*id)
	}
	return buo
}

// SetAdult sets the adult edge to Adult.
func (buo *BooksUpdateOne) SetAdult(a *Adult) *BooksUpdateOne {
	return buo.SetAdultID(a.ID)
}

// SetKidID sets the kid edge to Kid by id.
func (buo *BooksUpdateOne) SetKidID(id int) *BooksUpdateOne {
	buo.mutation.SetKidID(id)
	return buo
}

// SetNillableKidID sets the kid edge to Kid by id if the given value is not nil.
func (buo *BooksUpdateOne) SetNillableKidID(id *int) *BooksUpdateOne {
	if id != nil {
		buo = buo.SetKidID(*id)
	}
	return buo
}

// SetKid sets the kid edge to Kid.
func (buo *BooksUpdateOne) SetKid(k *Kid) *BooksUpdateOne {
	return buo.SetKidID(k.ID)
}

// SetRoomamountID sets the roomamount edge to Roomamount by id.
func (buo *BooksUpdateOne) SetRoomamountID(id int) *BooksUpdateOne {
	buo.mutation.SetRoomamountID(id)
	return buo
}

// SetNillableRoomamountID sets the roomamount edge to Roomamount by id if the given value is not nil.
func (buo *BooksUpdateOne) SetNillableRoomamountID(id *int) *BooksUpdateOne {
	if id != nil {
		buo = buo.SetRoomamountID(*id)
	}
	return buo
}

// SetRoomamount sets the roomamount edge to Roomamount.
func (buo *BooksUpdateOne) SetRoomamount(r *Roomamount) *BooksUpdateOne {
	return buo.SetRoomamountID(r.ID)
}

// SetRoomID sets the room edge to Room by id.
func (buo *BooksUpdateOne) SetRoomID(id int) *BooksUpdateOne {
	buo.mutation.SetRoomID(id)
	return buo
}

// SetNillableRoomID sets the room edge to Room by id if the given value is not nil.
func (buo *BooksUpdateOne) SetNillableRoomID(id *int) *BooksUpdateOne {
	if id != nil {
		buo = buo.SetRoomID(*id)
	}
	return buo
}

// SetRoom sets the room edge to Room.
func (buo *BooksUpdateOne) SetRoom(r *Room) *BooksUpdateOne {
	return buo.SetRoomID(r.ID)
}

// SetCustomerID sets the customer edge to Customer by id.
func (buo *BooksUpdateOne) SetCustomerID(id int) *BooksUpdateOne {
	buo.mutation.SetCustomerID(id)
	return buo
}

// SetNillableCustomerID sets the customer edge to Customer by id if the given value is not nil.
func (buo *BooksUpdateOne) SetNillableCustomerID(id *int) *BooksUpdateOne {
	if id != nil {
		buo = buo.SetCustomerID(*id)
	}
	return buo
}

// SetCustomer sets the customer edge to Customer.
func (buo *BooksUpdateOne) SetCustomer(c *Customer) *BooksUpdateOne {
	return buo.SetCustomerID(c.ID)
}

// Mutation returns the BooksMutation object of the builder.
func (buo *BooksUpdateOne) Mutation() *BooksMutation {
	return buo.mutation
}

// ClearAdult clears the adult edge to Adult.
func (buo *BooksUpdateOne) ClearAdult() *BooksUpdateOne {
	buo.mutation.ClearAdult()
	return buo
}

// ClearKid clears the kid edge to Kid.
func (buo *BooksUpdateOne) ClearKid() *BooksUpdateOne {
	buo.mutation.ClearKid()
	return buo
}

// ClearRoomamount clears the roomamount edge to Roomamount.
func (buo *BooksUpdateOne) ClearRoomamount() *BooksUpdateOne {
	buo.mutation.ClearRoomamount()
	return buo
}

// ClearRoom clears the room edge to Room.
func (buo *BooksUpdateOne) ClearRoom() *BooksUpdateOne {
	buo.mutation.ClearRoom()
	return buo
}

// ClearCustomer clears the customer edge to Customer.
func (buo *BooksUpdateOne) ClearCustomer() *BooksUpdateOne {
	buo.mutation.ClearCustomer()
	return buo
}

// Save executes the query and returns the updated entity.
func (buo *BooksUpdateOne) Save(ctx context.Context) (*Books, error) {

	var (
		err  error
		node *Books
	)
	if len(buo.hooks) == 0 {
		node, err = buo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BooksMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			buo.mutation = mutation
			node, err = buo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(buo.hooks) - 1; i >= 0; i-- {
			mut = buo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, buo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (buo *BooksUpdateOne) SaveX(ctx context.Context) *Books {
	b, err := buo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return b
}

// Exec executes the query on the entity.
func (buo *BooksUpdateOne) Exec(ctx context.Context) error {
	_, err := buo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (buo *BooksUpdateOne) ExecX(ctx context.Context) {
	if err := buo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (buo *BooksUpdateOne) sqlSave(ctx context.Context) (b *Books, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   books.Table,
			Columns: books.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: books.FieldID,
			},
		},
	}
	id, ok := buo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Books.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := buo.mutation.Checkin(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: books.FieldCheckin,
		})
	}
	if value, ok := buo.mutation.Checkout(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: books.FieldCheckout,
		})
	}
	if buo.mutation.AdultCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   books.AdultTable,
			Columns: []string{books.AdultColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: adult.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.AdultIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   books.AdultTable,
			Columns: []string{books.AdultColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: adult.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if buo.mutation.KidCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   books.KidTable,
			Columns: []string{books.KidColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: kid.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.KidIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   books.KidTable,
			Columns: []string{books.KidColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: kid.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if buo.mutation.RoomamountCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   books.RoomamountTable,
			Columns: []string{books.RoomamountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: roomamount.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.RoomamountIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   books.RoomamountTable,
			Columns: []string{books.RoomamountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: roomamount.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if buo.mutation.RoomCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   books.RoomTable,
			Columns: []string{books.RoomColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: room.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.RoomIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   books.RoomTable,
			Columns: []string{books.RoomColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: room.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if buo.mutation.CustomerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   books.CustomerTable,
			Columns: []string{books.CustomerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: customer.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.CustomerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   books.CustomerTable,
			Columns: []string{books.CustomerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: customer.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	b = &Books{config: buo.config}
	_spec.Assign = b.assignValues
	_spec.ScanValues = b.scanValues()
	if err = sqlgraph.UpdateNode(ctx, buo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{books.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return b, nil
}