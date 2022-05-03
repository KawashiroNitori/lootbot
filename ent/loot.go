// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/KawashiroNitori/lootbot/ent/loot"
	"github.com/KawashiroNitori/lootbot/internal/macro"
)

// Loot is the model entity for the Loot schema.
type Loot struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// PlayerName holds the value of the "player_name" field.
	PlayerName string `json:"player_name,omitempty"`
	// PlayerServer holds the value of the "player_server" field.
	PlayerServer string `json:"player_server,omitempty"`
	// PartyID holds the value of the "party_id" field.
	PartyID string `json:"party_id,omitempty"`
	// Role holds the value of the "role" field.
	Role macro.Role `json:"role,omitempty"`
	// Job holds the value of the "job" field.
	Job macro.Job `json:"job,omitempty"`
	// Category holds the value of the "category" field.
	Category macro.Category `json:"category,omitempty"`
	// ItemID holds the value of the "item_id" field.
	ItemID int64 `json:"item_id,omitempty"`
	// ItemName holds the value of the "item_name" field.
	ItemName string `json:"item_name,omitempty"`
	// IsObtained holds the value of the "is_obtained" field.
	IsObtained bool `json:"is_obtained,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// ObtainedAt holds the value of the "obtained_at" field.
	ObtainedAt *time.Time `json:"obtained_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Loot) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case loot.FieldCategory:
			values[i] = new(macro.Category)
		case loot.FieldJob:
			values[i] = new(macro.Job)
		case loot.FieldRole:
			values[i] = new(macro.Role)
		case loot.FieldIsObtained:
			values[i] = new(sql.NullBool)
		case loot.FieldID, loot.FieldItemID:
			values[i] = new(sql.NullInt64)
		case loot.FieldPlayerName, loot.FieldPlayerServer, loot.FieldPartyID, loot.FieldItemName:
			values[i] = new(sql.NullString)
		case loot.FieldCreatedAt, loot.FieldUpdatedAt, loot.FieldObtainedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Loot", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Loot fields.
func (l *Loot) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case loot.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			l.ID = int64(value.Int64)
		case loot.FieldPlayerName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field player_name", values[i])
			} else if value.Valid {
				l.PlayerName = value.String
			}
		case loot.FieldPlayerServer:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field player_server", values[i])
			} else if value.Valid {
				l.PlayerServer = value.String
			}
		case loot.FieldPartyID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field party_id", values[i])
			} else if value.Valid {
				l.PartyID = value.String
			}
		case loot.FieldRole:
			if value, ok := values[i].(*macro.Role); !ok {
				return fmt.Errorf("unexpected type %T for field role", values[i])
			} else if value != nil {
				l.Role = *value
			}
		case loot.FieldJob:
			if value, ok := values[i].(*macro.Job); !ok {
				return fmt.Errorf("unexpected type %T for field job", values[i])
			} else if value != nil {
				l.Job = *value
			}
		case loot.FieldCategory:
			if value, ok := values[i].(*macro.Category); !ok {
				return fmt.Errorf("unexpected type %T for field category", values[i])
			} else if value != nil {
				l.Category = *value
			}
		case loot.FieldItemID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field item_id", values[i])
			} else if value.Valid {
				l.ItemID = value.Int64
			}
		case loot.FieldItemName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field item_name", values[i])
			} else if value.Valid {
				l.ItemName = value.String
			}
		case loot.FieldIsObtained:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_obtained", values[i])
			} else if value.Valid {
				l.IsObtained = value.Bool
			}
		case loot.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				l.CreatedAt = value.Time
			}
		case loot.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				l.UpdatedAt = value.Time
			}
		case loot.FieldObtainedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field obtained_at", values[i])
			} else if value.Valid {
				l.ObtainedAt = new(time.Time)
				*l.ObtainedAt = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Loot.
// Note that you need to call Loot.Unwrap() before calling this method if this Loot
// was returned from a transaction, and the transaction was committed or rolled back.
func (l *Loot) Update() *LootUpdateOne {
	return (&LootClient{config: l.config}).UpdateOne(l)
}

// Unwrap unwraps the Loot entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (l *Loot) Unwrap() *Loot {
	tx, ok := l.config.driver.(*txDriver)
	if !ok {
		panic("ent: Loot is not a transactional entity")
	}
	l.config.driver = tx.drv
	return l
}

// String implements the fmt.Stringer.
func (l *Loot) String() string {
	var builder strings.Builder
	builder.WriteString("Loot(")
	builder.WriteString(fmt.Sprintf("id=%v", l.ID))
	builder.WriteString(", player_name=")
	builder.WriteString(l.PlayerName)
	builder.WriteString(", player_server=")
	builder.WriteString(l.PlayerServer)
	builder.WriteString(", party_id=")
	builder.WriteString(l.PartyID)
	builder.WriteString(", role=")
	builder.WriteString(fmt.Sprintf("%v", l.Role))
	builder.WriteString(", job=")
	builder.WriteString(fmt.Sprintf("%v", l.Job))
	builder.WriteString(", category=")
	builder.WriteString(fmt.Sprintf("%v", l.Category))
	builder.WriteString(", item_id=")
	builder.WriteString(fmt.Sprintf("%v", l.ItemID))
	builder.WriteString(", item_name=")
	builder.WriteString(l.ItemName)
	builder.WriteString(", is_obtained=")
	builder.WriteString(fmt.Sprintf("%v", l.IsObtained))
	builder.WriteString(", created_at=")
	builder.WriteString(l.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(l.UpdatedAt.Format(time.ANSIC))
	if v := l.ObtainedAt; v != nil {
		builder.WriteString(", obtained_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteByte(')')
	return builder.String()
}

// Loots is a parsable slice of Loot.
type Loots []*Loot

func (l Loots) config(cfg config) {
	for _i := range l {
		l[_i].config = cfg
	}
}
