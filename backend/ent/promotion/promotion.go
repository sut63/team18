// Code generated by entc, DO NOT EDIT.

package promotion

const (
	// Label holds the string label denoting the promotion type in the database.
	Label = "promotion"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldPromotionName holds the string denoting the promotion_name field in the database.
	FieldPromotionName = "promotion_name"
	// FieldDiscount holds the string denoting the discount field in the database.
	FieldDiscount = "discount"

	// EdgeReserves holds the string denoting the reserves edge name in mutations.
	EdgeReserves = "reserves"
	// EdgeDatarooms holds the string denoting the datarooms edge name in mutations.
	EdgeDatarooms = "datarooms"

	// Table holds the table name of the promotion in the database.
	Table = "promotions"
	// ReservesTable is the table the holds the reserves relation/edge.
	ReservesTable = "reserve_rooms"
	// ReservesInverseTable is the table name for the ReserveRoom entity.
	// It exists in this package in order to avoid circular dependency with the "reserveroom" package.
	ReservesInverseTable = "reserve_rooms"
	// ReservesColumn is the table column denoting the reserves relation/edge.
	ReservesColumn = "promotion_id"
	// DataroomsTable is the table the holds the datarooms relation/edge.
	DataroomsTable = "data_rooms"
	// DataroomsInverseTable is the table name for the DataRoom entity.
	// It exists in this package in order to avoid circular dependency with the "dataroom" package.
	DataroomsInverseTable = "data_rooms"
	// DataroomsColumn is the table column denoting the datarooms relation/edge.
	DataroomsColumn = "promotion_id"
)

// Columns holds all SQL columns for promotion fields.
var Columns = []string{
	FieldID,
	FieldPromotionName,
	FieldDiscount,
}

var (
	// PromotionNameValidator is a validator for the "promotion_name" field. It is called by the builders before save.
	PromotionNameValidator func(string) error
	// DiscountValidator is a validator for the "discount" field. It is called by the builders before save.
	DiscountValidator func(float64) error
)