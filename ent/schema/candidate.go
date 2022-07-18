package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Candidate represents a single candidate-user of the system
type Candidate struct {
	ent.Schema
}

// Fields of the Candidate
func (Candidate) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("name"),
		field.String("email").Unique(),
		field.Enum("status").
			Values("no-quiz-assigned").
			Default("no-quiz-assigned"),
	}
}

// Edges of the Candidate
func (Candidate) Edges() []ent.Edge {
	return nil
}
