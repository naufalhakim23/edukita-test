package payload

type CreateAssignmentResponse struct {
	ID string `json:"id"`
}

type UpdateAssignmentResponse struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	DueDate     string  `json:"due_date"`
	TotalPoints float64 `json:"total_points"`
	IsPublished bool    `json:"is_published"`
}

type GetAssignmentResponse struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	DueDate     string  `json:"due_date"`
	TotalPoints float64 `json:"total_points"`
	IsPublished bool    `json:"is_published"`
}

type CreateSubmissionResponse struct {
	ID string `json:"id"`
}

type GetSubmissionResponse struct {
	ID           string   `json:"id"`
	AssignmentID string   `json:"assignment_id"`
	StudentID    string   `json:"student_id"`
	SubmittedAt  string   `json:"submitted_at"`
	Content      string   `json:"content"`
	FileURL      string   `json:"file_url"`
	Grade        *float64 `json:"grade"`
	Feedback     *string  `json:"feedback"`
	GradedAt     *string  `json:"graded_at"`
	GradedBy     *string  `json:"graded_by"`
}

type UpdateSubmissionResponse struct {
	ID           string   `json:"id"`
	AssignmentID string   `json:"assignment_id"`
	StudentID    string   `json:"student_id"`
	SubmittedAt  string   `json:"submitted_at"`
	Content      string   `json:"content"`
	FileURL      string   `json:"file_url"`
	Grade        *float64 `json:"grade"`
	Feedback     *string  `json:"feedback"`
	GradedAt     *string  `json:"graded_at"`
	GradedBy     *string  `json:"graded_by"`
}

type GetAllSubmissionsResponse struct {
	Submissions []GetSubmissionResponse `json:"submissions"`
}
