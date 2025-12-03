package service

import (
	"errors"
)

type TaskRepository interface {
	Create(string) (int, error)
}

type TaskService struct {
	repo TaskRepository
}

func New(repo TaskRepository) *TaskService {
	return &TaskService{
		repo: repo,
	}
}

func (ts TaskService) AddTask(desc string) (int, error) {
	ts.repo.Create(desc)
	return 0, errors.New("")
}
