package tests

import (
	"bytes"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShorten(t *testing.T) {
	db := SetupTestDB(t)
	defer TeardownTestDB(db)

	app := SetupApp()

	t.Run("Positive: valid URL", func(t *testing.T) {
		body := []byte(`{"url": "https://example.com"}`)
		req := httptest.NewRequest("POST", "/api/v1/shorten", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req, -1)
		assert.Equal(t, 201, resp.StatusCode)
	})

	t.Run("Negative: missing URL", func(t *testing.T) {
		body := []byte(`{"url": ""}`)
		req := httptest.NewRequest("POST", "/api/v1/shorten", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req, -1)
		assert.Equal(t, 400, resp.StatusCode)
	})

	t.Run("Negative: invalid URL format", func(t *testing.T) {
		body := []byte(`{"url": "not-a-url"}`)
		req := httptest.NewRequest("POST", "/api/v1/shorten", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req, -1)
		assert.Equal(t, 422, resp.StatusCode)
	})
}
