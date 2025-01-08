package internal

import (
	"context"
	"database/sql"
	_ "modernc.org/sqlite"
	"time"
	"time-sink/internal/repository"
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
    seen INTEGER NOT NULL,
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

	existingApplication := repository.FindByNameAndCurrentDay(proc.Name, db)
	if existingApplication != nil {
		newDuration := calculateNewDuration(existingApplication)
		repository.UpdateDuration(*existingApplication.Id, newDuration, db)
		return
	}

	dto := processToApplicationDto(&proc)
	repository.SaveNew(dto, db)
}

func GetDailyRecords() []repository.ApplicationDto {

	db, err := openDb()
	defer db.Close()
	if err != nil {
		panic(err)
	}

	return repository.FindAllByCurrentDay(db)
}

func calculateNewDuration(existingApp *repository.ApplicationDto) int64 {
	now := time.Now()
	seen := existingApp.Seen

	return now.Unix() - seen
}

func processToApplicationDto(proc *Process) *repository.ApplicationDto {
	return &repository.ApplicationDto{
		Name: proc.Name,
		Seen: proc.Seen.Unix(),
	}
}

func openDb() (*sql.DB, error) {
	dbPath := *GetDbFilePath()
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		panic(err)
	}
	return db, err
}
