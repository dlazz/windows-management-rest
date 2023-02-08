package webserver

import (
	"net/http"
	"strings"

	"github.com/dlazz/windows-management-rest/internal/config"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func authMiddleware(ctx *gin.Context) {
	head := ctx.GetHeader("Authorization")
	key := strings.Replace(head, "Bearer ", "", -1)
	if !CheckTokenHash(key) {
		log.Error().Str("auth", "verify key").Msg("authorization token not valid")
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		ctx.Abort()
		return
	}
	ctx.Next()
}

func loggerMiddleware(ctx *gin.Context) {

	log.Info().Str("handler", ctx.HandlerName()).
		Str("client_ip", ctx.ClientIP()).
		Str("method", ctx.Request.Method).
		Msg("")

	ctx.Next()
}

func CheckTokenHash(key string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(config.Manager.Token), []byte(key))
	return err == nil
}
