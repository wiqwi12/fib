package server

import (
	"fib/internal/interface/http/handlers"
	"fib/internal/interface/http/middleware"
	"fib/pkg/logger"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	App  *fiber.App
	Port string
}

func NewServer(port string) *Server {

	app := fiber.New()

	app.Use(middleware.LoggingMiddleware)

	return &Server{
		App:  app,
		Port: port,
	}
}

func (s *Server) Run() error {

	lg := logger.NewLogger()
	lg.Info("Сервер запускается...") // Добавим отладочный вывод
	s.App.Post("/", handlers.SumHandler)
	err := s.App.Listen(s.Port)
	if err != nil {
		lg.Error("Ошибка при запуске сервера", "error", err) // Логируем ошибку
	}
	return err
}
