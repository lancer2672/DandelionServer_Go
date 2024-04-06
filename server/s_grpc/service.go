package sgrpc

import (
	"database/sql"

	db "github.com/lancer2672/DandelionServer_Go/db/sqlc"
	"github.com/lancer2672/DandelionServer_Go/internal/utils"
	"github.com/lancer2672/DandelionServer_Go/pb/service"
)

type MovieService struct {
	service.UnimplementedMovieServiceServer
	config utils.Config
	store  *db.Store
}

func NewMovieService(config utils.Config, conn *sql.DB) *MovieService {
	store := db.NewStore(conn)
	return &MovieService{
		config: config,
		store:  store,
	}

}
