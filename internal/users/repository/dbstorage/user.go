package localstorage

import (
	"context"
	"database/sql"
	"errors"
	"math"

	"github.com/IfuryI/WEB_BACK/internal/models"
	constants "github.com/IfuryI/WEB_BACK/pkg/const"
	"github.com/jackc/pgconn"
	pgx "github.com/jackc/pgx/v4"
	"golang.org/x/crypto/bcrypt"
)

// PgxPoolIface Интерфейс для драйвера БД
type PgxPoolIface interface {
	Begin(context.Context) (pgx.Tx, error)
	Exec(context.Context, string, ...interface{}) (pgconn.CommandTag, error)
	QueryRow(context.Context, string, ...interface{}) pgx.Row
	Query(context.Context, string, ...interface{}) (pgx.Rows, error)
	Ping(context.Context) error
}

func getHashedPassword(password string) (string, error) {
	hashedPasswordBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPasswordBytes), nil
}

// UserRepository структура репозитория юзера
type UserRepository struct {
	db PgxPoolIface
}

// NewUserRepository инициализация репозитория юзера
func NewUserRepository(database PgxPoolIface) *UserRepository {
	return &UserRepository{
		db: database,
	}
}

// CreateUser создание юзера
func (storage *UserRepository) CreateUser(user *models.User) error {
	hashedPassword, err := getHashedPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	sqlStatement := `
        INSERT INTO mdb.users (login, password, email)
        VALUES ($1, $2, $3)
    `

	_, errDB := storage.db.
		Exec(context.Background(), sqlStatement, user.Username, user.Password, user.Email)

	if errDB != nil {
		return errors.New("create Username Error")
	}

	return nil
}

// CheckEmailUnique проверка уникальности email`а
func (storage *UserRepository) CheckEmailUnique(newEmail string) error {
	sqlStatement := `
        SELECT COUNT(*) as count
        FROM mdb.users
        WHERE email=$1
    `

	var count int
	err := storage.db.
		QueryRow(context.Background(), sqlStatement, newEmail).
		Scan(&count)

	if err != nil || count != 0 {
		return errors.New("email is not unique")
	}

	return nil
}

// GetUserByUsername получить юзера
func (storage *UserRepository) GetUserByUsername(username string) (*models.User, error) {
	var user models.User

	sqlStatement := `
        SELECT login, password, email, img_src, movies_watched, reviews_count, subscribers_count, subscriptions_count
        FROM mdb.users
        WHERE login=$1
    `

	err := storage.db.
		QueryRow(context.Background(), sqlStatement, username).
		Scan(&user.Username, &user.Password,
			&user.Email, &user.Avatar,
			&user.MoviesWatched, &user.ReviewsNumber, &user.Subscribers, &user.Subscriptions)

	if err != nil {
		return nil, errors.New("username not found")
	}

	return &user, nil
}

// CheckPassword проверка пароля
func (storage *UserRepository) CheckPassword(password string, user *models.User) (bool, error) {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) == nil, nil
}

// UpdateUser обновить юзера
func (storage *UserRepository) UpdateUser(user *models.User, change models.User) (*models.User, error) {
	if user.Username != change.Username {
		return nil, errors.New("username doesn't match")
	}

	if change.Password != "" {
		newPassword, err := getHashedPassword(change.Password)
		if err != nil {
			return nil, err
		}

		user.Password = newPassword
	}

	if change.Email != "" {
		user.Email = change.Email
	}

	if change.Avatar != "" {
		user.Avatar = change.Avatar
	}

	if change.ReviewsNumber != nil {
		user.ReviewsNumber = change.ReviewsNumber
	}

	if change.MoviesWatched != nil {
		user.MoviesWatched = change.MoviesWatched
	}

	if change.Subscribers != nil {
		user.Subscribers = change.Subscribers
	}

	if change.Subscriptions != nil {
		user.Subscriptions = change.Subscriptions
	}

	sqlStatement := `
        UPDATE mdb.users
        SET (login, password, email, img_src, movies_watched, reviews_count, subscribers_count, subscriptions_count) =
            ($2, $3, $4, $5, $6, $7, $8, $9)
        WHERE login=$1
    `

	_, err := storage.db.
		Exec(context.Background(), sqlStatement, user.Username,
			user.Username, user.Password,
			user.Email, user.Avatar,
			user.MoviesWatched, user.ReviewsNumber, user.Subscribers, user.Subscriptions)

	if err != nil {
		return nil, errors.New("updating user error")
	}

	return user, nil
}

// CheckUnsubscribed проверка не подписан ли
func (storage *UserRepository) CheckUnsubscribed(subscriber string, user string) (bool, error) {
	sqlStatement := `
        SELECT COUNT(*) as count 
		FROM mdb.subscriptions
		WHERE user_1 = $1 AND user_2 = $2
    `

	var count int
	err := storage.db.
		QueryRow(context.Background(), sqlStatement, subscriber, user).
		Scan(&count)

	if err != nil || count != 0 {
		return false, err
	}

	return true, nil
}

