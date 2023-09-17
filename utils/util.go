package util

import (
	"fmt"

	"github.com/joho/godotenv"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func LoadDotEnv() error {
	err := godotenv.Load(".env")
	if err != nil {
		return fmt.Errorf("unable to load .env: %s", err)
	}
	return nil
}

func ToTitleCase(str *string) {
	titleCase := cases.Title(language.English)
	*str = titleCase.String(*str)
}
