package config

import "os"

type config struct {
	ADDRESS  string
	GIN_MODE string
}

func New() config {
	return config{
		ADDRESS:  getConfig("ADDRESS"),
		GIN_MODE: getConfig("GIN_MODE"),
	}
}

func getConfig(key string) string {
	if result, exists := os.LookupEnv(key); exists {
		return result
	}

	panic("There is no ENV data")
}
