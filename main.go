package main

import (
	"flag"
	"os"

	"github.com/dlazz/windows-management-rest/internal/config"
	"github.com/dlazz/windows-management-rest/internal/srvc"
	"github.com/dlazz/windows-management-rest/webserver"

	"github.com/kardianos/service"
	"github.com/rs/zerolog/log"
)

var (
	configPath string
	Version    string
)

func main() {
	flag.Parse()
	c, err := os.Open(configPath)
	if err != nil {
		log.Error().Err(err).Msg("unable to open configuration file")
		flag.PrintDefaults()
		os.Exit(1)
	}

	config.InitConfig(c)

	// Set the application version
	webserver.Version.Set(Version)
	log.Info().Str("version", Version).Msg("")

	s, err := service.New(srvc.New(webserver.Run), srvc.Config)
	if err != nil {
		log.Error().Err(err).Msg("")
	}
	logger, err := s.Logger(nil)
	if err != nil {
		log.Error().Err(err).Msg("")
	}
	if err := s.Run(); err != nil {
		logger.Error(err)
	}
}

func init() {
	flag.StringVar(&configPath, "config", "./config.json", "path to a json configuration file")
}
