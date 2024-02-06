package env

import "os"

func GetEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic(value)
	}

	return value
}
