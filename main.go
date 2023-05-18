package main

import (
	"flag"
	"io"
	"os"
	"time"

	"github.com/dlazz/windows-management-rest/internal/config"
	"github.com/dlazz/windows-management-rest/internal/logger"
	"github.com/dlazz/windows-management-rest/internal/srvc"
	"github.com/dlazz/windows-management-rest/webserver"

	_ "embed"

	"github.com/kardianos/service"
	"github.com/rs/zerolog/log"
)

var (
	configPath string
	Version    string
	logToFile  bool
)

func main() {
	flag.Parse()

	writers := []io.Writer{os.Stdout}

	if logToFile {
		date := time.Now().Format("20060102150405")
		logFileName := "wmr-" + date + ".log"
		logFile, err := os.Create(logFileName)
		if err != nil {
			log.Error().Err(err).Msg("")
		}
		defer logFile.Close()
		writers = append(writers, logFile)
	}
	logger.Init(writers...)

	c, err := os.Open(configPath)
	if err != nil {
		log.Error().Err(err).Msg("unable to open configuration file")
		flag.PrintDefaults()
		os.Exit(1)
	}

	config.InitConfig(c)

	// Set the application version
	webserver.Version.Set(Version)
	logger.Logger.Info().Str("version", Version).Msg("")

	s, err := service.New(srvc.New(webserver.Run), srvc.Config)
	if err != nil {
		logger.Logger.Error().Err(err).Msg("")
	}

	if err != nil {
		logger.Logger.Error().Err(err).Msg("")
	}
	if err := s.Run(); err != nil {
		logger.Logger.Error().Err(err).Msg("")
	}
}

func init() {
	flag.StringVar(&configPath, "config", "./config.json", "path to a json configuration file")
	flag.BoolVar(&logToFile, "log-to-file", false, "log to file")
}
