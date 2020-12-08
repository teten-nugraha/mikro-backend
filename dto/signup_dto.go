package dto

type SignupDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Fullname string `json:"fullname"`
	Role     string `json:"role"`
	IsActive bool   `json:"isActive"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}
