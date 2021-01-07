// Code generated by entc, DO NOT EDIT.

package typeroom

const (
	// Label holds the string label denoting the typeroom type in the database.
	Label = "type_room"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTypeName holds the string denoting the type_name field in the database.
	FieldTypeName = "type_name"

	// EdgeDatarooms holds the string denoting the datarooms edge name in mutations.
	EdgeDatarooms = "datarooms"

	// Table holds the table name of the typeroom in the database.
	Table = "type_rooms"
	// DataroomsTable is the table the holds the datarooms relation/edge.
	DataroomsTable = "data_rooms"
	// DataroomsInverseTable is the table name for the DataRoom entity.
	// It exists in this package in order to avoid circular dependency with the "dataroom" package.
	DataroomsInverseTable = "data_rooms"
	// DataroomsColumn is the table column denoting the datarooms relation/edge.
	DataroomsColumn = "typeroom_id"
)

// Columns holds all SQL columns for typeroom fields.
var Columns = []string{
	FieldID,
	FieldTypeName,
}

var (
	// TypeNameValidator is a validator for the "type_name" field. It is called by the builders before save.
	TypeNameValidator func(string) error
)