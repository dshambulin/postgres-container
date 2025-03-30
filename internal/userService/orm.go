package userService

import (
	"get-post/internal/taskService"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string             `json:"email"`
	Password string             `json:"password"`
	Tasks    []taskService.Task `gorm:"foreignKey:UserID" json:"tasks,omitempty"`
}
