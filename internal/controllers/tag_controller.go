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
// @Param input body models.TagRequest true "Tag data to create"
// @Success 201 {object} models.TagResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /tags [post]
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

// UpdateTag updates an existing tag.
// @Summary Update an existing tag
// @Description Updates an existing tag based on the provided data.
// @Tags tags
// @Accept json
// @Produce json
// @Param id path int true "Tag ID to update"
// @Param input body models.TagRequest true "Updated tag data"
// @Success 200 {string} string "Tag updated successfully"
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /tags/{id} [put]
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
// @Param id path int true "Tag ID to delete"
// @Success 200 {string} string "Tag deleted successfully"
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /tags/{id} [delete]
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
