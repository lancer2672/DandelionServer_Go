package database

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/lancer2672/DandelionServer_Go/utils"
)

var testStore Store

func TestMain(t *testing.M) {
	gin.SetMode(gin.TestMode)

	config, err := utils.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect db:", err)
	}
	testStore = *NewStore(conn)
	os.Exit(t.Run())
}
