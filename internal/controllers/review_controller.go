package controllers

import (
	"api-culinary-review/internal/models"
	"api-culinary-review/internal/usecases"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ReviewController interface
type ReviewController interface {
	GetAllReviews(c *gin.Context)
	GetReviewByID(c *gin.Context)
	CreateReview(c *gin.Context)
	UpdateReviewByID(c *gin.Context)
	DeleteReviewByID(c *gin.Context)
}

type reviewController struct {
	uc usecases.ReviewUsecase
}

// NewReviewController creates a new ReviewController instance
func NewReviewController(usecase usecases.ReviewUsecase) ReviewController {
	return &reviewController{
		uc: usecase,
	}
}

// GetAllReviews godoc
// @Summary Get all reviews
// @Description Get a list of all reviews
// @Tags reviews
// @Accept json
// @Produce json
// @Success 200 {array} models.Review
// @Failure 500 {object} map[string]interface{}
// @Router /api/reviews [get]
func (ctrl *reviewController) GetAllReviews(c *gin.Context) {
	reviews, err := ctrl.uc.GetAllReviews()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": reviews})
}

// GetReviewByID godoc
// @Summary Get review by ID
// @Description Get a review by its ID
// @Tags reviews
// @Accept json
// @Produce json
// @Param id path string true "Review ID"
// @Success 200 {object} models.Review
// @Failure 500 {object} map[string]interface{}
// @Router /api/reviews/{id} [get]
func (ctrl *reviewController) GetReviewByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	review, err := ctrl.uc.GetReviewByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": review})
}

// CreateReview godoc
// @Summary Create a new review
// @Description Create a new review for a recipe
// @Tags reviews
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param review body models.ReviewRequest true "Review Request"
// @Success 200 {object} models.Review
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security ApiKeyAuth
// @Router /api/reviews [post]
func (ctrl *reviewController) CreateReview(c *gin.Context) {
	var req models.ReviewRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	userIDUint, ok := userID.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse user ID"})
		return
	}

	req = models.ReviewRequest{
		UserID:   userIDUint,
		RecipeID: req.RecipeID,
		Content:  req.Content,
	}

	review, err := ctrl.uc.CreateReview(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": review})
}

// UpdateReviewByID godoc
// @Summary Update review by ID
// @Description Update an existing review by its ID
// @Tags reviews
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param id path string true "Review ID"
// @Param review body models.ReviewRequest true "Review Request"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security ApiKeyAuth
// @Router /api/reviews/{id} [put]
func (ctrl *reviewController) UpdateReviewByID(c *gin.Context) {
	var req models.ReviewRequest

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid review ID"})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	userIDUint, ok := userID.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse user ID"})
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Ensure the user is updating their own review
	review, err := ctrl.uc.GetReviewByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if review.UserID != userIDUint {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You can only update your own review"})
		return
	}

	req.UserID = userIDUint
	req.RecipeID = review.RecipeID

	if err := ctrl.uc.UpdateReviewByID(&req, uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Review updated successfully"})
}

// DeleteReviewByID godoc
// @Summary Delete review by ID
// @Description Delete a review by its ID
// @Tags reviews
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param id path string true "Review ID"
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security ApiKeyAuth
// @Router /api/reviews/{id} [delete]
func (ctrl *reviewController) DeleteReviewByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid review ID"})
		return
	}

	if err := ctrl.uc.DeleteReviewByID(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Review deleted successfully"})
}
