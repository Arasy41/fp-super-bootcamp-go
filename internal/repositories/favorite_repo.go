package repositories

import (
	"api-culinary-review/internal/models"
	"gorm.io/gorm"
)

type FavoriteRepository struct {
	DB *gorm.DB
}

func NewFavoriteRepository(db *gorm.DB) *FavoriteRepository {
	return &FavoriteRepository{DB: db}
}

func (repo *FavoriteRepository) Create(favorite *models.Favorite) error {
	return repo.DB.Create(favorite).Error
}

func (repo *FavoriteRepository) FindByID(id uint) (*models.Favorite, error) {
	var favorite models.Favorite
	err := repo.DB.First(&favorite, id).Error
	if err != nil {
		return nil, err
	}
	return &favorite, nil
}

func (repo *FavoriteRepository) Delete(id uint) error {
	return repo.DB.Delete(&models.Favorite{}, id).Error
}
