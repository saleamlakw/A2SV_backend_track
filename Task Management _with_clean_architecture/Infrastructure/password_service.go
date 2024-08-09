package Infrastructure

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

func HashPassword(Password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(Password), bcrypt.DefaultCost)
	print(hashedPassword)
	if err != nil {
		log.Panic(err)
	}
	return string(hashedPassword)
}
func VerifyPassword(userPassword string, providedPassword string) (bool, string) {
	msg := ""
	check := true
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	if err != nil {
		msg = "passowrd is incorrect"
		check = false
	}
	return check, msg
}