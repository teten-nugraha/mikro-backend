package dto

type SignupDTO struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Fullname string `json:"fullname" validate:"required"`
	Role     string `json:"role"`
	IsActive bool   `json:"isActive"`
	Email    string `json:"email" validate:"email"`
	Phone    string `json:"phone"`
}
