package main

import (
	"log"

	"123/handlers"
	middlewares "123/middleware"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Use(middlewares.LoggingMiddleware)
	app.Post("/", handlers.SumHandler)

	if err := app.Listen(":3000"); err != nil {
		log.Fatal(err)
	}
}
