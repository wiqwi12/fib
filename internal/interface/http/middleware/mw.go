package middleware

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(apiToken string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		requestToken := c.Get("API_TOKEN")

		if requestToken != apiToken {
			log.Print("wrong token")
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		return c.Next()
	}

}

// func LoggingMiddleware(lg *slog.Logger) fiber.Handler {
// 	// Инициализация логгера с конфигурацией
// 	log := slogfiber.NewWithConfig(lg, config.LoggerCfg)

// 	return func(c *fiber.Ctx) error {
// 		// Вызов следующего хэндлера в цепочке
// 		err := c.Next()
// 		if err != nil {
// 			return err
// 		}

// 		// Логируем информацию о запросе и ответе автоматически на основе конфигурации

// 		return nil
// 	}
// }
