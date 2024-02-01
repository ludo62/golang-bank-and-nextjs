package db_test

import (
	"database/sql"
	db "github/ludo62/bank_db/db/sqlc"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQuery *db.Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open("postgres", "postgres://root:secret@localhost:5432/bank_db?sslmode=disable")
	if err != nil {
		log.Fatal("Impossible de se connecter avec la base de données:", err)
	}
	testQuery = db.New(conn)

	os.Exit(m.Run())
}
