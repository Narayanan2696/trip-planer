package initializers

import (
	"os"

	"github.com/joho/godotenv"
)

func InitiateEnv() bool {
	err := godotenv.Load()
	if err != nil {
		path, _ := os.Getwd()
		if _, err := os.Stat(path + "../.env"); err == nil {
			return false
		}
	}
	return true
}
