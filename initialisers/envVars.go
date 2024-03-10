package initialisers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

/*
LoadEnvVariables loads environment variables from the .env file using godotenv.
It checks if the .env file exists and loads its content. If successful, it continues execution.
If the file cannot be loaded, it logs a fatal error and exits the application.
*/
func LoadEnvVariables() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading the .env file.")
		os.Exit(1)
	}

	checkEnvVariablesArePopulated()
}

/*
checkEnvVariablesArePopulated checks if each required environment variable is set before continuing the application execution.
Iterates over a list of required environment variables and checks if each one is set.
If any required variable is not set, it logs a fatal error and exits the application.
*/
func checkEnvVariablesArePopulated() {
	requiredEnvVars := []string{"DB_URL", "ENV", "PORT", "SITE_URL", "VERSION"}

	for _, envVal := range requiredEnvVars {
		if value := os.Getenv(envVal); value == "" {
			log.Fatalf("Environment Variable '%s' is not set.", envVal)
			os.Exit(1)
		}
	}

}
