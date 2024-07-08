package usecases

import (
	"api-culinary-review/internal/models"
	"api-culinary-review/internal/repositories"
	"time"
)

type FavoriteUsecase interface {
	GetByUserID(userID uint) ([]*models.Favorite, error)
	CreateFavorite(userID, recipeID uint) (*models.Favorite, error)
	DeleteFavorite(id uint) error
}

type favoriteUsecase struct {
	FavoriteRepository repositories.FavoriteRepository
}

func NewFavoriteUsecase(favoriteRepo repositories.FavoriteRepository) FavoriteUsecase {
	return &favoriteUsecase{
		FavoriteRepository: favoriteRepo,
	}
}

func (uc *favoriteUsecase) GetByUserID(userID uint) ([]*models.Favorite, error) {
	favorites, err := uc.FavoriteRepository.GetByUserID(userID)
	if err != nil {
		return nil, err
	}
	return favorites, nil
}

func (uc *favoriteUsecase) CreateFavorite(userID, recipeID uint) (*models.Favorite, error) {
	favorite := &models.Favorite{
		UserID:    userID,
		RecipeID:  recipeID,
		CreatedAt: time.Now(),
	}

	err := uc.FavoriteRepository.Create(favorite)
	if err != nil {
		return nil, err
	}

	return favorite, nil
}

func (uc *favoriteUsecase) DeleteFavorite(id uint) error {
	if id == 0 {
		return nil
	}
	return uc.FavoriteRepository.Delete(id)
}
