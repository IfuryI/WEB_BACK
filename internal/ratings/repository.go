package ratings

import "github.com/IfuryI/WEB_BACK/internal/models"

// Repository go:generate mockgen -destination=mocks/repository.go -package=mocks . Repository
type Repository interface {
	CreateRating(username string, movieID string, score int) error
	GetRating(username string, movieID string) (models.Rating, error)
	UpdateRating(username string, movieID string, score int) error
	DeleteRating(username string, movieID string) error
	GetFeed([]models.UserNoPassword) ([]models.RatingFeedItem, error)
}
