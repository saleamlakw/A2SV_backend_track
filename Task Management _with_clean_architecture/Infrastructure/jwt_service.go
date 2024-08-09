package Infrastructure

import (
	"os"
	"time"
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/saleamlakw/TaskManager/Domain"
)

func GenerateToken(foundUser Domain.User)(string,error) {
	claims := jwt.MapClaims{
		"userid":   foundUser.ID,
	    "username": foundUser.UserName,
		"role":     foundUser.Role,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	}
	var SecretKey = os.Getenv("SECRET_KEY")
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SecretKey))
	if err != nil {
		return "", fmt.Errorf("error generating token : %v", err)
	}
	return token, nil
}