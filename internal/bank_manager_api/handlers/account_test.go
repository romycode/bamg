package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/romycode/bank-manager/internal/bank_manager_api/models"
)

type mockAccountRepository struct {
	mock.Mock
}

func (m *mockAccountRepository) GetByUserId(usrID string) []models.Account {
	args := m.Called(usrID)
	return args.Get(0).([]models.Account)
}

func (m *mockAccountRepository) All() []models.Account {
	args := m.Called()
	return args.Get(0).([]models.Account)
}

func (m *mockAccountRepository) Save(u models.Account) {
	m.Called(u)
}

func (m *mockAccountRepository) Delete(id string) {
	m.Called(id)
}

func TestAccountHandler_GetAllAccounts(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/v1/accounts", strings.NewReader(""))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	m := &mockAccountRepository{}
	sut := NewAccountController(m)

	a := models.Account{ID: "id", UserID: "id", IBAN: "ES00", Credit: "0"}
	m.On("All").Return([]models.Account{a})

	sut.GetAllAccounts(c)

	m.AssertExpectations(t)
}

func TestAccountHandler_CreateAccount(t *testing.T) {
	e := echo.New()
	iban := models.NewIban()

	req := httptest.NewRequest(http.MethodPost, "/v1/accounts", strings.NewReader(fmt.Sprintf(`{"id":"id","iban":"%s","userId":"id","credit":"0"}`, iban)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	m := &mockAccountRepository{}
	sut := NewAccountController(m)

	a := models.Account{ID: "id", UserID: "id", IBAN: iban, Credit: "0"}
	m.On("Save", a)

	expected, _ := json.Marshal(a)

	if assert.NoError(t, sut.CreateAccount(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, string(expected), rec.Body.String())
	}

	m.AssertExpectations(t)
}

func TestAccountHandler_DeleteAccount(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/v1/account/150798", strings.NewReader(""))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("150798")

	m := &mockAccountRepository{}
	sut := NewAccountController(m)

	m.On("Delete", "150798")

	if assert.NoError(t, sut.DeleteAccount(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "", rec.Body.String())
	}

	m.AssertExpectations(t)
}
