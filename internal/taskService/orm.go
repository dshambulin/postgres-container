package taskService

import "gorm.io/gorm"

type Task struct {
	gorm.Model

	ID     int64  `gorm:"primaryKey"`
	Task   string `json:"task"`
	IsDone bool   `json:"is_done"`
}
