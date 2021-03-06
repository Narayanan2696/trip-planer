package initializers

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func InitiateEnv() bool {
	// fmt.Println("ENV: ", os.Environ())
	fmt.Println("app: ", os.Getenv("APP_ENV"))
	// if os.Getenv("APP_ENV") == "local" {
	err := godotenv.Load()
	if err != nil {
		return false
	}
	// }

	return true
}
