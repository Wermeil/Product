package tasksService

import "gorm.io/gorm"

type TaskRepository interface {
	GetAllTask() ([]Tasks, error)
	GetTaskByUserId(userId uint) ([]Tasks, error)
	CreateTask(tasks Tasks) (Tasks, error)
	GetTaskById(id string) (Tasks, error)
	DeleteTask(id string) error
	SaveTask(tasks Tasks) error
}
type TaskDd struct {
	db *gorm.DB
}

func (r *TaskDd) GetTaskByUserId(userId uint) ([]Tasks, error) {
	var task []Tasks
	if err := r.db.Find(&task, "user_id = ?", userId).Error; err != nil {
		return nil, err
	}
	return task, nil
}

func (r *TaskDd) SaveTask(tasks Tasks) error {
	if err := r.db.Save(&tasks).Error; err != nil {
		return err
	}
	return nil
}

func (r *TaskDd) GetAllTask() ([]Tasks, error) {
	var arr []Tasks
	if err := r.db.Find(&arr).Error; err != nil {
		return nil, err
	}
	return arr, nil
}

func (r *TaskDd) CreateTask(tasks Tasks) (Tasks, error) {
	if err := r.db.Create(&tasks).Error; err != nil {
		return Tasks{}, err
	}
	return tasks, nil
}

func (r *TaskDd) GetTaskById(id string) (Tasks, error) {
	var task Tasks
	if err := r.db.First(&task, "id = ?", id).Error; err != nil {
		return Tasks{}, err
	}
	return task, nil
}

func (r *TaskDd) DeleteTask(id string) error {
	if err := r.db.Delete(Tasks{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &TaskDd{db: db}
}
