package controllers

import (
	"api-culinary-review/internal/models"
	"api-culinary-review/internal/usecases"
	_ "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ProfileController interface
type ProfileController interface {
	CreateProfile(c *gin.Context)
	GetProfileByUserID(c *gin.Context)
	UpdateProfileByUserID(c *gin.Context)
}

type profileController struct {
	uc usecases.ProfileUsecase
}

// NewProfileController creates a new ProfileController instance
func NewProfileController(usecase usecases.ProfileUsecase) ProfileController {
	return &profileController{uc: usecase}
}

// CreateProfile godoc
// @Summary Create a new profile
// @Description Create a new user profile
// @Tags profiles
// @Accept multipart/form-data
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param fullName formData string true "Full Name"
// @Param bio formData string true "Bio"
// @Param avatar formData file true "Avatar File"
// @Success 201 {object} models.Profile
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security ApiKeyAuth
// @Router /api/profile [post]
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

// GetProfileByUserID godoc
// @Summary Get profile by user ID
// @Description Get the profile of the authenticated user
// @Tags profiles
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Success 200 {object} models.Profile
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security ApiKeyAuth
// @Router /api/profile/me [get]
func (ctrl *profileController) GetProfileByUserID(c *gin.Context) {
	userID := c.GetUint("userID")
	profile, err := ctrl.uc.GetProfileByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": profile})
}

// UpdateProfileByUserID godoc
// @Summary Update profile by user ID
// @Description Update the profile of the authenticated user
// @Tags profiles
// @Accept multipart/form-data
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param fullName formData string true "Full Name"
// @Param bio formData string true "Bio"
// @Param avatar formData file true "Avatar File"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security ApiKeyAuth
// @Router /api/profile [put]
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
