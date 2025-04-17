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
		LoginUser(ctx context.Context, requestBody *payload.LoginUserRequest) (response payload.LoginUserResponse, err error)
		GetUserByID(ctx context.Context, id string) (response payload.GetUserResponse, err error)
		LogoutUser(ctx context.Context, id string) (response payload.LogoutUserResponse, err error)
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
		user, err := s.Repository.User.GetUserByEmail(ctx, requestBody.Email, tx)
		if err != nil {
			s.Logger.Warnf(fmt.Sprintf("failed to get user by email: %s", err.Error()), zap.Error(err))
			return
		}
		if user.ID != uuid.Nil {
			err = pkg.NewBadRequestError("email already exists", nil)
			s.Logger.Warnf("email already exists: %s", user.Role, zap.Error(err))
			return
		}

		userID := uuid.New()
		user = model.User{
			BaseModel: model.BaseModel{
				ID:        userID,
				CreatedBy: userID,
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
			student := model.Student{
				UserID:         user.ID,
				StudentID:      uuid.NewString(),
				EnrollmentYear: time.Now().Year(),
				Program:        requestBody.Program,
			}
			student, err = s.Repository.User.CreateStudent(ctx, student, tx)
			if err != nil {
				s.Logger.Warnf(fmt.Sprintf("failed to create student:%s", err.Error()), zap.Error(err))
				return
			}
		default:
			err = pkg.NewBadRequestError("invalid role", nil)
			s.Logger.Warnf("invalid role: %s", user.Role, zap.Error(err))
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

func (s *UserService) LoginUser(ctx context.Context, requestBody *payload.LoginUserRequest) (response payload.LoginUserResponse, err error) {
	return response, repository.TransactionWrapper(ctx, s.Postgres, func(tx *sqlx.Tx) (err error) {
		user, err := s.Repository.User.GetUserByEmail(ctx, requestBody.Email, tx)
		if err != nil {
			s.Logger.Warnf(fmt.Sprintf("failed to get user by email: %s", err.Error()), zap.Error(err))
			return
		}

		if !user.CheckPassword(requestBody.Password) {
			err = pkg.NewBadRequestError("invalid password", nil)
			s.Logger.Warnf("invalid password: %s", user.Role, zap.Error(err))
			return
		}

		token, err := GenerateJWTToken(user, s.Config.Application.Secret, s.Config.Cookies.SSOExpired)
		if err != nil {
			s.Logger.Warnf(fmt.Sprintf("failed to generate token: %s", err.Error()), zap.Error(err))
			return err
		}

		now := time.Now()
		user.LastLogin = &now
		user.UpdatedBy = &user.ID
		user, err = s.Repository.User.UpdateUserByID(ctx, user, tx)
		if err != nil {
			s.Logger.Warnf(fmt.Sprintf("failed to update user: %s", err.Error()), zap.Error(err))
			return
		}

		response.Token = token

		return
	})
}

func (s *UserService) GetUserByID(ctx context.Context, id string) (response payload.GetUserResponse, err error) {
	return response, repository.TransactionWrapper(ctx, s.Postgres, func(tx *sqlx.Tx) (err error) {
		user, err := s.Repository.User.GetUserByID(ctx, id, tx)
		if err != nil {
			s.Logger.Warnf(fmt.Sprintf("failed to get user by id: %s", err.Error()), zap.Error(err))
			return
		}

		switch user.Role {
		case pkg.ROLE_ADMIN:
		case pkg.ROLE_TEACHER:
			teacher, err := s.Repository.User.GetTeacherByID(ctx, user.ID.String(), tx)
			if err != nil {
				s.Logger.Warnf(fmt.Sprintf("failed to get teacher by id: %s", err.Error()), zap.Error(err))
				return err
			}
			response.Role.Department = teacher.Department
			response.Role.Title = teacher.Title
		case pkg.ROLE_STUDENT:
			student, err := s.Repository.User.GetStudentByID(ctx, user.ID.String(), tx)
			if err != nil {
				s.Logger.Warnf(fmt.Sprintf("failed to get student by id: %s", err.Error()), zap.Error(err))
				return err
			}
			response.Role.StudentID = student.StudentID
			response.Role.EnrollmentYear = student.EnrollmentYear
			response.Role.Program = student.Program
		default:
			err = pkg.NewBadRequestError("invalid role", nil)
			s.Logger.Warnf("invalid role: %s", user.Role, zap.Error(err))
			return
		}

		response.ID = user.ID.String()
		response.FirstName = user.FirstName
		response.LastName = user.LastName
		response.Email = user.Email
		response.IsActive = user.IsActive
		response.LastLogin = user.LastLogin.Format(time.RFC3339)
		response.CreatedAt = user.CreatedAt.Format(time.RFC3339)
		if user.UpdatedAt != nil {
			response.UpdatedAt = user.UpdatedAt.Format(time.RFC3339)
		}

		return
	})
}

func (s *UserService) LogoutUser(ctx context.Context, id string) (response payload.LogoutUserResponse, err error) {
	return response, repository.TransactionWrapper(ctx, s.Postgres, func(tx *sqlx.Tx) (err error) {
		user, err := s.Repository.User.GetUserByID(ctx, id, tx)
		if err != nil {
			s.Logger.Warnf(fmt.Sprintf("failed to get user by id: %s", err.Error()), zap.Error(err))
			return
		}

		now := time.Now()
		user.LastLogin = &now
		user.UpdatedBy = &user.ID
		user, err = s.Repository.User.UpdateUserByID(ctx, user, tx)
		if err != nil {
			s.Logger.Warnf(fmt.Sprintf("failed to update user: %s", err.Error()), zap.Error(err))
			return
		}

		response.ID = user.ID.String()
		return
	})
}
