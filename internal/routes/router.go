package routes

import (
	"api-culinary-review/internal/controllers"
	"api-culinary-review/internal/middlewares"
	"api-culinary-review/internal/repositories"
	"api-culinary-review/internal/usecases"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = []string{"Content-Type", "X-XSRF-TOKEN", "Accept", "Origin", "X-Requested-With", "Authorization"}

	// To be able to send tokens to the server.
	corsConfig.AllowCredentials = true
	// OPTIONS method for ReactJS
	corsConfig.AddAllowMethods("OPTIONS")

	router.Use(cors.New(corsConfig))

	profileRepo := repositories.NewProfileRepository(db)
	profileUc := usecases.NewProfileUsecase(profileRepo)
	profileCtrl := controllers.NewProfileController(profileUc)
	userRepo := repositories.NewUserRepository(db)
	userUc := usecases.NewUserUsecase(userRepo, profileRepo)
	userCtrl := controllers.NewUserController(userUc)

	tagRepo := repositories.NewTagRepository(db)
	tagUc := usecases.NewtagUsecase(tagRepo)
	tagCtrl := controllers.NewTagController(tagUc)

	recipeRepo := repositories.NewRecipeRepository(db)
	recipeUc := usecases.NewRecipeUsecase(recipeRepo)
	recipeCtrl := controllers.NewRecipeController(recipeUc, tagUc)

	reviewRepo := repositories.NewReviewRepository(db)
	reviewUc := usecases.NewReviewUsecase(reviewRepo)
	reviewCtrl := controllers.NewReviewController(reviewUc)

	favoriteRepo := repositories.NewFavoriteRepository(db)
	favoriteUc := usecases.NewFavoriteUsecase(favoriteRepo)
	favoriteCtrl := controllers.NewFavoriteController(favoriteUc)

	authGroup := router.Group("/api")
	authGroup.Use(middlewares.JWTAuthMiddleware())
	{
		authGroup.GET("/detail-user", userCtrl.GetUserByID)
		authGroup.PUT("/change-password", userCtrl.ChangePassword)

		authGroup.POST("/profile", profileCtrl.CreateProfile)
		authGroup.GET("/profile/me", profileCtrl.GetProfileByUserID)
		authGroup.PUT("/profile", profileCtrl.UpdateProfileByUserID)

		authGroup.POST("/recipes", recipeCtrl.CreateRecipe)
		authGroup.PUT("/recipes/:id", recipeCtrl.UpdateRecipe)
		authGroup.DELETE("/recipes/:id", recipeCtrl.DeleteRecipe)

		authGroup.POST("/reviews", reviewCtrl.CreateReview)
		authGroup.PUT("/reviews/:id", reviewCtrl.UpdateReviewByID)
		authGroup.DELETE("/reviews/:id", reviewCtrl.DeleteReviewByID)

		authGroup.POST("/favorite-recipe", favoriteCtrl.CreateFavorite)
		authGroup.GET("/favorite-recipe", favoriteCtrl.GetByUserID)
		authGroup.DELETE("/favorite-recipe", favoriteCtrl.DeleteFavorite)

		authGroup.POST("tags", tagCtrl.CreateTag)
		authGroup.PUT("/tags/:id", tagCtrl.UpdateTag)
		authGroup.DELETE("/tags/:id", tagCtrl.DeleteTag)
	}

	publicGroup := router.Group("/api")
	{
		publicGroup.GET("/recipes", recipeCtrl.GetRecipes)
		publicGroup.GET("/recipes/:id", recipeCtrl.GetRecipeByID)
		publicGroup.GET("/reviews", reviewCtrl.GetAllReviews)
		publicGroup.GET("/reviews/:id", reviewCtrl.GetReviewByID)
		publicGroup.POST("/register", userCtrl.Register)
		publicGroup.POST("/login", userCtrl.Login)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
