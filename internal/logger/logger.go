package logger

import (
	"io"

	"github.com/rs/zerolog"
)

var Logger zerolog.Logger

func new(writers ...io.Writer) zerolog.Logger {

	multi := zerolog.MultiLevelWriter(writers...)
	return zerolog.New(multi).With().Timestamp().Logger()

}

func Init(writers ...io.Writer) {
	Logger = new(writers...)
}
