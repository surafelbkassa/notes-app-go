package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/surafelbkassa/notes-app-go/models"
	"golang.org/x/crypto/bcrypt"
)

type SignupInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

var userstore = map[string]models.User{}

func main() {
	router := gin.Default()

	// Public routes
	router.POST("/signup", func(c *gin.Context) {
		var input SignupInput
		// c.JSON(http.StatusOK, gin.H{"message": "Signup endpoint"})
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "username and password required"})
			return
		}
		if _, exists := userstore[input.Username]; exists {
			c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
		}
		//Hash password
		hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
			return
		}
		//Store user
		userstore[input.Username] = models.User{
			Username:     input.Username,
			PasswordHash: string(hash),
		}
		c.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "username": input.Username})
	})
	router.POST("/login", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Login endpoint"})
	})

	// Protected routes
	protected := router.Group("/notes")
	protected.Use(AuthMiddleware()) // TODO: implement later
	{
		protected.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"notes": []string{}})
		})

		protected.POST("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "Note created"})
		})
	}

	log.Println("Server running on http://localhost:8080")
	router.Run(":8080")
}
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
