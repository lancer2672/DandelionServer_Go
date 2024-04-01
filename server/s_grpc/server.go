package sgrpc

import (
	"database/sql"

	db "github.com/lancer2672/DandelionServer_Go/db/sqlc"
	"github.com/lancer2672/DandelionServer_Go/internal/utils"
	"github.com/lancer2672/DandelionServer_Go/pb/service"
)

type Server struct {
	service.UnimplementedMovieServiceServer
	config utils.Config
	store  *db.Store
}

func NewServer(config utils.Config, conn *sql.DB) *Server {
	store := db.NewStore(conn)
	return &Server{
		config: config,
		store:  store,
	}

}
