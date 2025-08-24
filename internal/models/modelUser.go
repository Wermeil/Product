package models

type Users struct {
	ID       uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
	Tasks    []Tasks `gorm:"foreignKey:UserId" json:"tasks"`
}
