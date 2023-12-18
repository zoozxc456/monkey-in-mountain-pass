package extensions

import (
	_ "github.com/joho/godotenv/autoload"
	"os"
)

func GetEnvironment(key string) string {
	return os.Getenv(key)
}
