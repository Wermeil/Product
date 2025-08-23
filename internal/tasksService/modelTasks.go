package tasksService

type Tasks struct {
	ID       uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	TaskName string `json:"task_name"`
	IsDone   bool   `json:"is_done"`
	UserId   uint   `gorm:"column:user_id" json:"user_id"`
}
