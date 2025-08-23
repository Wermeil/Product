package userService

import (
	t "Ctrl/internal/tasksService"
)

type UserService interface {
	GetUser() ([]Users, error)
	CreateUser(user Users) (Users, error)
	GetUserById(id string) (Users, error)
	ChangeUserById(id string, user Users) error
	DeleteUserById(id string) error
	GetTasksForUser(userID uint) ([]t.Tasks, error)
}
type Repo struct {
	repo         UserRepository
	tasksService t.TasksService
}

func (s *Repo) GetTasksForUser(userID uint) ([]t.Tasks, error) {
	return s.tasksService.GetTaskByUserId(userID)
}

func NewUserService(userRepo UserRepository, tasksService t.TasksService) *Repo {
	return &Repo{
		repo:         userRepo,
		tasksService: tasksService, // ← сохраняем зависимость
	}
}

func (s *Repo) GetUser() ([]Users, error) {
	return s.repo.GetAllUser()
}

func (s *Repo) CreateUser(user Users) (Users, error) {
	return s.repo.CreateUser(user)
}

func (s *Repo) GetUserById(id string) (Users, error) {
	return s.repo.GetUserById(id)
}

func (s *Repo) ChangeUserById(id string, us Users) error {
	use, err := s.repo.GetUserById(id)
	if err != nil {
		return err
	}
	obj := Users{
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
