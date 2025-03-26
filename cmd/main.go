package main

import (
	"fmt"
	"github.com/rs/zerolog/log"
)

func main() {
	fmt.Println("test")
	log.Info().Msg("✅ Сервис запущен")
}
