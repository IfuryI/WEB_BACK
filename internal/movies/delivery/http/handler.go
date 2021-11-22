package http

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/IfuryI/WEB_BACK/internal/logger"
	"github.com/IfuryI/WEB_BACK/internal/models"
	"github.com/IfuryI/WEB_BACK/internal/movies"
	constants "github.com/IfuryI/WEB_BACK/pkg/const"
	"github.com/gin-gonic/gin"
)

// Handler структура хендлера
type Handler struct {
	useCase movies.UseCase
	Log     *logger.Logger
}

type moviesPageResponse struct {
	CurrentPage int             `json:"current_page"`
	PagesNumber int             `json:"pages_number"`
	MaxItems    int             `json:"max_items"`
	Movies      []*models.Movie `json:"movies"`
}

// NewHandler инициализация новго хендлера
func NewHandler(useCase movies.UseCase, Log *logger.Logger) *Handler {
	return &Handler{
		useCase: useCase,
		Log:     Log,
	}
}


// CreateMovie создание фильма
func (h *Handler) CreateMovie(ctx *gin.Context) {
	movieData := new(models.Movie)
	err := ctx.BindJSON(movieData)
	if err != nil {
		h.Log.LogWarning(ctx, "movie", "CreateMovie", err.Error())
		ctx.AbortWithStatus(http.StatusBadRequest) // 400
		return
	}

	err = h.useCase.CreateMovie(movieData)
	if err != nil {
		h.Log.LogError(ctx, "movie", "CreateMovie", err)
		ctx.AbortWithStatus(http.StatusInternalServerError) // 500
		return
	}

	ctx.Status(http.StatusCreated)
}

// GetMovie godoc
// @Summary "получение информации о фильме по Id"
// @Produce json
// @Param id path integer true "Movie ID"
// @Success 200 {object} models.Movie
// @Failure 404 "Фильм не найден"
// @Router /movies/{id} [get]
// @Security UserKeyAuth
// GetMovie получение информации о фильме
func (h *Handler) GetMovie(ctx *gin.Context) {
	auth, ok := ctx.Get(constants.AuthStatusKey)
	authBool := auth.(bool)
	username := ""
	if ok && authBool {
		user, ok := ctx.Get(constants.UserKey)
		if ok {
			userModel := user.(models.User)
			username = userModel.Username
		}
	}

	movie, err := h.useCase.GetMovie(ctx.Param("id"), username)
	if err != nil {
		h.Log.LogWarning(ctx, "movie", "GetMovie", err.Error())
		ctx.AbortWithStatus(http.StatusNotFound) // 404
		return
	}

	ctx.JSON(http.StatusOK, movie)
}

// GetMovies godoc
// @Summary "получение фильмов"
// @Produce json
// @Param category path integer true "category"
// @Param category path integer true "genre"
// @Success 200 {object} []models.Movie
// @Failure 500 "Ошибка получения фильмов"
// @Router /movies [get]
// @Security UserKeyAuth
// GetMovies получить фильмы
func (h *Handler) GetMovies(ctx *gin.Context) {
	category := ctx.Query("category")
	if category == "best" {
		h.GetBestMovies(ctx)
	} else if category == "genre" {
		h.GetMoviesByGenres(ctx)
	}
}

// GetBestMovies получить лучшие фильмы
func (h *Handler) GetBestMovies(ctx *gin.Context) {
	auth, ok := ctx.Get(constants.AuthStatusKey)
	authBool := auth.(bool)
	username := ""
	if ok && authBool {
		user, ok := ctx.Get(constants.UserKey)
		if ok {
			userModel := user.(models.User)
			username = userModel.Username
		}
	}

	page, err := strconv.Atoi(ctx.DefaultQuery("page", constants.PageDefault))
	if err != nil || page < 1 {
		h.Log.LogWarning(ctx, "movie", "GetBestMovies", err.Error())
		ctx.AbortWithStatus(http.StatusBadRequest) // 400
		return
	}

	pagesNumber, bestMovies, err := h.useCase.GetBestMovies(page, username)

	if err != nil {
		h.Log.LogError(ctx, "movie", "GetBestMovies", err)
		ctx.AbortWithStatus(http.StatusInternalServerError) // 500
		return
	}

	moviesResponse := moviesPageResponse{
		CurrentPage: page,
		PagesNumber: pagesNumber,
		MaxItems:    constants.MoviesPageSize,
		Movies:      bestMovies,
	}

	ctx.JSON(http.StatusOK, moviesResponse)
}

// GetGenres godoc
// @Summary "получение списка всех жанров"
// @Produce json
// @Success 200 {object} []string
// @Failure 500 "Ошибка получения"
// @Router /genres [get]
// @Security UserKeyAuth
// GetGenres получить доступные жанры
func (h *Handler) GetGenres(ctx *gin.Context) {
	genres, err := h.useCase.GetAllGenres()
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError) // 500
		return
	}

	ctx.JSON(http.StatusOK, genres)
}

