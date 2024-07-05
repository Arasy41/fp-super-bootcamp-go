package repositories

import (
	"api-culinary-review/internal/models"
	"api-culinary-review/pkg/utils"
	"mime/multipart"

	"gorm.io/gorm"
)

type RecipeRepository interface {
	FindAll() ([]models.GetAllRecipeResponse, error)
	FindByID(id uint) (*models.Recipe, error)
	Create(req *models.RecipeRequest, images []*multipart.FileHeader) (*models.Recipe, error)
	UpdateByID(req *models.RecipeRequest, id uint, images []*multipart.FileHeader) error
	DeleteByID(id uint) error
}

type recipeRepository struct {
	db *gorm.DB
}

func NewRecipeRepository(db *gorm.DB) RecipeRepository {
	return &recipeRepository{
		db: db,
	}
}

func (repo *recipeRepository) FindAll() ([]models.GetAllRecipeResponse, error) {
	var recipes []models.GetAllRecipeResponse
	err := repo.db.Preload("Images").Find(&recipes).Error
	return recipes, err
}

func (repo *recipeRepository) FindByID(id uint) (*models.Recipe, error) {
	var recipe models.Recipe
	err := repo.db.Preload("Tags").Preload("Images").Preload("Reviews").First(&recipe, id).Error
	return &recipe, err
}

func (repo *recipeRepository) Create(req *models.RecipeRequest, images []*multipart.FileHeader) (*models.Recipe, error) {
	var tags []models.Tag
	for _, tagID := range req.TagsIds {
		tags = append(tags, models.Tag{ID: tagID})
	}

	recipe := models.Recipe{
		Title:        req.Title,
		Description:  req.Description,
		Ingredients:  req.Ingredients,
		Instructions: req.Instructions,
		Tags:         tags,
	}

	// Upload images to Cloudinary and store URLs in recipe.Images
	for _, image := range images {
		imageURL, err := utils.UploadToCloudinary(image)
		if err != nil {
			return nil, err
		}
		recipe.Images = append(recipe.Images, models.Image{URL: imageURL})
	}

	err := repo.db.Create(&recipe).Error
	return &recipe, err
}

func (repo *recipeRepository) UpdateByID(req *models.RecipeRequest, id uint, images []*multipart.FileHeader) error {
	var recipe models.Recipe
	if err := repo.db.Preload("Images").First(&recipe, id).Error; err != nil {
		return err
	}

	// Update the fields
	recipe.Title = req.Title
	recipe.Description = req.Description
	recipe.Ingredients = req.Ingredients
	recipe.Instructions = req.Instructions

	// Convert req.TagsIds to []models.Tag
	var tags []models.Tag
	for _, tagID := range req.TagsIds {
		tags = append(tags, models.Tag{ID: tagID})
	}
	recipe.Tags = tags

	// Upload new images to Cloudinary and update URLs in recipe.Images
	for _, image := range images {
		imageURL, err := utils.UploadToCloudinary(image)
		if err != nil {
			return err
		}
		recipe.Images = append(recipe.Images, models.Image{URL: imageURL})
	}

	// Save the updated recipe
	if err := repo.db.Save(&recipe).Error; err != nil {
		return err
	}

	return nil
}

func (repo *recipeRepository) DeleteByID(id uint) error {
	var recipe models.Recipe

	// Delete images from Cloudinary
	for _, image := range recipe.Images {
		if err := utils.DeleteImageFromCloudinary(image.URL); err != nil {
			return err
		}
	}

	return repo.db.Where("id = ?", id).Delete(&recipe).Error
}
