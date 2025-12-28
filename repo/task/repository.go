package repository

import (
	"errors"
	"sync"
	"task_app/models"
)

type TaskRepo interface {
	Save(task models.Task) error
	GetByID(id string) (models.Task, error)
	Update(task models.Task) error
	DeleteByID(id string) error
	GetAll() ([]models.Task, error)
}

type taskRepo struct {
	mu      sync.RWMutex
	taskMap map[string]models.Task
}

func (r *taskRepo) Save(task models.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.taskMap[task.ID]; exists {
		return errors.New("task already exists")
	}

	r.taskMap[task.ID] = task
	return nil
}

func (r *taskRepo) GetByID(id string) (models.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	task, exists := r.taskMap[id]
	if !exists {
		return models.Task{}, errors.New("task not found")
	}

	return task, nil
}

func (r *taskRepo) Update(task models.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.taskMap[task.ID]; !exists {
		return errors.New("task not found")
	}

	r.taskMap[task.ID] = task
	return nil
}

func (r *taskRepo) DeleteByID(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.taskMap[id]; !exists {
		return errors.New("task not found")
	}

	delete(r.taskMap, id)
	return nil
}

func (r *taskRepo) GetAll() ([]models.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	tasks := make([]models.Task, 0, len(r.taskMap))
	for _, task := range r.taskMap {
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func NewTaskRepo() TaskRepo {
	return &taskRepo{
		taskMap: make(map[string]models.Task),
	}
}
