package db

import (
	"os"
)

const (
	createFlag    = os.O_CREATE  // Flag to create file if not exists
	readWriteFlag = os.O_RDWR    // Flag to read and write
	fileMode      = 0644         // Permissions to read and write on the file
	path          = "tasks.json" // Storage/database path
)

func Connect() (*os.File, error) {
	db, err := os.OpenFile(path, readWriteFlag|createFlag, fileMode)
	if err != nil {
		return nil, err
	}
	return db, nil
}
