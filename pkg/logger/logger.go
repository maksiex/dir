package logger

import (
	"github.com/fatih/color"
	"github.com/maksiex/dir/internal/constants"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func InitLogger() {
	log.Logger = zerolog.New(zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: "15:04:05",
	}).With().Timestamp().Logger()
}

func LoggerInfoCommon(info string) {
	s := color.CyanString(info)
	log.Info().Msg(s)
}

func LoggerInfoRegular(info string, returnedError error) {
	i := &Logger{
		Info:          info,
		ReturnedError: returnedError,
	}
	log.Info().Msg(i.Info)
	log.Error().Err(i.ReturnedError).Msg("")
}

func LoggerInfoFrame(info, extended string) {
	log.Info().Msg(constants.SuccessStyles.Render(info))
	if extended != "" {
		s := color.YellowString(extended)
		log.Info().Msg(s)
	}
}

func LoggerErrorCommon(error string) {
	s := color.RedString(error)
	log.Info().Msg(s)
}
