package webserver

import (
	"net/http"

	"github.com/dlazz/windows-management-rest/internal/config"
	"github.com/dlazz/windows-management-rest/internal/module"
	_ "github.com/dlazz/windows-management-rest/internal/module/iis"
	_ "github.com/dlazz/windows-management-rest/internal/module/services"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	// gin-swagger middleware
)

func Run() error {
	if !config.Manager.Webserver.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.New()
	router.Use(loggerMiddleware)
	router.GET("/version", getVersion)
	router.GET("/configuration", getConfiguration(router))
	router.GET("/healthcheck", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"ok": true}) })
	api := router.Group("/api")
	api.Use(authMiddleware)
	loadModule(api)

	if err := router.Run(":" + config.Manager.Webserver.Port); err != nil {
		return err
	}
	return nil

}

func loadModule(router *gin.RouterGroup) {
	if config.Manager.Modules != nil {
		for _, mod := range config.Manager.Modules {
			log.Debug().Str("module", mod).Msg("loading...")
			module.Store[mod].Handle(router)
			log.Debug().Str("module", mod).Msg("loaded")
		}
	}
}
