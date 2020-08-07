package genv

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func Must(key string) string {
	value := os.Getenv(key)

	if strings.TrimSpace(value) == "" {
		log.Fatalf("Fatal: Environment variable %s must be set", key)
	}

	return value
}

func MustInt(key string) int {
	value := Must(key)
	trimmed := strings.TrimSpace(value)
	i, err := strconv.Atoi(trimmed)
	if err != nil {
		log.Fatalf("Fatal: Environment variable %s must be a number: %s", key, err)
	}

	return i
}

func MustInts(key string) []int {
	s := Must(key)

	parts := strings.Split(s, ",")

	var ints []int

	for _, part := range parts {
		trimmed := strings.TrimSpace(part)
		i, err := strconv.Atoi(trimmed)
		if err != nil {
			log.Fatalf("Fatal: Environment variable %s must be comma-separated list a numbers: %s", key, err)
		}

		ints = append(ints, i)
	}

	return ints
}
