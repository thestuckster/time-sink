package repository

import (
	"errors"
	"gorm.io/gorm"
	"log"
	_ "modernc.org/sqlite"
	"time"
	"time-sink/internal/helpers"
)

type Application struct {
	gorm.Model
	Name     string
	Seen     int64
	Duration int64
}

func autoMigrate(db *gorm.DB) {
	db.AutoMigrate(&Application{})
}

func SaveApplication(application Application, db *gorm.DB) {

	//autoMigrate(db)

	result := db.Save(&application)
	if result.Error != nil {
		panic(result.Error)
	}
}

func GetApplicationById(id int, db *gorm.DB) *Application {
	var application Application
	result := db.First(&application, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil
		}
	}

	return &application
}

func GetApplicationByNameForToday(name string, db *gorm.DB) *Application {
	start := time.Now()
	end := start.AddDate(0, 0, 1)

	return GetApplicationByNameAndDates(name, start, end, db)
}

func GetApplicationByNameAndDates(name string, start, end time.Time, db *gorm.DB) *Application {
	startUnix := helpers.GetStartOfDayUnixEpoch(start)
	endUnix := helpers.GetStartOfDayUnixEpoch(end)

	var application Application

	log.Printf("INFO: Searching for application named \"%s\" with seen date between %d and %d", name, startUnix, endUnix)
	result := db.Where("name = ? AND seen >= ? AND seen <= ?", name, startUnix, endUnix).First(&application)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			log.Printf("WARNING: No application with name %s was found.", name)
			return nil
		}
	}
	log.Printf("INFO: Found application with name %s\n %+v", name, application)
	return &application
}

func GetAllApplicationsByDates(start, end time.Time, db *gorm.DB) []Application {
	var applications []Application

	startUnix := helpers.GetStartOfDayUnixEpoch(start)
	endUnix := helpers.GetStartOfDayUnixEpoch(end)

	log.Printf("Searching for all applications with seen between %d and %d", startUnix, endUnix)
	result := db.Where("seen >= ? AND seen <= ?", startUnix, endUnix).Find(&applications)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			log.Printf("WARNING: No applications were found. Returning empty array")
			return make([]Application, 0)
		}

	}

	log.Printf("INFO: Found %d applications", len(applications))
	return applications
}
