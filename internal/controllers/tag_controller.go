package controllers

import (
	"api-culinary-review/internal/models"
	"api-culinary-review/internal/usecases"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// TagController is the interface that defines the methods for handling tag-related operations.
type TagController interface {
	GetAllTags(c *gin.Context)
	CreateTag(c *gin.Context)
	UpdateTag(c *gin.Context)
	DeleteTag(c *gin.Context)
}

type tagController struct {
	tagUsecase usecases.TagUsecase
}

// NewTagController creates a new instance of TagController.
func NewTagController(tagUsecase usecases.TagUsecase) TagController {
	return &tagController{
		tagUsecase: tagUsecase,
	}
}

// CreateTag creates a new tag.
// @Summary Create a new tag
// @Description Creates a new tag with the provided name.
// @Tags tags
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param input body models.TagRequest true "Tag data to create"
// @Success 201 {object} models.TagResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Security ApiKeyAuth
// @Router /api/tags [post]
func (ctrl *tagController) CreateTag(c *gin.Context) {
	var tagInput models.TagRequest
	if err := c.ShouldBindJSON(&tagInput); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	tag, err := ctrl.tagUsecase.CreateTag(tagInput.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, tag)
}

// GetAllTags godoc
// @Summary Get all tags
// @Description Get a list of all tags
// @Tags tags
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Success 200 {array} models.Tag
// @Failure 500 {object} ErrorResponse
// @Security ApiKeyAuth
// @Router /api/tags [get]
func (ctrl *tagController) GetAllTags(c *gin.Context) {
	tags, err := ctrl.tagUsecase.GetAllTags()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tags)
}

// UpdateTag updates an existing tag.
// @Summary Update an existing tag
// @Description Updates an existing tag based on the provided data.
// @Tags tags
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param id path int true "Tag ID to update"
// @Param input body models.TagRequest true "Updated tag data"
// @Success 200 {string} string "Tag updated successfully"
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Security ApiKeyAuth
// @Router /api/tags/{id} [put]
func (ctrl *tagController) UpdateTag(c *gin.Context) {
	tagID := c.Param("id")
	var tagInput models.TagRequest
	if err := c.ShouldBindJSON(&tagInput); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	// Convert tag ID to uint
	id, err := strconv.ParseUint(tagID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid tag ID"})
		return
	}

	tag := models.Tag{
		ID:   uint(id),
		Name: tagInput.Name,
	}

	err = ctrl.tagUsecase.UpdateTag(&tag)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, "Tag updated successfully")
}

// DeleteTag deletes a tag by its ID.
// @Summary Delete a tag by ID
// @Description Deletes a tag by its ID.
// @Tags tags
// @Param Authorization header string true "Bearer Token"
// @Param id path int true "Tag ID to delete"
// @Success 200 {string} string "Tag deleted successfully"
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Security ApiKeyAuth
// @Router /api/tags/{id} [delete]
func (ctrl *tagController) DeleteTag(c *gin.Context) {
	tagID := c.Param("id")
	id, err := strconv.ParseUint(tagID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid tag ID"})
		return
	}

	err = ctrl.tagUsecase.DeleteTag(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, "Tag deleted successfully")
}

// ErrorResponse represents a JSON structure for error responses.
type ErrorResponse struct {
	Error string `json:"error"`
}
