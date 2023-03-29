package main

import (
	"github.com/DeepjyotiSarmah/go_jwt_authentication_gin_gonic/controllers"
	"github.com/DeepjyotiSarmah/go_jwt_authentication_gin_gonic/initializers"
	"github.com/DeepjyotiSarmah/go_jwt_authentication_gin_gonic/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)

	r.Run() // listen and serve on 0.0.0.0:8080
}
