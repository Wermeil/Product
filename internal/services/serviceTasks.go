package services

import (
	"Ctrl/internal/database"
	"Ctrl/internal/models"
)

type TasksService interface {
	GetAllTask() ([]models.Tasks, error)
	CreateTask(tasks models.Tasks) (models.Tasks, error)
	GetTaskById(id string) (models.Tasks, error)
	DeleteTask(id string) error
	ChangeTask(id string, tasks models.Tasks) (models.Tasks, error)
	GetTaskByUserId(userId uint) ([]models.Tasks, error)
}

type TaskRepo struct {
	repo database.TaskRepository
}

func (s *TaskRepo) GetTaskByUserId(userId uint) ([]models.Tasks, error) {
	return s.repo.GetTaskByUserId(userId)
}

func (s *TaskRepo) GetAllTask() ([]models.Tasks, error) {
	return s.repo.GetAllTask()
}

func (s *TaskRepo) CreateTask(tasks models.Tasks) (models.Tasks, error) {
	return s.repo.CreateTask(tasks)
}

func (s *TaskRepo) GetTaskById(id string) (models.Tasks, error) {
	return s.repo.GetTaskById(id)
}

func (s *TaskRepo) DeleteTask(id string) error {
	return s.repo.DeleteTask(id)
}

func (s *TaskRepo) ChangeTask(id string, tasks models.Tasks) (models.Tasks, error) {
	task, err := s.repo.GetTaskById(id)
	if err != nil {
		return models.Tasks{}, err
	}
	response := models.Tasks{
		ID:       task.ID,
		TaskName: tasks.TaskName,
		IsDone:   tasks.IsDone,
		UserId:   task.UserId,
	}
	if err := s.repo.SaveTask(response); err != nil {
		return models.Tasks{}, err
	}
	return response, nil
}

func NewTaskService(r database.TaskRepository) TasksService {
	return &TaskRepo{repo: r}
}
