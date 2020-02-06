package initializers

import (
	"github.com/joho/godotenv"
)

func InitiateEnv() bool {
	err := godotenv.Load()
	if err != nil {
		return false
	}
	return true
}
