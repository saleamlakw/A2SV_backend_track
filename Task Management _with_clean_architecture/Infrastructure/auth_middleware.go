package Infrastructure

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)



func AuthMiddleware(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "Authorization Header is required"})
		c.Abort()
		return
	}
	authParts := strings.Split(authHeader, " ")
	if len(authParts) != 2 || strings.ToLower(authParts[0]) != "bearer" {
		fmt.Println(len(authParts) != 2, strings.ToLower(authParts[0]))
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization Header"})
		c.Abort()
		return
	}
	token, err := jwt.Parse(authParts[1], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing mmethod %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil || !token.Valid {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "Invalid jwt"})
		c.Abort()
		return
	}
	if claims,ok:=token.Claims.(jwt.MapClaims);ok{
		// fmt.Println("--------",float64(time.Now().Unix()),claims["exp"].(float64))
		if float64(time.Now().Unix())>claims["exp"].(float64){
			c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "token expired"})
			c.Abort()
			return
		}
		c.Set("role",claims["role"].(string))
		c.Set("userid",claims["userid"].(string))
		c.Next()
	}else{
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}

func IsAdminMiddleware(c *gin.Context){
	role, exists := c.Get("role")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "role not found"})
		c.Abort()
		return
	}

	urole, ok := role.(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid role"})
		c.Abort()
		return
	}


	if urole != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access forbidden"})
		c.Abort()
		return
	}

	c.Next()
}



func IsAuthorizedToViewUserDetails(c *gin.Context){
	userid, exists := c.Get("userid")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "userid not found"})
			c.Abort()
			return
		}

		uid, ok := userid.(string)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid userid"})
			c.Abort()
			return
		}

		userUserID := c.Param("user_id")
		if uid != userUserID {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access"})
			c.Abort()
			return
		}
		c.Next()
	}