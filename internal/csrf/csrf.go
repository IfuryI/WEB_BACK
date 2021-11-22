package csrf

import (
	"github.com/gin-gonic/gin"
	constants "github.com/IfuryI/WEB_BACK/pkg/const"
	uuid "github.com/satori/go.uuid"
)

// CreateCsrfToken создание CSRF токена
func CreateCsrfToken(ctx *gin.Context) {
	csrfToken := uuid.NewV4().String()

	ctx.Header("X-CSRF-Token", csrfToken)
	ctx.SetCookie("X-CSRF-Cookie",
		csrfToken,
		int(constants.CsrfExpires),
		"/",
		constants.Host,
		false,
		false,
	)
}
