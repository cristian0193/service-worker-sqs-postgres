package utils

import (
	"fmt"
	"os"
	"strconv"
)

func GetString(name string) (string, error) {
	v, ok := os.LookupEnv(name)
	if !ok {
		return "", fmt.Errorf("env var %s not found", name)
	}
	return v, nil
}

func GetInt(name string) (int, error) {
	v, ok := os.LookupEnv(name)
	if !ok {
		return 0, fmt.Errorf("env var %s not found", name)
	}
	intV, err := strconv.Atoi(v)
	if err != nil {
		return 0, fmt.Errorf("en var %s must be a number", name)
	}
	return intV, nil
}
