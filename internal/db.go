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
    seen REAL NOT NULL,
    duration INTEGER
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

	existingProc := getExistingUsageRecord(db, proc)
	if existingProc == nil {
		saveNewProcess(db, proc)
		return
	}

	updateExistingProcess(db, existingProc.Id)
}

func GetDailyRecords(date string) []ProcessUsageDbDto {

	processes := make([]ProcessUsageDbDto, 0)

	db, err := openDb()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM applications WHERE seen = ?", date)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var proc ProcessUsageDbDto
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

func getExistingUsageRecord(db *sql.DB, proc Process) *Process {

	var existingProc ProcessUsageDbDto

	standardDate := ToStandardDateFormat(proc.Seen)
	err := db.QueryRow(`SELECT * FROM applications WHERE name = ? AND seen >= julianday(?) AND seen < julianday(?, "+1 day")`, proc.Name, standardDate, standardDate).
		Scan(&existingProc.Id, &existingProc.Name, &existingProc.Seen, &existingProc.Duration)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		} else {
			panic(err)
		}
	}

	return &Process{
		Name:     existingProc.Name,
		Seen:     FromRealDate(existingProc.Seen),
		Duration: existingProc.Duration,
	}
}

func saveNewProcess(db *sql.DB, proc Process) {

	statement, err := db.Prepare("INSERT INTO applications (name, seen, duration) VALUES (?, ?, 0)")
	if err != nil {
		panic(err)
	}
	defer statement.Close()

	seen := ToRealDate(proc.Seen)
	_, err = statement.Exec(proc.Name, seen)
	if err != nil {
		panic(err)
	}
}

func updateExistingProcess(db *sql.DB, rowId int) {

	existingRecord := getExistingRecordById(db, rowId)
	updatedDuration := calculateDuration(existingRecord)

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

func getExistingRecordById(db *sql.DB, id int) *ProcessUsageDbDto {
	statement, err := db.Prepare("SELECT * FROM applications WHERE id = ?")
	if err != nil {
		panic(err)
	}
	defer statement.Close()

	var existingProc ProcessUsageDbDto
	err = statement.QueryRow(id).
		Scan(&existingProc.Id, &existingProc.Name, &existingProc.Seen, &existingProc.Duration)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		} else {
			panic(err)
		}
	}

	return &existingProc
}

func calculateDuration(dto *ProcessUsageDbDto) int64 {
	then := FromRealDate(dto.Seen).Unix()
	return time.Now().Unix() - then
}
