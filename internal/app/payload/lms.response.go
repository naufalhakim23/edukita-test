package payload

type CreateCourseResponse struct {
	ID string `json:"id"`
}

type UpdateCourseResponse struct {
	ID          string `json:"id"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	IsActive    bool   `json:"is_active"`
}

type GetCourseResponse struct {
	ID          string `json:"id"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	IsActive    bool   `json:"is_active"`
}

type GetAllCoursesResponse struct {
	Courses []GetCourseResponse `json:"courses"`
}

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
	CreatedAt    string   `json:"created_at"`
	CreatedBy    string   `json:"created_by"`
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

type (
	GetAllSubmissionsByCourseID struct {
		CourseID    string                     `json:"course_id"`
		Title       string                     `json:"title"`
		Description string                     `json:"description"`
		DueDate     string                     `json:"due_date"`
		CreatedBy   string                     `json:"created_by"`
		CreatedAt   string                     `json:"created_at"`
		Assignments []AssignmentAndSubmissions `json:"assignments"`
	}
	AssignmentAndSubmissions struct {
		AssignmentID string                  `json:"assignment_id"`
		Title        string                  `json:"title"`
		Description  string                  `json:"description"`
		DueDate      string                  `json:"due_date"`
		TotalPoints  float64                 `json:"total_points"`
		IsPublished  bool                    `json:"is_published"`
		CreatedAt    string                  `json:"created_at"`
		CreatedBy    string                  `json:"created_by"`
		Submissions  []GetSubmissionResponse `json:"submissions"`
	}
)
