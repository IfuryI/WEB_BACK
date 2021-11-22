package movies

import "github.com/IfuryI/WEB_BACK/internal/models"

// UseCase go:generate mockgen -destination=mocks/usecase.go -package=mocks . UseCase
type UseCase interface {
	CreateMovie(movie *models.Movie) error

	GetMovie(id string, username string) (*models.Movie, error)

	GetBestMovies(page int, username string) (int, []*models.Movie, error)

	GetAllGenres() ([]string, error)

	GetMoviesByGenres(genres []string, page int, username string) (int, []*models.Movie, error)

	MarkWatched(user models.User, id int) error

	MarkUnwatched(user models.User, id int) error

	GetSimilar(id string) ([]models.Movie, error)
}
