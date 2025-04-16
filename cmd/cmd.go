package cmd

import (
	"edukita-teaching-grading/configs"
	"edukita-teaching-grading/internal/app/repository"
	"edukita-teaching-grading/internal/app/server"
	"edukita-teaching-grading/internal/app/service"
	"edukita-teaching-grading/internal/pkg"
	"edukita-teaching-grading/pkg/driver"

	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
)

func Run() {
	config, err := configs.LoadConfigurations(".env")
	if err != nil {
		logrus.Fatalf("failed to load configurations: %v", err)
		return
	}

	logger := zap.Must(zap.NewProduction()).Sugar()
	defer logger.Sync()

	psql, err := driver.NewDatabaseDriver(driver.PostgreSQLOption{
		DatabaseName: config.Postgresql.Name,
		URL:          config.Postgresql.URL,
	})
	if err != nil {
		logrus.Fatalf("failed to connect to database: %v", err)
		return
	} else {
		logrus.Infof("connected to database: %s", config.Postgresql.Name)
	}

	options := pkg.OptionsApplication{
		Config:   config,
		Postgres: psql,
		Logger:   logger,
	}

	repo := repositoryConnector(repository.RepositoryOption{
		OptionsApplication: options,
	})

	svc := serviceConnector(service.ServiceOption{
		OptionsApplication: options,
		Repository:         repo,
	})

	app := server.NewServer(options, svc, repo)
	app.ServerRun()
}

func repositoryConnector(opt repository.RepositoryOption) *repository.Repository {
	userRepo := repository.InitiateUserRepository(opt)
	return &repository.Repository{
		User: userRepo,
	}
}

func serviceConnector(opt service.ServiceOption) *service.Service {
	userService := service.InitiateUserService(opt)
	lmsService := service.InitiateLearningManagementService(opt)
	return &service.Service{
		User:               userService,
		LearningManagement: lmsService,
	}
}
