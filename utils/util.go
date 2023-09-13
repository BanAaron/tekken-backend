package util

import (
	"fmt"

	"github.com/joho/godotenv"
)

func LoadDotEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(fmt.Errorf("unable to load .env: %s", err))
	}
}
