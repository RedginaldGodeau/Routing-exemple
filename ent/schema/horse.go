package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type Horse struct {
	ent.Schema
}

// Fields of the User.
func (Horse) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
	}
}

// Edges of the User.
func (Horse) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).Ref("horses").Unique(),
	}
}
