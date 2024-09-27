package server

import (
	"fib/internal/interface/http/handlers"
	"fib/internal/interface/http/middleware"
	"fib/internal/interface/http/server/config"
	"fib/pkg/logger"

	"github.com/gofiber/fiber/v2"
	slogfiber "github.com/samber/slog-fiber"
)

type Server struct {
	App   *fiber.App
	Port  string
	token string
}

var lg = logger.NewLogger()

var Log = slogfiber.NewWithConfig(lg, config.LoggerCfg)

func NewServer(port, token string) *Server {

	// lg := logger.NewLogger()

	app := fiber.New()

	app.Use(Log)

	// app.Use(middleware.LoggingMiddleware(lg))

	return &Server{
		App:   app,
		Port:  port,
		token: token,
	}
}

func (s *Server) Run() error {

	lg := logger.NewLogger()
	lg.Info("Сервер запускается...") // Добавим отладочный вывод
	s.App.Post("/", middleware.AuthMiddleware(s.token), handlers.SumHandler)
	err := s.App.Listen(s.Port)
	if err != nil {
		lg.Error("Ошибка при запуске сервера", "error", err) // Логируем ошибку
	}
	return err
}