// GetMoviesByGenres получить фильмы по жанрам
func (h *Handler) GetMoviesByGenres(ctx *gin.Context) {
	auth, ok := ctx.Get(constants.AuthStatusKey)
	authBool := auth.(bool)
	username := ""
	if ok && authBool {
		user, ok := ctx.Get(constants.UserKey)
		if ok {
			userModel := user.(models.User)
			username = userModel.Username
		}
	}

	genresQuery := ctx.Query("filter")
	if genresQuery == "" {
		ctx.AbortWithStatus(http.StatusBadRequest) // 400
		return
	}

	genres := strings.Split(genresQuery, " ")

	page, err := strconv.Atoi(ctx.DefaultQuery("page", constants.PageDefault))
	if err != nil || page < 1 {
		ctx.AbortWithStatus(http.StatusBadRequest) // 400
		return
	}

	pagesNumber, moviesList, err := h.useCase.GetMoviesByGenres(genres, page, username)

	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError) // 500
		return
	}

	moviesResponse := moviesPageResponse{
		CurrentPage: page,
		PagesNumber: pagesNumber,
		MaxItems:    constants.MoviesPageSize,
		Movies:      moviesList,
	}

	ctx.JSON(http.StatusOK, moviesResponse)
}

// MarkWatched godoc
// @Summary "Установка у юзера статус просмотренно для фильма"
// @Produce json
// @Param id path integer true "Movie ID"
// @Success 200 "Все ок"
// @Failure 500 "Ошибка простановки статуса"
// @Router /movies/{id}/watch [post]
// @Security UserKeyAuth
// MarkWatched отметить просмотренным
func (h *Handler) MarkWatched(ctx *gin.Context) {
	user, ok := ctx.Get(constants.UserKey)
	if !ok {
		err := fmt.Errorf("%s", "Failed to retrieve user from context")
		h.Log.LogWarning(ctx, "movies", "MarkWatched", err.Error())
		ctx.AbortWithStatus(http.StatusBadRequest) // 400
		return
	}

	userModel, ok := user.(models.User)
	if !ok {
		err := fmt.Errorf("%s", "Failed to cast user to model")
		h.Log.LogError(ctx, "movies", "MarkWatched", err)
		ctx.AbortWithStatus(http.StatusInternalServerError) // 500
		return
	}

	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		h.Log.LogError(ctx, "movies", "MarkWatched", err)
		ctx.AbortWithStatus(http.StatusInternalServerError) // 500
		return
	}
	err = h.useCase.MarkWatched(userModel, idInt)
	if err != nil {
		h.Log.LogError(ctx, "movies", "MarkWatched", err)
		ctx.AbortWithStatus(http.StatusInternalServerError) // 500
		return
	}

	ctx.Status(http.StatusOK)
}

// MarkUnwatched godoc
// @Summary "Удаление у юзера статус просмотренно для фильма"
// @Produce json
// @Param id path integer true "Movie ID"
// @Success 200 "Все ок"
// @Failure 500 "Ошибка удаления статуса"
// @Router /movies/{id}/watch [delete]
// @Security UserKeyAuth
// MarkUnwatched отметить непросмотренным
func (h *Handler) MarkUnwatched(ctx *gin.Context) {
	user, ok := ctx.Get(constants.UserKey)
	if !ok {
		err := fmt.Errorf("%s", "Failed to retrieve user from context")
		h.Log.LogWarning(ctx, "movies", "MarkUnwatched", err.Error())
		ctx.AbortWithStatus(http.StatusBadRequest) // 400
		return
	}

	userModel, ok := user.(models.User)
	if !ok {
		err := fmt.Errorf("%s", "Failed to cast user to model")
		h.Log.LogError(ctx, "movies", "MarkUnwatched", err)
		ctx.AbortWithStatus(http.StatusInternalServerError) // 500
		return
	}

	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		h.Log.LogError(ctx, "movies", "MarkUnwatched", err)
		ctx.AbortWithStatus(http.StatusInternalServerError) // 500
		return
	}
	err = h.useCase.MarkUnwatched(userModel, idInt)
	if err != nil {
		h.Log.LogError(ctx, "movies", "MarkUnwatched", err)
		ctx.AbortWithStatus(http.StatusInternalServerError) // 500
		return
	}

	ctx.Status(http.StatusOK)
}

// GetSimilar godoc
// @Summary "получение списка похожих фильмов"
// @Produce json
// @Param id path integer true "Movie ID"
// @Success 200 {object} []models.Movie
// @Failure 500 "Ошибка получения"
// @Router /movies/{id}/similar [get]
// @Security UserKeyAuth
// GetSimilar получить похожие
func (h *Handler) GetSimilar(ctx *gin.Context) {
	similarMovies, err := h.useCase.GetSimilar(ctx.Param("id"))
	if err != nil {
		h.Log.LogError(ctx, "movies", "GetSimilar", err)
		ctx.AbortWithStatus(http.StatusInternalServerError) // 500
		return
	}

	ctx.JSON(http.StatusOK, similarMovies)
}
