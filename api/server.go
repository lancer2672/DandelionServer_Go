package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/lancer2672/DandelionServer_Go/db/sqlc"
	"github.com/lancer2672/DandelionServer_Go/utils"
)

type Server struct {
	config utils.Config
	store  *db.Store
	router *gin.Engine
}

func NewServer(config utils.Config, store *db.Store) *Server {
	server := &Server{
		config: config,
		store:  store,
	}
	server.setUpRouter()
	return server
}
func (server *Server) setUpRouter() {
	router := gin.Default()

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
