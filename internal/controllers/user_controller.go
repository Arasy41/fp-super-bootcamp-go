package controllers

import (
	"api-culinary-review/internal/models"
	"api-culinary-review/internal/usecases"
	"api-culinary-review/pkg/jwt"
	"api-culinary-review/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	Register(c *gin.Context)
	GetUserByID(c *gin.Context)
	Login(c *gin.Context)
	ChangePassword(c *gin.Context)
}

type userController struct {
	UserUsecase usecases.UserUsecase
}

func NewUserController(userUC usecases.UserUsecase) UserController {
	return &userController{
		UserUsecase: userUC,
	}
}

func (ctrl *userController) Register(c *gin.Context) {
	var userInput models.UserRequest
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
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

func (ctrl *userController) GetUserByID(c *gin.Context) {
	userID := c.GetUint("userID")
	user, err := ctrl.UserUsecase.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}

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
