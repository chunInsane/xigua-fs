package models

import (
	"time"
	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	Username  string         `gorm:"unique;not null" json:"username"`
	Password  string         `gorm:"not null" json:"-"`
	Email     string         `gorm:"unique" json:"email"`
	IsAdmin   bool           `gorm:"default:false" json:"is_admin"`
	StorageQuota int64       `gorm:"default:1073741824" json:"storage_quota"` // 1GB默认
	StorageUsed  int64       `gorm:"default:0" json:"storage_used"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}