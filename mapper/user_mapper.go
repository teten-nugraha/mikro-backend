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

func ToUserEntity(dto dto.UserDTO) domain.User {
	return domain.User{
		Username: dto.Username,
		Fullname: dto.Fullname,
		Role: dto.Role,
		IsActive: dto.IsActive,
	}
}

func ToUserEntityFromSignupDto(dto dto.SignupDTO) domain.User {
	return domain.User{
		Username: dto.Username,
		Password: dto.Password,
		Fullname: dto.Fullname,
		Role: dto.Role,
		IsActive: dto.IsActive,
		Email: dto.Email,
		Phone: dto.Email,
	}
}

func ToUserDTOs(users []domain.User) []dto.UserDTO {
	userdtos := make([]dto.UserDTO, len(users))

	for i, itm := range users {
		userdtos[i] = ToUserDTO(itm)
	}

	return userdtos
}
