package route

import (
	"github.com/gin-gonic/gin"
	"github.com/lancer2672/DandelionServer_Go/api/handler"
	"github.com/lancer2672/DandelionServer_Go/api/middleware"
	"github.com/lancer2672/DandelionServer_Go/api/service"
	db "github.com/lancer2672/DandelionServer_Go/db/sqlc"
)

func SetUpRouter(db *db.Store) *gin.Engine {
	router := gin.Default()

	movieService := service.NewMovieService(db)
	genreService := service.NewGenreService(db)

	movieHandler := handler.NewMovieHandler(movieService)
	genreHandler := handler.NewGenreHandler(genreService)

	router.Use(middleware.ErrorHandler())

	router.GET("/movie/:id", movieHandler.GetMovie)
	router.GET("/movies/seri/:id", movieHandler.GetMoviesBySerie)
	router.GET("/movies/genre/:id", movieHandler.GetMoviesByGenre)

	router.GET("/genres", genreHandler.GetListGenres)

	return router
}
