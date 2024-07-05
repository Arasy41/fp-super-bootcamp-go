package usecases

import (
	"api-culinary-review/internal/models"
	"api-culinary-review/internal/repositories"
)

type TagUsecase struct {
	TagRepository repositories.TagRepository
}

func NewTagUsecase(tagRepo repositories.TagRepository) *TagUsecase {
	return &TagUsecase{
		TagRepository: tagRepo,
	}
}

func (uc *TagUsecase) CreateTag(name string) (*models.Tag, error) {
	tag := &models.Tag{
		Name: name,
	}

	err := uc.TagRepository.Create(tag)
	if err != nil {
		return nil, err
	}

	return tag, nil
}

func (uc *TagUsecase) GetTagByName(name string) (*models.Tag, error) {
	return uc.TagRepository.FindByName(name)
}

func (uc *TagUsecase) UpdateTag(tag *models.Tag) error {
	return uc.TagRepository.Update(tag)
}

func (uc *TagUsecase) DeleteTag(id uint) error {
	// Implement deletion logic as needed
	return uc.TagRepository.Delete(id)
}
