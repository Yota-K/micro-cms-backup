package config

import "os"

func Init() (string, string) {
	API_KEY := os.Getenv("API_KEY")
	ENDPOINT := os.Getenv("ENDPOINT")
	return API_KEY, ENDPOINT
}
