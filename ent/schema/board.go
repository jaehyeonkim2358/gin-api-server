package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Board holds the schema definition for the Board entity.
type Board struct {
	ent.Schema
}

// Fields of the Board.
func (Board) Fields() []ent.Field {
	return []ent.Field{
		field.String("title"),
		field.Text("content"),
	}
}

// Edges of the Board.
func (Board) Edges() []ent.Edge {
	return nil
}
