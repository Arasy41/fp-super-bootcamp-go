package repositories

import (
	"api-culinary-review/internal/models"

	"github.com/jinzhu/gorm"
)

type FavoriteRepository interface {
	GetByUserID(userID uint) ([]*models.Favorite, error)
	Create(favorite *models.Favorite) error
	FindByID(id uint) (*models.Favorite, error)
	Delete(id uint) error
}

type favoriteRepository struct {
	DB *gorm.DB
}

func NewFavoriteRepository(db *gorm.DB) FavoriteRepository {
	return &favoriteRepository{DB: db}
}

func (repo *favoriteRepository) GetByUserID(userID uint) ([]*models.Favorite, error) {
	var favorites []*models.Favorite
	err := repo.DB.Where("user_id = ?", userID).Find(&favorites).Error
	if err != nil {
		return nil, err
	}
	return favorites, nil
}

func (repo *favoriteRepository) Create(favorite *models.Favorite) error {
	return repo.DB.Create(favorite).Error
}

func (repo *favoriteRepository) FindByID(id uint) (*models.Favorite, error) {
	var favorite models.Favorite
	err := repo.DB.First(&favorite, id).Error
	if err != nil {
		return nil, err
	}
	return &favorite, nil
}

func (repo *favoriteRepository) Delete(id uint) error {
	return repo.DB.Delete(&models.Favorite{}, id).Error
}
