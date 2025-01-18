package db

import (
	"database/sql"
	"embed"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

//go:embed migrations/create.sql
var migrations embed.FS

func Read(path string) string {
	b, err := migrations.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func Create() *sql.DB {
	dbPath := os.Getenv("TOKEN_DB_PATH")
	if dbPath == "" {
		dbPath = "tokens.db"
	}
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal("Error opening database:", err)
	}

	if _, err := db.Exec(Read("migrations/create.sql")); err != nil {
		log.Fatal("Error running migrations", err)
	}

	DB = db
	return DB
}

func Execute(db *sql.DB, statement string) sql.Result {
	res, err := db.Exec(statement)
	if err != nil {
		log.Fatal(err)
	}

	return res
}
