package routes

import (
	"acwj/controllers"
	"acwj/db"
	"acwj/models"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var jwtKey = []byte("your_secret_key")

type Claims struct {
	Id int `json:"id"`
	jwt.StandardClaims
}

func UserRoutes(r *gin.Engine) {
	r.GET("/sign_user", signUser)
	r.GET("/get_user", controllers.ValidateTokenMiddleware(), getUser)
}

func signUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Where("email = ?", user.Email).First(&user).Error; err != nil {
		// Check if the error is due to the user not being found
		if err == gorm.ErrRecordNotFound {
			// User does not exist, create a new one
			if err := db.DB.Create(&user).Error; err != nil {
				c.JSON(500, gin.H{"error": err.Error()})
				return
			}
			// c.JSON(200, gin.H{"message": "User created successfully"})
		} else {
			// An error occurred while querying the database
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
	} else {
		// User already exists
		// c.JSON(400, gin.H{"error": "User already exists"})
	}

	// if err := db.DB.Create(&user).Error; err != nil {
	// 	c.JSON(500, gin.H{"error": err.Error()})
	// 	return
	// }

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

	c.JSON(http.StatusOK, gin.H{"token": tokenString, "data": user})
}

func getUser(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{"data": user})
}
