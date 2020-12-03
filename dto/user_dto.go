package dto

import uuid "github.com/satori/go.uuid"

type UserDTO struct {
	ID       uuid.UUID `json:"id,string,omitempty"`
	Username string    `json:"username"`
	Fullname string    `json:"fullname"`
	Role     string    `json:"role"`
	IsActive bool      `json:"isActive"`
}
