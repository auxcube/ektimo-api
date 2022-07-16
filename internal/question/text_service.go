package question

import (
	"context"

	"github.com/auxcube/ektimo-api/ent"
	"github.com/auxcube/ektimo-api/ent/textquestion"
	"github.com/google/uuid"
)

type TextService struct {
	repo textRepository
}

type textRepository interface {
	Query() *ent.TextQuestionQuery
}

func NewTextService(repo textRepository) *TextService {
	return &TextService{
		repo,
	}
}

func (s *TextService) FindAll(ctx context.Context) ([]*ent.TextQuestion, error) {
	tqs, err := s.repo.
		Query().
		All(ctx)
	if err != nil {
		return nil, err
	}
	return tqs, nil
}

func (s *TextService) Find(ctx context.Context, id uuid.UUID) (*ent.TextQuestion, error) {
	tq, err := s.repo.
		Query().
		Where(textquestion.ID(id)).
		Only(ctx)
	if err != nil {
		return nil, err
	}
	return tq, nil
}
