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
	router.GET("/configuration", getConfiuration)
	router.GET("/healthcheck", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"ok": true}) })
	router.GET("/modules", getModules)
	router.GET("/handlers", getHandlers(router))
	router.Use(authMiddleware)
	loadModule(router)

	if err := router.Run(":" + config.Manager.Webserver.Port); err != nil {
		return err
	}
	return nil

}

func loadModule(router *gin.Engine) {
	if config.Manager.Modules != nil {
		for _, mod := range config.Manager.Modules {
			if _, ok := module.Store[mod]; ok {
				log.Debug().Str("module", mod).Msg("loading module")
				module.Store[mod].Handle(router)
			}
		}
	}
}
