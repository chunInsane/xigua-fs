package models

import (
	"time"
	"gorm.io/gorm"
)

type Folder struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	UserID    uint           `gorm:"not null;index" json:"user_id"`
	Name      string         `gorm:"not null" json:"name"`
	ParentID  *uint          `gorm:"index" json:"parent_id,omitempty"`
	Path      string         `gorm:"not null" json:"path"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	
	User      User           `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Parent    *Folder        `gorm:"foreignKey:ParentID" json:"parent,omitempty"`
	Children  []Folder       `gorm:"foreignKey:ParentID" json:"children,omitempty"`
	Files     []FileInfo     `gorm:"foreignKey:FolderID" json:"files,omitempty"`
}

type FileInfo struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	UserID      uint           `gorm:"not null;index" json:"user_id"`
	FolderID    *uint          `gorm:"index" json:"folder_id,omitempty"`
	OriginalName string        `gorm:"not null" json:"original_name"`
	FileName    string         `gorm:"not null" json:"file_name"`
	FilePath    string         `gorm:"not null" json:"file_path"`
	FileSize    int64          `gorm:"not null" json:"file_size"`
	FileType    string         `gorm:"not null" json:"file_type"`
	MimeType    string         `json:"mime_type"`
	IsImage     bool           `gorm:"default:false" json:"is_image"`
	IsVideo     bool           `gorm:"default:false" json:"is_video"`
	Width       int            `json:"width,omitempty"`
	Height      int            `json:"height,omitempty"`
	Duration    float64        `json:"duration,omitempty"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	
	User        User           `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Folder      *Folder        `gorm:"foreignKey:FolderID" json:"folder,omitempty"`
}