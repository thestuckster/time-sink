package internal

import (
	"errors"
	"os"
)

// on windows puts in C:/<users>/<user>/time-sink.db
func CreateDbFile() {
	if !dbFileExists() {
		path := *GetDbFilePath()
		file, err := os.Create(path)
		if err != nil {
			panic(err)
		}

		defer file.Close()
	}
}

func GetDbFilePath() *string {
	path, err := os.UserHomeDir()
	path += "/time-sink.db"
	if err != nil {
		panic(err)
	}

	return &path
}

func dbFileExists() bool {

	path := *GetDbFilePath()
	_, err := os.Stat(path)
	if errors.Is(err, os.ErrNotExist) {
		return false
	} else if err != nil {
		panic(err)
	}

	return true
}
