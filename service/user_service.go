package service

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/teten-nugraha/mikro-backend/domain"
	"github.com/teten-nugraha/mikro-backend/helpers"
	"github.com/teten-nugraha/mikro-backend/respository"
	"time"
)

type UserService struct {
	UserRepository respository.UserRepository
}

func ProviderUserService(u respository.UserRepository) UserService {
	return UserService{
		UserRepository: u,
	}
}

func (u *UserService) FindById(id uint) domain.User {
	return u.UserRepository.FindById(id)
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

func (u *UserService) GenerateToken(username string) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["level"] = "application"
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte("secret"))

	return t, err

}