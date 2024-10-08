package usecases

import (
	"api-culinary-review/internal/models"
	"api-culinary-review/internal/repositories"
	"api-culinary-review/pkg/utils"
	"time"
)

type UserUsecase interface {
	CreateUser(username, password, email string) (*models.User, error)
	GetUserByID(id uint) (*models.User, error)
	GetUserByEmailOrUsername(emailOrUsername string) (*models.User, error)
	CheckUserEmail(email string) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id uint) error
}

type userUsecase struct {
	UserRepository    repositories.UserRepository
	ProfileRepository repositories.ProfileRepository
}

func NewUserUsecase(userRepo repositories.UserRepository, profileRepo repositories.ProfileRepository) UserUsecase {
	return &userUsecase{
		UserRepository:    userRepo,
		ProfileRepository: profileRepo,
	}
}

func (uc *userUsecase) CreateUser(username, password, email string) (*models.User, error) {
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Username:  username,
		Password:  hashedPassword,
		Email:     email,
		CreatedAt: time.Now(),
	}

	err = uc.UserRepository.Create(user)
	if err != nil {
		return nil, err
	}

	profile := models.Profile{
		UserID:   user.ID,
		FullName: username,
	}

	err = uc.ProfileRepository.CreateProfile(&profile)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (uc *userUsecase) GetUserByID(id uint) (*models.User, error) {
	return uc.UserRepository.FindByID(id)
}

func (uc *userUsecase) UpdateUser(user *models.User) error {
	return uc.UserRepository.Update(user)
}

func (uc *userUsecase) DeleteUser(id uint) error {
	return uc.UserRepository.Delete(id)
}

func (uc *userUsecase) GetUserByEmailOrUsername(emailOrUsername string) (*models.User, error) {
	return uc.UserRepository.GetUserByEmailOrUsername(emailOrUsername)
}

func (uc *userUsecase) CheckUserEmail(email string) (*models.User, error) {
	return uc.UserRepository.CheckUserEmail(email)
}
