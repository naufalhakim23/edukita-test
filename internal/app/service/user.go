package service

import (
	"context"
	"fmt"
	"time"

	"edukita-teaching-grading/internal/app/model"
	"edukita-teaching-grading/internal/app/payload"
	"edukita-teaching-grading/internal/app/repository"
	"edukita-teaching-grading/internal/pkg"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type (
	IUserService interface {
		RegisterUser(ctx context.Context, requestBody payload.RegisterUserRequest) (response *payload.RegisterUserResponse, err error)
	}
	UserService struct {
		ServiceOption
	}
)

func InitiateUserService(opt ServiceOption) IUserService {
	return &UserService{
		ServiceOption: opt,
	}
}

func (s *UserService) RegisterUser(ctx context.Context, requestBody payload.RegisterUserRequest) (response *payload.RegisterUserResponse, err error) {
	return response, repository.TransactionWrapper(ctx, s.Postgres, func(tx *sqlx.Tx) (err error) {
		now := time.Now()
		user := model.User{
			BaseModel: model.BaseModel{
				ID:        uuid.New(),
				CreatedBy: requestBody.Email,
				CreatedAt: now,
			},
			Email:     requestBody.Email,
			FirstName: requestBody.FirstName,
			LastName:  requestBody.LastName,
			Role:      requestBody.Role,
			IsActive:  true,
		}
		user.SetPassword(requestBody.Password, s.Config.Application.CostBcrypt)
		user, err = s.Repository.User.CreateUser(ctx, user, tx)
		if err != nil {
			s.Logger.Warnf(fmt.Sprintf("failed to create user: %s", err.Error()), zap.Error(err))
			return
		}

		switch user.Role {
		case pkg.ROLE_ADMIN:
		case pkg.ROLE_TEACHER:
			teacher := model.Teacher{
				UserID:     user.ID,
				Department: "",
				Title:      "",
			}
			teacher, err = s.Repository.User.CreateTeacher(ctx, teacher, tx)
			if err != nil {
				s.Logger.Warnf(fmt.Sprintf("failed to create teacher:%s", err.Error()), zap.Error(err))
				return
			}
		case pkg.ROLE_STUDENT:
			// student := model.Student{
			// 	UserID:         user.ID,
			// 	StudentID:      uuid.NewString(),
			// 	EnrollmentYear: time.Now().Year(),
			// 	// Program: ,
			// }
		default:
			err = pkg.NewBadRequestError("invalid role", nil)
			return
		}

		response = &payload.RegisterUserResponse{
			ID:        user.ID.String(),
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Role:      user.Role,
		}

		return
	})
}
