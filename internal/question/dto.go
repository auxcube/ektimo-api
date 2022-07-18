package question

import "github.com/auxcube/ektimo-api/ent"

type Questions struct {
	Text []*ent.TextQuestion `json:"text,omitempty"`
}
