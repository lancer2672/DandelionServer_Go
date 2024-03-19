package main

import (
	"database/sql"
	"log"
	"net"

	"github.com/lancer2672/DandelionServer_Go/api/server"
	"github.com/lancer2672/DandelionServer_Go/pb"
	"github.com/lancer2672/DandelionServer_Go/server/sgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/lancer2672/DandelionServer_Go/utils"
	_ "github.com/lib/pq"
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
	runGrpcServer(serverConfig, conn)
	// runGinServer(serverConfig, conn)

}

func runGrpcServer(config utils.Config, conn *sql.DB) {
	server := sgrpc.NewServer(config, conn)
	grpcServer := grpc.NewServer()
	pb.RegisterDandelionServer(grpcServer, server)

	//allow clients to see avaiable grpc server ~ self document
	reflection.Register(grpcServer)
	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal("Cannot create listener GRPC")
	}
	log.Println("GRPC Server started")
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("Cannot start GRPC")
	}

}
func runGinServer(config utils.Config, conn *sql.DB) {
	server := server.NewServer(config, conn)
	err := server.Start()
	if err != nil {
		log.Fatal("Cannot start server", err)
	}
}
