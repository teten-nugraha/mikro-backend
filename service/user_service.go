package service

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
	"github.com/teten-nugraha/mikro-backend/dto"
	"github.com/teten-nugraha/mikro-backend/helpers"
	"github.com/teten-nugraha/mikro-backend/mapper"
	"github.com/teten-nugraha/mikro-backend/respository"
	"github.com/teten-nugraha/mikro-backend/util"
)

type UserServiceContract interface {

	/**
	 * Mencari User by ID equal
	 */
	FindById(id uint) dto.UserDTO

	/*
	 * Mencari User by username equal
	 */
	FindByUsername(username string) dto.UserDTO

	/**
	 * Method untuk melakukan pencocokna password dari client dengan password yang ada di database
	 */
	CheckLogin(username, password string) (bool, error)

	/**
	 * Method untuk generate token sesudah user berhasil melakukan login
	 */
	GenerateToken(username string) (string, error)

	/**
	* Sign Up User
	 */
	SignUp(signupDto dto.SignupDTO) dto.UserDTO
}

type UserService struct {
	UserRepository respository.UserRepository
}

func ProviderUserService(u respository.UserRepository) UserService {
	return UserService{
		UserRepository: u,
	}
}

func (u *UserService) FindById(id uint) dto.UserDTO {
	return mapper.ToUserDTO(u.UserRepository.FindById(id))
}

func (u *UserService) FindByUsername(username string) dto.UserDTO {
	return mapper.ToUserDTO(u.UserRepository.FindByUsername(username))
}

func (u *UserService) CheckLogin(username, password string) (bool, error) {

	user := u.UserRepository.FindByUsername(username)

	match, err := helpers.CheckPasswordHash(password, user.Password)
	if !match {
		logrus.Warn("Hash and Password doesnt match")
		return false, err
	}

	return true, nil
}

func (u *UserService) GenerateToken(username string) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["level"] = "application"
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(util.SECRET))

	return t, err

}

func (u *UserService) SignUp(signupDto dto.SignupDTO) (dto.UserDTO, error) {

	hash, _ := helpers.HashPassword(signupDto.Password)

	signupDto.Password = hash
	signupDto.Role = util.MIKRO_USER
	signupDto.IsActive = true

	user, err := u.UserRepository.SaveOrUpdate(signupDto)
	if err != nil {
		return mapper.ToUserDTO(user), err
	}

	return mapper.ToUserDTO(user), nil
}
