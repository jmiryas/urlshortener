package tests

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRedirect(t *testing.T) {
	db := SetupTestDB(t)
	defer TeardownTestDB(db)

	app := SetupApp()

	t.Run("Positive: valid redirect", func(t *testing.T) {
		body := []byte(`{"url": "https://redirect.me"}`)
		req := httptest.NewRequest("POST", "/api/v1/shorten", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req, -1)

		var got map[string]interface{}
		json.NewDecoder(resp.Body).Decode(&got)
		token := got["short_token"].(string)

		req2 := httptest.NewRequest("GET", "/api/v1/"+token, nil)
		redirectResp, _ := app.Test(req2, -1)

		assert.Contains(t, []int{301, 302}, redirectResp.StatusCode)
		assert.Equal(t, "https://redirect.me", redirectResp.Header.Get("Location"))
	})

	t.Run("Negative: unknown token", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/api/v1/doesnotexist", nil)
		resp, _ := app.Test(req, -1)
		assert.Equal(t, 404, resp.StatusCode)
	})
}
