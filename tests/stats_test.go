package tests

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStats(t *testing.T) {
	db := SetupTestDB(t)
	defer TeardownTestDB(db)

	app := SetupApp()

	// buat short url
	body := []byte(`{"url": "https://analytics.com/page"}`)
	req := httptest.NewRequest("POST", "/api/v1/shorten", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req, -1)

	var got map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&got)
	token := got["short_token"].(string)

	// simulate clicks
	for i := 0; i < 3; i++ {
		app.Test(httptest.NewRequest("GET", "/api/v1/"+token, nil), -1)
	}

	t.Run("Positive: click counter increments", func(t *testing.T) {
		statsReq := httptest.NewRequest("GET", "/api/v1/stats/"+token, nil)
		statsResp, _ := app.Test(statsReq, -1)

		var stats map[string]interface{}
		json.NewDecoder(statsResp.Body).Decode(&stats)

		clicks := stats["click_count"].(float64)
		assert.GreaterOrEqual(t, clicks, float64(3))
	})

	t.Run("Negative: stats unknown token", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/api/v1/stats/doesnotexist", nil)
		resp, _ := app.Test(req, -1)
		assert.Equal(t, 404, resp.StatusCode)
	})
}
