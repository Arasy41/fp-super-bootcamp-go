package controllers

import (
	"api-culinary-review/internal/models"
	"api-culinary-review/internal/usecases"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// FavoriteController is the interface that defines the methods for handling favorite-related operations.
type FavoriteController interface {
	GetByUserID(c *gin.Context)
	CreateFavorite(c *gin.Context)
	DeleteFavorite(c *gin.Context)
}

type favoriteController struct {
	favoriteUsecase usecases.FavoriteUsecase
}

// NewFavoriteController creates a new instance of FavoriteController.
func NewFavoriteController(favoriteUC usecases.FavoriteUsecase) FavoriteController {
	return &favoriteController{
		favoriteUsecase: favoriteUC,
	}
}

// GetByUserID retrieves favorites by user ID.
// @Summary Retrieve favorites by user ID
// @Description Retrieves favorites associated with the specified user ID.
// @Tags favorites
// @Accept json
// @Produce json
// @Param user_id path int true "User ID"
// @Success 200 {array} models.Favorite
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/favorites/{user_id} [get]
func (ctrl *favoriteController) GetByUserID(c *gin.Context) {
	userIDStr := c.Param("user_id")
	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid user ID"})
		return
	}

	favorites, err := ctrl.favoriteUsecase.GetByUserID(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, favorites)
}

// CreateFavorite creates a new favorite.
// @Summary Create a new favorite
// @Description Creates a new favorite for a user with the specified recipe ID.
// @Tags favorites
// @Accept json
// @Produce json
// @Param input body models.FavoriteRequest true "Favorite data to create"
// @Success 201 {object} models.Favorite
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/favorites [post]
func (ctrl *favoriteController) CreateFavorite(c *gin.Context) {
	var favoriteInput models.FavoriteRequest
	if err := c.ShouldBindJSON(&favoriteInput); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	favorite, err := ctrl.favoriteUsecase.CreateFavorite(favoriteInput.UserID, favoriteInput.RecipeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, favorite)
}

// DeleteFavorite deletes a favorite by its ID.
// @Summary Delete a favorite by ID
// @Description Deletes a favorite by its ID.
// @Tags favorites
// @Param id path int true "Favorite ID"
// @Success 200 {string} string "Favorite deleted successfully"
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/favorites/{id} [delete]
func (ctrl *favoriteController) DeleteFavorite(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid ID"})
		return
	}

	err = ctrl.favoriteUsecase.DeleteFavorite(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, "Favorite deleted successfully")
}
