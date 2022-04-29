// Code generated by entc, DO NOT EDIT.

package loot

import (
	"fmt"
	"time"

	"github.com/KawashiroNitori/lootbot/internal/macro"
)

const (
	// Label holds the string label denoting the loot type in the database.
	Label = "loot"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldPlayerName holds the string denoting the player_name field in the database.
	FieldPlayerName = "player_name"
	// FieldPlayerServer holds the string denoting the player_server field in the database.
	FieldPlayerServer = "player_server"
	// FieldPartyID holds the string denoting the party_id field in the database.
	FieldPartyID = "party_id"
	// FieldRole holds the string denoting the role field in the database.
	FieldRole = "role"
	// FieldJob holds the string denoting the job field in the database.
	FieldJob = "job"
	// FieldCategory holds the string denoting the category field in the database.
	FieldCategory = "category"
	// FieldItemID holds the string denoting the item_id field in the database.
	FieldItemID = "item_id"
	// FieldItemName holds the string denoting the item_name field in the database.
	FieldItemName = "item_name"
	// FieldIsObtained holds the string denoting the is_obtained field in the database.
	FieldIsObtained = "is_obtained"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldObtainedAt holds the string denoting the obtained_at field in the database.
	FieldObtainedAt = "obtained_at"
	// Table holds the table name of the loot in the database.
	Table = "loots"
)

// Columns holds all SQL columns for loot fields.
var Columns = []string{
	FieldID,
	FieldPlayerName,
	FieldPlayerServer,
	FieldPartyID,
	FieldRole,
	FieldJob,
	FieldCategory,
	FieldItemID,
	FieldItemName,
	FieldIsObtained,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldObtainedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// PlayerNameValidator is a validator for the "player_name" field. It is called by the builders before save.
	PlayerNameValidator func(string) error
	// PlayerServerValidator is a validator for the "player_server" field. It is called by the builders before save.
	PlayerServerValidator func(string) error
	// DefaultIsObtained holds the default value on creation for the "is_obtained" field.
	DefaultIsObtained bool
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)

// RoleValidator is a validator for the "role" field enum values. It is called by the builders before save.
func RoleValidator(r macro.Role) error {
	switch r.String() {
	case "D1", "D2", "D3", "D4", "MT", "ST", "H1", "H2":
		return nil
	default:
		return fmt.Errorf("loot: invalid enum value for role field: %q", r)
	}
}

// JobValidator is a validator for the "job" field enum values. It is called by the builders before save.
func JobValidator(j macro.Job) error {
	switch j.String() {
	case "GLD", "PGL", "MRD", "LNC", "ARC", "CNJ", "THM", "CRP", "BSM", "ARM", "GSM", "LTW", "WVR", "ALC", "CUL", "MIN", "BTN", "FSH", "PLD", "MNK", "WAR", "DRG", "BRD", "WHM", "BLM", "ACN", "SMN", "SCH", "ROG", "NIN", "MCH", "DRK", "AST", "SAM", "RDM", "BLU", "GNB", "DNC", "RPR", "SGE":
		return nil
	default:
		return fmt.Errorf("loot: invalid enum value for job field: %q", j)
	}
}

// CategoryValidator is a validator for the "category" field enum values. It is called by the builders before save.
func CategoryValidator(c macro.Category) error {
	switch c.String() {
	case "weapon", "coffer", "coating", "tomestone", "roborant", "spool", "mount", "orchestraroll", "companion":
		return nil
	default:
		return fmt.Errorf("loot: invalid enum value for category field: %q", c)
	}
}
