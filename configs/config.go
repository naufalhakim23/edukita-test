package configs

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type (
	Config struct {
		Application Application
		Postgresql  Postgresql
	}
	Application struct {
		Name        string
		Env         string
		Port        int
		Secret      string
		StaticToken string
		SwaggerPath string
	}
	Postgresql struct {
		Name string
		URL  string
	}
)

func LoadConfigurations(fileName string) (*Config, error) {
	err := godotenv.Load(fileName)
	if err != nil {
		logrus.WithError(err).Warn("file not found")
	}
	app := Application{
		Name:        GetEnv("APP_NAME", "edukita-teaching-grading"),
		Env:         GetEnv("APP_ENV", "local"),
		Port:        getEnvAsInt("APP_PORT", 8080),
		Secret:      GetEnv("APP_SECRET", "supersecretsecret"),
		StaticToken: GetEnv("APP_STATIC_TOKEN", "supersecretsecret"),
		SwaggerPath: GetEnv("APP_SWAGGER_PATH", ""),
	}
	psql := Postgresql{
		Name: GetEnv("POSTGRES_NAME", "edukita-teaching-grading"),
		URL:  GetEnv("POSTGRES_URL", "localhost:5432"),
	}
	cfg := Config{
		Application: app,
		Postgresql:  psql,
	}
	return &cfg, nil
}

func GetEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func getEnvAsInt(name string, defaultVal int) int {
	valStr := GetEnv(name, "")
	if value, err := strconv.Atoi(valStr); err == nil {
		return value
	}

	return defaultVal
}
