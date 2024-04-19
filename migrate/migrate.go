package main

import (
	initializers "go-expense-app/initilizers"
	"go-expense-app/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDatabase()
}

func main() {
	initializers.DB.AutoMigrate(&models.Transaction{})
	initializers.DB.AutoMigrate(&models.User{})
}
