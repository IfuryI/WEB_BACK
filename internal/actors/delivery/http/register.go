package http

import (
	"github.com/gin-gonic/gin"
	"github.com/IfuryI/WEB_BACK/internal/actors"
	"github.com/IfuryI/WEB_BACK/internal/logger"
	"github.com/IfuryI/WEB_BACK/internal/middleware"
)

// RegisterHTTPEndpoints Зарегестрировать хендлеры
func RegisterHTTPEndpoints(router *gin.RouterGroup, actorsUC actors.UseCase, auth middleware.Auth, Log *logger.Logger) {
	handler := NewHandler(actorsUC, Log)

	// router.POST("/actors", auth.CheckAuth(true), handler.CreateActor)
	router.GET("/actors/:actor_id", auth.CheckAuth(false), handler.GetActor)
	// router.PUT("/actors/:actor_id", auth.CheckAuth(true), handler.EditActor)
	router.POST("/actors/:actor_id/like", auth.CheckAuth(true), handler.LikeActor)
	router.DELETE("/actors/:actor_id/like", auth.CheckAuth(true), handler.UnlikeActor)
}
