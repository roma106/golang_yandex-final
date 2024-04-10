package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

func AddJWT(username string) {
	const hmacSampleSecret = "secret_signature"
	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": username,
		"nbf":  now.Add(time.Minute).Unix(),
		"exp":  now.Add(time.Hour * 2).Unix(),
		"iat":  now.Unix(),
	})

	tokenString, err := token.SignedString([]byte(hmacSampleSecret))
	if err != nil {
		panic(err)
	}

	fmt.Println(tokenString)
}
