package services

import (
	"context"
	"database/sql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
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

func SaveApplication(name string) {
	dbPath := *internal.GetDbFilePath()
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	log.Printf("INFO: Searching for application with name %s", name)
	application := repository.GetApplicationByNameForToday(name, db)
	if application != nil {
		updateDuration(application)
	} else {
		log.Printf("WARN: No application with name %s found... creating new DB entry", name)
		application = &repository.Application{
			Name:     name,
			Seen:     time.Now().Unix(),
			Duration: 0,
		}
	}

	log.Printf("INFO: Saving application %v", application)
	repository.SaveApplication(*application, db)
}

func GetDailyApplications() []repository.Application {
	dbPath := *internal.GetDbFilePath()
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	start := time.Now()
	end := start.AddDate(0, 0, 1)
	return repository.GetAllApplicationsByDates(start, end, db)
}

func GetAllApplicationsBetweenDates(start, end time.Time) []repository.Application {
	dbPath := *internal.GetDbFilePath()
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return repository.GetAllApplicationsByDates(start, end, db)
}

func GetAll() []repository.Application {
	dbPath := *internal.GetDbFilePath()
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return repository.GetAllApplications(db)
}

func updateDuration(application *repository.Application) {
	now := time.Now().Unix()
	newDuration := now - application.Seen

	log.Printf("INFO: Updating duration from %d to %d", application.Duration, newDuration)
	application.Duration = newDuration
}
