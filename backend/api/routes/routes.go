package routes

import (
	"github.com/fajardwntara/vow-connect/api/auth"
	"github.com/fajardwntara/vow-connect/api/domain/user"
	"github.com/fajardwntara/vow-connect/api/handlers"
	"github.com/gin-gonic/gin"
)

func UserRouteRegistry(router *gin.Engine, repo user.UserRepository) {
	userHandler := handlers.NewUserHandler(repo)

	generalGroup := router.Group("/api")
	{
		generalGroup.POST("/login", auth.Login)
	}

	userGroup := router.Group("/api/users")
	{
		userGroup.GET("/all", userHandler.GetAllUsers)
		userGroup.GET("/:id", userHandler.GetOne)
		userGroup.POST("/add", userHandler.Create)
		userGroup.DELETE("/delete/:id", userHandler.Delete)
	}

}
