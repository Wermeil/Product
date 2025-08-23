package userService

import "gorm.io/gorm"

type UserRepository interface {
	GetAllUser() ([]Users, error)
	CreateUser(user Users) (Users, error)
	GetUserById(id string) (Users, error)
	DeleteUserById(id string) error
	PatchUser(user Users) error
}
type UserDb struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserDb{db: db}
}

func (r *UserDb) PatchUser(user Users) error {
	if err := r.db.Save(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserDb) GetAllUser() ([]Users, error) {
	var arr []Users
	err := r.db.Find(&arr).Error
	if err != nil {
		return nil, err
	}
	return arr, nil
}
func (r *UserDb) CreateUser(user Users) (Users, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return Users{}, nil
	}
	return user, nil
}

func (r *UserDb) GetUserById(id string) (Users, error) {
	var arr Users
	if err := r.db.First(&arr, "id = ?", id).Error; err != nil {
		return Users{}, err
	}
	return arr, nil
}

func (r *UserDb) DeleteUserById(id string) error {
	if err := r.db.Delete(Users{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
