package utils

import (
	"log"
	"os"
	"strconv"
)

func GetEnvStr(key string) string {
	v := os.Getenv(key)
	if v == "" {
		log.Fatal("Environment variable %s doesn't exist", key)
	}
	return v
}

func GetEnvInt(key string) int {
	s := GetEnvStr(key)
	v, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return v
}

func GetEnvBool(key string) bool {
	s := GetEnvStr(key)
	v, err := strconv.ParseBool(s)
	if err != nil {
		log.Fatal(err)
	}
	return v
}
