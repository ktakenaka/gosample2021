package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Office holds the schema definition for the Office entity.
type Office struct {
	ent.Schema
}

// Fields of the Office.
func (Office) Fields() []ent.Field {
	return []ent.Field{
		field.String("code").
			NotEmpty().
			MaxLen(10).
			Unique(),
		field.String("name").
			NotEmpty().
			MaxLen(50),
	}
}

// Edges of the Office.
func (Office) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("samples", Sample.Type).
			StorageKey(edge.Column("office_id")).
			Annotations(entsql.Annotation{
				OnDelete: entsql.Cascade,
			}),
	}
}

// Indexes of the Office.
func (Office) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("code"),
	}
}
