package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"notes-api/config"
	"notes-api/middleware"
)

func Register(c *gin.Context) {

	var input struct {
		Email    string
		Password string
	}

	c.BindJSON(&input)

	_, err := db.DB.Exec(
		"INSERT INTO users (email, password) VALUES ($1, $2)",
		input.Email, input.Password,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user exists"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "registered"})
}

func Login(c *gin.Context) {

	var input struct {
		Email    string
		Password string
	}

	c.BindJSON(&input)

	var id int
	var password string

	err := db.DB.QueryRow(
		"SELECT id, password FROM users WHERE email=$1",
		input.Email,
	).Scan(&id, &password)

	if err != nil || password != input.Password {
		c.JSON(401, gin.H{"error": "invalid credentials"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": id,
	})

	tokenStr, _ := token.SignedString(middleware.JwtKey)

	c.JSON(200, gin.H{"token": tokenStr})
}