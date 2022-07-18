//go:generate mockgen -source=service.go -destination=mocks/service.go -package=mocks
package candidate

import (
	"context"

	"github.com/auxcube/ektimo-api/ent"
)

type candidateRepository interface {
	Query() *ent.CandidateQuery
}

type Service struct {
	// TODO: mock this for unit tests
	repo candidateRepository
}

// NewService returns a new Service
func NewService(db candidateRepository) *Service {
	return &Service{repo: db}
}

// List returns a list of all candidates
func (s *Service) List(ctx context.Context) ([]*ent.Candidate, error) {
	candidates, err := s.repo.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	return candidates, nil
}
