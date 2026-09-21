package handlers

import (
	"net/http"
	"notes-api/models"
	"notes-api/repository"
	"notes-api/utils"
	

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context){
	var user models.User

	if err := c.ShouldBindJSON(&user);
	err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
		return
	}

	err = repository.CreateUser(user.Email, string(hashedPassword))
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "user registered"})
}

func Login(c *gin.Context){
	var user models.User

	if err := c.ShouldBindJSON(&user);
	err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json"})
		return
	}

	id, hashedPassword, err := repository.GetUserByEmail(user.Email)
	if err != nil{
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid email or password"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(user.Password))
	if err != nil{
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid email or password"})
		return
	}

	token,err := utils.GenerateJWT(id, user.Email)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}