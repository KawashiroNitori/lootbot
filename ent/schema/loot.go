package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
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
		field.String("player_name").NotEmpty(),
		field.String("player_server").NotEmpty(),
		field.String("party_id"),
		field.Enum("role").GoType(macro.Role(0)),
		field.Enum("job").GoType(macro.Job(0)),
		field.Enum("category").GoType(macro.Category(0)),
		field.Int64("item_id"),
		field.String("item_name"),
		field.Bool("is_obtained").Default(false),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.Time("obtained_at").Optional().Nillable(),
	}
}

// Edges of the Loot.
func (Loot) Edges() []ent.Edge {
	return nil
}

func (Loot) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("party_id", "item_id"),
	}
}
