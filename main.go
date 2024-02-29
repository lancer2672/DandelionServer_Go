package main

import (
	"database/sql"
	"log"

	"github.com/lancer2672/DandelionServer_Go/api"
	db "github.com/lancer2672/DandelionServer_Go/db/sqlc"
	"github.com/lancer2672/DandelionServer_Go/utils"
	_ "github.com/lib/pq"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to database", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(config, store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot start server", err)
	}

}
