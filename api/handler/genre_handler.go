package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	req "github.com/lancer2672/DandelionServer_Go/api/request"
	"github.com/lancer2672/DandelionServer_Go/api/service"
	db "github.com/lancer2672/DandelionServer_Go/db/sqlc"

	"github.com/lancer2672/DandelionServer_Go/utils"
)

type GenreHandler struct {
	service service.GenreServiceInterface
}

type GenreHandlerInterface interface {
	GetListGenres(ctx *gin.Context)
}

func NewGenreHandler(service service.GenreServiceInterface) GenreHandlerInterface {
	return &GenreHandler{
		service: service,
	}
}

func (handler *GenreHandler) GetListGenres(ctx *gin.Context) {
	var req req.GetListGenresRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.FormatErrorRes(err))
		return
	}

	arg := db.GetListGenresParams(req)
	genres, err := handler.service.GetListGenres(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.FormatErrorRes(err))
		return
	}
	utils.RespondSuccess(ctx, genres)
	// ctx.JSON(http.StatusOK, gin.H{
	// 	"code":    http.StatusOK,
	// 	"data":    genres,
	// 	"message": "success",
	// })
}
