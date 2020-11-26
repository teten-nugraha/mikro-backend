package respository

import (
	"github.com/jinzhu/gorm"
	"github.com/teten-nugraha/mikro-backend/domain"
)

type UserRepositoryContract interface {

	/**
	 * Mencari User by ID equal
	 */
    FindById(id uint) domain.User

    /*
     * Mencari User by username equal
     */
    FindByUsername(username string) domain.User
}

type UserRepository struct {
	DB *gorm.DB
}

func ProviderUserRepository(DB *gorm.DB) UserRepository {
	return UserRepository{DB: DB}
}

func (u *UserRepository) FindById(id uint) domain.User {
	var user domain.User

	u.DB.First(&user, id)

	return user
}

func (u *UserRepository) FindByUsername(username string) domain.User {
	var user domain.User
	u.DB.Where("username = ?", username).Find(&user)

	return user
}
