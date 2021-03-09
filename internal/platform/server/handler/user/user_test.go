package user

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/romycode/bank-manager/internal/bank_manager_api/models"
)

type mockUserRepository struct {
	mock.Mock
}

func (m *mockUserRepository) All() []models.UserInfo {
	args := m.Called()
	return args.Get(0).([]models.UserInfo)
}

func (m *mockUserRepository) Save(u models.User) {
	m.Called(u)
}

func (m *mockUserRepository) Delete(id string) {
	m.Called(id)
}

func TestUserHandler_GetAllUsers(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()

	m := &mockUserRepository{}
	u := models.User{ID: "id", Name: "name", Email: "email"}
	a := models.Account{ID: "id", UserID: "id", IBAN: "ES00", Credit: "0"}
	m.On("All").Return([]models.UserInfo{{
		User:     u,
		Accounts: []models.Account{a},
	}})

	r.GET("/v1/users", GetAllUsers(m))

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

func TestUserHandler_CreateUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()

	m := &mockUserRepository{}
	u := models.User{ID: "id", Name: "name", Email: "email"}
	m.On("Save", u)

	r.POST("/v1/users", CreateUser(m))

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

func TestUserHandler_DeleteUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()

	m := &mockUserRepository{}
	m.On("Delete", "150798")

	r.DELETE("/v1/users/:id", DeleteUser(m))

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
