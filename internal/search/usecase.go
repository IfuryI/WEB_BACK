package search

import "github.com/IfuryI/WEB_BACK/internal/models"

// UseCase интерфейс
type UseCase interface {
	Search(query string) (models.SearchResult, error)
}
