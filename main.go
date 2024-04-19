package main

import (
	"go-expense-app/controllers"
	initializers "go-expense-app/initilizers"
	"go-expense-app/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDatabase()
}

func main() {
	router := gin.Default()

	router.POST("/register", controllers.Register)

	router.POST("/login", controllers.Login)

	router.POST("/logout", controllers.Logout)

	router.GET("/validate", middleware.RequireAuth, controllers.Validate)

	router.POST("/transactions", controllers.CreateTransaction)

	router.GET("/transactions", controllers.GetTransactions)

	router.GET("/transactions/:id", controllers.GetTransaction)

	router.PATCH("/transactions/:id", controllers.UpdateTransaction)

	router.DELETE("/transactions/:id", controllers.DeleteTransaction)

	router.Run()

}
