package db_test

import (
	"database/sql"
	db "github/ludo62/bank_db/db/sqlc"
	"github/ludo62/bank_db/utils"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQuery *db.Queries

func TestMain(m *testing.M) {
	config, err := utils.LoadConfig("../../")
	if err != nil {
		log.Fatal("Impossible de charger la configuration:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Impossible de se connecter avec la base de données:", err)
	}
	testQuery = db.New(conn)

	os.Exit(m.Run())
}
