package reviews

import "github.com/IfuryI/WEB_BACK/internal/models"

// UseCase go:generate mockgen -destination=mocks/usecase.go -package=mocks . UseCase
type UseCase interface {
	CreateReview(user *models.User, review *models.Review) error

	GetReviewsByUser(username string) ([]*models.Review, error)

	GetReviewsByMovie(movieID string, page int) (int, []*models.Review, error)

	GetUserReviewForMovie(username string, movieID string) (*models.Review, error)

	EditUserReviewForMovie(user *models.User, review *models.Review) error

	DeleteUserReviewForMovie(user *models.User, movieID string) error

	DeleteReview(admin string, username string, movieID int) error
}
