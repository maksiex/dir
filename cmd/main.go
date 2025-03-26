package main

import (
	"fmt"
	"github.com/rs/zerolog/log"
)

func main() {
	fmt.Println("Initial")
	log.Info().Msg("✅ Сервис запущен")
}
