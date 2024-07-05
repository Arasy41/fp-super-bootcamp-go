package controllers

import (
	"api-culinary-review/internal/models"
	"api-culinary-review/internal/usecases"
	_ "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProfileController interface {
	CreateProfile(c *gin.Context)
	GetProfileByUserID(c *gin.Context)
	UpdateProfileByUserID(c *gin.Context)
}

type profileController struct {
	uc usecases.ProfileUsecase
}

func NewProfileController(usecase usecases.ProfileUsecase) ProfileController {
	return &profileController{uc: usecase}
}

func (ctrl *profileController) CreateProfile(c *gin.Context) {
	fullName := c.PostForm("fullName")
	bio := c.PostForm("bio")

	if fullName == "" || bio == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Full name and bio are required"})
		return
	}

	fileHeader, err := c.FormFile("avatar")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get avatar file"})
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

	req := &models.ProfileRequest{
		FullName: fullName,
		Bio:      bio,
	}

	profile, err := ctrl.uc.CreateProfile(req, userIDUint, fileHeader)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, profile)
}

func (ctrl *profileController) GetProfileByUserID(c *gin.Context) {
	userID := c.GetUint("userID")
	profile, err := ctrl.uc.GetProfileByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": profile})
}

func (ctrl *profileController) UpdateProfileByUserID(c *gin.Context) {
	fullName := c.PostForm("fullName")
	bio := c.PostForm("bio")

	if fullName == "" || bio == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Full name and bio are required"})
		return
	}

	fileHeader, err := c.FormFile("avatar")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get avatar file"})
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

	req := &models.ProfileRequest{
		FullName: fullName,
		Bio:      bio,
	}

	err = ctrl.uc.UpdateProfileByID(req, userIDUint, fileHeader)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Profile updated successfully"})
}
