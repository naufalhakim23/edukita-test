package service

import (
	"context"
	"edukita-teaching-grading/internal/app/model"
	"edukita-teaching-grading/internal/app/payload"
	"edukita-teaching-grading/internal/app/repository"
	"edukita-teaching-grading/internal/pkg"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type (
	ILearningManagementService interface {
		CreateCourse(ctx context.Context, requestBody *payload.CreateCourseRequest) (response payload.CreateCourseResponse, err error)
		GetCourseByID(ctx context.Context, id string) (response payload.GetCourseResponse, err error)
		GetCourseByCode(ctx context.Context, code string) (response payload.GetCourseResponse, err error)
		GetAllCourses(ctx context.Context) (response payload.GetAllCoursesResponse, err error)
		UpdateCourseByID(ctx context.Context, id string, requestBody *payload.UpdateCourseRequest) (response payload.UpdateCourseResponse, err error)

		CreateAssignment(ctx context.Context, requestBody *payload.CreateAssignmentRequest) (response payload.CreateAssignmentResponse, err error)
		GetAssignmentByID(ctx context.Context, id string) (response payload.GetAssignmentResponse, err error)
		UpdateAssignmentByID(ctx context.Context, id string, requestBody *payload.UpdateAssignmentRequest) (response payload.UpdateAssignmentResponse, err error)

		CreateSubmission(ctx context.Context, id string, requestBody *payload.CreateSubmissionRequest) (response payload.CreateSubmissionResponse, err error)
		GetSubmissionByID(ctx context.Context, id string) (response payload.GetSubmissionResponse, err error)
		UpdateSubmissionByID(ctx context.Context, id string, requestBody *payload.UpdateSubmissionRequest) (response payload.UpdateSubmissionResponse, err error)
		GetAllSubmissionsByCourseID(ctx context.Context, courseID string, userID string) (response payload.GetAllSubmissionsByCourseID, err error)
		GetAllSubmissionsByAssignmentID(ctx context.Context, assignmentID string, userID string) (response payload.GetAllSubmissionsResponse, err error)
		GetAllSubmissionsByUserID(ctx context.Context, id string) (response payload.GetAllSubmissionsResponse, err error)
	}
	LearningManagementService struct {
		ServiceOption
	}
)

func InitiateLearningManagementService(opt ServiceOption) ILearningManagementService {
	return &LearningManagementService{
		ServiceOption: opt,
	}
}

func (s *LearningManagementService) CreateCourse(ctx context.Context, requestBody *payload.CreateCourseRequest) (response payload.CreateCourseResponse, err error) {
	return response, repository.TransactionWrapper(ctx, s.Postgres, func(tx *sqlx.Tx) (err error) {
		user, err := s.Repository.User.GetUserByID(ctx, requestBody.CreatedBy, tx)
		if err != nil {
			s.Logger.Warnf(fmt.Sprintf("failed to get user by id: %s", err.Error()), zap.Error(err))
			return
		}

		switch user.Role {
		case pkg.ROLE_ADMIN:
		case pkg.ROLE_TEACHER:
			_, err := s.Repository.User.GetTeacherByID(ctx, user.ID.String(), tx)
			if err != nil {
				s.Logger.Warnf(fmt.Sprintf("failed to get teacher by id: %s", err.Error()), zap.Error(err))
				return err
			}
		default:
			err = pkg.NewBadRequestError("invalid role", nil)
			s.Logger.Warnf("invalid role: %s", user.Role, zap.Error(err))
			return
		}

		now := time.Now()
		startDate, _ := time.Parse(time.RFC3339, requestBody.StartDate)
		endDate, _ := time.Parse(time.RFC3339, requestBody.EndDate)
		course := model.Course{
			BaseModel: model.BaseModel{
				ID:        uuid.New(),
				CreatedBy: user.ID,
				CreatedAt: now,
			},
			Code:        requestBody.Code,
			Name:        requestBody.Name,
			Description: requestBody.Description,
			StartDate:   &startDate,
			EndDate:     &endDate,
			IsActive:    true,
		}
		course, err = s.Repository.LearningManagement.CreateCourse(ctx, course, tx)
		if err != nil {
			s.Logger.Warnf(fmt.Sprintf("failed to create course: %s", err.Error()), zap.Error(err))
			return
		}

		response.ID = course.ID.String()

		return
	})
}

func (s *LearningManagementService) GetCourseByID(ctx context.Context, id string) (response payload.GetCourseResponse, err error) {
	return response, repository.TransactionWrapper(ctx, s.Postgres, func(tx *sqlx.Tx) (err error) {
		course, err := s.Repository.LearningManagement.GetCourseByID(ctx, id, tx)
		if err != nil {
			s.Logger.Warnf(fmt.Sprintf("failed to get course by id: %s", err.Error()), zap.Error(err))
			return
		}

		response.ID = course.ID.String()
		response.Code = course.Code
		response.Name = course.Name
		response.Description = course.Description
		response.StartDate = course.StartDate.Format(time.RFC3339)
		response.EndDate = course.EndDate.Format(time.RFC3339)
		response.IsActive = course.IsActive
		return
	})
}

func (s *LearningManagementService) GetCourseByCode(ctx context.Context, code string) (response payload.GetCourseResponse, err error) {
	return response, repository.TransactionWrapper(ctx, s.Postgres, func(tx *sqlx.Tx) (err error) {
		course, err := s.Repository.LearningManagement.GetCourseByCode(ctx, code, tx)
		if err != nil {
			s.Logger.Warnf(fmt.Sprintf("failed to get course by code: %s", err.Error()), zap.Error(err))
			return
		}

		response.ID = course.ID.String()
		response.Code = course.Code
		response.Name = course.Name
		response.Description = course.Description
		response.StartDate = course.StartDate.Format(time.RFC3339)
		response.EndDate = course.EndDate.Format(time.RFC3339)
		response.IsActive = course.IsActive
		return
	})
}

func (s *LearningManagementService) GetAllCourses(ctx context.Context) (response payload.GetAllCoursesResponse, err error) {
	return response, repository.TransactionWrapper(ctx, s.Postgres, func(tx *sqlx.Tx) (err error) {
		courses, err := s.Repository.LearningManagement.GetAllCourses(ctx, tx)
		if err != nil {
			s.Logger.Warnf(fmt.Sprintf("failed to get all courses: %s", err.Error()), zap.Error(err))
			return
		}

		response.Courses = make([]payload.GetCourseResponse, len(courses))
		for i, course := range courses {
			response.Courses[i].ID = course.ID.String()
			response.Courses[i].Code = course.Code
			response.Courses[i].Name = course.Name
			response.Courses[i].Description = course.Description
			if course.StartDate != nil {
				response.Courses[i].StartDate = course.StartDate.Format(time.RFC3339)
			}
			if course.EndDate != nil {
				response.Courses[i].EndDate = course.EndDate.Format(time.RFC3339)
			}
			response.Courses[i].IsActive = course.IsActive
		}
		return
	})
}

func (s *LearningManagementService) UpdateCourseByID(ctx context.Context, id string, requestBody *payload.UpdateCourseRequest) (response payload.UpdateCourseResponse, err error) {
	return response, repository.TransactionWrapper(ctx, s.Postgres, func(tx *sqlx.Tx) (err error) {
		user, err := s.Repository.User.GetUserByID(ctx, requestBody.UserID, tx)
		if err != nil {
			s.Logger.Warnf(fmt.Sprintf("failed to get user by id: %s", err.Error()), zap.Error(err))
			return
		}

		switch user.Role {
		case pkg.ROLE_ADMIN:
		case pkg.ROLE_TEACHER:
			_, err := s.Repository.User.GetTeacherByID(ctx, user.ID.String(), tx)
			if err != nil {
				s.Logger.Warnf(fmt.Sprintf("failed to get teacher by id: %s", err.Error()), zap.Error(err))
				return err
			}
		default:
			err = pkg.NewBadRequestError("invalid role", nil)
			s.Logger.Warnf("invalid role: %s", user.Role, zap.Error(err))
			return
		}

		course, err := s.Repository.LearningManagement.GetCourseByCode(ctx, requestBody.Code, tx)
		if err != nil {
			s.Logger.Warnf(fmt.Sprintf("failed to get course by code: %s", err.Error()), zap.Error(err))
			return
		}

		now := time.Now()
		startDate, _ := time.Parse(time.RFC3339, requestBody.StartDate)
		endDate, _ := time.Parse(time.RFC3339, requestBody.EndDate)
		course.Name = requestBody.Name
		course.Code = requestBody.Code
		course.Description = requestBody.Description
		course.StartDate = &startDate
		course.EndDate = &endDate
		course.IsActive = true
		course.UpdatedBy = &user.ID
		course.UpdatedAt = &now

		course, err = s.Repository.LearningManagement.UpdateCourseByID(ctx, course, tx)
		if err != nil {
			s.Logger.Warnf(fmt.Sprintf("failed to update course: %s", err.Error()), zap.Error(err))
			return
		}

		response.ID = course.ID.String()
		response.Code = course.Code
		response.Name = course.Name
		response.Description = course.Description
		response.StartDate = course.StartDate.Format(time.RFC3339)
		response.EndDate = course.EndDate.Format(time.RFC3339)
		response.IsActive = course.IsActive
		return
	})
}

func (s *LearningManagementService) CreateAssignment(ctx context.Context, requestBody *payload.CreateAssignmentRequest) (response payload.CreateAssignmentResponse, err error) {
	return response, repository.TransactionWrapper(ctx, s.Postgres, func(tx *sqlx.Tx) (err error) {
		user, err := s.Repository.User.GetUserByID(ctx, requestBody.CreatedBy, tx)
		if err != nil {
			s.Logger.Warnf(fmt.Sprintf("failed to get user by id: %s", err.Error()), zap.Error(err))
			return
		}

		switch user.Role {
		case pkg.ROLE_ADMIN:
		case pkg.ROLE_TEACHER:
			_, err := s.Repository.User.GetTeacherByID(ctx, user.ID.String(), tx)
			if err != nil {
				s.Logger.Warnf(fmt.Sprintf("failed to get teacher by id: %s", err.Error()), zap.Error(err))
				return err
			}
		default:
			err = pkg.NewBadRequestError("invalid role", nil)
			s.Logger.Warnf("invalid role: %s", user.Role, zap.Error(err))
			return
		}

		course, err := s.Repository.LearningManagement.GetCourseByID(ctx, requestBody.CourseID, tx)
		if err != nil {
			s.Logger.Warnf(fmt.Sprintf("failed to get course by id: %s", err.Error()), zap.Error(err))
			return
		}

		now := time.Now()
		assignment := model.Assignment{
			BaseModel: model.BaseModel{
				ID:        uuid.New(),
				CreatedBy: user.ID,
				CreatedAt: now,
			},
			Title:       requestBody.Title,
			Description: requestBody.Description,
			Content:     requestBody.Content,
			DueDate:     now,
			TeacherID:   user.ID,
			CourseID:    course.ID,
			TotalPoints: requestBody.TotalPoints,
			IsPublished: true,
		}
		assignment, err = s.Repository.LearningManagement.CreateAssignment(ctx, assignment, tx)
		if err != nil {
			s.Logger.Warnf(fmt.Sprintf("failed to create assignment: %s", err.Error()), zap.Error(err))
			return
		}

		response.ID = assignment.ID.String()
		return
	})
}

func (s *LearningManagementService) GetAssignmentByID(ctx context.Context, id string) (response payload.GetAssignmentResponse, err error) {
	return response, repository.TransactionWrapper(ctx, s.Postgres, func(tx *sqlx.Tx) (err error) {
		assignment, err := s.Repository.LearningManagement.GetAssignmentByID(ctx, id, tx)
		if err != nil {
			s.Logger.Warnf(fmt.Sprintf("failed to get assignment by id: %s", err.Error()), zap.Error(err))
			return
		}

		response.ID = assignment.ID.String()
		response.Title = assignment.Title
		response.Description = assignment.Description
		response.DueDate = assignment.DueDate.Format(time.RFC3339)
		response.TotalPoints = assignment.TotalPoints
		response.IsPublished = assignment.IsPublished
		return
	})
}

func (s *LearningManagementService) UpdateAssignmentByID(ctx context.Context, id string, requestBody *payload.UpdateAssignmentRequest) (response payload.UpdateAssignmentResponse, err error) {
	return response, repository.TransactionWrapper(ctx, s.Postgres, func(tx *sqlx.Tx) (err error) {
		user, err := s.Repository.User.GetUserByID(ctx, requestBody.UserID, tx)
		if err != nil {
			s.Logger.Warnf(fmt.Sprintf("failed to get user by id: %s", err.Error()), zap.Error(err))
			return
		}
		assignment, err := s.Repository.LearningManagement.GetAssignmentByID(ctx, id, tx)
		if err != nil {
			s.Logger.Warnf(fmt.Sprintf("failed to get assignment by id: %s", err.Error()), zap.Error(err))
			return
		}

		switch assignment.TeacherID {
		case uuid.Nil:
			err = pkg.NewBadRequestError("invalid teacher id", nil)
			s.Logger.Warnf("invalid teacher id: %s", assignment.TeacherID, zap.Error(err))
			return
		}

		now := time.Now()
		assignment.Title = requestBody.Title
		assignment.Description = requestBody.Description
		assignment.DueDate = now
		assignment.TotalPoints = requestBody.TotalPoints
		assignment.IsPublished = requestBody.IsPublished
		assignment.UpdatedBy = &user.ID
		assignment.UpdatedAt = &now
		assignment, err = s.Repository.LearningManagement.UpdateAssignmentByID(ctx, assignment, tx)
		if err != nil {
			s.Logger.Warnf(fmt.Sprintf("failed to update assignment: %s", err.Error()), zap.Error(err))
			return
		}

		response.ID = assignment.ID.String()
		response.Title = assignment.Title
		response.Description = assignment.Description
		response.DueDate = assignment.DueDate.Format(time.RFC3339)
		response.TotalPoints = assignment.TotalPoints
		response.IsPublished = assignment.IsPublished
		return
	})
}

func (s *LearningManagementService) CreateSubmission(ctx context.Context, id string, requestBody *payload.CreateSubmissionRequest) (response payload.CreateSubmissionResponse, err error) {
	return response, repository.TransactionWrapper(ctx, s.Postgres, func(tx *sqlx.Tx) (err error) {
		user, err := s.Repository.User.GetUserByID(ctx, requestBody.CreatedBy, tx)
		if err != nil {
			s.Logger.Warnf(fmt.Sprintf("failed to get user by id: %s", err.Error()), zap.Error(err))
			return
		}

		assignment, err := s.Repository.LearningManagement.GetAssignmentByID(ctx, requestBody.AssignmentID, tx)
		if err != nil {
			s.Logger.Warnf(fmt.Sprintf("failed to get assignment by id: %s", err.Error()), zap.Error(err))
			return
		}

		switch user.Role {
		case pkg.ROLE_STUDENT:
			student, err := s.Repository.User.GetStudentByID(ctx, user.ID.String(), tx)
			if err != nil {
				s.Logger.Warnf(fmt.Sprintf("failed to get student by id: %s", err.Error()), zap.Error(err))
				return err
			}

			submission := model.Submission{
				BaseModel: model.BaseModel{
					ID:        uuid.New(),
					CreatedBy: user.ID,
					CreatedAt: time.Now(),
				},
				AssignmentID: assignment.ID,
				StudentID:    student.UserID,
				SubmittedAt:  time.Now(),
				TeacherID:    assignment.TeacherID,
				Content:      requestBody.Content,
			}
			if requestBody.FileURL != "" {
				submission.FileURL = &requestBody.FileURL
			}

			submission, err = s.Repository.LearningManagement.CreateSubmission(ctx, submission, tx)
			if err != nil {
				s.Logger.Warnf(fmt.Sprintf("failed to create submission: %s", err.Error()), zap.Error(err))
				return err
			}

			response.ID = submission.ID.String()
		default:
			err = pkg.NewBadRequestError("invalid role", nil)
			s.Logger.Warnf("invalid role: %s", user.Role, zap.Error(err))
			return

		}

		return
	})
}

func (s *LearningManagementService) GetSubmissionByID(ctx context.Context, id string) (response payload.GetSubmissionResponse, err error) {
	return response, repository.TransactionWrapper(ctx, s.Postgres, func(tx *sqlx.Tx) (err error) {
		submission, err := s.Repository.LearningManagement.GetSubmissionByID(ctx, id, tx)
		if err != nil {
			s.Logger.Warnf(fmt.Sprintf("failed to get submission by id: %s", err.Error()), zap.Error(err))
			return
		}
		response.ID = submission.ID.String()
		response.AssignmentID = submission.AssignmentID.String()
		response.StudentID = submission.StudentID.String()
		response.SubmittedAt = submission.SubmittedAt.Format(time.RFC3339)
		response.Content = submission.Content
		response.CreatedAt = submission.CreatedAt.Format(time.RFC3339)
		response.CreatedBy = submission.CreatedBy.String()
		if submission.FileURL != nil {
			response.FileURL = *submission.FileURL
		}

		if submission.Grade != nil {
			response.Grade = submission.Grade
		}
		if submission.Feedback != nil {
			response.Feedback = submission.Feedback
		}
		if submission.GradedAt != nil {
			gradedAt := submission.GradedAt.Format(time.RFC3339)
			response.GradedAt = &gradedAt
		}

		if submission.GradedBy != nil {
			response.GradedBy = submission.GradedBy
		}
		return
	})
}

func (s *LearningManagementService) UpdateSubmissionByID(ctx context.Context, id string, requestBody *payload.UpdateSubmissionRequest) (response payload.UpdateSubmissionResponse, err error) {
	return response, repository.TransactionWrapper(ctx, s.Postgres, func(tx *sqlx.Tx) (err error) {
		user, err := s.Repository.User.GetUserByID(ctx, requestBody.UserID, tx)
		if err != nil {
			s.Logger.Warnf(fmt.Sprintf("failed to get user by id: %s", err.Error()), zap.Error(err))
			return
		}

		submission, err := s.Repository.LearningManagement.GetSubmissionByID(ctx, id, tx)
		if err != nil {
			s.Logger.Warnf(fmt.Sprintf("failed to get submission by id: %s", err.Error()), zap.Error(err))
			return
		}

		now := time.Now()
		switch user.Role {
		case pkg.ROLE_TEACHER:
			_, err := s.Repository.User.GetTeacherByID(ctx, user.ID.String(), tx)
			if err != nil {
				s.Logger.Warnf(fmt.Sprintf("failed to get teacher by id: %s", err.Error()), zap.Error(err))
				return err
			}
			if requestBody.Grade != 0 {
				submission.Grade = &requestBody.Grade
				submission.GradedAt = &now
				userID := user.ID.String()
				submission.GradedBy = &userID
			}
			if requestBody.Feedback != "" {
				submission.Feedback = &requestBody.Feedback
			}
			submission.UpdatedBy = &user.ID
			submission.UpdatedAt = &now
		case pkg.ROLE_STUDENT:
			student, err := s.Repository.User.GetStudentByID(ctx, user.ID.String(), tx)
			if err != nil {
				s.Logger.Warnf(fmt.Sprintf("failed to get student by id: %s", err.Error()), zap.Error(err))
				return err
			}
			submission.StudentID = student.UserID
			submission.Content = requestBody.Content
			submission.UpdatedBy = &user.ID
			submission.UpdatedAt = &now
			if requestBody.FileURL != "" {
				submission.FileURL = &requestBody.FileURL
			}
		default:
			err = pkg.NewBadRequestError("invalid role", nil)
			s.Logger.Warnf("invalid role: %s", user.Role, zap.Error(err))
			return
		}

		submission, err = s.Repository.LearningManagement.UpdateSubmissionByID(ctx, submission, tx)
		if err != nil {
			s.Logger.Warnf(fmt.Sprintf("failed to update submission: %s", err.Error()), zap.Error(err))
			return
		}

		response.ID = submission.ID.String()
		response.AssignmentID = submission.AssignmentID.String()
		response.StudentID = submission.StudentID.String()
		response.SubmittedAt = submission.SubmittedAt.Format(time.RFC3339)

		response.Content = submission.Content
		if submission.FileURL != nil {
			response.FileURL = *submission.FileURL
		}

		if submission.Grade != nil {
			response.Grade = submission.Grade
		}
		if submission.Feedback != nil {
			response.Feedback = submission.Feedback
		}
		if submission.GradedAt != nil {
			gradedAt := submission.GradedAt.Format(time.RFC3339)
			response.GradedAt = &gradedAt
		}
		if submission.GradedBy != nil {
			response.GradedBy = submission.GradedBy
		}
		return
	})
}

func (s *LearningManagementService) GetAllSubmissionsByCourseID(ctx context.Context, courseID string, userID string) (response payload.GetAllSubmissionsByCourseID, err error) {
	return response, repository.TransactionWrapper(ctx, s.Postgres, func(tx *sqlx.Tx) (err error) {
		user, err := s.Repository.User.GetUserByID(ctx, userID, tx)
		if err != nil {
			s.Logger.Warnf(fmt.Sprintf("failed to get user by id: %s", err.Error()), zap.Error(err))
			return
		}

		switch user.Role {
		case pkg.ROLE_TEACHER:
			_, err := s.Repository.User.GetTeacherByID(ctx, user.ID.String(), tx)
			if err != nil {
				s.Logger.Warnf(fmt.Sprintf("failed to get teacher by id: %s", err.Error()), zap.Error(err))
				return err
			}
		case pkg.ROLE_STUDENT:
			_, err := s.Repository.User.GetStudentByID(ctx, user.ID.String(), tx)
			if err != nil {
				s.Logger.Warnf(fmt.Sprintf("failed to get student by id: %s", err.Error()), zap.Error(err))
				return err
			}
		default:
			err = pkg.NewBadRequestError("invalid role", nil)
			s.Logger.Warnf("invalid role: %s", user.Role, zap.Error(err))
			return
		}

		course, err := s.Repository.LearningManagement.GetCourseByID(ctx, courseID, tx)
		if err != nil {
			s.Logger.Warnf(fmt.Sprintf("failed to get course by id: %s", err.Error()), zap.Error(err))
			return
		}

		assignments, err := s.Repository.LearningManagement.GetAllAssignmentsByCourseID(ctx, course.ID.String(), tx)
		if err != nil {
			s.Logger.Warnf(fmt.Sprintf("failed to get assignment by id: %s", err.Error()), zap.Error(err))
			return
		}

		for _, assignment := range assignments {
			submissions, err := s.Repository.LearningManagement.GetAllSubmissionsByAssignmentID(ctx, assignment.ID.String(), tx)
			if err != nil {
				s.Logger.Warnf(fmt.Sprintf("failed to get submission by id: %s", err.Error()), zap.Error(err))
				return err
			}
			response.CourseID = course.ID.String()
			response.Title = course.Name
			response.Description = course.Description
			response.DueDate = course.EndDate.Format(time.RFC3339)
			response.CreatedAt = course.CreatedAt.Format(time.RFC3339)
			response.CreatedBy = course.CreatedBy.String()
			response.Assignments = make([]payload.AssignmentAndSubmissions, len(assignments))
			for i, assignment := range assignments {
				response.Assignments[i].AssignmentID = assignment.ID.String()
				response.Assignments[i].Title = assignment.Title
				response.Assignments[i].Description = assignment.Description
				response.Assignments[i].DueDate = assignment.DueDate.Format(time.RFC3339)
				response.Assignments[i].TotalPoints = assignment.TotalPoints
				response.Assignments[i].IsPublished = assignment.IsPublished
				response.Assignments[i].CreatedAt = assignment.CreatedAt.Format(time.RFC3339)
				response.Assignments[i].CreatedBy = assignment.CreatedBy.String()
				response.Assignments[i].Submissions = make([]payload.GetSubmissionResponse, len(submissions))
				for j, submission := range submissions {
					response.Assignments[i].Submissions[j].ID = submission.ID.String()
					response.Assignments[i].Submissions[j].AssignmentID = submission.AssignmentID.String()
					response.Assignments[i].Submissions[j].StudentID = submission.StudentID.String()
					response.Assignments[i].Submissions[j].SubmittedAt = submission.SubmittedAt.Format(time.RFC3339)
					response.Assignments[i].Submissions[j].Content = submission.Content
					response.Assignments[i].Submissions[j].CreatedAt = submission.CreatedAt.Format(time.RFC3339)
					response.Assignments[i].Submissions[j].CreatedBy = submission.CreatedBy.String()
					if submission.FileURL != nil {
						response.Assignments[i].Submissions[j].FileURL = *submission.FileURL
					}
					if submission.Grade != nil {
						response.Assignments[i].Submissions[j].Grade = submission.Grade
					}
					if submission.Feedback != nil {
						response.Assignments[i].Submissions[j].Feedback = submission.Feedback
					}
					if submission.GradedAt != nil {
						gradedAt := submission.GradedAt.Format(time.RFC3339)
						response.Assignments[i].Submissions[j].GradedAt = &gradedAt
					}
					if submission.GradedBy != nil {
						response.Assignments[i].Submissions[j].GradedBy = submission.GradedBy
					}
				}
			}
		}
		return
	})
}

func (s *LearningManagementService) GetAllSubmissionsByAssignmentID(ctx context.Context, assignmentID string, userID string) (response payload.GetAllSubmissionsResponse, err error) {
	return response, repository.TransactionWrapper(ctx, s.Postgres, func(tx *sqlx.Tx) (err error) {
		user, err := s.Repository.User.GetUserByID(ctx, userID, tx)
		if err != nil {
			s.Logger.Warnf(fmt.Sprintf("failed to get user by id: %s", err.Error()), zap.Error(err))
			return
		}

		switch user.Role {
		case pkg.ROLE_TEACHER:
			_, err := s.Repository.User.GetTeacherByID(ctx, user.ID.String(), tx)
			if err != nil {
				s.Logger.Warnf(fmt.Sprintf("failed to get teacher by id: %s", err.Error()), zap.Error(err))
				return err
			}
		default:
			err = pkg.NewBadRequestError("invalid role", nil)
			s.Logger.Warnf("invalid role: %s", user.Role, zap.Error(err))
			return
		}

		submissions, err := s.Repository.LearningManagement.GetAllSubmissionsByAssignmentID(ctx, assignmentID, tx)
		if err != nil {
			s.Logger.Warnf(fmt.Sprintf("failed to get submission by id: %s", err.Error()), zap.Error(err))
			return
		}

		response.Submissions = make([]payload.GetSubmissionResponse, len(submissions))
		for i, submission := range submissions {
			response.Submissions[i].ID = submission.ID.String()
			response.Submissions[i].AssignmentID = submission.AssignmentID.String()
			response.Submissions[i].StudentID = submission.StudentID.String()
			response.Submissions[i].SubmittedAt = submission.SubmittedAt.Format(time.RFC3339)
			response.Submissions[i].Content = submission.Content
			response.Submissions[i].CreatedAt = submission.CreatedAt.Format(time.RFC3339)
			response.Submissions[i].CreatedBy = submission.CreatedBy.String()
			if submission.FileURL != nil {
				response.Submissions[i].FileURL = *submission.FileURL
			}

			if submission.Grade != nil {
				response.Submissions[i].Grade = submission.Grade
			}
			if submission.Feedback != nil {
				response.Submissions[i].Feedback = submission.Feedback
			}
			if submission.GradedAt != nil {
				gradedAt := submission.GradedAt.Format(time.RFC3339)
				response.Submissions[i].GradedAt = &gradedAt
			}
			if submission.GradedBy != nil {
				response.Submissions[i].GradedBy = submission.GradedBy
			}
		}

		return
	})
}

func (s *LearningManagementService) GetAllSubmissionsByUserID(ctx context.Context, userID string) (response payload.GetAllSubmissionsResponse, err error) {
	return response, repository.TransactionWrapper(ctx, s.Postgres, func(tx *sqlx.Tx) (err error) {
		user, err := s.Repository.User.GetUserByID(ctx, userID, tx)
		if err != nil {
			s.Logger.Warnf(fmt.Sprintf("failed to get user by id: %s", err.Error()), zap.Error(err))
			return
		}

		var submissions []model.Submission
		switch user.Role {
		case pkg.ROLE_TEACHER:
			assignment, err := s.Repository.LearningManagement.GetAssignmentByTeacherID(ctx, user.ID.String(), tx)
			if err != nil {
				s.Logger.Warnf(fmt.Sprintf("failed to get assignment by id: %s", err.Error()), zap.Error(err))
				return err
			}

			submissions, err = s.Repository.LearningManagement.GetAllSubmissionsByAssignmentID(ctx, assignment.ID.String(), tx)
			if err != nil {
				s.Logger.Warnf(fmt.Sprintf("failed to get submission by id: %s", err.Error()), zap.Error(err))
				return err
			}

		case pkg.ROLE_STUDENT:
			student, err := s.Repository.User.GetStudentByID(ctx, user.ID.String(), tx)
			if err != nil {
				s.Logger.Warnf(fmt.Sprintf("failed to get student by id: %s", err.Error()), zap.Error(err))
				return err
			}
			submissions, err = s.Repository.LearningManagement.GetAllSubmissionsByAssignmentID(ctx, student.StudentID, tx)
			if err != nil {
				s.Logger.Warnf(fmt.Sprintf("failed to get submission by id: %s", err.Error()), zap.Error(err))
				return err
			}
		default:
			err = pkg.NewBadRequestError("invalid role", nil)
			s.Logger.Warnf("invalid role: %s", user.Role, zap.Error(err))
			return
		}

		response.Submissions = make([]payload.GetSubmissionResponse, len(submissions))
		for i, submission := range submissions {
			response.Submissions[i].ID = submission.ID.String()
			response.Submissions[i].AssignmentID = submission.AssignmentID.String()
			response.Submissions[i].StudentID = submission.StudentID.String()
			response.Submissions[i].SubmittedAt = submission.SubmittedAt.Format(time.RFC3339)
			response.Submissions[i].Content = submission.Content
			if submission.FileURL != nil {
				response.Submissions[i].FileURL = *submission.FileURL
			}

			if submission.Grade != nil {
				response.Submissions[i].Grade = submission.Grade
			}
			if submission.Feedback != nil {
				response.Submissions[i].Feedback = submission.Feedback
			}
			if submission.GradedAt != nil {
				gradedAt := submission.GradedAt.Format(time.RFC3339)
				response.Submissions[i].GradedAt = &gradedAt
			}
			if submission.GradedBy != nil {
				response.Submissions[i].GradedBy = submission.GradedBy
			}
		}

		return
	})
}
