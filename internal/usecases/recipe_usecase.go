package usecases

import (
	"api-culinary-review/internal/models"
	"api-culinary-review/internal/repositories"
	"mime/multipart"
)

type RecipeUsecase interface {
	GetAllRecipes() ([]models.GetAllRecipeResponse, error)
	GetRecipeByID(id uint) (*models.Recipe, error)
	CreateRecipe(req *models.RecipeRequest, images []*multipart.FileHeader) (*models.Recipe, error)
	UpdateRecipeByID(req *models.RecipeRequest, id uint, images []*multipart.FileHeader) error
	DeleteRecipeByID(id uint) error
}

type recipeUsecase struct {
	repo repositories.RecipeRepository
}

func NewRecipeUsecase(repo repositories.RecipeRepository) RecipeUsecase {
	return &recipeUsecase{
		repo: repo,
	}
}

func (uc *recipeUsecase) GetAllRecipes() ([]models.GetAllRecipeResponse, error) {
	return uc.repo.FindAll()
}

func (uc *recipeUsecase) GetRecipeByID(id uint) (*models.Recipe, error) {
	return uc.repo.FindByID(id)
}

func (uc *recipeUsecase) CreateRecipe(req *models.RecipeRequest, images []*multipart.FileHeader) (*models.Recipe, error) {
	return uc.repo.Create(req, images)
}

func (uc *recipeUsecase) UpdateRecipeByID(req *models.RecipeRequest, id uint, images []*multipart.FileHeader) error {
	return uc.repo.UpdateByID(req, id, images)
}

func (uc *recipeUsecase) DeleteRecipeByID(id uint) error {
	return uc.repo.DeleteByID(id)
}
