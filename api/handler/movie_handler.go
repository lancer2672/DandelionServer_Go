package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lancer2672/DandelionServer_Go/api/service"
	"github.com/lancer2672/DandelionServer_Go/utils"
)

type MovieHandler struct {
	service service.MovieServiceInterface
}

type MovieHandlerInterface interface {
	GetMovie(*gin.Context)
}

func NewMovieHanlder(service service.MovieServiceInterface) MovieHandlerInterface {
	return &MovieHandler{
		service: service,
	}
}

func (handler *MovieHandler) GetMovie(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorFormatter(err))
		return
	}

	movie, err := handler.service.GetMovie(ctx, int32(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorFormatter(err))
		return
	}
	ctx.JSON(http.StatusOK, movie)
}
