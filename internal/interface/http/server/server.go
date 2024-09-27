package server

import (
	"fib/internal/interface/http/handlers"
	"fib/internal/interface/http/middleware"
	"fib/pkg/logger"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	App   *fiber.App
	Port  string
	token string
	lg    logger.MyLogger
}

func NewServer(port, token string) *Server {
	app := fiber.New()
	lg := logger.NewSlogLogger()
	app.Use(middleware.AuthMiddleware(lg, token))

	return &Server{
		App:   app,
		Port:  port,
		token: token,
		lg:    lg,
	}

}

func (s *Server) Run() error {
	s.lg.Info("Сервер запускается...")

	s.App.Post("/", middleware.LoggingMw(s.lg), handlers.SumHandler)

	err := s.App.Listen(s.Port)
	if err != nil {
		s.lg.Error("Ошибка при запуске сервера", err)
	}
	return err
}
