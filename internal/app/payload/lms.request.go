package payload

type CreateAssignmentRequest struct {
	CreatedBy   string  `json:"created_by" validate:"required"`
	Title       string  `json:"title" validate:"required"`
	Description string  `json:"description" validate:"required"`
	TotalPoints float64 `json:"total_points" validate:"required"`
}

type UpdateAssignmentRequest struct {
	UserID      string  `json:"user_id" validate:"required"`
	Title       string  `json:"title" validate:"required"`
	Description string  `json:"description" validate:"required"`
	TotalPoints float64 `json:"total_points" validate:"required"`
	IsPublished bool    `json:"is_published" validate:"required"`
}

type CreateSubmissionRequest struct {
	CreatedBy    string   `json:"created_by" validate:"required"`
	AssignmentID string   `json:"assignment_id" validate:"required"`
	Content      string   `json:"content"`
	FileURL      string   `json:"file_url"`
	Grade        *float64 `json:"grade"`
	Feedback     *string  `json:"feedback"`
}

type UpdateSubmissionRequest struct {
	UserID       string   `json:"user_id" validate:"required"`
	AssignmentID string   `json:"assignment_id" validate:"required"`
	Content      string   `json:"content"`
	FileURL      string   `json:"file_url"`
	Grade        *float64 `json:"grade"`
	Feedback     *string  `json:"feedback"`
}
