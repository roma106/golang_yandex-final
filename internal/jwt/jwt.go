package jwt

import (
	"calculator_final/internal/logger"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateJWT(username string) (string, error) {
	const hmacSampleSecret = "secret_signature"
	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": username,
		"nbf":  now.Unix(),
		"exp":  now.Add(time.Minute * 10).Unix(),
		"iat":  now.Unix(),
	})

	tokenString, err := token.SignedString([]byte(hmacSampleSecret))
	if err != nil {
		logger.Error(fmt.Sprintf("failed to sign token. Error: %v", err))
		return "", err
	}

	logger.Info(fmt.Sprintf("Created token for user: %v", username))
	return tokenString, nil
}

// обновляет время истечения jwt токена обратно на 15 минут
func UpdateToken(tokenString string) (string, error) {
	var claims jwt.MapClaims
	_, _, err := new(jwt.Parser).ParseUnverified(tokenString, &claims)
	if err != nil {
		return "", err
	}

	expirationTime := time.Now().Add(10 * time.Minute)
	claims["exp"] = expirationTime.Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString([]byte("secret_signature"))
	if err != nil {
		return "", err
	}
	logger.Info(fmt.Sprintf("Updated token for user: %v", claims["name"]))
	return tokenString, nil
}

// проверяет валидность jwt токена (время истечения)
func IsTokenExpired(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("secret_signature"), nil
	})
	if err != nil {
		logger.Error(fmt.Sprintf("failed to parse token. Error: %v", err))
		return false, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		expirationTime := time.Unix(int64(claims["exp"].(float64)), 0)
		if expirationTime.Before(time.Now()) {
			logger.Error(fmt.Sprintf("token expired. Error: %v", err))
			return true, nil
		}
		return false, nil
	}
	logger.Info(fmt.Sprintf("token for user %s is valid", token.Claims.(jwt.MapClaims)["name"]))
	return false, nil
}
