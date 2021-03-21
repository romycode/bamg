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

	noter "github.com/romycode/bank-manager/internal"
	"github.com/romycode/bank-manager/internal/platform/storage/storagemocks"
)

func TestUserHandler_GetAllUsers(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()

	u := noter.User{ID: "id", Name: "name", Email: "email"}
	a := noter.Account{ID: "id", UserID: "id", IBAN: "ES00", Credit: "0"}

	m := new(storagemocks.UserRepository)
	m.On("All", mock.Anything).Return([]noter.UserInfo{{
		User:     u,
		Accounts: []noter.Account{a},
	}})

	r.GET("/v1/users", FetchAllUsersHandler(m))

	t.Run("given a valid request it returns all users", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/v1/users", strings.NewReader(""))
		req.Header.Set("Content-Type", gin.MIMEJSON)
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		expectedResponse := "[{\"id\":\"id\",\"name\":\"name\",\"email\":\"email\",\"accounts\":[{\"id\":\"id\",\"userId\":\"id\",\"iban\":\"ES00\",\"credit\":\"0\"}]}]"

		m.AssertExpectations(t)
		assert.Equal(t, http.StatusOK, res.StatusCode)
		assert.Equal(t, expectedResponse, rec.Body.String())
	})
}
