package usecases

import (
	"api-culinary-review/internal/models"
	"api-culinary-review/internal/repositories"
	"api-culinary-review/pkg/utils"
	"mime/multipart"
)

type ProfileUsecase interface {
	CreateProfile(req *models.ProfileRequest, userID uint, file *multipart.FileHeader) (*models.Profile, error)
	GetProfileByUserID(userID uint) (*models.Profile, error)
	UpdateProfileByID(req *models.ProfileRequest, userID uint, file *multipart.FileHeader) error
}

type profileUsecase struct {
	repo repositories.ProfileRepository
}

func NewProfileUsecase(repo repositories.ProfileRepository) ProfileUsecase {
	return &profileUsecase{repo: repo}
}

func (uc *profileUsecase) CreateProfile(req *models.ProfileRequest, userID uint, file *multipart.FileHeader) (*models.Profile, error) {
	avatarURL, err := utils.UploadToCloudinary(file)
	if err != nil {
		return nil, err
	}

	profile := &models.Profile{
		UserID:    userID,
		FullName:  req.FullName,
		Bio:       req.Bio,
		AvatarURL: avatarURL,
	}

	if err := uc.repo.CreateProfile(profile); err != nil {
		return nil, err
	}

	return profile, nil
}

func (uc *profileUsecase) GetProfileByUserID(userID uint) (*models.Profile, error) {
	return uc.repo.GetProfileByUserID(userID)
}

func (uc *profileUsecase) UpdateProfileByID(req *models.ProfileRequest, userID uint, file *multipart.FileHeader) error {
	profile, err := uc.repo.GetProfileByUserID(userID)
	if err != nil {
		return err
	}

	avatarURL, err := utils.UploadToCloudinary(file)
	if err != nil {
		return err
	}

	profile.FullName = req.FullName
	profile.Bio = req.Bio
	profile.AvatarURL = avatarURL

	return uc.repo.UpdateProfile(profile)
}
