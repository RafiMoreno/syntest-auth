package main

import (
	"auth-service/controllers"
	"auth-service/initializers"
	"auth-service/middleware"
	"log"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.SetupDatabase()

	if initializers.DB == nil {
        log.Fatal("Database connection is not initialized")
    }

	initializers.MigrateDB(initializers.DB)
}

func main() {
    router := gin.Default() 

    routerGroup := router.Group("/api/")
    {
        routerGroup.POST("/sign-up", controllers.SignUp)
		routerGroup.POST("/login", controllers.Login)
		routerGroup.GET("/validate",middleware.RequireAuth, controllers.Validate)
    }

    router.Run(":8082") 
}
