package tests

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"r-cha/goblog/controllers"
	"r-cha/goblog/db"
	"r-cha/goblog/models"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	// Setup
	db.Connect()

	// Run test
	code := m.Run()

	// Teardown
	os.Exit(code)
}

func MockServe(req *http.Request) *httptest.ResponseRecorder {
	r := controllers.NewRoutes()
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	return w
}

func TestListPosts(t *testing.T) {
	// List posts
	req, _ := http.NewRequest("GET", "/v1/posts/", nil)
	w := MockServe(req)

	assert.Equal(t, w.Code, 200)
}

func TestCreatePost(t *testing.T) {
	// Create a new post
	body := strings.NewReader(`
		{
			"title": "Post",
			"author": {
				"name": "Robert"
			},
			"text": "this is a test"
		}
	`)
	req, _ := http.NewRequest("POST", "/v1/posts/", body)
	w := MockServe(req)

	assert.Equal(t, w.Code, 200)
	res, _ := io.ReadAll(w.Body)
	assert.Contains(t, string(res), "Robert")
}

func TestGetPost(t *testing.T) {
	// Create a new post
	body := strings.NewReader(`
		{
			"title": "Post To Get",
			"author": {
				"name": "Still Robert"
			},
			"text": "this is another test"
		}
	`)
	req, _ := http.NewRequest("POST", "/v1/posts/", body)
	w := MockServe(req)
	res, _ := io.ReadAll(w.Body)

	var result models.Post
	json.Unmarshal(res, &result)
	id := result.ID

	// Get it directly
	req, _ = http.NewRequest(
		"GET",
		fmt.Sprintf("/v1/posts/%s", fmt.Sprint(id)),
		body,
	)
	w = MockServe(req)

	assert.Equal(t, w.Code, 200)
}
