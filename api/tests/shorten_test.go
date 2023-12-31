// package tests

// import (
// 	"bytes"
// 	"encoding/json"
// 	"net/http"
// 	"testing"
// 	"time"

// 	"github.com/abefong54/shorten-url/routes"
// 	"github.com/gofiber/fiber/v2"
// 	"github.com/stretchr/testify/assert"
// )

// type ShortenURLRequest struct {
// 	URL         string        `json:"url"`
// 	CustomShort string        `json:"custom_short"`
// 	Expiry      time.Duration `json:"expiry"`
// }

// func TestShortenURL(t *testing.T) {

// 	server := fiber.New()

// 	reqData := new(ShortenURLRequest)
// 	reqData.URL = "https://www.google.com"
// 	reqData.CustomShort = "testing3"
// 	reqData.Expiry = 2000

// 	t.Log(reqData)

// 	body, _ := json.Marshal(reqData)

// 	// handler + route
// 	server.Post("/api/v1", routes.ShortenURL)

// 	// request + response
// 	req, _ := http.NewRequest(http.MethodPost, "/api/v1", bytes.NewBuffer(body))
// 	resp, _ := server.Test(req, -1)

// 	// logs + assertion
// 	t.Log(resp.StatusCode)
// 	t.Log(resp.Body)
// 	assert.Equal(t, 200, resp.StatusCode)

// }
