package configs

import (
	"github.com/joho/godotenv"
	"github.com/maksiex/dir/internal/constants"
	"github.com/maksiex/dir/pkg/logger"
	"path/filepath"
)

func LoadInitialConfig() {
	logger.InitLogger()

	envPath, _ := filepath.Abs(".env.local")
	err := godotenv.Load(envPath)
	if err != nil {
		logger.LoggerErrorCommon(constants.ELoadEnv)
	} else {
		logger.LoggerInfoCommon(constants.Start)
	}
}
