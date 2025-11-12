package env

import (
	"log"
	"os"
	"strconv"
)

func GetStringEnv(key string) string {
	value, err := getEnv(key)
	if err != nil {
		log.Printf("Environment variable %s is not set", key)
		os.Exit(1)
	}
	return value
}

func GetNumberEnv(key string) int {
	value, err := getEnv(key)
	if err != nil {
		log.Printf("Environment variable %s is not set", key)
		os.Exit(1)
	}
	intVal, err := strconv.Atoi(value)
	if err != nil {
		log.Printf("Numeric conversion for %s failed", key)
		os.Exit(1)
	}
	return intVal
}