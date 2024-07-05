package usecases

import (
	"api-culinary-review/internal/models"
	"api-culinary-review/internal/repositories"
	"time"
)

type FavoriteUsecase struct {
	FavoriteRepository repositories.FavoriteRepository
}

func NewFavoriteUsecase(favoriteRepo repositories.FavoriteRepository) *FavoriteUsecase {
	return &FavoriteUsecase{
		FavoriteRepository: favoriteRepo,
	}
}

func (uc *FavoriteUsecase) CreateFavorite(userID, recipeID uint) (*models.Favorite, error) {
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

func (uc *FavoriteUsecase) DeleteFavorite(id uint) error {
	if id == 0 {
		return nil
	}
	return uc.FavoriteRepository.Delete(id)
}
