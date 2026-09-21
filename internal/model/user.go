package model

import "time"

type User struct {
	ID       int64     `gorm:"primaryKey" json:"id"`
	Username string    `gorm:"column:username;not null" json:"username"`
	Email    string    `gorm:"unique;not null" json:"email"`
	CreateAt time.Time `json:"create_at"`
}
type CreateUserRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}
