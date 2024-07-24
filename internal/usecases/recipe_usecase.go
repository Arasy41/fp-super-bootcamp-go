package usecases

import (
	"api-culinary-review/internal/models"
	"api-culinary-review/internal/repositories"
	"api-culinary-review/pkg/utils"
	"fmt"
	"log"
	"mime/multipart"
)

type RecipeUsecase interface {
	CreateRecipe(images []*multipart.FileHeader, recipe *models.RecipeRequest, userID uint) (*models.Recipe, error)
	GetRecipeByID(id uint) (*models.Recipe, error)
	GetRecipes() ([]*models.Recipe, error)
	UpdateRecipe(id uint, images []*multipart.FileHeader, recipe *models.RecipeRequest) (*models.Recipe, error)
	DeleteRecipe(id uint) error
}

type recipeUsecase struct {
	recipeRepository repositories.RecipeRepository
}

func NewRecipeUsecase(recipeRepository repositories.RecipeRepository) RecipeUsecase {
	return &recipeUsecase{recipeRepository: recipeRepository}
}

func (r *recipeUsecase) GetRecipeByID(id uint) (*models.Recipe, error) {
	return r.recipeRepository.GetRecipeByID(id)
}

func (r *recipeUsecase) GetRecipes() ([]*models.Recipe, error) {
	return r.recipeRepository.GetRecipes()
}

func (r *recipeUsecase) CreateRecipe(images []*multipart.FileHeader, recipe *models.RecipeRequest, userID uint) (*models.Recipe, error) {
	newRecipe := &models.Recipe{
		Title:        recipe.Title,
		Description:  recipe.Description,
		Ingredients:  recipe.Ingredients,
		Instructions: recipe.Instructions,
		UserID:       userID,
	}

	// Create recipe first to get a valid ID
	createdRecipe, err := r.recipeRepository.CreateRecipe(newRecipe)
	if err != nil {
		return nil, err
	}

	// Create recipe tags
	if err := r.createRecipeTags(createdRecipe.ID, recipe.TagIDs); err != nil {
		return nil, err
	}

	// Upload images and set to createdRecipe
	for _, image := range images {
		uploadedImage, err := utils.UploadToCloudinary(image)
		if err != nil {
			return nil, err
		}

		// Tambahkan URL gambar ke dalam resep
		newRecipe.Images = append(newRecipe.Images, models.Image{URL: uploadedImage, RecipeID: createdRecipe.ID})
	}

	return r.recipeRepository.UpdateRecipe(createdRecipe)
}

func (r *recipeUsecase) UpdateRecipe(id uint, images []*multipart.FileHeader, recipe *models.RecipeRequest) (*models.Recipe, error) {
	existingRecipe, err := r.recipeRepository.GetRecipeByID(id)
	if err != nil {
		return nil, err
	}

	// Update recipe fields
	existingRecipe.Title = recipe.Title
	existingRecipe.Description = recipe.Description
	existingRecipe.Ingredients = recipe.Ingredients
	existingRecipe.Instructions = recipe.Instructions

	// Clear existing tags and create new ones
	if err := r.updateRecipeTags(id, recipe.TagIDs); err != nil {
		return nil, err
	}

	// Clear existing images and upload new ones
	if err := r.recipeRepository.DeleteRecipeImages(id); err != nil {
		return nil, err
	}

	for _, image := range images {
		uploadedImage, err := utils.UploadToCloudinary(image)
		if err != nil {
			return nil, err
		}

		// Add new image URLs to the recipe
		log.Printf("Uploading image URL: %s", uploadedImage)
		existingRecipe.Images = append(existingRecipe.Images, models.Image{URL: uploadedImage, RecipeID: existingRecipe.ID})
	}

	return r.recipeRepository.UpdateRecipe(existingRecipe)
}

// Helper function to create recipe tags
func (r *recipeUsecase) createRecipeTags(recipeID uint, tagIds []uint) error {
	for _, tagId := range tagIds {
		// Validate tagId exists in the tags table
		if exists, err := r.recipeRepository.RecipeTagExists(tagId); err != nil {
			return err
		} else if !exists {
			return fmt.Errorf("tag with ID %d does not exist", tagId)
		}

		if err := r.recipeRepository.CreateRecipeTag(recipeID, tagId); err != nil {
			return err
		}
	}
	return nil
}

func (r *recipeUsecase) updateRecipeTags(recipeID uint, tagIds []uint) error {
	// Hapus tag yang ada untuk resep tertentu
	if err := r.recipeRepository.DeleteRecipeTagsByRecipeID(recipeID); err != nil {
		return err
	}

	// Buat tag baru untuk resep tertentu
	for _, tagId := range tagIds {
		// Validasi tagId ada di tabel tags
		if exists, err := r.recipeRepository.RecipeTagExists(tagId); err != nil {
			return err
		} else if !exists {
			return fmt.Errorf("tag dengan ID %d tidak ada", tagId)
		}

		if err := r.recipeRepository.CreateRecipeTag(recipeID, tagId); err != nil {
			return err
		}
	}
	return nil
}

// Helper function to upload images and set to recipe
func (r *recipeUsecase) uploadRecipeImages(images []*multipart.FileHeader, recipe *models.Recipe) error {
	for _, image := range images {
		uploadedImage, err := utils.UploadToCloudinary(image)
		if err != nil {
			return err
		}

		// Tambahkan URL gambar ke dalam resep
		recipe.Images = append(recipe.Images, models.Image{URL: uploadedImage, RecipeID: recipe.ID})
	}

	// Simpan resep dengan URL gambar yang sudah diupdate ke database
	_, err := r.recipeRepository.UpdateRecipe(recipe)
	if err != nil {
		return err
	}

	return nil
}

func (r *recipeUsecase) DeleteRecipe(id uint) error {
	return r.recipeRepository.DeleteRecipe(id)
}
