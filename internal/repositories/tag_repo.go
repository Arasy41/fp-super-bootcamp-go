package repositories

import (
	"api-culinary-review/internal/models"
	"gorm.io/gorm"
)

type TagRepository struct {
	DB *gorm.DB
}

func NewTagRepository(db *gorm.DB) *TagRepository {
	return &TagRepository{DB: db}
}

func (repo *TagRepository) Create(tag *models.Tag) error {
	return repo.DB.Create(tag).Error
}

func (repo *TagRepository) FindByName(name string) (*models.Tag, error) {
	var tag models.Tag
	err := repo.DB.Where("name = ?", name).First(&tag).Error
	if err != nil {
		return nil, err
	}
	return &tag, nil
}

func (repo *TagRepository) Update(tag *models.Tag) error {
	return repo.DB.Save(tag).Error
}

func (repo *TagRepository) Delete(id uint) error {
	return repo.DB.Delete(&models.Tag{}, id).Error
}
