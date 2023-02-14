package module

import "github.com/gin-gonic/gin"

var Store = make(map[string]Module)

type Module interface {
	Handle(r *gin.RouterGroup)
}

func Add(name string, module Module) {
	Store[name] = module
}
