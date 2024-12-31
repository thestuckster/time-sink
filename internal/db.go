package internal

import (
	"context"
	"database/sql"
	"errors"
	_ "modernc.org/sqlite"
	"time"
)

func CreateTableIfNotExists() {
	dbPath := *GetDbFilePath()

	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		panic(err)
	}

	_, err = db.ExecContext(context.Background(),
		`CREATE TABLE IF NOT EXISTS applications (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(255) NOT NULL,
    seen TEXT NOT NULL,
    duration TEXT
		)`)
	if err != nil {
		return
	}

	defer db.Close()
}

func SaveSeenProcess(proc Process) {
	dbPath := *GetDbFilePath()
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	existingProc := checkForExistingRecord(db, proc)
	if existingProc == nil {
		saveNewProcess(db, proc)
		return
	}

	updatedDuration := FormatTimeToHHMMSS(ParseHHMMSS(existingProc.Duration).Add(time.Minute))
	updateExistingProcess(db, existingProc.Id, updatedDuration)
}

func GetDailyRecords(date string) []ProcessDto {

	processes := make([]ProcessDto, 0)

	db, err := openDb()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM applications WHERE seen = ?", date)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var proc ProcessDto
		err = rows.Scan(&proc.Id, &proc.Name, &proc.Seen, &proc.Duration)
		if err != nil {
			panic(err)
		}

		processes = append(processes, proc)
	}

	return processes
}

func openDb() (*sql.DB, error) {
	dbPath := *GetDbFilePath()
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		panic(err)
	}
	return db, err
}

func checkForExistingRecord(db *sql.DB, proc Process) *Process {

	var existingProc ProcessDto

	seenDate := ToStandardDateFormat(proc.Time)

	err := db.QueryRow("SELECT * FROM applications WHERE name = ? AND seen = ?", proc.Name, seenDate).
		Scan(&existingProc.Id, &existingProc.Name, &existingProc.Seen, &existingProc.Duration)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		} else {
			panic(err)
		}
	}

	return &Process{
		Id:       existingProc.Id,
		Name:     existingProc.Name,
		Seen:     existingProc.Seen,
		Duration: existingProc.Duration,
	}
}

func saveNewProcess(db *sql.DB, proc Process) {

	statement, err := db.Prepare("INSERT INTO applications (name, seen, duration) VALUES (?, ?, ?)")
	if err != nil {
		panic(err)
	}
	defer statement.Close()

	seen := ToStandardDateFormat(proc.Time)
	_, err = statement.Exec(proc.Name, seen, StartingDuration())
	if err != nil {
		panic(err)
	}
}

func updateExistingProcess(db *sql.DB, rowId int, updatedDuration string) {
	updateQuery := `
UPDATE applications 
SET duration = ? 
WHERE id = ?
`
	_, err := db.Exec(updateQuery, updatedDuration, rowId)
	if err != nil {
		panic(err)
	}
}
