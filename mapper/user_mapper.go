package mapper

import (
	"github.com/teten-nugraha/mikro-backend/domain"
	"github.com/teten-nugraha/mikro-backend/dto"
)

func ToUserDTO(user domain.User) dto.UserDTO {
	return dto.UserDTO{
		ID: user.ID,
		Username: user.Username,
		Fullname: user.Fullname,
		Role: user.Role,
		IsActive: user.IsActive,
	}
}

func ToProductDTOs(users []domain.User) []dto.UserDTO {
	userdtos := make([]dto.UserDTO, len(users))

	for i, itm := range users {
		userdtos[i] = ToUserDTO(itm)
	}

	return userdtos
}
