package config

import (
	"os"
)

// Config ...
type Config struct {
	Environment string

	CtxTimeout string
	GinMode    string
	HttpPort   string

	PostgresHost     string
	PostgresPort     string
	PostgresUser     string
	PostgresDatabase string
	PostgresPassword string

	RedisHost     string
	RedisPort     string
	RedisDatabase string
	RedisPassword string

	SigningKey          string
	AccessTokenTimeout  string
	RefreshTokenTimeout string

	AuthConfigPath string
	CSVFilePath    string
}

// Load loads environment vars and inflates Config
func Load() Config {
	c := Config{}

	c.Environment = getEnv("ENVIRONMENT", "production")

	c.GinMode = getEnv("GIN_MODE", "release")
	c.HttpPort = getEnv("HTTP_PORT", ":8888")
	c.CtxTimeout = getEnv("CTX_TIMEOUT", "5s")

	c.PostgresHost = getEnv("POSTGRES_HOST", "localhost")
	// c.PostgresHost = getEnv("POSTGRES_HOST", "postgres")
	c.PostgresPort = getEnv("POSTGRES_PORT", "5432")
	c.PostgresDatabase = getEnv("POSTGRES_DATABASE", "tender")
	c.PostgresUser = getEnv("POSTGRES_USER", "postgres")
	c.PostgresPassword = getEnv("POSTGRES_PASSWORD", "ebot")

	c.RedisHost = getEnv("REDIS_HOST", "localhost")
	// c.RedisHost = getEnv("REDIS_HOST", "redisdb")
	c.RedisPort = getEnv("REDIS_PORT", ":6379")
	c.RedisDatabase = getEnv("REDIS_DATABASE", "0")
	c.RedisPassword = getEnv("REDIS_PASSWORD", "")

	c.SigningKey = getEnv("SIGNING_KEY", "template-key")
	c.AccessTokenTimeout = getEnv("ACCESS_TOKEN_TIMEOUT", "10800")   // 3h
	c.RefreshTokenTimeout = getEnv("REFRESH_TOKEN_TIMEOUT", "86400") // 24h

	c.CSVFilePath = getEnv("CSV_FILE_PATH", "./config/policy.csv")
	c.AuthConfigPath = getEnv("AUTH_PATH", "./config/model.conf")

	return c
}

func getEnv(key string, defaultVaule string) string {
	value, exists := os.LookupEnv(key)
	if exists {
		return value
	}
	return defaultVaule
}
