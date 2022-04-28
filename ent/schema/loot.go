package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/KawashiroNitori/lootbot/internal/macro"
)

// Loot holds the schema definition for the Loot entity.
type Loot struct {
	ent.Schema
}

// Fields of the Loot.
func (Loot) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("name").NotEmpty(),
		field.Enum("role").GoType(macro.Role(0)),
		field.Enum("job").GoType(macro.Job(0)),
		field.Enum("category").GoType(macro.Category(0)),
	}
}

// Edges of the Loot.
func (Loot) Edges() []ent.Edge {
	return nil
}
