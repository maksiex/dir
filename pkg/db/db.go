package db

import (
	"github.com/maksiex/dir/internal/configs"
	"github.com/maksiex/dir/internal/constants"
	"github.com/maksiex/dir/pkg/logger"
	"github.com/maksiex/dir/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Db struct {
	Gorm *gorm.DB
}

func DirDb(c *configs.Config) *Db {
	db, err := gorm.Open(postgres.Open(c.Db.DSN), &gorm.Config{})
	if err != nil {
		logger.LoggerErrorCommon(constants.EDbConnection)
		return nil
	}
	err = db.AutoMigrate(
		&models.NewsFromMediastack{},
		&models.NewsFromCurrents{},
		&models.NewsFromGnews{},
		&models.NewsFromNewsapi{},
	)

	if err != nil {
		logger.LoggerErrorCommon(constants.EDbMigration)
		return nil
	}

	return &Db{Gorm: db}
}
