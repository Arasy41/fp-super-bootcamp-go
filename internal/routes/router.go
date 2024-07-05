package routes

import (
	"api-culinary-review/internal/controllers"
	"api-culinary-review/internal/middlewares"
	"api-culinary-review/internal/repositories"
	"api-culinary-review/internal/usecases"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	userRepo := repositories.NewUserRepository(db)
	userUc := usecases.NewUserUsecase(userRepo)
	userCtrl := controllers.NewUserController(userUc)

	profileRepo := repositories.NewProfileRepository(db)
	profileUc := usecases.NewProfileUsecase(profileRepo)
	profileCtrl := controllers.NewProfileController(profileUc)

	recipeRepo := repositories.NewRecipeRepository(db)
	recipeUc := usecases.NewRecipeUsecase(recipeRepo)
	recipeCtrl := controllers.NewRecipeController(recipeUc)

	reviewRepo := repositories.NewReviewRepository(db)
	reviewUc := usecases.NewReviewUsecase(reviewRepo)
	reviewCtrl := controllers.NewReviewController(reviewUc)

	authGroup := router.Group("/api")
	authGroup.Use(middlewares.JWTAuthMiddleware())
	{
		authGroup.GET("/detail-user", userCtrl.GetUserByID)
		authGroup.PUT("/change-password", userCtrl.ChangePassword)

		authGroup.POST("/profile", profileCtrl.CreateProfile)
		authGroup.GET("/profile", profileCtrl.GetProfileByUserID)
		authGroup.PUT("/profile", profileCtrl.UpdateProfileByUserID)

		authGroup.GET("/recipes", recipeCtrl.GetAllRecipes)
		authGroup.GET("/recipes/:id", recipeCtrl.GetRecipeByID)
		authGroup.POST("/recipes", recipeCtrl.CreateRecipe)
		authGroup.PUT("/recipes/:id", recipeCtrl.UpdateRecipeByID)
		authGroup.DELETE("/recipes/:id", recipeCtrl.DeleteRecipeByID)

		authGroup.GET("/reviews", reviewCtrl.GetAllReviews)
		authGroup.GET("/reviews/:id", reviewCtrl.GetReviewByID)
		authGroup.POST("/reviews", reviewCtrl.CreateReview)
		authGroup.PUT("/reviews/:id", reviewCtrl.UpdateReviewByID)
		authGroup.DELETE("/reviews/:id", reviewCtrl.DeleteReviewByID)
	}

	publicGroup := router.Group("/api")
	{
		publicGroup.POST("/register", userCtrl.Register)
		publicGroup.POST("/login", userCtrl.Login)
	}

	return router
}
