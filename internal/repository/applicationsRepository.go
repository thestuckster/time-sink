package repository

import (
	"database/sql"
	"errors"
	_ "modernc.org/sqlite"
	"time"
)

type ApplicationDto struct {
	Id       *int
	Name     string
	Seen     int64 // seconds of unix epoch
	Duration int64 // delta of Seen and now
}

func FindById(id int, db *sql.DB) *ApplicationDto {
	statement, err := db.Prepare("SELECT * FROM applications WHERE id=?")
	defer statement.Close()
	if err != nil {
		panic(err)
	}

	var record ApplicationDto
	err = statement.QueryRow(id).Scan(&record.Id, &record.Name, &record.Seen, &record.Duration)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		} else {
			panic(err)
		}
	}

	return &record
}

func FindByNameAndCurrentDay(name string, db *sql.DB) *ApplicationDto {

	var record ApplicationDto

	now := getToday()
	tomorrow := getTomorrowMidnight(now)

	err := db.QueryRow(`SELECT * FROM applications WHERE name = ? AND ? >= seen AND seen < ?`,
		name, now.UnixMilli(), tomorrow.UnixMilli()).
		Scan(&record.Id, &record.Name, &record.Seen, &record.Duration)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		} else {
			panic(err)
		}
	}

	return &record
}

func FindAllByCurrentDay(db *sql.DB) []ApplicationDto {

	records := make([]ApplicationDto, 0)

	now := getToday()
	tomorrow := getTomorrowMidnight(now)

	rows, err := db.Query(`SELECT * FROM applications WHERE seen >= ? AND seen < ?`, now.Unix(), tomorrow.Unix())
	defer rows.Close()
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var record ApplicationDto
		err = rows.Scan(&record.Id, &record.Name, &record.Seen, &record.Duration)
		if err != nil {
			panic(err)
		}

		records = append(records, record)
	}

	return records
}

func SaveNew(dto *ApplicationDto, db *sql.DB) {
	statement, err := db.Prepare(`INSERT INTO applications(name, seen, duration) VALUES (?, ?, ?)`)
	defer statement.Close()
	if err != nil {
		panic(err)
	}

	_, err = statement.Exec(dto.Name, dto.Seen, dto.Duration)
	if err != nil {
		panic(err)
	}
}

func UpdateDuration(id int, newDuration int64, db *sql.DB) {

	statement, err := db.Prepare("UPDATE applications SET duration = ? WHERE id = ?")
	defer statement.Close()
	if err != nil {
		panic(err)
	}

	_, err = statement.Exec(newDuration, id)
	if err != nil {
		panic(err)
	}

}

func getToday() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
}

func getTomorrowMidnight(now time.Time) time.Time {
	tomorrow := now.AddDate(0, 0, 1)
	return time.Date(tomorrow.Year(), tomorrow.Month(), tomorrow.Day(), 0, 0, 0, 0, tomorrow.Location())
}
