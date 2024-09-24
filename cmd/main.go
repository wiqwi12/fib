package main

import (
	"fib/internal/interface/http/handlers"
	"fib/internal/interface/http/middleware"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Use(middleware.LoggingMiddleware)
	app.Post("/", handlers.SumHandler)

	if err := app.Listen(":3000"); err != nil {
		log.Fatal(err)
	}
}
