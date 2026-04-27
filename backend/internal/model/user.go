package model

import (
	"time"
	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Username  string         `gorm:"unique;not null;index" json:"username"` // 建立索引提高查询效率
	Password  string         `gorm:"not null" json:"-"`                     // 永远不返回密码给前端
	Email     string         `gorm:"unique;not null" json:"email"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // 软删除
}