// Code generated by entc, DO NOT EDIT.

package checkout

const (
	// Label holds the string label denoting the checkout type in the database.
	Label = "checkout"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCheckoutDate holds the string denoting the checkout_date field in the database.
	FieldCheckoutDate = "checkout_date"

	// EdgeStatuss holds the string denoting the statuss edge name in mutations.
	EdgeStatuss = "statuss"
	// EdgeCounterstaffs holds the string denoting the counterstaffs edge name in mutations.
	EdgeCounterstaffs = "counterstaffs"
	// EdgeCheckins holds the string denoting the checkins edge name in mutations.
	EdgeCheckins = "checkins"

	// Table holds the table name of the checkout in the database.
	Table = "checkouts"
	// StatussTable is the table the holds the statuss relation/edge.
	StatussTable = "checkouts"
	// StatussInverseTable is the table name for the Status entity.
	// It exists in this package in order to avoid circular dependency with the "status" package.
	StatussInverseTable = "status"
	// StatussColumn is the table column denoting the statuss relation/edge.
	StatussColumn = "status_checkouts"
	// CounterstaffsTable is the table the holds the counterstaffs relation/edge.
	CounterstaffsTable = "checkouts"
	// CounterstaffsInverseTable is the table name for the CounterStaff entity.
	// It exists in this package in order to avoid circular dependency with the "counterstaff" package.
	CounterstaffsInverseTable = "counter_staffs"
	// CounterstaffsColumn is the table column denoting the counterstaffs relation/edge.
	CounterstaffsColumn = "staff_id"
	// CheckinsTable is the table the holds the checkins relation/edge.
	CheckinsTable = "checkouts"
	// CheckinsInverseTable is the table name for the CheckIn entity.
	// It exists in this package in order to avoid circular dependency with the "checkin" package.
	CheckinsInverseTable = "check_ins"
	// CheckinsColumn is the table column denoting the checkins relation/edge.
	CheckinsColumn = "check_in_checkouts"
)

// Columns holds all SQL columns for checkout fields.
var Columns = []string{
	FieldID,
	FieldCheckoutDate,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Checkout type.
var ForeignKeys = []string{
	"check_in_checkouts",
	"staff_id",
	"status_checkouts",
}