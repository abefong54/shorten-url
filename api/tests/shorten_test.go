package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/abefong54/shorten-url/routes"
	"github.com/gofiber/fiber/v2"
)

func TestGetShortenURL(t *testing.T) {

	server := fiber.New()

	// test data
	reqData := struct {
		URL         string        `json:"url"`
		CustomShort string        `json:"custom_short"`
		Expiry      time.Duration `json:"expiry"`
	}{
		URL:         "www.google.com",
		CustomShort: "mineplease",
		Expiry:      10,
	}

	body, _ := json.Marshal(reqData)

	// handler + route
	server.Get("/api/v1", routes.ShortenURL)

	// request + response
	req, _ := http.NewRequest(http.MethodGet, "/api/v1", bytes.NewBuffer(body))
	resp, _ := server.Test(req, -1)

	// logs + assertion
	t.Log(resp.StatusCode)
	t.Log(resp.Body)
	// assert.Equal(t, 200, resp.StatusCode)

}
