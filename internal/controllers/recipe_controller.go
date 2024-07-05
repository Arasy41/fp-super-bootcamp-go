package controllers

import (
	"api-culinary-review/internal/models"
	"api-culinary-review/internal/usecases"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RecipeController interface {
	GetAllRecipes(c *gin.Context)
	GetRecipeByID(c *gin.Context)
	CreateRecipe(c *gin.Context)
	UpdateRecipeByID(c *gin.Context)
	DeleteRecipeByID(c *gin.Context)
}

type recipeController struct {
	uc usecases.RecipeUsecase
}

func NewRecipeController(usecase usecases.RecipeUsecase) RecipeController {
	return &recipeController{
		uc: usecase,
	}
}

func (ctrl *recipeController) GetAllRecipes(c *gin.Context) {
	recipes, err := ctrl.uc.GetAllRecipes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": recipes})
}

func (ctrl *recipeController) GetRecipeByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	recipe, err := ctrl.uc.GetRecipeByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": recipe})
}

func (ctrl *recipeController) CreateRecipe(c *gin.Context) {
	var req models.RecipeRequest

	// Bind form-data ke struct RecipeRequest
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tagsIdStr := c.PostForm("tags")
	tagsId, err := strconv.ParseUint(tagsIdStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tag ID"})
		return
	}
	req.TagsIds = append(req.TagsIds, uint(tagsId))

	req = models.RecipeRequest{
		Title:        c.PostForm("title"),
		Description:  c.PostForm("description"),
		Ingredients:  c.PostForm("ingredients"),
		Instructions: c.PostForm("instructions"),
		TagsIds:      req.TagsIds,
	}

	// Handle file uploads
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	files := form.File["images"]

	// Konversi dari FileHeader ke []*multipart.FileHeader
	imageHeaders := files

	recipe, err := ctrl.uc.CreateRecipe(&req, imageHeaders)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": recipe})
}

func (ctrl *recipeController) UpdateRecipeByID(c *gin.Context) {
	var req models.RecipeRequest

	// Bind form-data ke struct RecipeRequest
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	tagsIdStr := c.PostForm("tags")
	tagsId, err := strconv.ParseUint(tagsIdStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tag ID"})
		return
	}
	req.TagsIds = append(req.TagsIds, uint(tagsId))

	req = models.RecipeRequest{
		Title:        c.PostForm("title"),
		Description:  c.PostForm("description"),
		Ingredients:  c.PostForm("ingredients"),
		Instructions: c.PostForm("instructions"),
		TagsIds:      req.TagsIds,
	}

	// Handle file uploads
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	files := form.File["images"]

	// Convert from FileHeader to []*multipart.FileHeader
	imageHeaders := files

	if err := ctrl.uc.UpdateRecipeByID(&req, uint(id), imageHeaders); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Recipe updated successfully"})
}

func (ctrl *recipeController) DeleteRecipeByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := ctrl.uc.DeleteRecipeByID(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Recipe deleted successfully"})
}
