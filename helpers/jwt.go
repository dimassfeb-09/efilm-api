package helpers

import (
	"errors"
	"github.com/dimassfeb-09/efilm-api.git/entity/web"
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

func ValidateTokenJWT(jwtToken string) (bool, *web.UserInfoResponse, error) {

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
		return false, nil, err
	}

	// Check if the token is valid
	if token.Valid {

		var userInfo web.UserInfoResponse

		// Access the claims
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			expiration := int64(claims["exp"].(float64)) // expired token
			currentDateTime := time.Now().Unix()
			if currentDateTime > expiration {
				return false, nil, errors.New("token is expired")
			}

			userInfo.UserID = int(claims["id"].(float64))
			userInfo.Username = claims["username"].(string)

		}
		return true, &userInfo, nil
	} else {
		return false, nil, errors.New("token is invalid")
	}
}
