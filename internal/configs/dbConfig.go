package configs

import (
	"fmt"
	"github.com/maksiex/dir/internal/constants"
	"github.com/maksiex/dir/pkg/logger"
	"os"
)

type Config struct {
	Db DBConfig
}

type DBConfig struct {
	DSN string
}

func LoadDBConfig() *Config {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSL"),
	)
	if dsn == "" {
		logger.LoggerErrorCommon(constants.EDsn)
	}
	return &Config{
		DBConfig{
			DSN: dsn,
		},
	}
}
