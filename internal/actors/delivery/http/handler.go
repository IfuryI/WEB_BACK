package http

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/IfuryI/WEB_BACK/internal/actors"
	"github.com/IfuryI/WEB_BACK/internal/logger"
	"github.com/IfuryI/WEB_BACK/internal/models"
	constants "github.com/IfuryI/WEB_BACK/pkg/const"
)

// Handler структура хендлера
type Handler struct {
	useCase actors.UseCase
	Log     *logger.Logger
}

// NewHandler Инициализация хендлера
func NewHandler(useCase actors.UseCase, Log *logger.Logger) *Handler {
	return &Handler{
		useCase: useCase,
		Log:     Log,
	}
}

// func (h *Handler) CreateActor(ctx *gin.Context) {
// 	user, ok := ctx.Get(constants.UserKey)
// 	if !ok {
// 		err := fmt.Errorf("%s", "Failed to retrieve user from context")
// 		h.Log.LogWarning(ctx, "actors", "CreateActor", err.Error())
// 		ctx.AbortWithStatus(http.StatusBadRequest) // 400
// 		return
// 	}

// 	userModel, ok := user.(models.User)
// 	if !ok {
// 		err := fmt.Errorf("%s", "Failed to cast user to model")
// 		h.Log.LogError(ctx, "actors", "CreateActor", err)
// 		ctx.AbortWithStatus(http.StatusInternalServerError) // 500
// 		return
// 	}

// 	actorModel := new(models.Actor)
// 	err := ctx.BindJSON(actorModel)
// 	if err != nil {
// 		h.Log.LogWarning(ctx, "actors", "CreateActor", err.Error())
// 		ctx.AbortWithStatus(http.StatusBadRequest) // 400
// 		return
// 	}

// 	err = h.useCase.CreateActor(userModel, *actorModel)
// 	if err != nil {
// 		h.Log.LogError(ctx, "actors", "CreateActor", err)
// 		ctx.AbortWithStatus(http.StatusInternalServerError) // 500
// 		return
// 	}
// }

// GetActor Функция получения информации об актере
func (h *Handler) GetActor(ctx *gin.Context) {
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

	id := ctx.Param("actor_id")

	actor, err := h.useCase.GetActor(id, username)
	if err != nil {
		h.Log.LogWarning(ctx, "actors", "GetActor", err.Error())
		ctx.AbortWithStatus(http.StatusBadRequest) // 400
		return
	}

	ctx.JSON(http.StatusOK, actor)
}

// func (h *Handler) EditActor(ctx *gin.Context) {
// 	user, ok := ctx.Get(constants.UserKey)
// 	if !ok {
// 		err := fmt.Errorf("%s", "Failed to retrieve user from context")
// 		h.Log.LogWarning(ctx, "actors", "EditActor", err.Error())
// 		ctx.AbortWithStatus(http.StatusBadRequest) // 400
// 		return
// 	}

// 	userModel, ok := user.(models.User)
// 	if !ok {
// 		err := fmt.Errorf("%s", "Failed to cast user to model")
// 		h.Log.LogError(ctx, "actors", "EditActor", err)
// 		ctx.AbortWithStatus(http.StatusInternalServerError) // 500
// 		return
// 	}

// 	id := ctx.Param("actor_id")

// 	change := new(models.Actor)
// 	err := ctx.BindJSON(change)
// 	if err != nil {
// 		h.Log.LogWarning(ctx, "actors", "EditActor", err.Error())
// 		ctx.AbortWithStatus(http.StatusBadRequest) // 400
// 		return
// 	}

// 	change.ID = id

// 	changed, err := h.useCase.EditActor(userModel, *change)
// 	if err != nil {
// 		h.Log.LogError(ctx, "actors", "EditActor", err)
// 		ctx.AbortWithStatus(http.StatusInternalServerError) // 500
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, changed)
// }

// LikeActor Поставить лайк актеру
func (h *Handler) LikeActor(ctx *gin.Context) {
	user, ok := ctx.Get(constants.UserKey)
	if !ok {
		err := fmt.Errorf("%s", "Failed to retrieve user from context")
		h.Log.LogWarning(ctx, "actors", "LikeActor", err.Error())
		ctx.AbortWithStatus(http.StatusBadRequest) // 400
		return
	}

	userModel, ok := user.(models.User)
	if !ok {
		err := fmt.Errorf("%s", "Failed to cast user to model")
		h.Log.LogError(ctx, "actors", "LikeActor", err)
		ctx.AbortWithStatus(http.StatusInternalServerError) // 500
		return
	}

	id := ctx.Param("actor_id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		h.Log.LogError(ctx, "actors", "LikeActor", err)
		ctx.AbortWithStatus(http.StatusInternalServerError) // 500
		return
	}
	err = h.useCase.LikeActor(userModel.Username, idInt)
	if err != nil {
		h.Log.LogError(ctx, "actors", "LikeActor", err)
		ctx.AbortWithStatus(http.StatusInternalServerError) // 500
		return
	}

	ctx.Status(http.StatusOK)
}

// UnlikeActor убрать лайк актера
func (h *Handler) UnlikeActor(ctx *gin.Context) {
	user, ok := ctx.Get(constants.UserKey)
	if !ok {
		err := fmt.Errorf("%s", "Failed to retrieve user from context")
		h.Log.LogWarning(ctx, "actors", "UnlikeActor", err.Error())
		ctx.AbortWithStatus(http.StatusBadRequest) // 400
		return
	}

	userModel, ok := user.(models.User)
	if !ok {
		err := fmt.Errorf("%s", "Failed to cast user to model")
		h.Log.LogError(ctx, "actors", "UnlikeActor", err)
		ctx.AbortWithStatus(http.StatusInternalServerError) // 500
		return
	}

	id := ctx.Param("actor_id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		h.Log.LogError(ctx, "actors", "UnlikeActor", err)
		ctx.AbortWithStatus(http.StatusInternalServerError) // 500
		return
	}
	err = h.useCase.UnlikeActor(userModel.Username, idInt)
	if err != nil {
		h.Log.LogError(ctx, "actors", "UnlikeActor", err)
		ctx.AbortWithStatus(http.StatusInternalServerError) // 500
		return
	}

	ctx.Status(http.StatusOK)
}
