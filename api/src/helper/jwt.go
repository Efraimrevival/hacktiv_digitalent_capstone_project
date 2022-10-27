package helper

import (
	"api/src/config/env"
	"errors"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

const JWTERROR = "login to process"

func GenerateJWT(id int) string {
	claims := jwt.MapClaims{
		"id":         id,
		"expires_at": time.Now().Add(1 * time.Hour).Unix(),
	}

	secret := env.GetSecretJWTEnv()
	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, _ := rawToken.SignedString([]byte(secret))
	return token
}

func GetAuthorization(c *gin.Context) string {
	return c.Request.Header.Get("Authorization")
}

func ExtractJWT(tokenStr string) (interface{}, error) {
	secret := env.GetSecretJWTEnv()
	token, _ := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return nil, errors.New(JWTERROR)
	}

	var mapClaims jwt.MapClaims
	if v, ok := token.Claims.(jwt.MapClaims); !ok || !token.Valid {
		return nil, errors.New(JWTERROR)
	} else {
		mapClaims = v
	}

	if exp, ok := mapClaims["expires_at"].(float64); !ok {
		return nil, errors.New(JWTERROR)
	} else {
		if int64(exp)-time.Now().Unix() <= 0 {
			return nil, errors.New(JWTERROR)
		}
	}

	if _, ok := mapClaims["id"].(float64); !ok {
		return nil, errors.New(JWTERROR)
	}

	return mapClaims, nil
}

func ValidateJWT(c *gin.Context) (interface{}, error) {
	header := GetAuthorization(c)
	bearer := strings.HasPrefix(strings.ToLower(header), "bearer")
	if !bearer {
		return nil, errors.New(JWTERROR)
	}

	token := strings.Split(header, " ")
	if len(token) != 2 {
		return nil, errors.New(JWTERROR)
	}

	return ExtractJWT(token[1])
}
