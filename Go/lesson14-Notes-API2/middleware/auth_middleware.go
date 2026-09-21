package middleware

import(
	"net/http"
	"notes-api/utils"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc{
	return func(c *gin.Context){
		authHeader := c.GetHeader("Authorization")
		if authHeader == ""{
			c.JSON(http.StatusUnauthorized, gin.H{"error": "error"})
			c.Abort()
			return 
		}
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer"{
			c.JSON(http.StatusUnauthorized, gin.H{"error":"invalid auth format"})
			c.Abort()
			return 
		}

		tokenStr := parts[1]

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error){
			return utils.JwtKey,nil
		})
		if err != nil || !token.Valid{
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return 
		}
		claims,ok := token.Claims.(jwt.MapClaims)
		if !ok{
			c.JSON(http.StatusUnauthorized, gin.H{"erorr":"invalid token claims"})
			c.Abort()
			return 
		}
		userId := int(claims["user_id"].(float64))
		c.Set("user_id",userId)
		c.Next()
	}
}