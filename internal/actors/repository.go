package actors

import "github.com/IfuryI/WEB_BACK/internal/models"

// Repository go:generate mockgen -destination=mocks/repository.go -package=mocks . Repository
type Repository interface {
	// CreateActor(models.Actor) error
	GetActorByID(id string, username string) (models.Actor, error)
	GetFavoriteActors(username string) ([]models.Actor, error)
	// EditActor(models.Actor) (models.Actor, error)
	LikeActor(username string, actorID int) error
	UnlikeActor(username string, actorID int) error
	SearchActors(query string) ([]models.Actor, error)
}
