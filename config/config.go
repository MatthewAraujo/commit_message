package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	OpenAi OpenAi
}

type OpenAi struct {
	API_KEY string
	Model   string
}

var Envs = initConfig()

func initConfig() Config {
	godotenv.Load()

	return Config{
		OpenAi: OpenAi{
			API_KEY: getEnv("OPENAI_API_KEY", "MY_SECRET_API_KEY"),
			Model:   getEnv("OPENAI_MODEL", "ChatModelGPT4o"),
		},
	}

}

// Gets the env by key or fallbacks
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
