package models

import (
	"time"
)

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Email     string    `json:"email" gorm:"unique;not null"`
	Password  string    `json:"-" gorm:"not null"`
	Tier      string    `json:"tier" gorm:"default:free"`
	CreatedAt time.Time `json:"created_at"`
}
