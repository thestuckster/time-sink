package repository

import (
	"errors"
	"gorm.io/gorm"
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

	autoMigrate(db)

	result := db.Create(&application)
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
	result := db.Where("name = ? AND seen >= start AND seen <= end", name, startUnix, endUnix).Find(&application)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil
		}
	}
	return &application
}

func GetAllApplicationsByDates(start, end time.Time, db *gorm.DB) []Application {
	var applications []Application

	startUnix := helpers.GetStartOfDayUnixEpoch(start)
	endUnix := helpers.GetStartOfDayUnixEpoch(end)

	result := db.Where("seen >= ? AND seen <= ?", startUnix, endUnix).Find(&applications)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return make([]Application, 0)
		}

	}
	return applications
}
