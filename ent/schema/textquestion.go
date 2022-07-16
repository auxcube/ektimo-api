package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// TextQuestion holds the schema definition for the TextQuestion entity.
type TextQuestion struct {
	ent.Schema
}

// Mixin of the TextQuestion.
func (TextQuestion) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
		QuestionMixin{},
	}
}

// Fields of the TextQuestion.
func (TextQuestion) Fields() []ent.Field {
	return []ent.Field{
		field.String("answer").
			Optional(),
	}
}

// Edges of the TextQuestion.
func (TextQuestion) Edges() []ent.Edge {
	return nil
}
