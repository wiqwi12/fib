package middleware

import (
	"context"
	"fib/pkg/logger"
	"fmt"
	"log/slog"
	"time"

	"github.com/gofiber/fiber/v2"
)

func LoggingMiddleware(c *fiber.Ctx) error {
	start := time.Now()
	bodyBytes := c.Body()

	lg := logger.NewLogger()

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

	lg.Log(context.Background(), slog.LevelInfo, logMessage)

	return nil
}
