// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/team18/app/ent/customer"
	"github.com/team18/app/ent/dataroom"
	"github.com/team18/app/ent/fixroom"
	"github.com/team18/app/ent/furnituredetail"
)

// FixRoom is the model entity for the FixRoom schema.
type FixRoom struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// FixDetail holds the value of the "fix_detail" field.
	FixDetail string `json:"fix_detail,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the FixRoomQuery when eager-loading is set.
	Edges       FixRoomEdges `json:"edges"`
	customer_id *int
	room_id     *int
	object_id   *int
}

// FixRoomEdges holds the relations/edges for other nodes in the graph.
type FixRoomEdges struct {
	// Customer holds the value of the customer edge.
	Customer *Customer
	// FurnitureDetail holds the value of the furnitureDetail edge.
	FurnitureDetail *FurnitureDetail
	// Room holds the value of the room edge.
	Room *DataRoom
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// CustomerOrErr returns the Customer value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e FixRoomEdges) CustomerOrErr() (*Customer, error) {
	if e.loadedTypes[0] {
		if e.Customer == nil {
			// The edge customer was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: customer.Label}
		}
		return e.Customer, nil
	}
	return nil, &NotLoadedError{edge: "customer"}
}

// FurnitureDetailOrErr returns the FurnitureDetail value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e FixRoomEdges) FurnitureDetailOrErr() (*FurnitureDetail, error) {
	if e.loadedTypes[1] {
		if e.FurnitureDetail == nil {
			// The edge furnitureDetail was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: furnituredetail.Label}
		}
		return e.FurnitureDetail, nil
	}
	return nil, &NotLoadedError{edge: "furnitureDetail"}
}

// RoomOrErr returns the Room value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e FixRoomEdges) RoomOrErr() (*DataRoom, error) {
	if e.loadedTypes[2] {
		if e.Room == nil {
			// The edge room was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: dataroom.Label}
		}
		return e.Room, nil
	}
	return nil, &NotLoadedError{edge: "room"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*FixRoom) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // fix_detail
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*FixRoom) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // customer_id
		&sql.NullInt64{}, // room_id
		&sql.NullInt64{}, // object_id
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the FixRoom fields.
func (fr *FixRoom) assignValues(values ...interface{}) error {
	if m, n := len(values), len(fixroom.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	fr.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field fix_detail", values[0])
	} else if value.Valid {
		fr.FixDetail = value.String
	}
	values = values[1:]
	if len(values) == len(fixroom.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field customer_id", value)
		} else if value.Valid {
			fr.customer_id = new(int)
			*fr.customer_id = int(value.Int64)
		}
		if value, ok := values[1].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field room_id", value)
		} else if value.Valid {
			fr.room_id = new(int)
			*fr.room_id = int(value.Int64)
		}
		if value, ok := values[2].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field object_id", value)
		} else if value.Valid {
			fr.object_id = new(int)
			*fr.object_id = int(value.Int64)
		}
	}
	return nil
}

// QueryCustomer queries the customer edge of the FixRoom.
func (fr *FixRoom) QueryCustomer() *CustomerQuery {
	return (&FixRoomClient{config: fr.config}).QueryCustomer(fr)
}

// QueryFurnitureDetail queries the furnitureDetail edge of the FixRoom.
func (fr *FixRoom) QueryFurnitureDetail() *FurnitureDetailQuery {
	return (&FixRoomClient{config: fr.config}).QueryFurnitureDetail(fr)
}

// QueryRoom queries the room edge of the FixRoom.
func (fr *FixRoom) QueryRoom() *DataRoomQuery {
	return (&FixRoomClient{config: fr.config}).QueryRoom(fr)
}

// Update returns a builder for updating this FixRoom.
// Note that, you need to call FixRoom.Unwrap() before calling this method, if this FixRoom
// was returned from a transaction, and the transaction was committed or rolled back.
func (fr *FixRoom) Update() *FixRoomUpdateOne {
	return (&FixRoomClient{config: fr.config}).UpdateOne(fr)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (fr *FixRoom) Unwrap() *FixRoom {
	tx, ok := fr.config.driver.(*txDriver)
	if !ok {
		panic("ent: FixRoom is not a transactional entity")
	}
	fr.config.driver = tx.drv
	return fr
}

// String implements the fmt.Stringer.
func (fr *FixRoom) String() string {
	var builder strings.Builder
	builder.WriteString("FixRoom(")
	builder.WriteString(fmt.Sprintf("id=%v", fr.ID))
	builder.WriteString(", fix_detail=")
	builder.WriteString(fr.FixDetail)
	builder.WriteByte(')')
	return builder.String()
}

// FixRooms is a parsable slice of FixRoom.
type FixRooms []*FixRoom

func (fr FixRooms) config(cfg config) {
	for _i := range fr {
		fr[_i].config = cfg
	}
}