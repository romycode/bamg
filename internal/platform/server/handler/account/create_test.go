package account

import (
	"fmt"
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

func TestAccountHandler_CreateAccount(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()

	iban := noter.NewIban()
	a := noter.Account{ID: "id", UserID: "id", IBAN: iban, Credit: "0"}
	m := new(storagemocks.AccountRepository)
	m.On("Save", mock.Anything, a)

	r.POST("/v1/accounts", CreateAccountHandler(m))

	t.Run("given a valid request it returns 201", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodPost, "/v1/accounts", strings.NewReader(fmt.Sprintf(`{"id":"id","iban":"%s","userId":"id","credit":"0"}`, iban)))
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		m.AssertExpectations(t)
		assert.Equal(t, http.StatusCreated, res.StatusCode)
	})
}
