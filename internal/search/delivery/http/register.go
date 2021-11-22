package http

import (
	"github.com/gin-gonic/gin"
	"github.com/IfuryI/WEB_BACK/internal/logger"
	"github.com/IfuryI/WEB_BACK/internal/search"
)

// RegisterHTTPEndpoints Зарегестрировать хендлеры
func RegisterHTTPEndpoints(router *gin.RouterGroup, searchUC search.UseCase, Log *logger.Logger) {
	handler := NewHandler(searchUC, Log)
	router.GET("/search", handler.Search)
}
