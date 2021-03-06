// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/team18/app/ent/counterstaff"
	"github.com/team18/app/ent/dataroom"
	"github.com/team18/app/ent/furniture"
	"github.com/team18/app/ent/furnituredetail"
	"github.com/team18/app/ent/furnituretype"
)

// FurnitureDetail is the model entity for the FurnitureDetail schema.
type FurnitureDetail struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// DateAdd holds the value of the "date_add" field.
	DateAdd time.Time `json:"date_add,omitempty"`
	// FurnitureAmount holds the value of the "furniture_amount" field.
	FurnitureAmount int `json:"furniture_amount,omitempty"`
	// FurnitureColour holds the value of the "furniture_colour" field.
	FurnitureColour string `json:"furniture_colour,omitempty"`
	// Detail holds the value of the "detail" field.
	Detail string `json:"detail,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the FurnitureDetailQuery when eager-loading is set.
	Edges        FurnitureDetailEdges `json:"edges"`
	staff_id     *int
	room_id      *int
	furniture_id *int
	type_id      *int
}

// FurnitureDetailEdges holds the relations/edges for other nodes in the graph.
type FurnitureDetailEdges struct {
	// Fixs holds the value of the fixs edge.
	Fixs []*FixRoom
	// Furnitures holds the value of the furnitures edge.
	Furnitures *Furniture
	// Counterstaffs holds the value of the counterstaffs edge.
	Counterstaffs *CounterStaff
	// Types holds the value of the types edge.
	Types *FurnitureType
	// Rooms holds the value of the rooms edge.
	Rooms *DataRoom
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [5]bool
}

// FixsOrErr returns the Fixs value or an error if the edge
// was not loaded in eager-loading.
func (e FurnitureDetailEdges) FixsOrErr() ([]*FixRoom, error) {
	if e.loadedTypes[0] {
		return e.Fixs, nil
	}
	return nil, &NotLoadedError{edge: "fixs"}
}

// FurnituresOrErr returns the Furnitures value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e FurnitureDetailEdges) FurnituresOrErr() (*Furniture, error) {
	if e.loadedTypes[1] {
		if e.Furnitures == nil {
			// The edge furnitures was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: furniture.Label}
		}
		return e.Furnitures, nil
	}
	return nil, &NotLoadedError{edge: "furnitures"}
}

// CounterstaffsOrErr returns the Counterstaffs value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e FurnitureDetailEdges) CounterstaffsOrErr() (*CounterStaff, error) {
	if e.loadedTypes[2] {
		if e.Counterstaffs == nil {
			// The edge counterstaffs was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: counterstaff.Label}
		}
		return e.Counterstaffs, nil
	}
	return nil, &NotLoadedError{edge: "counterstaffs"}
}

// TypesOrErr returns the Types value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e FurnitureDetailEdges) TypesOrErr() (*FurnitureType, error) {
	if e.loadedTypes[3] {
		if e.Types == nil {
			// The edge types was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: furnituretype.Label}
		}
		return e.Types, nil
	}
	return nil, &NotLoadedError{edge: "types"}
}

// RoomsOrErr returns the Rooms value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e FurnitureDetailEdges) RoomsOrErr() (*DataRoom, error) {
	if e.loadedTypes[4] {
		if e.Rooms == nil {
			// The edge rooms was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: dataroom.Label}
		}
		return e.Rooms, nil
	}
	return nil, &NotLoadedError{edge: "rooms"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*FurnitureDetail) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullTime{},   // date_add
		&sql.NullInt64{},  // furniture_amount
		&sql.NullString{}, // furniture_colour
		&sql.NullString{}, // detail
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*FurnitureDetail) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // staff_id
		&sql.NullInt64{}, // room_id
		&sql.NullInt64{}, // furniture_id
		&sql.NullInt64{}, // type_id
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the FurnitureDetail fields.
func (fd *FurnitureDetail) assignValues(values ...interface{}) error {
	if m, n := len(values), len(furnituredetail.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	fd.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field date_add", values[0])
	} else if value.Valid {
		fd.DateAdd = value.Time
	}
	if value, ok := values[1].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field furniture_amount", values[1])
	} else if value.Valid {
		fd.FurnitureAmount = int(value.Int64)
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field furniture_colour", values[2])
	} else if value.Valid {
		fd.FurnitureColour = value.String
	}
	if value, ok := values[3].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field detail", values[3])
	} else if value.Valid {
		fd.Detail = value.String
	}
	values = values[4:]
	if len(values) == len(furnituredetail.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field staff_id", value)
		} else if value.Valid {
			fd.staff_id = new(int)
			*fd.staff_id = int(value.Int64)
		}
		if value, ok := values[1].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field room_id", value)
		} else if value.Valid {
			fd.room_id = new(int)
			*fd.room_id = int(value.Int64)
		}
		if value, ok := values[2].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field furniture_id", value)
		} else if value.Valid {
			fd.furniture_id = new(int)
			*fd.furniture_id = int(value.Int64)
		}
		if value, ok := values[3].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field type_id", value)
		} else if value.Valid {
			fd.type_id = new(int)
			*fd.type_id = int(value.Int64)
		}
	}
	return nil
}

// QueryFixs queries the fixs edge of the FurnitureDetail.
func (fd *FurnitureDetail) QueryFixs() *FixRoomQuery {
	return (&FurnitureDetailClient{config: fd.config}).QueryFixs(fd)
}

// QueryFurnitures queries the furnitures edge of the FurnitureDetail.
func (fd *FurnitureDetail) QueryFurnitures() *FurnitureQuery {
	return (&FurnitureDetailClient{config: fd.config}).QueryFurnitures(fd)
}

// QueryCounterstaffs queries the counterstaffs edge of the FurnitureDetail.
func (fd *FurnitureDetail) QueryCounterstaffs() *CounterStaffQuery {
	return (&FurnitureDetailClient{config: fd.config}).QueryCounterstaffs(fd)
}

// QueryTypes queries the types edge of the FurnitureDetail.
func (fd *FurnitureDetail) QueryTypes() *FurnitureTypeQuery {
	return (&FurnitureDetailClient{config: fd.config}).QueryTypes(fd)
}

// QueryRooms queries the rooms edge of the FurnitureDetail.
func (fd *FurnitureDetail) QueryRooms() *DataRoomQuery {
	return (&FurnitureDetailClient{config: fd.config}).QueryRooms(fd)
}

// Update returns a builder for updating this FurnitureDetail.
// Note that, you need to call FurnitureDetail.Unwrap() before calling this method, if this FurnitureDetail
// was returned from a transaction, and the transaction was committed or rolled back.
func (fd *FurnitureDetail) Update() *FurnitureDetailUpdateOne {
	return (&FurnitureDetailClient{config: fd.config}).UpdateOne(fd)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (fd *FurnitureDetail) Unwrap() *FurnitureDetail {
	tx, ok := fd.config.driver.(*txDriver)
	if !ok {
		panic("ent: FurnitureDetail is not a transactional entity")
	}
	fd.config.driver = tx.drv
	return fd
}

// String implements the fmt.Stringer.
func (fd *FurnitureDetail) String() string {
	var builder strings.Builder
	builder.WriteString("FurnitureDetail(")
	builder.WriteString(fmt.Sprintf("id=%v", fd.ID))
	builder.WriteString(", date_add=")
	builder.WriteString(fd.DateAdd.Format(time.ANSIC))
	builder.WriteString(", furniture_amount=")
	builder.WriteString(fmt.Sprintf("%v", fd.FurnitureAmount))
	builder.WriteString(", furniture_colour=")
	builder.WriteString(fd.FurnitureColour)
	builder.WriteString(", detail=")
	builder.WriteString(fd.Detail)
	builder.WriteByte(')')
	return builder.String()
}

// FurnitureDetails is a parsable slice of FurnitureDetail.
type FurnitureDetails []*FurnitureDetail

func (fd FurnitureDetails) config(cfg config) {
	for _i := range fd {
		fd[_i].config = cfg
	}
}
