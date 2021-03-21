package user

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/mock"

	noter "github.com/romycode/bank-manager/internal"
	"github.com/romycode/bank-manager/internal/platform/storage/storagemocks"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUserHandler_CreateUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()

	m := new(storagemocks.UserRepository)
	u := noter.User{ID: "id", Name: "name", Email: "email"}
	m.On("Save", mock.Anything, u)

	r.POST("/v1/users", CreateUserHandler(m))

	t.Run("given a valid request it returns 201", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodPost, "/v1/users", strings.NewReader(`{"id": "id","name": "name","email": "email"}`))
		req.Header.Set("Content-Type", gin.MIMEJSON)
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		expected, _ := json.Marshal(u)

		m.AssertExpectations(t)
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, string(expected), rec.Body.String())
	})
}
