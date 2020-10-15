// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/Sujitnapa29/app/ent/roomamount"
	"github.com/facebookincubator/ent/dialect/sql"
)

// Roomamount is the model entity for the Roomamount schema.
type Roomamount struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Amount holds the value of the "Amount" field.
	Amount int `json:"Amount,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RoomamountQuery when eager-loading is set.
	Edges RoomamountEdges `json:"edges"`
}

// RoomamountEdges holds the relations/edges for other nodes in the graph.
type RoomamountEdges struct {
	// Books holds the value of the books edge.
	Books []*Books
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// BooksOrErr returns the Books value or an error if the edge
// was not loaded in eager-loading.
func (e RoomamountEdges) BooksOrErr() ([]*Books, error) {
	if e.loadedTypes[0] {
		return e.Books, nil
	}
	return nil, &NotLoadedError{edge: "books"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Roomamount) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // id
		&sql.NullInt64{}, // Amount
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Roomamount fields.
func (r *Roomamount) assignValues(values ...interface{}) error {
	if m, n := len(values), len(roomamount.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	r.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field Amount", values[0])
	} else if value.Valid {
		r.Amount = int(value.Int64)
	}
	return nil
}

// QueryBooks queries the books edge of the Roomamount.
func (r *Roomamount) QueryBooks() *BooksQuery {
	return (&RoomamountClient{config: r.config}).QueryBooks(r)
}

// Update returns a builder for updating this Roomamount.
// Note that, you need to call Roomamount.Unwrap() before calling this method, if this Roomamount
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Roomamount) Update() *RoomamountUpdateOne {
	return (&RoomamountClient{config: r.config}).UpdateOne(r)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (r *Roomamount) Unwrap() *Roomamount {
	tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Roomamount is not a transactional entity")
	}
	r.config.driver = tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Roomamount) String() string {
	var builder strings.Builder
	builder.WriteString("Roomamount(")
	builder.WriteString(fmt.Sprintf("id=%v", r.ID))
	builder.WriteString(", Amount=")
	builder.WriteString(fmt.Sprintf("%v", r.Amount))
	builder.WriteByte(')')
	return builder.String()
}

// Roomamounts is a parsable slice of Roomamount.
type Roomamounts []*Roomamount

func (r Roomamounts) config(cfg config) {
	for _i := range r {
		r[_i].config = cfg
	}
}
