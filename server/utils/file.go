package utils

import (
	"crypto/md5"
	"fmt"
	"mime"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func GenerateFileName(originalName string) string {
	ext := filepath.Ext(originalName)
	hash := md5.Sum([]byte(fmt.Sprintf("%s_%d", originalName, time.Now().UnixNano())))
	return fmt.Sprintf("%x%s", hash, ext)
}

func GetFileType(filename string) string {
	ext := strings.ToLower(filepath.Ext(filename))
	switch ext {
	case ".jpg", ".jpeg", ".png", ".gif", ".bmp", ".webp":
		return "image"
	case ".mp4", ".avi", ".mov", ".wmv", ".flv", ".webm", ".mkv":
		return "video"
	case ".mp3", ".wav", ".flac", ".aac", ".ogg":
		return "audio"
	case ".pdf":
		return "pdf"
	case ".txt", ".md":
		return "text"
	default:
		return "other"
	}
}

func IsImageFile(filename string) bool {
	return GetFileType(filename) == "image"
}

func IsVideoFile(filename string) bool {
	return GetFileType(filename) == "video"
}

func GetMimeType(filename string) string {
	ext := filepath.Ext(filename)
	return mime.TypeByExtension(ext)
}

func EnsureDir(dir string) error {
	return os.MkdirAll(dir, 0755)
}

func GetFileSize(filepath string) (int64, error) {
	fileInfo, err := os.Stat(filepath)
	if err != nil {
		return 0, err
	}
	return fileInfo.Size(), nil
}