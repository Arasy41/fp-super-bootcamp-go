// internal/repositories/image_repository.go

package repositories

import (
	"api-culinary-review/internal/models"
	"gorm.io/gorm"
)

type ImageRepository struct {
	DB *gorm.DB
}

func NewImageRepository(db *gorm.DB) *ImageRepository {
	return &ImageRepository{DB: db}
}

func (repo *ImageRepository) Create(image *models.Image) error {
	return repo.DB.Create(image).Error
}

func (repo *ImageRepository) FindByID(id uint) (*models.Image, error) {
	var image models.Image
	err := repo.DB.First(&image, id).Error
	if err != nil {
		return nil, err
	}
	return &image, nil
}

func (repo *ImageRepository) Update(image *models.Image) error {
	return repo.DB.Save(image).Error
}

func (repo *ImageRepository) Delete(id uint) error {
	return repo.DB.Delete(&models.Image{}, id).Error
}

// Implement other methods as needed
