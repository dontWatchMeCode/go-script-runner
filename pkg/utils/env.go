package utils

import (
	"errors"
	"os"
	"strconv"
)

func GetEnv(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return "", errors.New(key + " is not set")
	}
	return value, nil
}

func GetEnvAtoi(key string) (int, error) {
	value := os.Getenv(key)
	if value == "" {
		return 0, errors.New(key + " is not set")
	}

	valueInt, err := strconv.Atoi(value)
	if err != nil {
		return 0, errors.New(key + " conversion to int failed")
	}

	return valueInt, nil
}
