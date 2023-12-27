package main

import (
	"fmt"
	"log"
	"os"

	"github.com/abefong54/shorten-url/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

// list of all available API routes
func setupRoutes(app *fiber.App) {
	app.Get("/:url", routes.ResolveURL())
	app.Post("/api/v1", routes.ShortenURL())
}

// APP variable is an instance of the Fiber library
// we will pass it around
func main() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println(err)
	}

	app := fiber.New()

	app.Use(logger.New())

	setupRoutes(app)

	log.Fatal(app.Listen(os.Getenv("APP_PORT")))

}
