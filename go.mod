module github.com/gbburleigh/quick-card-tokenizer

go 1.23.5

require github.com/mattn/go-sqlite3 v1.14.17

replace github.com/gbburleigh/quick-card-tokenizer/token => ./pkg/token

replace github.com/gbburleigh/quick-card-tokenizer/db => ./pkg/db

replace github.com/gbburleigh/quick-card-tokenizer/db/migrations => ./pkg/db/migrations
