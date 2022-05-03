package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/rs/xid"
)

// Party holds the schema definition for the Party entity.
type Party struct {
	ent.Schema
}

// Fields of the Party.
func (Party) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			DefaultFunc(func() string {
				return xid.New().String()
			}),
		field.String("channel_id"),
	}
}

// Edges of the Party.
func (Party) Edges() []ent.Edge {
	return nil
}
