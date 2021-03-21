package user

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/romycode/bank-manager/internal/platform/storage/storagemocks"
)

func TestUserHandler_DeleteUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()

	m := new(storagemocks.UserRepository)
	m.On("Delete", mock.Anything, "150798")

	r.DELETE("/v1/users/:id", DeleteUserHandler(m))

	t.Run("given a valid request it returns 200", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodDelete, "/v1/users/150798", strings.NewReader(``))
		req.Header.Set("Content-Type", gin.MIMEJSON)
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		m.AssertExpectations(t)
		assert.Equal(t, http.StatusOK, rec.Code)
	})
}
