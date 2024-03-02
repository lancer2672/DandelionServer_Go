package route

import (
	"github.com/gin-gonic/gin"
	"github.com/lancer2672/DandelionServer_Go/api/handler"
	"github.com/lancer2672/DandelionServer_Go/api/service"
	db "github.com/lancer2672/DandelionServer_Go/db/sqlc"
)

func SetUpRouter(db *db.Store) *gin.Engine {
	router := gin.Default()

	movieService := service.NewMovieService(db)

	movieHandler := handler.NewMovieHanlder(movieService)

	router.GET("/movie/:id", movieHandler.GetMovie)

	return router
}
