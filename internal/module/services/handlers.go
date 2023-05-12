package services

import (
	"net/http"

	"github.com/dlazz/windows-management-rest/internal/module"
	"github.com/gin-gonic/gin"
)

const moduleName = "services"

type Module struct {
}

func NewModule() module.Module {
	return &Module{}
}

func (m *Module) Handle(r *gin.RouterGroup) {
	services := r.Group("/services")

	services.Handle(m.get())
	services.Handle(m.getAllServices())
	services.Handle(m.start())
	services.Handle(m.stop())
	services.Handle(m.restart())
}

// @BasePath /api/services
// GetServices godoc
// @Summary get service list
// @Schemes
// @Description returns all available windows services
// @Accept json
// @Produce json
// @Success 200 {json} message
// @Router /api/services [get]
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

// @BasePath /api/services/:name
// GetService godoc
// @Summary get
// @Schemes
// @Description returns a windows service from a given name
// @Produce json
// @Param name path string true "service name"
// @Param			Authorization	header		string	true	"Authentication header"
// @Success 200 {json} message
// @Router /api/services/{name} [get]
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

// @BasePath /
// StopService godoc
// @Summary post
// @Schemes
// @Description stop a windows service
// @Produce json
// @Param name path string true "service name"
// @Param			Authorization	header		string	true	"Authentication header"
// @Success 200 {json} message
// @Router /api/services/{name}/stop [post]
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

// @BasePath /
// StartService godoc
// @Summary post
// @Schemes
// @Description start a windows service
// @Produce json
// @Param name path string true "service name"
// @Param			Authorization	header		string	true	"Authentication header"
// @Success 200 {json} message
// @Router /api/services/{name}/start [post]
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

// @BasePath /
// RestartService godoc
// @Summary post
// @Schemes
// @Description restart a windows service
// @Produce json
// @Param name path string true "service name"
// @Param			Authorization	header		string	true	"Authentication header"
// @Success 200 {json} message
// @Router /api/services/{name}/restart [post]
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
