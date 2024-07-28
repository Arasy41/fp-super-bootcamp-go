package usecases

import (
	"api-culinary-review/internal/models"
	"api-culinary-review/internal/repositories"
)

type TagUsecase interface {
	CreateTag(name string) (*models.Tag, error)
	GetAllTags() ([]models.Tag, error)
	GetTagsByNames(tagNames []string) ([]models.Tag, error)
	UpdateTag(tag *models.Tag) error
	DeleteTag(id uint) error
}

type tagUsecase struct {
	TagRepository repositories.TagRepository
}

func NewtagUsecase(tagRepo repositories.TagRepository) TagUsecase {
	return &tagUsecase{
		TagRepository: tagRepo,
	}
}

func (uc *tagUsecase) CreateTag(name string) (*models.Tag, error) {
	tag := &models.Tag{
		Name: name,
	}

	err := uc.TagRepository.Create(tag)
	if err != nil {
		return nil, err
	}

	return tag, nil
}

func (uc *tagUsecase) GetAllTags() ([]models.Tag, error) {
	return uc.TagRepository.GetAllTags()
}

func (uc *tagUsecase) GetTagsByNames(tagNames []string) ([]models.Tag, error) {
	return uc.TagRepository.GetTagsByNames(tagNames)
}

func (uc *tagUsecase) UpdateTag(tag *models.Tag) error {
	return uc.TagRepository.Update(tag)
}

func (uc *tagUsecase) DeleteTag(id uint) error {
	// Implement deletion logic as needed
	return uc.TagRepository.Delete(id)
}
