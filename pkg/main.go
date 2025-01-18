package main

import "github.com/gbburleigh/quick-card-tokenizer/internal/db"

func main() {
	database := db.Create()
	defer database.Close()
}
