package webserver

import (
	"net/http"

	"github.com/dlazz/windows-management-rest/internal/config"
	"github.com/dlazz/windows-management-rest/internal/module"
	_ "github.com/dlazz/windows-management-rest/internal/module/iis"
	_ "github.com/dlazz/windows-management-rest/internal/module/reboot"
	_ "github.com/dlazz/windows-management-rest/internal/module/services"
	"github.com/dlazz/windows-management-rest/webserver/docs"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	// gin-swagger middleware
)

// @title Windows Management Rest
func Run() error {
	if !config.Manager.Webserver.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.New()
	docs.SwaggerInfo.BasePath = "/"
	router.Use(loggerMiddleware)
	router.GET("/version", getVersion)
	router.GET("/configuration", getConfiguration(router))
	router.GET("/healthcheck", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"ok": true}) })
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, swaggerConfig))
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
			for _, existing := range module.AvailableModules {
				if mod == existing {
					log.Debug().Str("module", mod).Msg("loading...")
					module.Store[mod].Handle(router)
					log.Debug().Str("module", mod).Msg("loaded")
				}
			}
		}
	}
}

func swaggerConfig(config *ginSwagger.Config) {
	config.Title = "Windows Management Rest"
}
