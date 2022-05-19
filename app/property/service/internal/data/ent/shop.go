// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/yongliu1992/smartpms/app/property/service/internal/data/ent/shop"
)

// Shop is the model entity for the Shop schema.
type Shop struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// FloorID holds the value of the "floor_id" field.
	FloorID string `json:"floor_id,omitempty"`
	// RoomNum holds the value of the "room_num" field.
	RoomNum int `json:"room_num,omitempty"`
	// Layer holds the value of the "layer" field.
	Layer int `json:"layer,omitempty"`
	// BuiltUpArea holds the value of the "built_up_area" field.
	BuiltUpArea float32 `json:"built_up_area,omitempty"`
	// CommunityID holds the value of the "community_id" field.
	CommunityID int `json:"community_id,omitempty"`
	// FeeRate holds the value of the "fee_rate" field.
	FeeRate float32 `json:"fee_rate,omitempty"`
	// RoomArea holds the value of the "room_area" field.
	RoomArea float32 `json:"room_area,omitempty"`
	// Rent holds the value of the "rent" field.
	Rent float32 `json:"rent,omitempty"`
	// Remark holds the value of the "remark" field.
	Remark string `json:"remark,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Shop) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case shop.FieldBuiltUpArea, shop.FieldFeeRate, shop.FieldRoomArea, shop.FieldRent:
			values[i] = new(sql.NullFloat64)
		case shop.FieldID, shop.FieldRoomNum, shop.FieldLayer, shop.FieldCommunityID:
			values[i] = new(sql.NullInt64)
		case shop.FieldFloorID, shop.FieldRemark:
			values[i] = new(sql.NullString)
		case shop.FieldCreatedAt, shop.FieldUpdatedAt, shop.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Shop", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Shop fields.
func (s *Shop) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case shop.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = int(value.Int64)
		case shop.FieldFloorID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field floor_id", values[i])
			} else if value.Valid {
				s.FloorID = value.String
			}
		case shop.FieldRoomNum:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field room_num", values[i])
			} else if value.Valid {
				s.RoomNum = int(value.Int64)
			}
		case shop.FieldLayer:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field layer", values[i])
			} else if value.Valid {
				s.Layer = int(value.Int64)
			}
		case shop.FieldBuiltUpArea:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field built_up_area", values[i])
			} else if value.Valid {
				s.BuiltUpArea = float32(value.Float64)
			}
		case shop.FieldCommunityID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field community_id", values[i])
			} else if value.Valid {
				s.CommunityID = int(value.Int64)
			}
		case shop.FieldFeeRate:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field fee_rate", values[i])
			} else if value.Valid {
				s.FeeRate = float32(value.Float64)
			}
		case shop.FieldRoomArea:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field room_area", values[i])
			} else if value.Valid {
				s.RoomArea = float32(value.Float64)
			}
		case shop.FieldRent:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field rent", values[i])
			} else if value.Valid {
				s.Rent = float32(value.Float64)
			}
		case shop.FieldRemark:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field remark", values[i])
			} else if value.Valid {
				s.Remark = value.String
			}
		case shop.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				s.CreatedAt = value.Time
			}
		case shop.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				s.UpdatedAt = value.Time
			}
		case shop.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				s.DeletedAt = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Shop.
// Note that you need to call Shop.Unwrap() before calling this method if this Shop
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Shop) Update() *ShopUpdateOne {
	return (&ShopClient{config: s.config}).UpdateOne(s)
}

// Unwrap unwraps the Shop entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Shop) Unwrap() *Shop {
	tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Shop is not a transactional entity")
	}
	s.config.driver = tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Shop) String() string {
	var builder strings.Builder
	builder.WriteString("Shop(")
	builder.WriteString(fmt.Sprintf("id=%v", s.ID))
	builder.WriteString(", floor_id=")
	builder.WriteString(s.FloorID)
	builder.WriteString(", room_num=")
	builder.WriteString(fmt.Sprintf("%v", s.RoomNum))
	builder.WriteString(", layer=")
	builder.WriteString(fmt.Sprintf("%v", s.Layer))
	builder.WriteString(", built_up_area=")
	builder.WriteString(fmt.Sprintf("%v", s.BuiltUpArea))
	builder.WriteString(", community_id=")
	builder.WriteString(fmt.Sprintf("%v", s.CommunityID))
	builder.WriteString(", fee_rate=")
	builder.WriteString(fmt.Sprintf("%v", s.FeeRate))
	builder.WriteString(", room_area=")
	builder.WriteString(fmt.Sprintf("%v", s.RoomArea))
	builder.WriteString(", rent=")
	builder.WriteString(fmt.Sprintf("%v", s.Rent))
	builder.WriteString(", remark=")
	builder.WriteString(s.Remark)
	builder.WriteString(", created_at=")
	builder.WriteString(s.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(s.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", deleted_at=")
	builder.WriteString(s.DeletedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Shops is a parsable slice of Shop.
type Shops []*Shop

func (s Shops) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}
