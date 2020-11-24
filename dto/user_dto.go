package dto

type UserDTO struct {
	ID       uint   `json:"id,string,omitempty"`
	Username string `json:"username"`
	Fullname string `json:"fullname"`
	Role     string `json:"role"`
	IsActive bool   `json:"isActive"`
}
