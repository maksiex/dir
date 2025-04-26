package configs

import (
	"github.com/joho/godotenv"
	"github.com/maksiex/dir/internal/constants"
	"github.com/maksiex/dir/pkg/logger"
	"os"
	"path/filepath"
)

func LoadInitialConfig() {
	logger.InitLogger()
	envPath, _ := filepath.Abs(".env")
	err := godotenv.Load(envPath)
	if err != nil {
		logger.LoggerErrorCommon(constants.ELoadEnv)
	} else {
		logger.LoggerInfoCommon(constants.Start)
	}

	aviaURL := os.Getenv("AVIA_URL")
	if aviaURL == "" {
		logger.LoggerErrorCommon(constants.EAviaUrl)
	}

	aviaApi := os.Getenv("AVIA_API_KEY")
	if aviaApi == "" {
		logger.LoggerErrorCommon(constants.EAviaApi)
	}

	if aviaURL != "" && aviaApi != "" {
		logger.LoggerInfoFrame(constants.SRunning, constants.ExtraS)
	}
}
