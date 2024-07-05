// internal/controllers/tag_controller.go

package controllers

import (
	"api-culinary-review/internal/models"
	"api-culinary-review/internal/usecases"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TagController struct {
	TagUsecase usecases.TagUsecase
}

func NewTagController(tagUC usecases.TagUsecase) *TagController {
	return &TagController{
		TagUsecase: tagUC,
	}
}

func (controller *TagController) CreateTag(c *gin.Context) {
	var tagInput models.Tag
	if err := c.ShouldBindJSON(&tagInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tag, err := controller.TagUsecase.CreateTag(tagInput.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, tag)
}

func (controller *TagController) GetTagByName(c *gin.Context) {
	tagName := c.Param("name")
	tag, err := controller.TagUsecase.GetTagByName(tagName)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tag not found"})
		return
	}

	c.JSON(http.StatusOK, tag)
}

func (controller *TagController) UpdateTag(c *gin.Context) {
	var tag models.Tag
	if err := c.ShouldBindJSON(&tag); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := controller.TagUsecase.UpdateTag(&tag)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tag updated successfully"})
}

func (controller *TagController) DeleteTag(c *gin.Context) {
	tagID := c.Param("id")
	err := controller.TagUsecase.DeleteTag(c.GetUint(tagID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tag deleted successfully"})
}

// Implement other tag-related controller methods as needed
