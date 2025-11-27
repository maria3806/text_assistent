package config

import "os"

func GetAPIKey() string {
	return os.Getenv("OPENAI_API_KEY")
}
