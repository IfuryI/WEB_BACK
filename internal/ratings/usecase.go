package ratings

import "github.com/IfuryI/WEB_BACK/internal/models"

// UseCase go:generate mockgen -destination=mocks/usecase.go -package=mocks . UseCase
type UseCase interface {
	CreateRating(userID string, movieID string, score int) error
	GetRating(userID string, movieID string) (models.Rating, error)
	UpdateRating(userID string, movieID string, score int) error
	DeleteRating(userID string, movieID string) error
}
