package routers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitRouter(t *testing.T) {
	router := InitRouter()

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("HEAD", "ping", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}
