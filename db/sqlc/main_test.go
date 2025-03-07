package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

var testQueries *Queries
var testDB *pgxpool.Pool

func TestMain(m *testing.M) {
	var err error
	
	// Connect to the database
	connString := "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
	testDB, err = pgxpool.New(context.Background(), connString)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	// Create a new Queries instance
	testQueries = New(testDB)
	
	os.Exit(m.Run())
}