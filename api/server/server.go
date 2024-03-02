package server

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	r "github.com/lancer2672/DandelionServer_Go/api/route"
	db "github.com/lancer2672/DandelionServer_Go/db/sqlc"
	"github.com/lancer2672/DandelionServer_Go/utils"
)

type Server struct {
	config utils.Config
	store  *db.Store
	router *gin.Engine
}

func NewServer(config utils.Config, conn *sql.DB) *Server {
	store := db.NewStore(conn)
	router := r.SetUpRouter(store)

	return &Server{
		config: config,
		store:  store,
		router: router,
	}
}

func (server *Server) Start() error {
	return server.router.Run(server.config.ServerAddress)
}
