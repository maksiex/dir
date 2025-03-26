package main

import (
	"fmt"
	"github.com/rs/zerolog/log"
)

func main() {
	fmt.Printf("Initial")
	log.Info().Msg("✅ Сервис запущен")
}
