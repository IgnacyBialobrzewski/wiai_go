package helpers

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

const dbFile = "db.sqlite"
const migrationFile = "migrations.sql"

func Establish() {
	var err error
	Db, err = sql.Open("sqlite3", dbFile)

	if err != nil {
		log.Fatalln("failed to open connection: %w", err)
	} else {
		log.Println("connected to database")
	}
}

func Migrate() {
	bytes, err := os.ReadFile(migrationFile)
	
	if err != nil {
		log.Fatalf("failed to read migrationFile: %s", err)
	}

	_, err = Db.Exec(string(bytes))

	if err != nil {
		log.Printf("failed to migrate: %+v\n", err)
	}
}