package database

import (
	"Ctrl/internal/models"
	"gorm.io/gorm"
)

type TaskRepository interface {
	GetAllTask() ([]models.Tasks, error)
	GetTaskByUserId(userId uint) ([]models.Tasks, error)
	CreateTask(tasks models.Tasks) (models.Tasks, error)
	GetTaskById(id string) (models.Tasks, error)
	DeleteTask(id string) error
	SaveTask(tasks models.Tasks) error
}
type TaskDd struct {
	db *gorm.DB
}

func (r *TaskDd) GetTaskByUserId(userId uint) ([]models.Tasks, error) {
	var task []models.Tasks
	if err := r.db.Find(&task, "user_id = ?", userId).Error; err != nil {
		return nil, err
	}
	return task, nil
}

func (r *TaskDd) SaveTask(tasks models.Tasks) error {
	if err := r.db.Save(&tasks).Error; err != nil {
		return err
	}
	return nil
}

func (r *TaskDd) GetAllTask() ([]models.Tasks, error) {
	var arr []models.Tasks
	if err := r.db.Find(&arr).Error; err != nil {
		return nil, err
	}
	return arr, nil
}

func (r *TaskDd) CreateTask(tasks models.Tasks) (models.Tasks, error) {
	if err := r.db.Create(&tasks).Error; err != nil {
		return models.Tasks{}, err
	}
	return tasks, nil
}

func (r *TaskDd) GetTaskById(id string) (models.Tasks, error) {
	var task models.Tasks
	if err := r.db.First(&task, "id = ?", id).Error; err != nil {
		return models.Tasks{}, err
	}
	return task, nil
}

func (r *TaskDd) DeleteTask(id string) error {
	if err := r.db.Delete(models.Tasks{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &TaskDd{db: db}
}
