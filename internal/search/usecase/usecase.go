package usecase

import (
	"strings"

	ac "github.com/IfuryI/WEB_BACK/internal/actors"
	"github.com/IfuryI/WEB_BACK/internal/models"
	mv "github.com/IfuryI/WEB_BACK/internal/movies"
	us "github.com/IfuryI/WEB_BACK/internal/users"
)

// SearchUseCase структура usecase поиска
type SearchUseCase struct {
	userRepository  us.UserRepository
	movieRepository mv.MovieRepository
	actorRepository ac.Repository
}

// NewSearchUseCase инициализация usecase поиска
func NewSearchUseCase(usRepo us.UserRepository, mvRepo mv.MovieRepository, acRepo ac.Repository) *SearchUseCase {
	return &SearchUseCase{
		userRepository:  usRepo,
		movieRepository: mvRepo,
		actorRepository: acRepo,
	}
}

// Search поиск
func (uc *SearchUseCase) Search(query string) (models.SearchResult, error) {
	query = strings.ToLower(query)
	movies, err := uc.movieRepository.SearchMovies(query)
	if err != nil {
		return models.SearchResult{}, err
	}
	actors, err := uc.actorRepository.SearchActors(query)
	if err != nil {
		return models.SearchResult{}, err
	}
	users, err := uc.userRepository.SearchUsers(query)
	if err != nil {
		return models.SearchResult{}, err
	}

	result := models.SearchResult{
		Movies: movies,
		Actors: actors,
		Users:  users,
	}

	return result, nil
}
