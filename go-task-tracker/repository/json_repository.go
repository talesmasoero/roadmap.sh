package repository

import (
	"encoding/json"
	"os"
	"task-cli/model"
	"time"
)

type JSONRepository struct {
	file *os.File
}

func New(file *os.File) *JSONRepository {
	return &JSONRepository{
		file: file,
	}
}

func (jr *JSONRepository) Create(desc string) (int, error) {
	task := model.Task{
		//ID:          id,
		Description: desc,
		Status:      model.StatusTodo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	jsonData, err := json.MarshalIndent(task, "", "  ")
	if err != nil {
		return 0, err
	}

	if _, err := jr.file.Write(jsonData); err != nil {
		return 0, err
	}
	return 0, nil // Last/Inserted/CreatedID, nil
}

func (jr *JSONRepository) getLastID() int {
	return 0 // LastID
}
