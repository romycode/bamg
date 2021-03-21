package account

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

func TestAccountHandler_GetAllAccounts(t *testing.T) {
	a := noter.Account{ID: "id", UserID: "id", IBAN: "ES00", Credit: "0"}

	m := new(storagemocks.AccountRepository)
	m.On("All", mock.Anything).Return([]noter.Account{a})

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/v1/accounts", FetchAllAccountsHandler(m))

	t.Run("given a valid request it returns all accounts", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/v1/accounts", strings.NewReader(""))
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		expectedResponse := "[{\"id\":\"id\",\"userId\":\"id\",\"iban\":\"ES00\",\"credit\":\"0\"}]"
		m.AssertExpectations(t)
		assert.Equal(t, http.StatusOK, res.StatusCode)
		assert.Equal(t, expectedResponse, rec.Body.String())
	})
}
