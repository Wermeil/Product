package services

import (
	"Ctrl/internal/database"
	"Ctrl/internal/models"
	"context"
	"encoding/json"
	"fmt"
	"time"
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
	redisService *database.RedisClient
}

func NewUserService(userRepo database.UserRepository, tasksService TasksService, redisServ *database.RedisClient) *Repo {
	return &Repo{
		repo:         userRepo,
		tasksService: tasksService, // ← сохраняем зависимость
		redisService: redisServ,
	}
}

func (s *Repo) GetTasksForUser(userID uint) ([]models.Tasks, error) {
	tasksRedisId := fmt.Sprintf("users:task:%v", userID)

	val, err := s.redisService.Get(context.Background(), tasksRedisId)
	if err == nil {
		var task []models.Tasks
		if err := json.Unmarshal([]byte(val), &task); err != nil {
			return []models.Tasks{}, err
		}
	}
	vag, err := s.tasksService.GetTaskByUserId(userID)
	if err != nil {
		return []models.Tasks{}, err
	}
	if err := s.redisService.SetJSON(context.Background(), tasksRedisId, vag, 10*time.Minute); err != nil {
		return []models.Tasks{}, err
	}
	return vag, nil
}

func (s *Repo) GetUser() ([]models.Users, error) {
	redisId := "users:all"

	val, err := s.redisService.Get(context.Background(), redisId)
	if err == nil {
		var user []models.Users
		if err := json.Unmarshal([]byte(val), &user); err != nil {
			return nil, err
		}
		return user, nil
	}

	user, err := s.repo.GetAllUser()
	if err != nil {
		return []models.Users{}, err
	}
	if err := s.redisService.SetJSON(context.Background(), redisId, user, 10*time.Minute); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Repo) CreateUser(user models.Users) (models.Users, error) {
	result, err := s.repo.CreateUser(user)
	if err != nil {
		return models.Users{}, err
	}
	return result, nil
}

func (s *Repo) GetUserById(id string) (models.Users, error) {
	cachedKey := fmt.Sprintf("users:%s", id)

	val, err := s.redisService.Get(context.Background(), cachedKey)
	if err == nil {
		var user models.Users
		if err := json.Unmarshal([]byte(val), &user); err != nil {
			return models.Users{}, err
		}
		return user, nil
	}

	user, err := s.repo.GetUserById(id)
	if err != nil {
		return models.Users{}, err
	}

	userJson, _ := json.Marshal(user)
	if err := s.redisService.Set(context.Background(), cachedKey, userJson, 10*time.Minute); err != nil {
		return models.Users{}, err
	}
	return user, nil
}

func (s *Repo) ChangeUserById(id string, us models.Users) error {
	use, err := s.repo.GetUserById(id)
	if err != nil {
		return err
	}
	if err := s.redisService.Delete(context.Background(), fmt.Sprintf("users:%v", use.ID)); err != nil {
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
	if err := s.redisService.Delete(context.Background(), fmt.Sprintf("users:%v", id)); err != nil {
		return err
	}
	return s.repo.DeleteUserById(id)
}
