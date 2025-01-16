package routes

import (
	"acwj/controllers"
	"acwj/db"
	"acwj/models"

	"github.com/gin-gonic/gin"
)

func StoreRoutes(r *gin.Engine) {
	r.GET("/get_all_stores", controllers.ValidateTokenMiddleware(), getAllStore)
	r.GET("/create_store", controllers.ValidateTokenMiddleware(), createStore)
	r.GET("/update_store", controllers.ValidateTokenMiddleware(), updateStore)
	r.GET("/delete_store", controllers.ValidateTokenMiddleware(), deleteStore)
}

func getAllStore(c *gin.Context) {
	var stores []models.Store
	if err := db.DB.Find(&stores).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": stores})
}

func createStore(c *gin.Context) {
	var store models.Store
	if err := c.ShouldBindJSON(&store); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := db.DB.Create(&store).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Store created successfully"})
}

func updateStore(c *gin.Context) {
	var store models.Store
	if err := c.ShouldBindJSON(&store); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Save(&store).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Store updated successfully"})
}

func deleteStore(c *gin.Context) {
	var store models.Store
	if err := c.ShouldBindJSON(&store); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := db.DB.Delete(&store).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Store deleted successfully"})
}
