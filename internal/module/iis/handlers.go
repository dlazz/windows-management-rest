package iis

import (
	"net/http"

	"github.com/dlazz/windows-management-rest/internal/module"
	"github.com/gin-gonic/gin"
)

const moduleName = "iis"

type Module struct {

	//handlersList []func() (httpMethod, relativePath string, handler gin.HandlerFunc)
}

func NewModule() module.Module {
	return &Module{}
}

func (m *Module) Handle(r *gin.Engine) {

	iis := r.Group("/iis")
	iis.Handle(m.getWebsite())
	iis.Handle(m.getWebsiteList())
	iis.Handle(m.stopWebsite())
	iis.Handle(m.startWebsite())
	iis.Handle(m.getAppPoolList())
	iis.Handle(m.startWebAppPool())
	iis.Handle(m.stopWebAppPool())
}

func (m *Module) getWebsiteList() (httpMethod, relativePath string, handler gin.HandlerFunc) {
	httpMethod = http.MethodGet
	relativePath = "/websites"

	return httpMethod, relativePath, func(ctx *gin.Context) {
		runner := NewRunner()
		sites, err := runner.GetWebsiteList()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"ok": false, "error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"ok": true, "message": sites})
	}
}

func (m *Module) getWebsite() (httpMethod, relativePath string, handler gin.HandlerFunc) {
	httpMethod = http.MethodGet
	relativePath = "/websites/:name"

	return httpMethod, relativePath, func(ctx *gin.Context) {
		name := ctx.Param("name")
		runner := NewRunner()
		site, err := runner.GetWebsite(name)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"ok": false, "error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"ok": true, "message": site})
	}
}

func (m *Module) stopWebsite() (httpMethod, relativePath string, handler gin.HandlerFunc) {
	httpMethod = http.MethodPost
	relativePath = "/websites/:name/stop"

	return httpMethod, relativePath, func(ctx *gin.Context) {
		name := ctx.Param("name")
		runner := NewRunner()
		sites, err := runner.StopWebsite(name)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"ok": false, "error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"ok": true, "message": sites})
	}
}

func (m *Module) startWebsite() (httpMethod, relativePath string, handler gin.HandlerFunc) {
	httpMethod = http.MethodPost
	relativePath = "/websites/:name/start"

	return httpMethod, relativePath, func(ctx *gin.Context) {
		name := ctx.Param("name")
		runner := NewRunner()
		sites, err := runner.StartWebsite(name)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"ok": false, "error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"ok": true, "message": sites})
	}
}

func (m *Module) getAppPoolList() (httpMethod, relativePath string, handler gin.HandlerFunc) {
	httpMethod = http.MethodGet
	relativePath = "/apppools"

	return httpMethod, relativePath, func(ctx *gin.Context) {
		runner := NewRunner()
		appPools, err := runner.GetAppPoolList()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"ok": false, "error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"ok": true, "message": appPools})
	}
}

func (m *Module) startWebAppPool() (httpMethod, relativePath string, handler gin.HandlerFunc) {
	httpMethod = http.MethodPost
	relativePath = "/apppools/:name/start"

	return httpMethod, relativePath, func(ctx *gin.Context) {
		name := ctx.Param("name")
		runner := NewRunner()
		msg, err := runner.StartWebAppPool(name)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"ok": false, "error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"ok": true, "message": msg})
	}
}

func (m *Module) stopWebAppPool() (httpMethod, relativePath string, handler gin.HandlerFunc) {
	httpMethod = http.MethodPost
	relativePath = "/apppools/:name/stop"

	return httpMethod, relativePath, func(ctx *gin.Context) {
		name := ctx.Param("name")
		runner := NewRunner()
		msg, err := runner.StopWebAppPool(name)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"ok": false, "error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"ok": true, "message": msg})
	}
}

func init() {
	module.Add(moduleName, NewModule())
}
