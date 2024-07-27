package repositories

import (
	"api-culinary-review/internal/models"

	"github.com/jinzhu/gorm"
)

type ReviewRepository interface {
	FindAll() ([]models.Review, error)
	FindByID(id uint) (*models.Review, error)
	Create(req *models.ReviewRequest) (*models.Review, error)
	UpdateByID(req *models.ReviewRequest, id uint) error
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
	err := repo.db.Preload("User").Preload("Recipe").Where("id = ?", id).First(&review).Error
	return &review, err
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

func (repo *reviewRepository) UpdateByID(req *models.ReviewRequest, id uint) error {
	var review models.Review
	if err := repo.db.Where("id = ?", id).First(&review).Error; err != nil {
		return err
	}
	review.Content = req.Content
	return repo.db.Save(&review).Error
}

func (repo *reviewRepository) DeleteByID(id uint) error {
	return repo.db.Delete(&models.Review{}).Where("id = ?", id).Error
}
