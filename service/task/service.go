package task

import (
	"errors"
	"task_app/models"
	task_repo "task_app/repo/task"
)

type Service interface {
	Create(task models.Task) error
	GetByID(id string) (models.Task, error)
	Update(task models.Task) error
	DeleteByID(id string) error
	GetAll() ([]models.Task, error)
}

type service struct {
	repo task_repo.TaskRepo
}

func NewService(repo task_repo.TaskRepo) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) Create(task models.Task) error {
	if task.ID == "" {
		return errors.New("task id cannot be empty")
	}

	return s.repo.Save(task)
}

func (s *service) GetByID(id string) (models.Task, error) {
	if id == "" {
		return models.Task{}, errors.New("id cannot be empty")
	}

	return s.repo.GetByID(id)
}

func (s *service) Update(task models.Task) error {
	if task.ID == "" {
		return errors.New("task id cannot be empty")
	}

	return s.repo.Update(task)
}

func (s *service) DeleteByID(id string) error {
	if id == "" {
		return errors.New("id cannot be empty")
	}

	return s.repo.DeleteByID(id)
}

func (s *service) GetAll() ([]models.Task, error) {
	return s.repo.GetAll()
}
