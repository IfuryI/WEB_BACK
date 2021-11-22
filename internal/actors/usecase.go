package actors

import "github.com/IfuryI/WEB_BACK/internal/models"

// UseCase go:generate mockgen -destination=mocks/usecase.go -package=mocks . UseCase
type UseCase interface {
	// CreateActor(user models.User, actor models.Actor) error
	GetActor(id string, username string) (models.Actor, error)
	// EditActor(user models.User, change models.Actor) (models.Actor, error)
	LikeActor(username string, actorID int) error
	UnlikeActor(username string, actorID int) error
}
