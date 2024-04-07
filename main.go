package main

import (
	"context"
	"database/sql"

	"net"
	"net/http"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/lancer2672/DandelionServer_Go/internal/helper"
	"github.com/lancer2672/DandelionServer_Go/internal/middleware"
	"github.com/lancer2672/DandelionServer_Go/internal/utils"
	"github.com/lancer2672/DandelionServer_Go/pb/service"
	sgrpc "github.com/lancer2672/DandelionServer_Go/server/s_grpc"
	shttp "github.com/lancer2672/DandelionServer_Go/server/s_http"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	serverConfig, err := utils.LoadConfig(".")
	if err != nil {
		log.Error().Err(err).Msg("Error loading config")
	}
	helper.ConfigHttpClient(serverConfig)
	conn, err := sql.Open(serverConfig.DBDriver, serverConfig.DBSource)
	if err != nil {
		log.Error().Err(err).Msg("Cannot connect to database")
		os.Exit(1)
	}
	// runDatabaseMigration(serverConfig.MigrationUrl, serverConfig.DBSource)
	// go runGinServer(serverConfig, conn)
	go runGatewayServer(serverConfig, conn)
	runGrpcServer(serverConfig, conn)

}

func runGrpcServer(config utils.Config, conn *sql.DB) {
	movieService := sgrpc.NewMovieService(config, conn)

	//allow clients to see avaiable grpc server ~ self document
	grpcServer := grpc.NewServer(grpc.ChainUnaryInterceptor(middleware.LoggerInterceptor, middleware.CheckApiKeyInterceptor))
	service.RegisterMovieServiceServer(grpcServer, movieService)
	reflection.Register(grpcServer)
	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Error().Err(err).Msg("Cannot create listener GRPC")
		os.Exit(1)
	}
	log.Info().Str("address", config.GRPCServerAddress).Msg("GRPC Server started")
	if err = grpcServer.Serve(listener); err != nil {
		log.Error().Err(err).Msg("Cannot start GRPC")
		os.Exit(1)
	}

}

func runGatewayServer(config utils.Config, conn *sql.DB) {
	movieService := sgrpc.NewMovieService(config, conn)
	grpcMux := runtime.NewServeMux(
		//disable to keep snake_case variable names
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				UseProtoNames: true,
			},
			UnmarshalOptions: protojson.UnmarshalOptions{
				DiscardUnknown: true,
			},
		}),
	)

	ctx, cancel := context.WithCancel(context.Background())
	//prevent system doing unnecessary works
	defer cancel()
	err := service.RegisterMovieServiceHandlerServer(ctx, grpcMux, movieService)
	if err != nil {
		log.Error().Err(err).Msg("Cannot create listener HTTP Gateway")
		os.Exit(1)
	}

	mux := http.NewServeMux()

	// mux.Handle("/", middleware.CorsMiddleware(middleware.Logger(middleware.CheckApiKey(grpcMux))))
	mux.HandleFunc("/movies/stream", shttp.StreamFile)
	listener, err := net.Listen("tcp", config.ServerAddress)
	if err != nil {
		log.Error().Err(err).Msg("Cannot create listener HTTP Gateway")
	}
	mux.Handle("/", middleware.Logger(grpcMux))
	// err = http.Serve(listener, mux)
	log.Info().Str("address", config.ServerAddress).Msg("HTTP_GRPC gateway Server started")
	if err = http.Serve(listener, mux); err != nil {
		log.Error().Err(err).Msg("Cannot start HTTP Gateway")
		os.Exit(1)
	}

}

func runDatabaseMigration(url string, dbSource string) {
	m, err := migrate.New(
		url,
		dbSource)
	if err != nil {
		log.Error().Err(err).Msg("Cannot create migrate instance")
		os.Exit(1)
	}
	if err = m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Error().Err(err).Msg("Failed to migrate up")
		os.Exit(1)
	}
	log.Info().Msg("Run db migration successfully")

}
