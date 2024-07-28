package repositories

import (
	"api-culinary-review/internal/models"

	"github.com/jinzhu/gorm"
)

type ReviewRepository interface {
	FindAll() ([]models.Review, error)
	FindByID(id uint) (*models.Review, error)
	Create(req *models.ReviewRequest) (*models.Review, error)
	UpdateReviewByID(review *models.Review, id uint) error
	DeleteByID(id uint) error
}

type reviewRepository struct {
	db *gorm.DB
}

func NewReviewRepository(db *gorm.DB) ReviewRepository {
	return &reviewRepository{
		db: db,
	}
}

func (repo *reviewRepository) FindAll() ([]models.Review, error) {
	var reviews []models.Review
	err := repo.db.Preload("User.Profile").Preload("Recipe.User").Find(&reviews).Error
	return reviews, err
}

func (repo *reviewRepository) FindByID(id uint) (*models.Review, error) {
	var review models.Review
	if err := repo.db.Preload("User").First(&review, id).Error; err != nil {
		return nil, err
	}
	return &review, nil
}

func (repo *reviewRepository) Create(req *models.ReviewRequest) (*models.Review, error) {
	review := models.Review{
		UserID:   req.UserID,
		RecipeID: req.RecipeID,
		Content:  req.Content,
	}
	err := repo.db.Create(&review).Error
	return &review, err
}

func (repo *reviewRepository) UpdateReviewByID(review *models.Review, id uint) error {
	return repo.db.Model(&models.Review{}).Where("id = ?", id).Updates(review).Error
}

func (repo *reviewRepository) DeleteByID(id uint) error {
	return repo.db.Delete(&models.Review{}).Where("id = ?", id).Error
}
