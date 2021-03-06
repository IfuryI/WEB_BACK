package usecase

import (
	"errors"

	"github.com/IfuryI/WEB_BACK/internal/models"
	"github.com/IfuryI/WEB_BACK/internal/movies"
	"github.com/IfuryI/WEB_BACK/internal/users"
	constants "github.com/IfuryI/WEB_BACK/pkg/const"
)

// MoviesUseCase структура usecase фильма
type MoviesUseCase struct {
	movieRepository movies.MovieRepository
	userRepository  users.UserRepository
}

// NewMoviesUseCase инициализация usecase фильма
func NewMoviesUseCase(repo movies.MovieRepository, userRepo users.UserRepository) *MoviesUseCase {
	return &MoviesUseCase{
		movieRepository: repo,
		userRepository:  userRepo,
	}
}

// CreateMovie создать фильм
func (moviesUC *MoviesUseCase) CreateMovie(movie *models.Movie) error {
	_, err := moviesUC.movieRepository.GetMovieByID(movie.ID, "")
	if err == nil {
		return errors.New("movie already exists")
	}
	return moviesUC.movieRepository.CreateMovie(movie)
}

// GetMovie получение информации о фильме
func (moviesUC *MoviesUseCase) GetMovie(id string, username string) (*models.Movie, error) {
	return moviesUC.movieRepository.GetMovieByID(id, username)
}

// GetBestMovies получить лучшие фильмы
func (moviesUC *MoviesUseCase) GetBestMovies(page int, username string) (int, []*models.Movie, error) {
	startIndex := (page - 1) * constants.MoviesPageSize
	return moviesUC.movieRepository.GetBestMovies(startIndex, username)
}

// GetAllGenres получить доступные жанры
func (moviesUC *MoviesUseCase) GetAllGenres() ([]string, error) {
	return moviesUC.movieRepository.GetAllGenres()
}

// GetMoviesByGenres получить фильмы по жанрам
func (moviesUC *MoviesUseCase) GetMoviesByGenres(genres []string, page int, username string) (int, []*models.Movie, error) {
	startIndex := (page - 1) * constants.MoviesPageSize
	return moviesUC.movieRepository.GetMoviesByGenres(genres, startIndex, username)
}

// MarkWatched отметить просмотренным
func (moviesUC *MoviesUseCase) MarkWatched(user models.User, id int) error {
	err := moviesUC.movieRepository.MarkWatched(user.Username, id)
	if err != nil {
		return err
	}
	// successful mark watched, must increment movies_watched for user
	newMoviesWatchNumber := *user.MoviesWatched + 1
	_, err = moviesUC.userRepository.UpdateUser(&user, models.User{
		Username:      user.Username,
		MoviesWatched: &newMoviesWatchNumber,
	})
	return err
}

// MarkUnwatched отметить непросмотренным
func (moviesUC *MoviesUseCase) MarkUnwatched(user models.User, id int) error {
	err := moviesUC.movieRepository.MarkUnwatched(user.Username, id)
	if err != nil {
		return err
	}
	// successful mark unwatched, must decrement movies_watched for user
	newMoviesWatchNumber := *user.MoviesWatched - 1
	_, err = moviesUC.userRepository.UpdateUser(&user, models.User{
		Username:      user.Username,
		MoviesWatched: &newMoviesWatchNumber,
	})
	return err
}

// GetSimilar получить похожие
func (moviesUC *MoviesUseCase) GetSimilar(id string) ([]models.Movie, error) {
	similarMovies, err := moviesUC.movieRepository.GetSimilar(id)
	if err != nil {
		return nil, nil
	}
	return similarMovies, nil
}
