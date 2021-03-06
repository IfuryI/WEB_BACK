package http

import (
	"github.com/gin-gonic/gin"
	"github.com/IfuryI/WEB_BACK/internal/logger"
	"github.com/IfuryI/WEB_BACK/internal/middleware"
	"github.com/IfuryI/WEB_BACK/internal/reviews"
	"github.com/IfuryI/WEB_BACK/internal/users"
)

// RegisterHTTPEndpoints Зарегестрировать хендлеры
func RegisterHTTPEndpoints(router *gin.RouterGroup, reviewsUC reviews.UseCase, usersUC users.UseCase,
	authMiddleware middleware.Auth, Log *logger.Logger) {
	handler := NewHandler(reviewsUC, usersUC, Log)

	router.POST("/users/reviews", authMiddleware.CheckAuth(true), handler.CreateReview)
	router.GET("/movies/:id/reviews", handler.GetMovieReviews)
	router.GET("/user/:username/reviews", handler.GetUserReviews)
	router.GET("/users/movies/:id/reviews", authMiddleware.CheckAuth(true), handler.GetUserReviewForMovie)
	router.PUT("/users/movies/:id/reviews", authMiddleware.CheckAuth(true), handler.EditUserReviewForMovie)
	router.DELETE("/users/movies/:id/reviews", authMiddleware.CheckAuth(true), handler.DeleteUserReviewForMovie)
	router.DELETE("/users/movies/:id/reviews/:username", authMiddleware.CheckAuth(true), handler.DeleteReview)
}
