package playlists

import "github.com/IfuryI/WEB_BACK/internal/models"

// UseCase go:generate mockgen -destination=mocks/usecase.go -package=mocks . UseCase
type UseCase interface {
	CreatePlaylist(username string, playlistName string, isShared bool) error

	GetPlaylist(playlistID int) (*models.Playlist, error)
	GetPlaylists(username string) ([]models.Playlist, error)
	GetPlaylistsInfo(username string, movieID int) ([]models.PlaylistsInfo, error)

	UpdatePlaylist(username string, playlistID int, playlistName string, isShared bool) error
	DeletePlaylist(username string, playlistID int) error

	AddMovieToPlaylist(username string, playlistID int, movieID int) error
	DeleteMovieFromPlaylist(username string, playlistID int, movieID int) error

	AddUserToPlaylist(username string, playlistID int, usernameToAdd string) error
	DeleteUserFromPlaylist(username string, playlistID int, usernameToDelete string) error
}
