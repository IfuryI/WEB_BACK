package users

import (
	"github.com/IfuryI/WEB_BACK/internal/models"
)

// UseCase go:generate mockgen -destination=mocks/usecase.go -package=mocks . UseCase
type UseCase interface {
	CreateUser(user *models.User) error

	Login(login, password string) bool

	GetUser(username string) (*models.User, error)

	UpdateUser(user *models.User, change models.User) (*models.User, error)

	Subscribe(subscriber string, user string) error

	Unsubscribe(subscriber string, user string) error

	GetSubscribers(page int, user string) (int, []models.UserNoPassword, error)

	IsSubscribed(subscriber string, user string) (bool, error)

	GetSubscriptions(page int, user string) (int, []models.UserNoPassword, error)

	GetFeed(username string) (models.Feed, error)

	DeleteUser(admin string, username string) error
}
