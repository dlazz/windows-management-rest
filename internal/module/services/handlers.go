package services

import (
	"net/http"

	"github.com/dlazz/windows-management-rest/internal/module"
	"github.com/gin-gonic/gin"
)

var moduleName = "services"

type Module struct {
}

func NewModule() module.Module {
	return &Module{}
}

func (m *Module) Handle(r *gin.Engine) {
	services := r.Group("/services")

	services.Handle(m.get())
	services.Handle(m.getAllServices())
	services.Handle(m.start())
	services.Handle(m.stop())
	services.Handle(m.restart())
}

func (m *Module) getAllServices() (httpMethod, relativePath string, handler gin.HandlerFunc) {
	httpMethod = http.MethodGet
	relativePath = "/"
	return httpMethod, relativePath, func(ctx *gin.Context) {
		runner := NewRunner()
		srvc, err := runner.GetServices()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"ok": false, "error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"ok": true, "message": srvc})
	}
}

func (m *Module) get() (httpMethod, relativePath string, handler gin.HandlerFunc) {
	httpMethod = http.MethodGet
	relativePath = "/:name"
	return httpMethod, relativePath, func(ctx *gin.Context) {
		name := ctx.Param("name")
		runner := NewRunner()
		srvc, err := runner.GetService(name)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"ok": false, "error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"ok": true, "message": srvc})
	}
}

func (m *Module) stop() (httpMethod, relativePath string, handler gin.HandlerFunc) {
	httpMethod = http.MethodPost
	relativePath = "/:name/stop"
	return httpMethod, relativePath, func(ctx *gin.Context) {
		name := ctx.Param("name")
		runner := NewRunner()
		if err := runner.StopService(name); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"ok": false, "error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"ok": true, "message": "stopped"})
	}
}

func (m *Module) start() (httpMethod, relativePath string, handler gin.HandlerFunc) {
	httpMethod = http.MethodPost
	relativePath = "/:name/start"
	return httpMethod, relativePath, func(ctx *gin.Context) {
		name := ctx.Param("name")
		runner := NewRunner()
		if err := runner.StartService(name); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"ok": false, "error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"ok": true, "message": "started"})
	}
}

func (m *Module) restart() (httpMethod, relativePath string, handler gin.HandlerFunc) {
	httpMethod = http.MethodPost
	relativePath = "/:name/restart"
	return httpMethod, relativePath, func(ctx *gin.Context) {
		name := ctx.Param("name")
		runner := NewRunner()
		if err := runner.RestartService(name); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"ok": false, "error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"ok": true, "message": "restarted"})
	}
}

func init() {
	module.Add(moduleName, NewModule())
}
