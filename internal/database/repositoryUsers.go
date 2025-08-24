package database

import (
	"Ctrl/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetAllUser() ([]models.Users, error)
	CreateUser(user models.Users) (models.Users, error)
	GetUserById(id string) (models.Users, error)
	DeleteUserById(id string) error
	PatchUser(user models.Users) error
}
type UserDb struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserDb{db: db}
}

func (r *UserDb) PatchUser(user models.Users) error {
	if err := r.db.Save(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserDb) GetAllUser() ([]models.Users, error) {
	var arr []models.Users
	err := r.db.Find(&arr).Error
	if err != nil {
		return nil, err
	}
	return arr, nil
}
func (r *UserDb) CreateUser(user models.Users) (models.Users, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return models.Users{}, nil
	}
	return user, nil
}

func (r *UserDb) GetUserById(id string) (models.Users, error) {
	var arr models.Users
	if err := r.db.First(&arr, "id = ?", id).Error; err != nil {
		return models.Users{}, err
	}
	return arr, nil
}

func (r *UserDb) DeleteUserById(id string) error {
	if err := r.db.Delete(models.Users{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
