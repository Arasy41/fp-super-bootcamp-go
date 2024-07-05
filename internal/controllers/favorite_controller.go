package controllers

// import (
// 	"api-culinary-review/internal/models"
// 	"api-culinary-review/internal/usecases"
// 	"github.com/gin-gonic/gin"
// 	"net/http"
// )

// type FavoriteController struct {
// 	FavoriteUsecase usecases.FavoriteUsecase
// }

// func NewFavoriteController(favoriteUC usecases.FavoriteUsecase) *FavoriteController {
// 	return &FavoriteController{
// 		FavoriteUsecase: favoriteUC,
// 	}
// }

// func (controller *FavoriteController) CreateFavorite(c *gin.Context) {
// 	var favoriteInput models.FavoriteInput
// 	if err := c.ShouldBindJSON(&favoriteInput); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	favorite, err := controller.FavoriteUsecase.CreateFavorite(favoriteInput.UserID, favoriteInput.RecipeID)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusCreated, favorite)
// }

// func (controller *FavoriteController) DeleteFavorite(c *gin.Context) {
// 	favoriteID := c.Param("id")
// 	err := controller.FavoriteUsecase.DeleteFavorite(c.GetUint(favoriteID))
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "Favorite deleted successfully"})
// }