// Subscribe подписаться на юзера
func (storage *UserRepository) Subscribe(subscriber string, user string) error {
	sqlStatement := `
        INSERT INTO mdb.subscriptions(user_1, user_2)
		VALUES ($1, $2)
    `
	_, err := storage.db.
		Exec(context.Background(), sqlStatement, subscriber, user)

	return err
}

// Unsubscribe отписаться от юзера
func (storage *UserRepository) Unsubscribe(subscriber string, user string) error {
	sqlStatement := `
        DELETE FROM mdb.subscriptions
		WHERE user_1 = $1 AND user_2 = $2
    `
	_, err := storage.db.
		Exec(context.Background(), sqlStatement, subscriber, user)

	return err
}

// DeleteUser отписаться от юзера
func (storage *UserRepository) DeleteUser(username string) error {
	sqlStatement := `
        DELETE FROM mdb.users
		WHERE login = $1
    `
	_, err := storage.db.
		Exec(context.Background(), sqlStatement, username)

	return err
}

// GetModels получить модели пользователей
func (storage *UserRepository) GetModels(ids []string, limit, offset int) ([]models.UserNoPassword, error) {
	users := make([]models.UserNoPassword, 0)

	sqlStatement := `
        SELECT login, email, img_src, movies_watched, reviews_count
        FROM mdb.users
        WHERE login=ANY($1) 
		ORDER BY login
		LIMIT $2 OFFSET $3
    `
	rows, err := storage.db.Query(context.Background(), sqlStatement, ids, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		user := models.UserNoPassword{}
		var moviesWatched uint
		var reviewsNumber uint
		err = rows.Scan(&user.Username, &user.Email, &user.Avatar, &moviesWatched, &reviewsNumber)
		if err != nil {
			return nil, err
		}
		user.MoviesWatched = &moviesWatched
		user.ReviewsNumber = &reviewsNumber
		users = append(users, user)
	}

	return users, nil
}

// GetSubscribers получить подписчиков
func (storage *UserRepository) GetSubscribers(startIndex int, user string) (int, []models.UserNoPassword, error) {
	subs := make([]string, 0)

	sqlStatement := `
        SELECT user_1
        FROM mdb.subscriptions
		WHERE user_2 = $1
		ORDER BY user_1
		LIMIT $2 OFFSET $3
    `

	rows, err := storage.db.Query(context.Background(), sqlStatement, user, constants.SubsPageSize, startIndex)
	if err != nil {
		return 0, nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var sub string
		err = rows.Scan(&sub)
		if err != nil {
			return 0, nil, err
		}

		subs = append(subs, sub)
	}

	var rowsCount int
	sqlStatement = `
        SELECT COUNT(*) as count
        FROM mdb.users
        WHERE login=ANY($1)
    `
	err = storage.db.QueryRow(context.Background(), sqlStatement, subs).Scan(&rowsCount)
	if err != nil {
		return 0, nil, err
	}

	users, err := storage.GetModels(subs, constants.SubsPageSize, startIndex)

	if err != nil {
		return 0, nil, err
	}

	pagesNumber := int(math.Ceil(float64(rowsCount) / constants.SubsPageSize))

	return pagesNumber, users, nil
}

// GetSubscriptions получить подписки
func (storage *UserRepository) GetSubscriptions(startIndex int, user string) (int, []models.UserNoPassword, error) {
	subs := make([]string, 0)

	sqlStatement := `
        SELECT user_2
        FROM mdb.subscriptions
		WHERE user_1 = $1
		ORDER BY user_2
		LIMIT $2 OFFSET $3
    `

	rows, err := storage.db.Query(context.Background(), sqlStatement, user, constants.SubsPageSize, startIndex)
	if err != nil {
		return 0, nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var sub string
		err = rows.Scan(&sub)
		if err != nil {
			return 0, nil, err
		}

		subs = append(subs, sub)
	}

	var rowsCount int
	sqlStatement = `
        SELECT COUNT(*) as count
        FROM mdb.users
        WHERE login=ANY($1)
    `
	err = storage.db.QueryRow(context.Background(), sqlStatement, subs).Scan(&rowsCount)

	if err != nil {
		return 0, nil, err
	}
	users, err := storage.GetModels(subs, constants.SubsPageSize, startIndex)

	if err != nil {
		return 0, nil, err
	}
	pagesNumber := int(math.Ceil(float64(rowsCount) / constants.SubsPageSize))

	return pagesNumber, users, nil
}

// SearchUsers поиск по юзерам
func (storage *UserRepository) SearchUsers(query string) ([]models.User, error) {
	sqlSearchUsers := `
		SELECT login, img_src
		FROM mdb.users
		WHERE lower(login) LIKE '%' || $1 || '%'
	`

	rows, err := storage.db.Query(context.Background(), sqlSearchUsers, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		user := models.User{}
		err = rows.Scan(&user.Username, &user.Avatar)
		if err != nil && err != sql.ErrNoRows {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}
