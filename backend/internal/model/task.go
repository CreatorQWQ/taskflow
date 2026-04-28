package model

import (
	"time"
)

type Task struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Title       string    `gorm:"not null" json:"title"`
	Content     string    `json:"content"`
	Status      string    `gorm:"default:'pending'" json:"status"` // pending 或 completed
	UserID      uint      `gorm:"index" json:"user_id"`            // 外键，并建立索引
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}