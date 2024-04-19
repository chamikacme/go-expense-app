package controllers

import (
	initializers "go-expense-app/initilizers"
	"go-expense-app/models"

	"github.com/gin-gonic/gin"
)

func CreateTransaction(c *gin.Context) {

	var transaction models.Transaction

	user, _ := c.Get("user")

	c.BindJSON(&transaction)

	transaction.UserID = user.(models.User).ID

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

	user, _ := c.Get("user")

	var transactions []models.Transaction

	result := initializers.DB.Where("user_id = ?", user.(models.User).ID).Find(&transactions)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(200, transactions)
}

func GetTransaction(c *gin.Context) {

	user, _ := c.Get("user")

	var transaction models.Transaction

	result := initializers.DB.Where("user_id = ? AND id = ?", user.(models.User).ID, c.Param("id")).First(&transaction)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(200, transaction)
}

func UpdateTransaction(c *gin.Context) {

	user, _ := c.Get("user")

	var transaction models.Transaction

	result := initializers.DB.Where("user_id = ? AND id = ?", user.(models.User).ID, c.Param("id")).First(&transaction)

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

	user, _ := c.Get("user")

	var transaction models.Transaction

	result := initializers.DB.Where("user_id = ? AND id = ?", user.(models.User).ID, c.Param("id")).First(&transaction)

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
