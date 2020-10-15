// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/Sujitnapa29/app/ent/room"
	"github.com/facebookincubator/ent/dialect/sql"
)

// Room is the model entity for the Room schema.
type Room struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Type holds the value of the "Type" field.
	Type string `json:"Type,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RoomQuery when eager-loading is set.
	Edges RoomEdges `json:"edges"`
}

// RoomEdges holds the relations/edges for other nodes in the graph.
type RoomEdges struct {
	// Books holds the value of the books edge.
	Books []*Books
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// BooksOrErr returns the Books value or an error if the edge
// was not loaded in eager-loading.
func (e RoomEdges) BooksOrErr() ([]*Books, error) {
	if e.loadedTypes[0] {
		return e.Books, nil
	}
	return nil, &NotLoadedError{edge: "books"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Room) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // Type
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Room fields.
func (r *Room) assignValues(values ...interface{}) error {
	if m, n := len(values), len(room.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	r.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field Type", values[0])
	} else if value.Valid {
		r.Type = value.String
	}
	return nil
}

// QueryBooks queries the books edge of the Room.
func (r *Room) QueryBooks() *BooksQuery {
	return (&RoomClient{config: r.config}).QueryBooks(r)
}

// Update returns a builder for updating this Room.
// Note that, you need to call Room.Unwrap() before calling this method, if this Room
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Room) Update() *RoomUpdateOne {
	return (&RoomClient{config: r.config}).UpdateOne(r)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (r *Room) Unwrap() *Room {
	tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Room is not a transactional entity")
	}
	r.config.driver = tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Room) String() string {
	var builder strings.Builder
	builder.WriteString("Room(")
	builder.WriteString(fmt.Sprintf("id=%v", r.ID))
	builder.WriteString(", Type=")
	builder.WriteString(r.Type)
	builder.WriteByte(')')
	return builder.String()
}

// Rooms is a parsable slice of Room.
type Rooms []*Room

func (r Rooms) config(cfg config) {
	for _i := range r {
		r[_i].config = cfg
	}
}
