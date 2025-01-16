package main

import (
	"acwj/db"
	"acwj/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Connect()
	//db.Migrate()

	r := gin.Default()

	routes.UserRoutes(r)
	routes.StoreRoutes(r)
	routes.OrderRoutes(r)

	r.Run(":8080")
}
