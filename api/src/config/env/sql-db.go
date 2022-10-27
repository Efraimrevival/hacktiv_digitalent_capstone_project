package env

import (
	"os"

	"github.com/joho/godotenv"
)

type sqlEnv struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func GetSQLEnv() sqlEnv {
	// load env file
	if err := godotenv.Load("app.env"); err != nil {
		panic("failed to load env")
	}

	return sqlEnv{
		Host:     os.Getenv("SQL_HOST"),
		Port:     os.Getenv("SQL_PORT"),
		User:     os.Getenv("SQL_USER"),
		Password: os.Getenv("SQL_PASSWORD"),
		DBName:   os.Getenv("SQL_DB_NAME"),
	}
}
