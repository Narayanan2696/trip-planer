package initializers

import (
	"os"

	"github.com/joho/godotenv"
)

func InitiateEnv() bool {
	path, _ := os.Getwd()
	path = path + "../.env"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := godotenv.Load()
		if err != nil {
			return false
		}
	}

	return true
}
