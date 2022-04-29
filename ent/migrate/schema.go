// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// LootsColumns holds the columns for the "loots" table.
	LootsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "player_name", Type: field.TypeString},
		{Name: "player_server", Type: field.TypeString},
		{Name: "party_id", Type: field.TypeInt64},
		{Name: "role", Type: field.TypeEnum, Enums: []string{"D1", "D2", "D3", "D4", "MT", "ST", "H1", "H2"}},
		{Name: "job", Type: field.TypeEnum, Enums: []string{"GLD", "PGL", "MRD", "LNC", "ARC", "CNJ", "THM", "CRP", "BSM", "ARM", "GSM", "LTW", "WVR", "ALC", "CUL", "MIN", "BTN", "FSH", "PLD", "MNK", "WAR", "DRG", "BRD", "WHM", "BLM", "ACN", "SMN", "SCH", "ROG", "NIN", "MCH", "DRK", "AST", "SAM", "RDM", "BLU", "GNB", "DNC", "RPR", "SGE"}},
		{Name: "category", Type: field.TypeEnum, Enums: []string{"weapon", "coffer", "coating", "tomestone", "roborant", "spool", "mount", "orchestraroll", "companion"}},
		{Name: "item_id", Type: field.TypeInt64},
		{Name: "item_name", Type: field.TypeString},
		{Name: "is_obtained", Type: field.TypeBool, Default: false},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "obtained_at", Type: field.TypeTime, Nullable: true},
	}
	// LootsTable holds the schema information for the "loots" table.
	LootsTable = &schema.Table{
		Name:       "loots",
		Columns:    LootsColumns,
		PrimaryKey: []*schema.Column{LootsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "loot_party_id_item_id",
				Unique:  false,
				Columns: []*schema.Column{LootsColumns[3], LootsColumns[7]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		LootsTable,
	}
)

func init() {
}
