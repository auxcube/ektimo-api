package candidate_test

// import (
// 	"context"
// 	"errors"
// 	"testing"

// 	"github.com/auxcube/ektimo-api/ent"
// 	ent_candidate "github.com/auxcube/ektimo-api/ent/candidate"
// 	"github.com/auxcube/ektimo-api/internal/candidate"
// 	"github.com/auxcube/ektimo-api/internal/candidate/mocks"
// 	"github.com/golang/mock/gomock"
// 	"github.com/google/uuid"
// 	"github.com/stretchr/testify/assert"
// )

// func setup(t *testing.T) (*candidate.Service, *mocks.MockCandidateRepository, *assert.Assertions, *gomock.Controller) {
// 	ctrl := gomock.NewController(t)
// 	repoMock := mocks.NewMockCandidateRepository(ctrl)
// 	service := candidate.NewService(repoMock)
// 	assert := assert.New(t)
// 	return service, repoMock, assert, ctrl
// }

// func TestService_List_returnsAllCandidates(t *testing.T) {
// 	t.Parallel()
// 	service, repo, assert, ctrl := setup(t)
// 	defer ctrl.Finish()

// 	expected := []*ent.Candidate{
// 		{
// 			ID:     uuid.New(),
// 			Name:   "Mike Wheeler",
// 			Email:  "mike@mike.com",
// 			Status: ent_candidate.StatusNoQuizAssigned,
// 		},
// 		{
// 			ID:     [16]byte{},
// 			Name:   "Will Byers",
// 			Email:  "will@will.com",
// 			Status: ent_candidate.StatusNoQuizAssigned,
// 		},
// 	}

// 	query := mocks.NewMockCandidateQuery(ctrl)
// 	query.EXPECT().All(gomock.Any()).Return(expected, nil)
// 	repo.EXPECT().Query().Return(query)

// 	actual, err := service.List(context.Background())
// 	assert.Nil(err)
// 	assert.Equal(expected, actual)
// }

// func TestService_List_failsOnError(t *testing.T) {
// 	t.Parallel()
// 	service, repo, assert, ctrl := setup(t)
// 	defer ctrl.Finish()

// 	expected := errors.New("some error")

// 	query := mocks.NewMockCandidateQuery(ctrl)
// 	query.EXPECT().All(gomock.Any()).Return(nil, expected)
// 	repo.EXPECT().Query().Return(query)

// 	actual, err := service.List(context.Background())
// 	assert.Nil(actual)
// 	assert.Equal(expected, err)
// }
