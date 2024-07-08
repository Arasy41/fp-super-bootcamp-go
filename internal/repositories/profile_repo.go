// repositories/profile_repository.go

package repositories

import (
	"api-culinary-review/internal/models"

	"github.com/jinzhu/gorm"
)

type ProfileRepository interface {
	CreateProfile(profile *models.Profile) error
	GetProfileByUserID(userID uint) (*models.Profile, error)
	UpdateProfile(profile *models.Profile) error
}

type profileRepository struct {
	db *gorm.DB
}

func NewProfileRepository(db *gorm.DB) ProfileRepository {
	return &profileRepository{db: db}
}

func (r *profileRepository) CreateProfile(profile *models.Profile) error {
	return r.db.Create(profile).Error
}

func (r *profileRepository) GetProfileByUserID(userID uint) (*models.Profile, error) {
	var profile models.Profile
	err := r.db.Where("user_id = ?", userID).First(&profile).Error
	return &profile, err
}

func (r *profileRepository) UpdateProfile(profile *models.Profile) error {
	return r.db.Save(profile).Error
}
