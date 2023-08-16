package helpers

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateTokenJWT(ID int, username string, role string) (string, error) {
	secretKeyJWTEnv := os.Getenv("SECRET_KEY_JWT")
	if secretKeyJWTEnv == "" {
		log.Fatal("SECRET_KEY_JWT not found")
	}

	secretKey := []byte(secretKeyJWTEnv)

	// Create the JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       ID,
		"iss":      "eFilm APIs",
		"iat":      time.Now().Unix(),
		"nbf":      time.Now().Unix(),
		"exp":      time.Now().Add(time.Hour * 365).Unix(),
		"role":     role,
		"username": username,
	})

	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", errors.New("error creating token")
	}

	return signedToken, nil
}

func ValidateTokenJWT(jwtToken string) (bool, error) {

	secretKeyJWTEnv := os.Getenv("SECRET_KEY_JWT")
	if secretKeyJWTEnv == "" {
		log.Fatal("SECRET_KEY_JWT not found")
	}

	secretKey := []byte(secretKeyJWTEnv)

	// Parse the token
	token, err := jwt.Parse(jwtToken, func(t *jwt.Token) (interface{}, error) {
		if t.Method != jwt.SigningMethodHS256 {
			return nil, errors.New("invalid signing method")
		}
		return secretKey, nil
	})

	if err != nil {
		return false, err
	}

	// Check if the token is valid
	if token.Valid {
		// Access the claims
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			expiration := int64(claims["exp"].(float64)) // expired token
			currentDateTime := time.Now().Unix()
			if currentDateTime > expiration {
				return false, errors.New("token is expired")
			}
		}
		return true, nil
	} else {
		return false, errors.New("token is invalid")
	}
}
