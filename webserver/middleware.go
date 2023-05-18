package webserver

import (
	"net/http"
	"strings"

	"github.com/dlazz/windows-management-rest/internal/config"
	"github.com/dlazz/windows-management-rest/internal/logger"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func authMiddleware(ctx *gin.Context) {
	head := ctx.GetHeader("Authorization")
	key := strings.Replace(head, "Bearer ", "", -1)
	if !checkTokenHash(key) {
		logger.Logger.Error().Str("auth", "verify key").Msg("authorization token not valid")
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		ctx.Abort()
		return
	}
	ctx.Next()
}

func loggerMiddleware(ctx *gin.Context) {

	logger.Logger.Info().Str("handler", ctx.HandlerName()).
		Str("client_ip", ctx.ClientIP()).
		Str("method", ctx.Request.Method).
		Str("path", ctx.Request.URL.Path).
		Int("status", ctx.Writer.Status()).
		Msg("")

	ctx.Next()
}

func checkTokenHash(key string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(config.Manager.Token), []byte(key))
	return err == nil
}
