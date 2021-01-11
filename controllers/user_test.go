package controllers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/romycode/bank-manager/models"
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

func TestUserController_GetAllUsers(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/v1/users", strings.NewReader(""))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	m := &mockUserRepository{}
	sut := NewUserController(m)

	u := models.User{ID: "id", Name: "name", Email: "email"}
	a := models.Account{ID: "id", UserID: "id", IBAN: "ES00", Credit: "0"}
	m.On("All").Return([]models.UserInfo{{
		User:     u,
		Accounts: []models.Account{a},
	}})

	sut.GetAllUsers(c)

	m.AssertExpectations(t)
}

func TestUserController_CreateUser(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/v1/users", strings.NewReader(`{"id": "id","name": "name","email": "email"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	m := &mockUserRepository{}
	sut := NewUserController(m)

	u := models.User{ID: "id", Name: "name", Email: "email"}
	m.On("Save", u)

	expected, _ := json.Marshal(u)

	if assert.NoError(t, sut.CreateUser(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, string(expected), rec.Body.String())
	}

	m.AssertExpectations(t)
}

func TestUserController_DeleteUser(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/v1/users/150798", strings.NewReader(`{"id": "id","name": "name","email": "email"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("150798")

	m := &mockUserRepository{}
	sut := NewUserController(m)

	m.On("Delete", "150798")

	if assert.NoError(t, sut.DeleteUser(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "", rec.Body.String())
	}

	m.AssertExpectations(t)
}
