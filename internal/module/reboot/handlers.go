package reboot

import (
	"net/http"

	"github.com/dlazz/windows-management-rest/internal/module"
	"github.com/gin-gonic/gin"
)

const moduleName = "reboot"

type Module struct {
}

func NewModule() module.Module {
	return &Module{}
}

func (m *Module) Handle(r *gin.RouterGroup) {
	services := r.Group("/reboot")

	services.Handle(m.reboot())

}

// @BasePath /api/reboot/
// GetService godoc
// @Summary get
// @Schemes
// @Description reboot the computer
// @Produce json
// @Param			Authorization	header		string	true	"Authentication header"
// @Success 200 {json} message
// @Router /api/reboot [get]
func (m *Module) reboot() (httpMethod, relativePath string, handler gin.HandlerFunc) {
	httpMethod = http.MethodGet
	relativePath = "/"
	return httpMethod, relativePath, func(ctx *gin.Context) {
		runner := NewRunner()
		srvc, err := runner.reboot()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"ok": false, "error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"ok": true, "message": srvc})
	}
}

func init() {
	module.Add(moduleName, NewModule())
}
