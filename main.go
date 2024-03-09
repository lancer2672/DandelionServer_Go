package main

import (
	"database/sql"
	"log"

	"github.com/lancer2672/DandelionServer_Go/api/server"

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
	server := server.NewServer(serverConfig, conn)
    if err != nil {
        log.Fatal(err)
    }

	err = server.Start()

	if err != nil {
		log.Fatal("Cannot start server", err)
	}

}
