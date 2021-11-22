package ratings

import (
	"github.com/gin-gonic/gin"
	"github.com/IfuryI/WEB_BACK/internal/logger"
	"github.com/IfuryI/WEB_BACK/internal/middleware"
	"github.com/IfuryI/WEB_BACK/internal/ratings"
)

// RegisterHTTPEndpoints Зарегестрировать хендлеры
func RegisterHTTPEndpoints(router *gin.RouterGroup, ratingsUC ratings.UseCase, authMiddleware middleware.Auth,
	Log *logger.Logger) {
	handler := NewHandler(ratingsUC, Log)

	router.POST("/ratings", authMiddleware.CheckAuth(true), handler.CreateRating)
	router.GET("/ratings/:movie_id", authMiddleware.CheckAuth(true), handler.GetRating)
	router.PUT("/ratings", authMiddleware.CheckAuth(true), handler.UpdateRating)
	router.DELETE("/ratings/:movie_id", authMiddleware.CheckAuth(true), handler.DeleteRating)
}
