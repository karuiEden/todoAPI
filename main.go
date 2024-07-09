package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"todoAPI/controllers"
	"todoAPI/models"
)

func main() {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"POST", "GET", "PUT", "PATCH", "DELETE"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization", "Accept", "User-Agent", "Cache-Control", "Pragma"}
	config.ExposeHeaders = []string{"Content-Length"}
	config.AllowCredentials = true
	router.Use(cors.New(config))
	models.SetupDB()
	router.GET("/tasks", controllers.GetAllTodos)
	router.GET("/tasks/:id", controllers.GetTask)
	//router.DELETE("/tasks/:id", controllers.DeleteTask)
	router.POST("/tasks", controllers.CreateTask)
	router.PATCH("/tasks/:id", controllers.CompleteTask)
	err := router.Run("localhost:8080")
	if err != nil {
		return
	}
}
