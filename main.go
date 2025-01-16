package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID    int    `gorm:"primary_key"`
	Name  string `json:"name"`
	Email string `gorm:"unique" json:"email"`
}

func (u *User) TableName() string {
	return "users"
}

var jwtKey = []byte("your_secret_key")

type Claims struct {
	Id int `json:"id"`
	jwt.StandardClaims
}

func main() {
	dsn := "host=autorack.proxy.rlwy.net user=postgres password=moGcokUQZOQzFYefiAcYnWxnhGsgLxmm dbname=railway port=41861 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	fmt.Println("Connected to the database successfully!")
	r := gin.Default()

	//db.AutoMigrate(&User{})

	r.GET("/signup", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if err := db.Create(&user).Error; err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		// Generate JWT token
		expirationTime := time.Now().Add(15 * time.Minute)
		claims := &Claims{
			Id: user.ID,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expirationTime.Unix(),
			},
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(jwtKey)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		// c.JSON(201, user)
		c.JSON(http.StatusOK, gin.H{"token": tokenString})
	})

	r.GET("/validate", validateTokenMiddleware(db), func(c *gin.Context) {
		user, _ := c.Get("user")
		c.JSON(http.StatusOK, user)
	})

	// Start the server
	r.Run(":8080") // Listen and serve on localhost:8080
}

func validateTokenMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := strings.Split(c.GetHeader("Authorization"), " ")[1]
		// print(strings.Split(tokenString, " ")[1])
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
			c.Abort()
			return
		}
		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
				c.Abort()
				return
			}
			c.JSON(http.StatusBadRequest, gin.H{"error": "could not parse token"})
			c.Abort()
			return
		}
		if !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}
		var user User
		if err := db.Where("id = ?", claims.Id).First(&user).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
			c.Abort()
			return
		}
		c.Set("user", user)
		c.Next()
	}
}
