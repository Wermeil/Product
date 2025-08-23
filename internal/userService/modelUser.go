package userService

import "Ctrl/internal/tasksService"

type Users struct {
	ID       uint                 `gorm:"primaryKey;autoIncrement" json:"id"`
	Email    string               `json:"email"`
	Password string               `json:"password"`
	Tasks    []tasksService.Tasks `gorm:"foreignKey:UserId" json:"tasks"`
}
