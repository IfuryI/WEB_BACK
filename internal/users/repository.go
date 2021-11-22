package users

import (
	"github.com/IfuryI/WEB_BACK/internal/models"
)

// UserRepository go:generate mockgen -destination=mocks/repository.go -package=mocks . UserRepository
type UserRepository interface {
	CreateUser(user *models.User) error

	GetUserByUsername(username string) (*models.User, error)

	CheckPassword(password string, user *models.User) (bool, error)

	UpdateUser(user *models.User, change models.User) (*models.User, error)

	CheckEmailUnique(newEmail string) error

	CheckUnsubscribed(subscriber string, user string) (bool, error)

	Subscribe(subscriber string, user string) error

	Unsubscribe(subscriber string, user string) error

	GetModels(ids []string, limit, offset int) ([]models.UserNoPassword, error)

	GetSubscribers(startIndex int, user string) (int, []models.UserNoPassword, error)

	GetSubscriptions(startIndex int, user string) (int, []models.UserNoPassword, error)

	SearchUsers(query string) ([]models.User, error)

	DeleteUser(username string) error
}
