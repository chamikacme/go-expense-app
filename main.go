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

	router.POST("/transactions", middleware.RequireAuth, controllers.CreateTransaction)

	router.GET("/transactions", middleware.RequireAuth, controllers.GetTransactions)

	router.GET("/transactions/:id", middleware.RequireAuth, controllers.GetTransaction)

	router.PATCH("/transactions/:id", middleware.RequireAuth, controllers.UpdateTransaction)

	router.DELETE("/transactions/:id", middleware.RequireAuth, controllers.DeleteTransaction)

	router.Run()

}
