package routes

// import (
// 	"bytes"
// 	"encoding/json"
// 	"net/http"
// 	"testing"

// 	"github.com/gofiber/fiber/v2"
// 	"github.com/stretchr/testify/assert"
// )

// func TestResolveURL(t *testing.T) {
// 	server := fiber.New()
// 	server.Get(":/url", ResolveURL)
// 	// mock data
// 	happyPathInput := "www.google.com"
// 	reqBody, _ := json.Marshal(happyPathInput)
// 	req, _ := http.NewRequest(http.MethodGet, ":/url", bytes.NewBuffer(reqBody))
// 	req.Header.Set("Content-Type", "application/json")

// 	t.Log("hi")
// 	resp, _ := server.Test(req, -1)
// 	defer resp.Body.Close()

// 	assert.Equal(t, 301, resp.StatusCode)

// }
