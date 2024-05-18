package config

type Config struct {
	Port        string
	Environment string
}

func LoadConfig() Config {
	return Config{
		Port:        "8080",
		Environment: "development",
	}
}