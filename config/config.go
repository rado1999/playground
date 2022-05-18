package config

import "os"

func New() config {
	return config{
		ADDRESS:  getEnv("ADDRESS"),
		GIN_MODE: getEnv("GIN_MODE"),
		HOST:     getEnv("HOST"),
		PORT:     getEnv("PORT"),
		USER:     getEnv("USER"),
		PASSWORD: getEnv("PASSWORD"),
		DBNAME:   getEnv("DBNAME"),
		SSLMODE:  getEnv("SSLMODE"),
	}
}

type config struct {
	ADDRESS  string
	GIN_MODE string
	HOST     string
	PORT     string
	USER     string
	PASSWORD string
	DBNAME   string
	SSLMODE  string
}

func getEnv(key string) string {
	if result, exists := os.LookupEnv(key); exists {
		return result
	}

	panic("There is no ENV data")
}
