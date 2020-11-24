package service

import (
	"github.com/teten-nugraha/mikro-backend/domain"
	"github.com/teten-nugraha/mikro-backend/respository"
)

type UserService struct {
	UserRepository respository.UserRepository
}

func ProviderUserService(u respository.UserRepository) UserService {
	return UserService{
		UserRepository: u,
	}
}

func (u *UserService) FindByUsername(username string) domain.User {
	return u.UserRepository.FindByUsername(username)
}

