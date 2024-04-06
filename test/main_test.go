package test

import (
	"context"
	"database/sql"
	"net"

	"github.com/lancer2672/DandelionServer_Go/internal/utils"
	"github.com/lancer2672/DandelionServer_Go/pb/service"
	sgrpc "github.com/lancer2672/DandelionServer_Go/server/s_grpc"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	lis = bufconn.Listen(bufSize)
	config, err := utils.LoadConfig("../.")
	if err != nil {
		log.Error().Err(err).Msg("Error loading config")
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Error().Err(err).Msg("Cannot connect to database")
	}
	movieService := sgrpc.NewMovieService(config, conn)

	//allow clients to see avaiable grpc server ~ self document
	grpcServer := grpc.NewServer()
	service.RegisterMovieServiceServer(grpcServer, movieService)
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatal().Err(err).Msg("Server exited with error")
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}
