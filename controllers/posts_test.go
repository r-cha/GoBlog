package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListPosts(t *testing.T) {
	r := NewRoutes()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/posts/", nil)
	r.ServeHTTP(w, req) // returns a 500 after panicking

	assert.Equal(t, 200, w.Code)
}
