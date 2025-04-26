package main

import (
	"github.com/maksiex/dir/internal/configs"
	"github.com/maksiex/dir/pkg/db"
)

func main() {
	configs.LoadInitialConfig()
	dbConfig := configs.LoadDBConfig()
	db.DirDb(dbConfig)
}
