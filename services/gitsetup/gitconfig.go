package gitsetup

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type GoDotEnvLoader interface {
	Load(filenames ...string) error
}

type OSGetter interface {
	Getenv(key string) string
}

var goDotEnvLoader GoDotEnvLoader = godotenvLoader{}
var osGetter OSGetter = osEnvGetter{}

type godotenvLoader struct{}

func (godotenvLoader) Load(filenames ...string) error {
	return godotenv.Load(filenames...)
}

type osEnvGetter struct{}

func (osEnvGetter) Getenv(key string) string {
	return os.Getenv(key)
}

// InitEnv initializes the environment variables.
func InitEnv() {
	loadEnv()
	checkTemplateURL()
}

func loadEnv() {
	if err := goDotEnvLoader.Load(".env"); err != nil {
		log.Printf("Error loading .env file: %v", err)
	}
	if osGetter.Getenv("TEMPLATE_URL") == "" {
		panic("TEMPLATE_URL must be set in the environment")
	}
}

func checkTemplateURL() {
	if osGetter.Getenv("TEMPLATE_URL") == "" {
		panic("TEMPLATE_URL must be set in the environment")
	}
}

// LoadEnv is an accessible function for loading env variables.
func LoadEnv() {
	loadEnv()
}

// CheckTemplateURL is an accessible function for checking the TEMPLATE_URL variable.
func CheckTemplateURL() {
	checkTemplateURL()
}

// DefaultRepoConfig provides a default repository configuration.
func DefaultRepoConfig(repoName string, description string) RepoConfig {
	templateURL := osGetter.Getenv("TEMPLATE_URL")
	if templateURL == "" {
		panic("TEMPLATE_URL must be set in the environment")
	}

	return RepoConfig{
		Name:        repoName,
		Description: description,
		Private:     true,
		AutoInit:    true,
		TemplateURL: templateURL,
	}
}
