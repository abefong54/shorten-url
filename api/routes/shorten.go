package routes

import (
	"time"

	"github.com/abefong54/shorten-url/helpers"
	"github.com/gofiber/fiber"
)

// Format structure for request and response
type request struct {
	URL         string        `json:"url"`
	CustomShort string        `json:"custom_short"`
	Expiry      time.Duration `json:"expiry"`
}
type response struct {
	URL             string        `json:"url`
	CustomShort     string        `json:"custom_short`
	Expiry          time.Duration `json:"expiry`
	XRateRemaining  int           `json:"rate_limit`
	XRateLimitReset time.Duration `json:"rate_limit_reset`
}

// function that takes in a URL
// and generates a shorter string version
func ShortenURL(c *fiber.Ctx) error {

	body := new(request)

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Could not parse JSON"})
	}

	// implement rate limiting, check the IP of client

	// check if input is a real URL
	if !govalidator.IsURL(body.URL) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid URL"})
	}

	// check for domain errors ( is url accessible publically)
	if !helpers.RemoveDomainError(body.URL) {
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{"error": "Domain error!"})
	}

	// enforce https, SSL
	body.URL = helpers.EnforceHTTP(body.URL)
}
