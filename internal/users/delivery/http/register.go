package http

import (
	"github.com/gin-gonic/gin"
	"github.com/IfuryI/WEB_BACK/internal/logger"
	"github.com/IfuryI/WEB_BACK/internal/middleware"
	"github.com/IfuryI/WEB_BACK/internal/proto"
	"github.com/IfuryI/WEB_BACK/internal/services/sessions"
	"github.com/IfuryI/WEB_BACK/internal/users"
)

// RegisterHTTPEndpoints Зарегестрировать хендлеры
func RegisterHTTPEndpoints(router *gin.RouterGroup, usersUC users.UseCase, sessions sessions.Delivery,
	authMiddleware middleware.Auth, fileServer proto.FileServerHandlerClient, Log *logger.Logger) {
	handler := NewHandler(usersUC, sessions, fileServer, Log)

	router.POST("/users", handler.CreateUser)
	router.DELETE("/users/admin/:username", authMiddleware.CheckAuth(true), handler.DeleteUser)
	router.POST("/users/avatar", authMiddleware.CheckAuth(true), handler.UploadAvatar)
	router.GET("/users", authMiddleware.CheckAuth(true), handler.GetCurrentUser)
	router.GET("/user/:username", handler.GetUser)
	router.PUT("/users", authMiddleware.CheckAuth(true), handler.UpdateUser)
	router.DELETE("/sessions", authMiddleware.CheckAuth(true), handler.Logout)
	router.POST("/sessions", handler.Login)

	router.GET("/subscriptions/:username", handler.GetSubscriptions)
	router.POST("/subscriptions/:username", authMiddleware.CheckAuth(true), handler.Subscribe)
	router.DELETE("/subscriptions/:username", authMiddleware.CheckAuth(true), handler.Unsubscribe)
	router.GET("/subscriptions/:username/check", authMiddleware.CheckAuth(true), handler.IsSubscribed)
	router.GET("/subscribers/:username", handler.GetSubscribers)
	router.GET("/feed", authMiddleware.CheckAuth(true), handler.GetFeed)
}
