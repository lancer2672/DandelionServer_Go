package main

import (
	"context"
	"database/sql"
	"log"
	"net"
	"net/http"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/lancer2672/DandelionServer_Go/api/server"
	"github.com/lancer2672/DandelionServer_Go/pb/service"
	"github.com/lancer2672/DandelionServer_Go/server/sgrpc"
	"github.com/lancer2672/DandelionServer_Go/utils"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	serverConfig, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}
	conn, err := sql.Open(serverConfig.DBDriver, serverConfig.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to database", err)
	}
	runDatabaseMigration(serverConfig.MigrationUrl, serverConfig.DBSource)
	go runGatewayServer(serverConfig, conn)
	runGrpcServer(serverConfig, conn)
	// runGinServer(serverConfig, conn)

}

func runGrpcServer(config utils.Config, conn *sql.DB) {
	server := sgrpc.NewServer(config, conn)
	grpcServer := grpc.NewServer()
	service.RegisterDandelionServer(grpcServer, server)

	//allow clients to see avaiable grpc server ~ self document
	reflection.Register(grpcServer)
	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal("Cannot create listener GRPC", err)
	}
	log.Println("GRPC Server started", config.GRPCServerAddress)
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("Cannot start GRPC", err)
	}

}

func runGatewayServer(config utils.Config, conn *sql.DB) {
	server := sgrpc.NewServer(config, conn)

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
	err := service.RegisterDandelionHandlerServer(ctx, grpcMux, server)
	if err != nil {
		log.Fatal("Cannot create register server")
	}
	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)
	listener, err := net.Listen("tcp", config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot create listener HTTP Gateway", err)
	}
	log.Println("HTTP gateway Server started", config.ServerAddress)
	err = http.Serve(listener, mux)
	if err != nil {
		log.Fatal("Cannot start HTTP Gateway", err)
	}

}

func runGinServer(config utils.Config, conn *sql.DB) {
	server := server.NewServer(config, conn)
	err := server.Start()
	if err != nil {
		log.Fatal("Cannot start server", err)
	}
}

func runDatabaseMigration(url string, dbSource string) {
	m, err := migrate.New(
		url,
		dbSource)
	if err != nil {
		log.Fatal("cannot create migrate instance ", err)
	}
	if err = m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("failed to migrate up", err)
	}
	log.Println("run db migration successfully")

}
