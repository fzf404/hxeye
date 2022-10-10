package model

import (
	"gorm.io/gorm"
)

// User 用户
type User struct {
	gorm.Model
	Name     string
	Password string
	Super    bool
}

// UserDto 用户数据传输
type UserDto struct {
	UserID  uint   `json:"userid"`
	Name    string `json:"name"`
	Super   bool   `json:"Super"`
}
