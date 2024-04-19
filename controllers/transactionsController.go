package controllers

import (
	initializers "go-expense-app/initilizers"
	"go-expense-app/models"

	"github.com/gin-gonic/gin"
)

func CreateTransaction(c *gin.Context) {

	var transaction models.Transaction

	c.BindJSON(&transaction)

	result := initializers.DB.Create(&transaction)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(200, transaction)
}

func GetTransactions(c *gin.Context) {

	var transactions []models.Transaction

	result := initializers.DB.Find(&transactions)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(200, transactions)
}

func GetTransaction(c *gin.Context) {

	var transaction models.Transaction

	result := initializers.DB.First(&transaction, c.Param("id"))

	if result.Error != nil {
		c.JSON(400, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(200, transaction)
}

func UpdateTransaction(c *gin.Context) {

	var transaction models.Transaction

	result := initializers.DB.First(&transaction, c.Param("id"))

	if result.Error != nil {
		c.JSON(400, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.BindJSON(&transaction)

	result = initializers.DB.Save(&transaction)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(200, transaction)
}

func DeleteTransaction(c *gin.Context) {

	var transaction models.Transaction

	result := initializers.DB.First(&transaction, c.Param("id"))

	if result.Error != nil {
		c.JSON(400, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	result = initializers.DB.Delete(&transaction)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Transaction deleted successfully",
	})
}
