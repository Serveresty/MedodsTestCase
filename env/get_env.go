package env

import "os"

// Метод получения параметров из .env
func GetEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic(value)
	}

	return value
}
