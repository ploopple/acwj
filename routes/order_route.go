package routes

import (
	"acwj/controllers"
	"acwj/db"
	"acwj/models"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(r *gin.Engine) {
	r.GET("/get_all_store_orders", controllers.ValidateTokenMiddleware(), getAllStoreOrders)
	r.GET("/get_all_user_orders", controllers.ValidateTokenMiddleware(), getAllUserOrders)
	r.GET("/create_order", controllers.ValidateTokenMiddleware(), createOrder)
}

func getAllStoreOrders(c *gin.Context) {
	var orders []models.Order

	if err := db.DB.Where("storeId = ?", c.Query("storeId")).Find(&orders).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": orders})
}

func getAllUserOrders(c *gin.Context) {
	var orders []models.Order

	if err := db.DB.Where("uId = ?", c.Query("uId")).Find(&orders).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": orders})
}

func createOrder(c *gin.Context) {
	var order models.Order

	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := db.DB.Create(&order).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "order created successfully"})
}
