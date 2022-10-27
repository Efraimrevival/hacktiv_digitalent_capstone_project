package env

import (
	"os"

	"github.com/joho/godotenv"
)

func GetSecretJWTEnv() string {
	// load env file
	if err := godotenv.Load("app.env"); err != nil {
		panic("failed to load env")
	}

	return os.Getenv("SECRET_JWT")
}
