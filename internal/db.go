package internal

import (
	"context"
	"database/sql"
	_ "modernc.org/sqlite"
	"time"
)

func CreateTableIfNotExists() {
	dbPath := *GetDbFilePath()

	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		panic(err)
	}

	db.ExecContext(context.Background(),
		`CREATE TABLE IF NOT EXISTS applications (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(255) NOT NULL,
    seen REAL DEFAULT (julianday('now'))
		)`)

	defer db.Close()
}

func SaveSeenProcess(proc Process) {
	dbPath := *GetDbFilePath()
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		panic(err)
	}

	statement, err := db.Prepare("INSERT INTO applications (name, seen) VALUES (?, ?)")
	if err != nil {
		panic(err)
	}
	defer statement.Close()

	_, err = statement.Exec(proc.Name, goTimeToSqlRealFormat(proc.Seen))
	if err != nil {
		panic(err)
	}
}

func goTimeToSqlRealFormat(t time.Time) string {
	return t.UTC().Format("2006-01-02 15:04:05")
}
