package movies

import "github.com/IfuryI/WEB_BACK/internal/models"

// MovieRepository go:generate mockgen -destination=mocks/repository.go -package=mocks . MovieRepository
type MovieRepository interface {
	CreateMovie(movie *models.Movie) error

	GetMovieByID(id string, username string) (*models.Movie, error)

	GetBestMovies(startIndex int, username string) (int, []*models.Movie, error)

	GetAllGenres() ([]string, error)

	GetMoviesByGenres(genres []string, startIndex int, username string) (int, []*models.Movie, error)

	MarkWatched(username string, id int) error

	MarkUnwatched(username string, id int) error

	SearchMovies(query string) ([]models.Movie, error)

	GetSimilar(id string) ([]models.Movie, error)
}
