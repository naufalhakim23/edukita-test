package payload

type CreateAssignmentRequest struct {
	CreatedBy   string  `json:"created_by"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	TotalPoints float64 `json:"total_points"`
}

type UpdateAssignmentRequest struct {
	UserID      string  `json:"user_id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	TotalPoints float64 `json:"total_points"`
	IsPublished bool    `json:"is_published"`
}

type CreateSubmissionRequest struct {
	CreatedBy    string   `json:"created_by"`
	AssignmentID string   `json:"assignment_id"`
	Content      string   `json:"content"`
	FileURL      string   `json:"file_url"`
	Grade        *float64 `json:"grade"`
	Feedback     string   `json:"feedback"`
}

type UpdateSubmissionRequest struct {
	UserID       string   `json:"user_id"`
	AssignmentID string   `json:"assignment_id"`
	Content      string   `json:"content"`
	FileURL      string   `json:"file_url"`
	Grade        *float64 `json:"grade"`
	Feedback     string   `json:"feedback"`
}
