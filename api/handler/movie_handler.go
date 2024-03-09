package handler

import (
	"net/http"
	"strconv"

	req "github.com/lancer2672/DandelionServer_Go/api/request"

	"github.com/gin-gonic/gin"
	"github.com/lancer2672/DandelionServer_Go/api/service"
	db "github.com/lancer2672/DandelionServer_Go/db/sqlc"
	"github.com/lancer2672/DandelionServer_Go/utils"
)

type MovieHandler struct {
	service service.MovieServiceInterface
}

type MovieHandlerInterface interface {
	GetMovie(ctx *gin.Context)
	CreateMovie(ctx *gin.Context)
	GetMoviesByGenre(ctx *gin.Context)
	GetMoviesBySerie(ctx *gin.Context)
}

func NewMovieHandler(service service.MovieServiceInterface) MovieHandlerInterface {
	return &MovieHandler{
		service: service,
	}
}

func (handler *MovieHandler) GetMovie(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.FormatErrorRes(err))
		return
	}

	movie, err := handler.service.GetMovie(ctx, int32(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.FormatErrorRes(err))
		return
	}
	ctx.JSON(http.StatusOK, movie)
}

func (handler *MovieHandler) GetMoviesByGenre(ctx *gin.Context) {
	var req req.MoviesByGenreRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.FormatErrorRes(err))
		return
	}
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.FormatErrorRes(err))
		return
	}

	arg := db.GetMoviesByGenreParams(req)
	movies, err := handler.service.GetMoviesByGenre(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.FormatErrorRes(err))
		return
	}
	utils.RespondSuccess(ctx, movies)
}

func (handler *MovieHandler) GetMoviesBySerie(ctx *gin.Context) {
	var req req.MoviesBySeriesRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.FormatErrorRes(err))
		return
	}
	arg := db.GetMoviesBySerieParams(req)

	movies, err := handler.service.GetMoviesBySerie(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.FormatErrorRes(err))
		return
	}
	utils.RespondSuccess(ctx, movies)

}

func (handler *MovieHandler) SearchMovies(ctx *gin.Context) {
	var req req.SearchMoviesRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.FormatErrorRes(err))
		return
	}

	arg := db.SearchMoviesParams(req)
	movies, err := handler.service.SearchMovies(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.FormatErrorRes(err))
		return
	}
	ctx.JSON(http.StatusOK, movies)
}
func (handler *MovieHandler) CreateMovie(ctx *gin.Context) {
	var req req.CreateMovieRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.FormatErrorRes(err))
		return
	}

	arg := db.CreateMovieParams(req)
	err := handler.service.CreateMovie(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.FormatErrorRes(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Message":"Create successfully",
	})
}
