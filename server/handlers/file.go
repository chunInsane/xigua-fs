package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"xigua-server/config"
	"xigua-server/models"
	"xigua-server/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type FileHandler struct {
	DB     *gorm.DB
	Config *config.Config
}

func NewFileHandler(db *gorm.DB, cfg *config.Config) *FileHandler {
	return &FileHandler{
		DB:     db,
		Config: cfg,
	}
}

func (h *FileHandler) Upload(c *gin.Context) {
	userID, _ := c.Get("user_id")
	folderIDStr := c.PostForm("folder_id")
	
	var user models.User
	if err := h.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
		return
	}

	var folderID *uint
	if folderIDStr != "" {
		parsedFolderID, err := strconv.ParseUint(folderIDStr, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid folder_id"})
			return
		}
		
		// 验证文件夹是否存在且属于当前用户
		var folder models.Folder
		if err := h.DB.Where("id = ? AND user_id = ?", parsedFolderID, userID).First(&folder).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Folder not found"})
			return
		}
		
		id := uint(parsedFolderID)
		folderID = &id
	}

	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file provided"})
		return
	}
	defer file.Close()

	if header.Size > h.Config.MaxFileSize {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("File too large. Max size: %d bytes", h.Config.MaxFileSize),
		})
		return
	}

	if user.StorageUsed+header.Size > user.StorageQuota {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Storage quota exceeded"})
		return
	}

	userDir := filepath.Join(h.Config.StoragePath, fmt.Sprintf("user_%d", userID))
	if err := utils.EnsureDir(userDir); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user directory"})
		return
	}

	fileName := utils.GenerateFileName(header.Filename)
	filePath := filepath.Join(userDir, fileName)

	dst, err := os.Create(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create file"})
		return
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		os.Remove(filePath)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	fileInfo := models.FileInfo{
		UserID:       userID.(uint),
		FolderID:     folderID,
		OriginalName: header.Filename,
		FileName:     fileName,
		FilePath:     filePath,
		FileSize:     header.Size,
		FileType:     utils.GetFileType(header.Filename),
		MimeType:     utils.GetMimeType(header.Filename),
		IsImage:      utils.IsImageFile(header.Filename),
		IsVideo:      utils.IsVideoFile(header.Filename),
	}

	if err := h.DB.Create(&fileInfo).Error; err != nil {
		os.Remove(filePath)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file info"})
		return
	}

	user.StorageUsed += header.Size
	h.DB.Save(&user)

	// 预加载文件夹信息
	h.DB.Preload("Folder").First(&fileInfo, fileInfo.ID)

	c.JSON(http.StatusOK, gin.H{
		"message": "File uploaded successfully",
		"file":    fileInfo,
	})
}

func (h *FileHandler) Download(c *gin.Context) {
	userID, _ := c.Get("user_id")
	fileID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file ID"})
		return
	}

	var fileInfo models.FileInfo
	query := h.DB.Where("id = ?", fileID)
	
	isAdmin, _ := c.Get("is_admin")
	if !isAdmin.(bool) {
		query = query.Where("user_id = ?", userID)
	}
	
	if err := query.First(&fileInfo).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}

	if _, err := os.Stat(fileInfo.FilePath); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found on disk"})
		return
	}

	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", fileInfo.OriginalName))
	c.Header("Content-Type", fileInfo.MimeType)
	c.File(fileInfo.FilePath)
}

func (h *FileHandler) Preview(c *gin.Context) {
	userID, _ := c.Get("user_id")
	fileID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file ID"})
		return
	}

	var fileInfo models.FileInfo
	query := h.DB.Where("id = ?", fileID)
	
	isAdmin, _ := c.Get("is_admin")
	if !isAdmin.(bool) {
		query = query.Where("user_id = ?", userID)
	}
	
	if err := query.First(&fileInfo).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}

	if !fileInfo.IsImage && !fileInfo.IsVideo {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File type not supported for preview"})
		return
	}

	if _, err := os.Stat(fileInfo.FilePath); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found on disk"})
		return
	}

	// 设置缓存头以优化加载
	c.Header("Cache-Control", "public, max-age=3600")
	c.Header("Content-Type", fileInfo.MimeType)
	
	// 支持跨域请求
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET")
	c.Header("Access-Control-Allow-Headers", "Authorization")
	
	c.File(fileInfo.FilePath)
}

func (h *FileHandler) List(c *gin.Context) {
	userID, _ := c.Get("user_id")
	
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	fileType := c.Query("type")
	folderIDStr := c.Query("folder_id")
	
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 20
	}
	
	offset := (page - 1) * limit
	
	query := h.DB.Where("user_id = ?", userID)
	
	// 文件夹过滤
	if folderIDStr == "" {
		// 获取根目录下的文件
		query = query.Where("folder_id IS NULL")
	} else {
		folderID, err := strconv.ParseUint(folderIDStr, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid folder_id"})
			return
		}
		query = query.Where("folder_id = ?", folderID)
	}
	
	if fileType != "" {
		query = query.Where("file_type = ?", fileType)
	}
	
	var total int64
	query.Model(&models.FileInfo{}).Count(&total)
	
	var files []models.FileInfo
	if err := query.Preload("Folder").Order("created_at DESC").Limit(limit).Offset(offset).Find(&files).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch files"})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"files": files,
		"total": total,
		"page":  page,
		"limit": limit,
	})
}

func (h *FileHandler) Delete(c *gin.Context) {
	userID, _ := c.Get("user_id")
	fileID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file ID"})
		return
	}

	var fileInfo models.FileInfo
	query := h.DB.Where("id = ? AND user_id = ?", fileID, userID)
	
	if err := query.First(&fileInfo).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}

	if err := os.Remove(fileInfo.FilePath); err != nil && !os.IsNotExist(err) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete file from disk"})
		return
	}

	if err := h.DB.Delete(&fileInfo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete file record"})
		return
	}

	var user models.User
	if err := h.DB.First(&user, userID).Error; err == nil {
		user.StorageUsed -= fileInfo.FileSize
		if user.StorageUsed < 0 {
			user.StorageUsed = 0
		}
		h.DB.Save(&user)
	}

	c.JSON(http.StatusOK, gin.H{"message": "File deleted successfully"})
}

type MoveFileRequest struct {
	FolderID *uint `json:"folder_id"`
}

func (h *FileHandler) Move(c *gin.Context) {
	userID, _ := c.Get("user_id")
	fileID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file ID"})
		return
	}

	var req MoveFileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var fileInfo models.FileInfo
	if err := h.DB.Where("id = ? AND user_id = ?", fileID, userID).First(&fileInfo).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}

	// 验证目标文件夹
	if req.FolderID != nil {
		var folder models.Folder
		if err := h.DB.Where("id = ? AND user_id = ?", *req.FolderID, userID).First(&folder).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Target folder not found"})
			return
		}
	}

	fileInfo.FolderID = req.FolderID

	if err := h.DB.Save(&fileInfo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to move file"})
		return
	}

	// 预加载文件夹信息
	h.DB.Preload("Folder").First(&fileInfo, fileInfo.ID)

	c.JSON(http.StatusOK, gin.H{
		"message": "File moved successfully",
		"file":    fileInfo,
	})
}