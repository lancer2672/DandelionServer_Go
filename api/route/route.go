package route

import (
	"github.com/gin-gonic/gin"
	"github.com/lancer2672/DandelionServer_Go/api/handler"
	"github.com/lancer2672/DandelionServer_Go/api/middleware"
	"github.com/lancer2672/DandelionServer_Go/api/service"
	db "github.com/lancer2672/DandelionServer_Go/db/sqlc"
	_ "github.com/lancer2672/DandelionServer_Go/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetUpRouter(db *db.Store) *gin.Engine {
	router := gin.Default()

	movieService := service.NewMovieService(db)
	genreService := service.NewGenreService(db)

	movieHandler := handler.NewMovieHandler(movieService)
	genreHandler := handler.NewGenreHandler(genreService)
	uploadHandler := handler.NewUploadHandler()

	router.Use(middleware.ErrorHandler())
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.MaxMultipartMemory = 8
	// @Summary Add a new pet to the store
	// @Description get string by ID
	// @ID get-string-by-int
	// @Accept  json
	// @Produce  json
	// @Param   some_id     path    int     true        "Some ID"
	// @Success 200 {string} string  "ok"
	// @Router /string/{some_id} [get]
	router.POST("/upload", uploadHandler.UploadFile)
	router.GET("/movie/:id", movieHandler.GetMovie)
	router.POST("/movie/:id", movieHandler.CreateMovie)
	router.GET("/movies/seri/:id", movieHandler.GetMoviesBySerie)
	router.GET("/movies/genre/:id", movieHandler.GetMoviesByGenre)

	router.GET("/genres", genreHandler.GetListGenres)

	return router
}
