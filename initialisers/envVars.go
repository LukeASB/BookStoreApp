package initialisers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading the .env file.")
		os.Exit(1)
	}

	checkEnvVariablesArePopulated()
}

// Check if each required env var is set before continuing app execution.
func checkEnvVariablesArePopulated() {
	requiredEnvVars := []string{"DB_URL", "ENV", "PORT", "SITE_URL", "VERSION"}

	for _, envVal := range requiredEnvVars {
		if value := os.Getenv(envVal); value == "" {
			log.Fatalf("Environment Variable '%s' is not set.", envVal)
			os.Exit(1)
		}
	}

}
