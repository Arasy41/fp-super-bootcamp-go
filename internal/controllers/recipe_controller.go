package controllers

import (
	"api-culinary-review/internal/models"
	"api-culinary-review/internal/usecases"
	"api-culinary-review/pkg/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// RecipeController interface defines the methods for handling recipe-related operations.
type RecipeController interface {
	CreateRecipe(c *gin.Context)
	GetRecipeByID(c *gin.Context)
	GetRecipes(c *gin.Context)
	UpdateRecipe(c *gin.Context)
	DeleteRecipe(c *gin.Context)
}

type recipeController struct {
	recipeUsecase usecases.RecipeUsecase
	tagUsecase    usecases.TagUsecase
}

// NewRecipeController creates a new instance of RecipeController.
func NewRecipeController(recipeUsecase usecases.RecipeUsecase, tagUsecase usecases.TagUsecase) RecipeController {
	return &recipeController{
		recipeUsecase: recipeUsecase,
		tagUsecase:    tagUsecase,
	}
}

// CreateRecipe creates a new recipe.
// @Summary Create a new recipe
// @Description Creates a new recipe with the provided details.
// @Tags recipes
// @Accept multipart/form-data
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param title formData string true "Title of the recipe"
// @Param description formData string true "Description of the recipe"
// @Param ingredients formData string true "Ingredients of the recipe"
// @Param instructions formData string true "Instructions of the recipe"
// @Param images formData file true "Images of the recipe"
// @Param tag_names formData string true "Tag names in JSON array format"
// @Success 201 {object} models.Recipe
// @Security ApiKeyAuth
// @Router /api/recipes [post]
func (c *recipeController) CreateRecipe(ctx *gin.Context) {
	recipeRequest := &models.RecipeRequest{}

	// Extract fields from form-data
	title := ctx.Request.FormValue("title")
	description := ctx.Request.FormValue("description")
	ingredients := ctx.Request.FormValue("ingredients")
	instructions := ctx.Request.FormValue("instructions")
	images, _ := ctx.Request.MultipartForm.File["images"]

	// Check if form-data is empty
	if title == "" || description == "" || ingredients == "" || instructions == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "All fields are required"})
		return
	}

	tagNamesStr := ctx.PostForm("tag_names")
	fmt.Println("Received tag_names:", tagNamesStr) // Debugging log

	var tagNames []string
	if err := json.Unmarshal([]byte(tagNamesStr), &tagNames); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tag_names format"})
		return
	}

	// Get tag IDs based on tag names
	tags, err := c.tagUsecase.GetTagsByNames(tagNames)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var tagIDs []uint
	for _, tag := range tags {
		tagIDs = append(tagIDs, tag.ID)
	}

	// Upload images to Cloudinary and get URLs
	var imageUrls []string
	for _, file := range images {
		uploadedImage, err := utils.UploadToCloudinary(file)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload image"})
			return
		}
		imageUrls = append(imageUrls, uploadedImage)
	}

	userID, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	userIDUint, ok := userID.(uint)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse user ID"})
		return
	}

	// Populate recipeRequest struct
	recipeRequest.Title = title
	recipeRequest.Description = description
	recipeRequest.Ingredients = ingredients
	recipeRequest.Instructions = instructions
	recipeRequest.Images = images
	recipeRequest.TagIDs = tagIDs

	recipe, err := c.recipeUsecase.CreateRecipe(recipeRequest.Images, recipeRequest, userIDUint)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, recipe)
}

// UpdateRecipe updates an existing recipe.
// @Summary Update an existing recipe
// @Description Updates an existing recipe with the provided details.
// @Tags recipes
// @Accept multipart/form-data
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param id path int true "Recipe ID"
// @Param title formData string true "Title of the recipe"
// @Param description formData string true "Description of the recipe"
// @Param ingredients formData string true "Ingredients of the recipe"
// @Param instructions formData string true "Instructions of the recipe"
// @Param images formData file true "Images of the recipe"
// @Param tag_names formData string true "Tag names in JSON array format"
// @Success 200 {object} models.Recipe
// @Security ApiKeyAuth
// @Router /api/recipes/{id} [put]
func (c *recipeController) UpdateRecipe(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	recipeRequest := &models.RecipeRequest{}

	// Extract fields from form-data
	title := ctx.Request.FormValue("title")
	description := ctx.Request.FormValue("description")
	ingredients := ctx.Request.FormValue("ingredients")
	instructions := ctx.Request.FormValue("instructions")
	images, _ := ctx.Request.MultipartForm.File["images"]

	// Parsing tag_names
	tagNamesStr := ctx.PostForm("tag_names")
	var tagNames []string
	if err := json.Unmarshal([]byte(tagNamesStr), &tagNames); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tag_names format"})
		return
	}

	// Get tag IDs based on tag names
	tags, err := c.tagUsecase.GetTagsByNames(tagNames)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var tagIDs []uint
	for _, tag := range tags {
		tagIDs = append(tagIDs, tag.ID)
	}

	// Populate recipeRequest struct
	recipeRequest.Title = title
	recipeRequest.Description = description
	recipeRequest.Ingredients = ingredients
	recipeRequest.Instructions = instructions
	recipeRequest.TagIDs = tagIDs // Set tag IDs

	recipe, err := c.recipeUsecase.UpdateRecipe(uint(id), images, recipeRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, recipe)
}

// DeleteRecipe deletes an existing recipe.
// @Summary Delete a recipe
// @Description Deletes a recipe by ID.
// @Tags recipes
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param id path int true "Recipe ID"
// @Success 204 {object} nil
// @Security ApiKeyAuth
// @Router /api/recipes/{id} [delete]
func (c *recipeController) DeleteRecipe(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = c.recipeUsecase.DeleteRecipe(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Recipe deleted successfully",
	})
}

// GetRecipeByID retrieves a recipe by its ID.
// @Summary Get a recipe by ID
// @Description Retrieves a recipe by its ID.
// @Tags recipes
// @Produce json
// @Param id path int true "Recipe ID"
// @Success 200 {object} models.Recipe
// @Router /api/recipes/{id} [get]
func (c *recipeController) GetRecipeByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	recipe, err := c.recipeUsecase.GetRecipeByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Recipe not found"})
		return
	}

	ctx.JSON(http.StatusOK, recipe)
}

// GetRecipes retrieves all recipes.
// @Summary Get all recipes
// @Description Retrieves all recipes.
// @Tags recipes
// @Produce json
// @Success 200 {object} []models.Recipe
// @Router /api/recipes [get]
func (c *recipeController) GetRecipes(ctx *gin.Context) {
	recipes, err := c.recipeUsecase.GetRecipes()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, recipes)
}
