package account

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/gin-gonic/gin"

	"github.com/romycode/bank-manager/internal/platform/storage/storagemocks"
)

func TestAccountHandler_DeleteAccount(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()

	m := new(storagemocks.AccountRepository)
	m.On("Delete", mock.Anything, "150798")

	r.DELETE("/v1/accounts/:id", DeleteAccountHandler(m))

	t.Run("given a valid request it delete the account related to id", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodDelete, "/v1/accounts/150798", strings.NewReader(""))
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		m.AssertExpectations(t)
		assert.Equal(t, http.StatusOK, res.StatusCode)
	})
}
