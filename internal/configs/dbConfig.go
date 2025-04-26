package configs

import (
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
	dsn := os.Getenv("DSN")
	if dsn == "" {
		logger.LoggerErrorCommon(constants.EDsn)
	}
	return &Config{
		DBConfig{
			DSN: dsn,
		},
	}
}
