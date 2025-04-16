package cmd

import (
	"edukita-teaching-grading/configs"
	"edukita-teaching-grading/internal/app/repository"
	"edukita-teaching-grading/internal/app/service"
	"edukita-teaching-grading/pkg/driver"

	"github.com/sirupsen/logrus"
)

func Run() {
	config, err := configs.LoadConfigurations(".env")
	if err != nil {
		logrus.Fatalf("failed to load configurations: %v", err)
		return
	}

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

	_ = repositoryConnector(repository.RepositoryOption{
		DB: psql,
	})
}

func repositoryConnector(opt repository.RepositoryOption) *repository.Repository {
	userRepo := repository.InitiateUserRepository(opt)
	return &repository.Repository{
		User: userRepo,
	}
}

func serviceConnector(opt service.ServiceOption) *service.Service {
	// userService := service.InitiateUserService(opt)
	return &service.Service{}
}
