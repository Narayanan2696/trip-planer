package initializers

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func InitiateEnv() bool {
	fmt.Println("app: ", os.Getenv("APP_ENV"))
	fmt.Println("DB_PROVIDER: ", os.Getenv("DB_PROVIDER"))
	err := godotenv.Load()
	if err != nil {
		path, _ := os.Getwd()
		if _, err := os.Stat(path + "../.env"); err == nil {
			return false
		}
	}
	return true
}
