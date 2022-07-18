package question

import (
	"context"

	"github.com/auxcube/ektimo-api/ent"
)

type Service struct {
	textService textQuestionService
}

type textQuestionService interface {
	FindAll(ctx context.Context) ([]*ent.TextQuestion, error)
}

func NewService(ts textQuestionService) *Service {
	return &Service{
		ts,
	}
}

func (s *Service) FindAll(ctx context.Context) (Questions, error) {
	textQuestions, err := s.textService.FindAll(ctx)
	if err != nil {
		return Questions{}, err
	}
	// TODO: add mcq and coding question queries here
	return Questions{
		Text: textQuestions,
	}, nil
}
