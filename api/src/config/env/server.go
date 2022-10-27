package env

import (
	"os"

	"github.com/joho/godotenv"
)

func GetServerEnv() string {
	// load env file
	if err := godotenv.Load("app.env"); err != nil {
		panic("failed to load env")
	}

	return os.Getenv("SERVER_PORT")
}
