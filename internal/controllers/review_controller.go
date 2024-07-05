package controllers

import (
	"api-culinary-review/internal/models"
	"api-culinary-review/internal/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

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

func NewReviewController(usecase usecases.ReviewUsecase) ReviewController {
	return &reviewController{
		uc: usecase,
	}
}

func (ctrl *reviewController) GetAllReviews(c *gin.Context) {
	reviews, err := ctrl.uc.GetAllReviews()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": reviews})
}

func (ctrl *reviewController) GetReviewByID(c *gin.Context) {
	id := c.Param("id")
	review, err := ctrl.uc.GetReviewByID(c.GetUint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": review})
}

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

func (ctrl *reviewController) UpdateReviewByID(c *gin.Context) {
	var req models.ReviewRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")

	if err := ctrl.uc.UpdateReviewByID(&req, c.GetUint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Review updated successfully"})
}

func (ctrl *reviewController) DeleteReviewByID(c *gin.Context) {
	id := c.Param("id")

	if err := ctrl.uc.DeleteReviewByID(c.GetUint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Review deleted successfully"})
}
