package usecases

import (
	"api-culinary-review/internal/models"
	"api-culinary-review/internal/repositories"
)

type ImageUsecase struct {
	ImageRepository repositories.ImageRepository
}

func NewImageUsecase(imageRepo repositories.ImageRepository) *ImageUsecase {
	return &ImageUsecase{
		ImageRepository: imageRepo,
	}
}

func (uc *ImageUsecase) GetImageByID(id uint) (*models.Image, error) {
	return uc.ImageRepository.FindByID(id)
}

func (uc *ImageUsecase) DeleteImage(id uint) error {
	// Implement deletion logic as needed
	return uc.ImageRepository.Delete(id)
}
