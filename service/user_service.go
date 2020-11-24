package service

import (
	"fmt"
	"github.com/teten-nugraha/mikro-backend/domain"
	"github.com/teten-nugraha/mikro-backend/helpers"
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

func (u *UserService) CheckLogin(username, password string) (bool, error) {

	user := u.UserRepository.FindByUsername(username)

	match, err := helpers.CheckPasswordHash(password, user.Password)
	if !match {
		fmt.Println("hash and password doesnt match")
		return false, err
	}

	return true,nil
}
