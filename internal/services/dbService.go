package services

import (
	"context"
	"database/sql"
	_ "modernc.org/sqlite"
	"time"
	"time-sink/internal"
	"time-sink/internal/repository"
)

func CreateTableIfNotExists() {
	dbPath := *internal.GetDbFilePath()

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

func SaveSeenProcess(proc internal.Process) {
	dbPath := *internal.GetDbFilePath()
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

func GetDailyRecords() []repository.ApplicationRecordEntity {

	db, err := openDb()
	defer db.Close()
	if err != nil {
		panic(err)
	}

	return repository.FindAllByCurrentDay(db)
}

func GetRecordsInRange(start, end time.Time) []repository.ApplicationRecordEntity {
	db, err := openDb()
	defer db.Close()
	if err != nil {
		panic(err)
	}

	return repository.FindAllInRange(db, start.Unix(), end.Unix())
}

func calculateNewDuration(existingApp *repository.ApplicationRecordEntity) int64 {
	now := time.Now()
	seen := existingApp.Seen

	return now.Unix() - seen
}

func processToApplicationDto(proc *internal.Process) *repository.ApplicationRecordEntity {
	return &repository.ApplicationRecordEntity{
		Name: proc.Name,
		Seen: proc.Seen.Unix(),
	}
}

func openDb() (*sql.DB, error) {
	dbPath := *internal.GetDbFilePath()
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		panic(err)
	}
	return db, err
}
