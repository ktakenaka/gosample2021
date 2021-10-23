package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Office holds the schema definition for the Office entity.
type Office struct {
	ent.Schema
}

// Fields of the Office.
func (Office) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty().
			MaxLen(25),
	}
}

// Edges of the Office.
func (Office) Edges() []ent.Edge {
	return nil
}
