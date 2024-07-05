package usecases

import (
	"api-culinary-review/internal/models"
	"api-culinary-review/internal/repositories"
)

type ReviewUsecase interface {
	GetAllReviews() ([]models.Review, error)
	GetReviewByID(id uint) (*models.Review, error)
	CreateReview(req *models.ReviewRequest) (*models.Review, error)
	UpdateReviewByID(req *models.ReviewRequest, id uint) error
	DeleteReviewByID(id uint) error
}

type reviewUsecase struct {
	repo repositories.ReviewRepository
}

func NewReviewUsecase(repo repositories.ReviewRepository) ReviewUsecase {
	return &reviewUsecase{
		repo: repo,
	}
}

func (uc *reviewUsecase) GetAllReviews() ([]models.Review, error) {
	return uc.repo.FindAll()
}

func (uc *reviewUsecase) GetReviewByID(id uint) (*models.Review, error) {
	return uc.repo.FindByID(id)
}

func (uc *reviewUsecase) CreateReview(req *models.ReviewRequest) (*models.Review, error) {
	return uc.repo.Create(req)
}

func (uc *reviewUsecase) UpdateReviewByID(req *models.ReviewRequest, id uint) error {
	return uc.repo.UpdateByID(req, id)
}

func (uc *reviewUsecase) DeleteReviewByID(id uint) error {
	return uc.repo.DeleteByID(id)
}
