package webserver

import (
	"net/http"

	"github.com/dlazz/windows-management-rest/internal/config"
	"github.com/dlazz/windows-management-rest/internal/module"
	"github.com/gin-gonic/gin"
)

var Version = version{}

type version struct {
	version string
}

func (v *version) Set(version string) {
	v.version = version
}

func (v *version) Get() string {
	return v.version
}

func getVersion(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"version": Version.Get()})
}

func getModules() map[string][]string {
	m := []string{}
	for k := range module.Store {
		m = append(m, k)
	}
	return map[string][]string{
		"available": m,
		"loaded":    config.Manager.Modules,
	}

}

func getHandlers(router *gin.Engine) []map[string]string {

	handlerList := []map[string]string{}
	handlers := router.Routes()
	for _, item := range handlers {
		handlerList = append(handlerList, map[string]string{
			"handler": item.Path,
			"method":  item.Method,
		})
	}
	return handlerList

}

func getConfiguration(router *gin.Engine) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		type OverallConfig struct {
			Version       string                `json:"version"`
			Configuration *config.Configuration `json:"configuration"`
			Modules       map[string][]string   `json:"modules"`
			Handlers      []map[string]string   `json:"handlers"`
		}
		o := OverallConfig{
			Configuration: config.Manager,
			Modules:       getModules(),
			Version:       Version.Get(),
			Handlers:      getHandlers(router),
		}
		ctx.JSON(http.StatusOK, o)
	}
}
