package util

import (
	"fmt"

	"github.com/joho/godotenv"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func LoadDotEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(fmt.Errorf("unable to load .env: %s", err))
	}
}
