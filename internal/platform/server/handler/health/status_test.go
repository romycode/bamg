package health

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandler_Create(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/health", CheckHandler())

	t.Run("given a valid request it returns 200", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/health", strings.NewReader(""))
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		expectedResponse := "{\"data\":\"¡Alive!\"}"
		assert.Equal(t, http.StatusOK, res.StatusCode)
		assert.Equal(t, expectedResponse, rec.Body.String())
	})
}
