// Code generated by entimport, DO NOT EDIT.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{field.Uint("id").SchemaType(map[string]string{"postgres": "bigserial"}), field.String("name").Optional(), field.Int("age").Optional(), field.Time("birthday").Optional(), field.Time("created_at").Optional(

	// Edges of the User.
	), field.Time("updated_at").Optional(), field.Time("deleted_at").Optional()}

}

func (User) Edges() []ent.Edge {
	return nil
}
