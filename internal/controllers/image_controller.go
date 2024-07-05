package controllers

// import (
//     "api-culinary-review/internal/models"
//     "api-culinary-review/internal/usecases"
//     "github.com/gin-gonic/gin"
//     "net/http"
// )

// type ImageController struct {
//     ImageUsecase usecases.ImageUsecase
// }

// func NewImageController(imageUC usecases.ImageUsecase) *ImageController {
//     return &ImageController{
//         ImageUsecase: imageUC,
//     }
// }

// func (controller *ImageController) GetImageByID(c *gin.Context) {
//     imageID := c.Param("id")
//     image, err := controller.ImageUsecase.GetImageByID(imageID)
//     if err != nil {
//         c.JSON(http.StatusNotFound, gin.H{"error": "Image not found"})
//         return
//     }

//     c.JSON(http.StatusOK, image)
// }

// func (controller *ImageController) DeleteImage(c *gin.Context) {
//     imageID := c.Param("id")
//     err := controller.ImageUsecase.DeleteImage(imageID)
//     if err != nil {
//         c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//         return
//     }

//     c.JSON(http.StatusOK, gin.H{"message": "Image deleted successfully"})
// }

// // Implement other image-related controller methods as needed
