package env

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type EnvVal interface {
	int64 | string
}

func init() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalf("Could not load .env file, %v", err)
	}
}

func GetEnvVar[E EnvVal](key string, fallback E) E {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	switch any(fallback).(type) {
	case int:
		intVal, err := strconv.Atoi(val)
		if err != nil {
			return fallback
		}
		return any(intVal).(E)
	}

	return any(val).(E)
}
