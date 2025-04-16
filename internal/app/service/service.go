package service

import (
	"edukita-teaching-grading/internal/app/repository"
	"edukita-teaching-grading/internal/pkg"
)

type ServiceOption struct {
	pkg.OptionsApplication
	Repository *repository.Repository
}

type Service struct {
	// User IUserService
}
