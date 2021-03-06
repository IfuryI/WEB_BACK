package http

import (
	"github.com/gin-gonic/gin"
	"github.com/IfuryI/WEB_BACK/internal/logger"
	"github.com/IfuryI/WEB_BACK/internal/middleware"
	"github.com/IfuryI/WEB_BACK/internal/movies"
)

// RegisterHTTPEndpoints Зарегестрировать хендлеры
func RegisterHTTPEndpoints(router *gin.RouterGroup, moviesUC movies.UseCase, auth middleware.Auth, Log *logger.Logger) {
	handler := NewHandler(moviesUC, Log)

	router.POST("/movies", handler.CreateMovie)
	router.GET("/movies", auth.CheckAuth(false), handler.GetMovies)
	router.GET("/movies/:id", auth.CheckAuth(false), handler.GetMovie)
	router.POST("/movies/:id/watch", auth.CheckAuth(true), handler.MarkWatched)
	router.DELETE("/movies/:id/watch", auth.CheckAuth(true), handler.MarkUnwatched)
	router.GET("/movies/:id/similar", handler.GetSimilar)
	router.GET("/genres", handler.GetGenres)
}
