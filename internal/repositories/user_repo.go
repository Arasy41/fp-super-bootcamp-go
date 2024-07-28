package repositories

import (
	"api-culinary-review/internal/models"

	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	Create(user *models.User) error
	FindByID(id uint) (*models.User, error)
	Update(user *models.User) error
	GetUserByEmailOrUsername(emailOrUsername string) (*models.User, error)
	CheckUserEmail(email string) (*models.User, error)
	Delete(id uint) error
}

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{DB: db}
}

func (r *userRepository) Create(user *models.User) error {
	return r.DB.Create(user).Error
}

func (r *userRepository) FindByID(id uint) (*models.User, error) {
	var user models.User
	err := r.DB.Preload("Profile").Preload("Reviews").Preload("Favorites").First(&user, id).Error
	return &user, err
}

func (r *userRepository) Update(user *models.User) error {
	return r.DB.Save(user).Error
}

func (r *userRepository) GetUserByEmailOrUsername(emailOrUsername string) (*models.User, error) {
	var user models.User
	err := r.DB.Where("email = ? OR username = ?", emailOrUsername, emailOrUsername).First(&user).Error
	return &user, err
}

func (repo *userRepository) CheckUserEmail(email string) (*models.User, error) {
	var user models.User
	if err := repo.DB.Where("email = ?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Delete(id uint) error {
	return r.DB.Delete(&models.User{}, id).Error
}
