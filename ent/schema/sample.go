package schema

import (
	"net/url"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/lucsky/cuid"
)

// Sample holds the schema definition for the Sample entity.
type Sample struct {
	ent.Schema
}

// Fields of the Sample.
func (Sample) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			DefaultFunc(func() string {
				return cuid.New()
			}),
		field.Int("office_id"),
		field.String("code"). // TODO: Implement unique constraint within same office
			NotEmpty().
			MaxLen(20),
		field.Enum("size").
			Values("small", "medium", "big"),
		field.Float("amount").
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(6,2)",
			}),
		field.Text("memo").
			Nillable().
			Optional(),
		field.JSON("url", &url.URL{}).
			Optional(),
		field.Bool("active").
			Default(true),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the Sample.
func (Sample) Edges() []ent.Edge {
	return nil
}

// Indexes of the Sample.
func (Sample) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("created_at"),
		index.Fields("office_id", "code").Unique(),
	}
}
