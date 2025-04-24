package payload

type RegisterUserResponse struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Role      string `json:"role"`
}

type LoginUserResponse struct {
	Token string `json:"token"`
}

type GetUserResponse struct {
	ID        string           `json:"id"`
	FirstName string           `json:"first_name"`
	LastName  string           `json:"last_name"`
	Email     string           `json:"email"`
	UserRole  RoleUserResponse `json:"user_role"`
	IsActive  bool             `json:"is_active"`
	LastLogin string           `json:"last_login"`
	CreatedAt string           `json:"created_at"`
	UpdatedAt string           `json:"updated_at"`
}

type RoleUserResponse struct {
	Role           string `json:"role"`
	StudentID      string `json:"student_id,omitempty"`
	EnrollmentYear int    `json:"enrollment_year,omitempty"`
	Program        string `json:"program,omitempty"`
	Department     string `json:"department,omitempty"`
	Title          string `json:"title,omitempty"`
}

type LogoutUserResponse struct {
	ID string `json:"id"`
}
