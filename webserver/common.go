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

func getModules(ctx *gin.Context) {
	m := []string{}
	for k := range module.Store {
		m = append(m, k)
	}
	ctx.JSON(http.StatusOK, gin.H{"available": m, "loaded": config.Manager.Modules})
}

func getConfiuration(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, config.Manager)
}

func getHandlers(router *gin.Engine) gin.HandlerFunc {
	type H struct {
		Handler string `json:"handler"`
		Method  string `json:"method"`
	}

	return func(ctx *gin.Context) {
		handlerList := []*H{}
		handlers := router.Routes()
		for _, item := range handlers {
			handlerList = append(handlerList, &H{
				Handler: item.Path,
				Method:  item.Method,
			})
		}
		ctx.JSON(http.StatusOK, handlerList)
	}
}
