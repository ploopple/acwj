package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type User struct { 
    Username string `json:"username"` 
    Password string `json:"password"` 
}

func main() {
  r := gin.Default()

  psqlInfo := "host=localhost port=5432 user=exampleuser password=examplepass dbname=exampledb sslmode=disable"
  db, err := sql.Open("postgres", psqlInfo) 
  if err != nil { 
    log.Fatal(err) 
    } 
    defer db.Close() 
    err = db.Ping() 
    if err != nil { 
        log.Fatal(err) 
        } 
        fmt.Println("Successfully connected to the database!")

  r.GET("/ping", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "pong",
    })
  })

   r.POST("/ping", func(c *gin.Context) {
    var user User 
    if err := c.ShouldBindJSON(&user); 
    err != nil { 
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) 
        return 
    }
    c.JSON(http.StatusOK, gin.H{
      "message": "pong",
    })
  })
  

  r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
