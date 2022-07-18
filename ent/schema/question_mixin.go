package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

// QuestionMixin implements the ent.Mixin for sharing
// entity details fields with Question type entities.
type QuestionMixin struct {
	// We embed the `mixin.Schema` to avoid
	// implementing the rest of the methods.
	mixin.Schema
}

func (QuestionMixin) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.String("human_id").
			NotEmpty().
			Unique(),
		field.String("text").
			NotEmpty(),
	}
}
