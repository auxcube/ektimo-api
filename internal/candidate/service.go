package candidate

import (
	"context"

	"github.com/auxcube/ektimo-api/ent"
	"github.com/auxcube/ektimo-api/pkg/database"
)

type Service struct {
	// TODO: mock this for unit tests
	db *database.Client
}

// NewService returns a new Service
func NewService(db *database.Client) *Service {
	return &Service{db: db}
}

func (s *Service) List(ctx context.Context) ([]*ent.Candidate, error) {
	candidates, err := s.db.Candidate.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	return candidates, nil
}
