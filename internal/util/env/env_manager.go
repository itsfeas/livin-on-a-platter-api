package env

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func LoadEnvFile(path string) error {
	home, err := os.Getwd()
	if err != nil {
		return err
	}
	if err := godotenv.Load(home + path); err != nil {
		return fmt.Errorf("error loading .env file")
	}
	return nil
}
