package controllers

import (
	"api-culinary-review/internal/models"
	"api-culinary-review/internal/usecases"
	"api-culinary-review/pkg/jwt"
	"api-culinary-review/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserController interface
type UserController interface {
	Register(c *gin.Context)
	GetUserByID(c *gin.Context)
	Login(c *gin.Context)
	ChangePassword(c *gin.Context)
}

type userController struct {
	UserUsecase usecases.UserUsecase
}

// NewUserController creates a new UserController instance
func NewUserController(userUC usecases.UserUsecase) UserController {
	return &userController{
		UserUsecase: userUC,
	}
}

// Register godoc
// @Summary Register a new user
// @Description Create a new user account
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.UserRequest true "User Data"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/register [post]
func (ctrl *userController) Register(c *gin.Context) {
	var userInput models.UserRequest
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	existingUser, err := ctrl.UserUsecase.GetUserByEmailOrUsername(userInput.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if existingUser != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Email already registered"})
	}

	user, err := ctrl.UserUsecase.CreateUser(userInput.Username, userInput.Password, userInput.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":  "User registered successfully",
		"username": user.Username,
	})
}

// GetUserByID godoc
// @Summary Get user by ID
// @Description Get user information by user ID
// @Tags users
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security ApiKeyAuth
// @Router /api/detail-user [get]
func (ctrl *userController) GetUserByID(c *gin.Context) {
	userID := c.GetUint("userID")
	user, err := ctrl.UserUsecase.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}

// Login godoc
// @Summary User login
// @Description Authenticate user and get a JWT token
// @Tags users
// @Accept json
// @Produce json
// @Param login body models.LoginInput true "Login Data"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/login [post]
func (ctrl *userController) Login(c *gin.Context) {
	var input models.LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := ctrl.UserUsecase.GetUserByEmailOrUsername(input.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !utils.CheckPasswordHash(input.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	token, err := jwt.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

// ChangePassword godoc
// @Summary Change user password
// @Description Change the password for the authenticated user
// @Tags users
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param password body models.InputChangePassword true "Change Password Data"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security ApiKeyAuth
// @Router /api/change-password [put]
func (ctrl *userController) ChangePassword(c *gin.Context) {
	var input models.InputChangePassword
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.GetUint("userID")
	user, err := ctrl.UserUsecase.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !utils.CheckPasswordHash(input.OldPassword, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Old password is incorrect"})
		return
	}

	user.Password, _ = utils.HashPassword(input.NewPassword)
	if err := ctrl.UserUsecase.UpdateUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password changed successfully"})
}
