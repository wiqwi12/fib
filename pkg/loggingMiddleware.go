package middleware

import (
	"context"
	"fmt"

	"log/slog"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
)

func nLogger() *slog.Logger {
	return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))
}

func LoggingMiddleware(c *fiber.Ctx) error {
	start := time.Now()
	bodyBytes := c.Body()
	l := nLogger()

	err := c.Next()
	if err != nil {
		return err
	}

	statusCode := c.Response().StatusCode()
	responseBody := c.Response().Body()
	duration := time.Since(start).Milliseconds()

	logMessage := fmt.Sprintf("TIME: %s, REQ: %s, RES: %s, CODE: %d, DUR: %d ms",
		start.Format(time.RFC3339),
		string(bodyBytes),
		string(responseBody),
		statusCode,
		duration,
	)

	l.Log(context.Background(), slog.LevelInfo, logMessage)

	return nil
}
