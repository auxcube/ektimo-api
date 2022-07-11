package health

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/auxcube/ektimo-api/internal/health/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func setup(t *testing.T) (*mocks.Mockdatabase, *gin.Engine, func()) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	dbMock := mocks.NewMockdatabase(ctrl)
	router := gin.New()
	healthController := NewController(dbMock)
	healthController.RegisterRoutes(router)
	return dbMock, router, ctrl.Finish
}

func TestController_handleHealth_ok(t *testing.T) {
	dbMock, router, teardown := setup(t)
	defer teardown()

	req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	dbMock.EXPECT().Ping(req.Context()).Return(nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	expected := Response{
		Status: StatusHealthy,
	}
	var got Response
	resp := w.Result()
	defer resp.Body.Close()
	bytes, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(bytes, &got)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, expected, got)
}

func TestController_handleHealth_errorConnectionRefused(t *testing.T) {
	dbMock, router, teardown := setup(t)
	defer teardown()

	req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	dbMock.EXPECT().Ping(req.Context()).Return(errors.New("dial tcp 127.0.0.1:5432: connect: connection refused"))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	expected := Response{
		Status: StatusUnhealthy,
		Error:  "dial tcp 127.0.0.1:5432: connect: connection refused",
	}
	var got Response
	resp := w.Result()
	defer resp.Body.Close()
	bytes, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(bytes, &got)

	assert.Equal(t, 500, resp.StatusCode)
	assert.Equal(t, expected, got)
}
