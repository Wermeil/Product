package services

import (
	"Ctrl/internal/database"
	"Ctrl/internal/models"
)

type UserService interface {
	GetUser() ([]models.Users, error)
	CreateUser(user models.Users) (models.Users, error)
	GetUserById(id string) (models.Users, error)
	ChangeUserById(id string, user models.Users) error
	DeleteUserById(id string) error
	GetTasksForUser(userID uint) ([]models.Tasks, error)
}
type Repo struct {
	repo         database.UserRepository
	tasksService TasksService
}

func (s *Repo) GetTasksForUser(userID uint) ([]models.Tasks, error) {
	return s.tasksService.GetTaskByUserId(userID)
}

func NewUserService(userRepo database.UserRepository, tasksService TasksService) *Repo {
	return &Repo{
		repo:         userRepo,
		tasksService: tasksService, // ← сохраняем зависимость
	}
}

func (s *Repo) GetUser() ([]models.Users, error) {
	return s.repo.GetAllUser()
}

func (s *Repo) CreateUser(user models.Users) (models.Users, error) {
	return s.repo.CreateUser(user)
}

func (s *Repo) GetUserById(id string) (models.Users, error) {
	return s.repo.GetUserById(id)
}

func (s *Repo) ChangeUserById(id string, us models.Users) error {
	use, err := s.repo.GetUserById(id)
	if err != nil {
		return err
	}
	obj := models.Users{
		ID:       use.ID,
		Email:    us.Email,
		Password: us.Password,
	}
	if err := s.repo.PatchUser(obj); err != nil {
		return err
	}
	return nil
}

func (s *Repo) DeleteUserById(id string) error {
	return s.repo.DeleteUserById(id)
}
