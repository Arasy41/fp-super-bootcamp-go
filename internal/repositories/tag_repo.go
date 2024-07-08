package repositories

import (
	"api-culinary-review/internal/models"

	"github.com/jinzhu/gorm"
)

type TagRepository interface {
	Create(tag *models.Tag) error
	GetTagsByNames(names []string) ([]models.Tag, error)
	Update(tag *models.Tag) error
	Delete(id uint) error
}

type tagRepository struct {
	DB *gorm.DB
}

func NewTagRepository(db *gorm.DB) TagRepository {
	return &tagRepository{DB: db}
}

func (repo *tagRepository) Create(tag *models.Tag) error {
	return repo.DB.Create(tag).Error
}

func (r *tagRepository) GetTagsByNames(names []string) ([]models.Tag, error) {
	var tags []models.Tag
	if err := r.DB.Where("name IN (?)", names).Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil
}

func (repo *tagRepository) Update(tag *models.Tag) error {
	return repo.DB.Save(tag).Error
}

func (repo *tagRepository) Delete(id uint) error {
	return repo.DB.Delete(&models.Tag{}, id).Error
}
