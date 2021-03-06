// Code generated by entc, DO NOT EDIT.

package room

const (
	// Label holds the string label denoting the room type in the database.
	Label = "room"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"

	// EdgeBooks holds the string denoting the books edge name in mutations.
	EdgeBooks = "books"

	// Table holds the table name of the room in the database.
	Table = "rooms"
	// BooksTable is the table the holds the books relation/edge.
	BooksTable = "books"
	// BooksInverseTable is the table name for the Books entity.
	// It exists in this package in order to avoid circular dependency with the "books" package.
	BooksInverseTable = "books"
	// BooksColumn is the table column denoting the books relation/edge.
	BooksColumn = "room_books"
)

// Columns holds all SQL columns for room fields.
var Columns = []string{
	FieldID,
	FieldType,
}
