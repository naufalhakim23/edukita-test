package server

import (
	"fmt"

	"edukita-teaching-grading/internal/app/handler"
	"edukita-teaching-grading/internal/app/repository"
	"edukita-teaching-grading/internal/app/service"
	"edukita-teaching-grading/internal/pkg"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
)

type IServer interface {
	ServerRun()
}

type server struct {
	Option     pkg.OptionsApplication
	Service    *service.Service
	Logger     *zap.SugaredLogger
	Repository *repository.Repository
}

// NewServer create object server
func NewServer(opt pkg.OptionsApplication, svc *service.Service, repo *repository.Repository) IServer {
	return &server{
		Option:     opt,
		Service:    svc,
		Logger:     opt.Logger,
		Repository: repo,
	}
}

func (s *server) ServerRun() {
	pkg.SwaggerInfo(s.Option.Config)

	f := fiber.New()

	f.Use(recover.New())
	// CORS
	f.Use(cors.New(cors.Config{
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
	}))

	Router(handler.HandlerOptions{
		OptionsApplication: s.Option,
		Service:            s.Service,
		Repository:         s.Repository,
	}, f)

	address := fmt.Sprintf(":%v", s.Option.Config.Application.Port)

	// Start the server
	if err := f.Listen(address); err != nil {
		logrus.Fatalf("failed to listen: %v", err)
	}

	// Gracefully shut down the server when the application is shutting down
	if err := f.Shutdown(); err != nil {
		logrus.Errorf("Failed to shut down server gracefully: %v", err)
	} else {
		logrus.Info("Server shut down gracefully")
	}
}
