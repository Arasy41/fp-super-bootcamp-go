package repositories

import (
	"api-culinary-review/internal/models"

	"github.com/jinzhu/gorm"
)

type RecipeRepository interface {
	CreateRecipe(recipe *models.Recipe) (*models.Recipe, error)
	GetRecipeByID(id uint) (*models.Recipe, error)
	GetRecipes() ([]*models.Recipe, error)
	UpdateRecipe(recipe *models.Recipe) (*models.Recipe, error)
	DeleteRecipe(id uint) error
	CreateRecipeTag(recipeId uint, tagId uint) error
	RecipeTagExists(tagId uint) (bool, error)
	DeleteRecipeTagsByRecipeID(recipeID uint) error
	DeleteRecipeImages(recipeID uint) error
}

type recipeRepository struct {
	db *gorm.DB
}

func NewRecipeRepository(db *gorm.DB) RecipeRepository {
	return &recipeRepository{db: db}
}

func (r *recipeRepository) CreateRecipe(recipe *models.Recipe) (*models.Recipe, error) {
	err := r.db.Create(recipe).Error
	return recipe, err
}

func (r *recipeRepository) GetRecipeByID(id uint) (*models.Recipe, error) {
	var recipe models.Recipe
	err := r.db.Preload("User").
		Preload("Tags").
		Preload("Images").
		Preload("Reviews").
		First(&recipe, id).Error
	if err != nil {
		return nil, err
	}
	return &recipe, nil
}

func (r *recipeRepository) GetRecipes() ([]*models.Recipe, error) {
	var recipes []*models.Recipe
	err := r.db.Preload("Tags").Preload("Images").Find(&recipes).Error
	return recipes, err
}

func (r *recipeRepository) UpdateRecipe(recipe *models.Recipe) (*models.Recipe, error) {
	err := r.db.Save(recipe).Error
	return recipe, err
}

func (r *recipeRepository) DeleteRecipe(id uint) error {
	return r.db.Delete(&models.Recipe{}, id).Error
}

func (r *recipeRepository) CreateRecipeTag(recipeId uint, tagId uint) error {
	recipeTag := &models.RecipeTag{
		RecipeID: recipeId,
		TagID:    tagId,
	}
	return r.db.Create(recipeTag).Error
}

func (r *recipeRepository) RecipeTagExists(tagId uint) (bool, error) {
	var count int64
	err := r.db.Model(&models.Tag{}).Where("id = ?", tagId).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *recipeRepository) DeleteRecipeTagsByRecipeID(recipeID uint) error {
	return r.db.Where("recipe_id = ?", recipeID).Delete(&models.RecipeTag{}).Error
}

func (r *recipeRepository) DeleteRecipeImages(recipeID uint) error {
	return r.db.Where("recipe_id = ?", recipeID).Delete(&models.Image{}).Error
}
