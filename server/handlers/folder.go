package handlers

import (
	"net/http"
	"strconv"
	"strings"
	"xigua-server/config"
	"xigua-server/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type FolderHandler struct {
	DB     *gorm.DB
	Config *config.Config
}

func NewFolderHandler(db *gorm.DB, cfg *config.Config) *FolderHandler {
	return &FolderHandler{
		DB:     db,
		Config: cfg,
	}
}

type CreateFolderRequest struct {
	Name     string `json:"name" binding:"required,min=1,max=255"`
	ParentID *uint  `json:"parent_id,omitempty"`
}

type UpdateFolderRequest struct {
	Name string `json:"name" binding:"required,min=1,max=255"`
}

type MoveFolderRequest struct {
	ParentID *uint `json:"parent_id"`
}

func (h *FolderHandler) Create(c *gin.Context) {
	userID, _ := c.Get("user_id")
	
	var req CreateFolderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 验证文件夹名称
	if strings.TrimSpace(req.Name) == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Folder name cannot be empty"})
		return
	}

	// 构建文件夹路径
	var path string
	if req.ParentID != nil {
		// 验证父文件夹是否存在且属于当前用户
		var parentFolder models.Folder
		if err := h.DB.Where("id = ? AND user_id = ?", *req.ParentID, userID).First(&parentFolder).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Parent folder not found"})
			return
		}
		path = parentFolder.Path + "/" + req.Name
	} else {
		path = "/" + req.Name
	}

	// 检查同一目录下是否已存在同名文件夹
	var existingFolder models.Folder
	query := h.DB.Where("user_id = ? AND name = ?", userID, req.Name)
	if req.ParentID != nil {
		query = query.Where("parent_id = ?", *req.ParentID)
	} else {
		query = query.Where("parent_id IS NULL")
	}
	
	if err := query.First(&existingFolder).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Folder with this name already exists"})
		return
	}

	folder := models.Folder{
		UserID:   userID.(uint),
		Name:     strings.TrimSpace(req.Name),
		ParentID: req.ParentID,
		Path:     path,
	}

	if err := h.DB.Create(&folder).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create folder"})
		return
	}

	// 预加载关联数据
	h.DB.Preload("Parent").First(&folder, folder.ID)

	c.JSON(http.StatusOK, gin.H{
		"message": "Folder created successfully",
		"folder":  folder,
	})
}

func (h *FolderHandler) List(c *gin.Context) {
	userID, _ := c.Get("user_id")
	parentIDStr := c.Query("parent_id")
	
	query := h.DB.Where("user_id = ?", userID)
	
	if parentIDStr == "" {
		// 获取根目录下的文件夹
		query = query.Where("parent_id IS NULL")
	} else {
		parentID, err := strconv.ParseUint(parentIDStr, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parent_id"})
			return
		}
		query = query.Where("parent_id = ?", parentID)
	}
	
	var folders []models.Folder
	if err := query.Order("name ASC").Find(&folders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch folders"})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"folders": folders})
}

func (h *FolderHandler) Update(c *gin.Context) {
	userID, _ := c.Get("user_id")
	folderID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid folder ID"})
		return
	}

	var req UpdateFolderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var folder models.Folder
	if err := h.DB.Where("id = ? AND user_id = ?", folderID, userID).First(&folder).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Folder not found"})
		return
	}

	// 检查同一目录下是否已存在同名文件夹
	var existingFolder models.Folder
	query := h.DB.Where("user_id = ? AND name = ? AND id != ?", userID, req.Name, folderID)
	if folder.ParentID != nil {
		query = query.Where("parent_id = ?", *folder.ParentID)
	} else {
		query = query.Where("parent_id IS NULL")
	}
	
	if err := query.First(&existingFolder).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Folder with this name already exists"})
		return
	}

	oldPath := folder.Path
	folder.Name = strings.TrimSpace(req.Name)
	
	// 更新路径
	if folder.ParentID != nil {
		var parentFolder models.Folder
		h.DB.First(&parentFolder, *folder.ParentID)
		folder.Path = parentFolder.Path + "/" + folder.Name
	} else {
		folder.Path = "/" + folder.Name
	}

	if err := h.DB.Save(&folder).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update folder"})
		return
	}

	// 更新所有子文件夹的路径
	h.updateChildrenPaths(oldPath, folder.Path, userID.(uint))

	c.JSON(http.StatusOK, gin.H{
		"message": "Folder updated successfully",
		"folder":  folder,
	})
}

func (h *FolderHandler) Delete(c *gin.Context) {
	userID, _ := c.Get("user_id")
	folderID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid folder ID"})
		return
	}

	var folder models.Folder
	if err := h.DB.Where("id = ? AND user_id = ?", folderID, userID).First(&folder).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Folder not found"})
		return
	}

	// 检查文件夹是否为空
	var childCount int64
	h.DB.Model(&models.Folder{}).Where("parent_id = ?", folderID).Count(&childCount)
	
	var fileCount int64
	h.DB.Model(&models.FileInfo{}).Where("folder_id = ?", folderID).Count(&fileCount)
	
	if childCount > 0 || fileCount > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot delete non-empty folder"})
		return
	}

	if err := h.DB.Delete(&folder).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete folder"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Folder deleted successfully"})
}

func (h *FolderHandler) Move(c *gin.Context) {
	userID, _ := c.Get("user_id")
	folderID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid folder ID"})
		return
	}

	var req MoveFolderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var folder models.Folder
	if err := h.DB.Where("id = ? AND user_id = ?", folderID, userID).First(&folder).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Folder not found"})
		return
	}

	// 验证目标父文件夹
	if req.ParentID != nil {
		var parentFolder models.Folder
		if err := h.DB.Where("id = ? AND user_id = ?", *req.ParentID, userID).First(&parentFolder).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Target parent folder not found"})
			return
		}

		// 防止循环引用
		if h.isDescendant(uint(*req.ParentID), uint(folderID)) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot move folder to its descendant"})
			return
		}
	}

	// 检查目标位置是否已存在同名文件夹
	var existingFolder models.Folder
	query := h.DB.Where("user_id = ? AND name = ? AND id != ?", userID, folder.Name, folderID)
	if req.ParentID != nil {
		query = query.Where("parent_id = ?", *req.ParentID)
	} else {
		query = query.Where("parent_id IS NULL")
	}
	
	if err := query.First(&existingFolder).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Folder with this name already exists in target location"})
		return
	}

	oldPath := folder.Path
	folder.ParentID = req.ParentID
	
	// 更新路径
	if req.ParentID != nil {
		var parentFolder models.Folder
		h.DB.First(&parentFolder, *req.ParentID)
		folder.Path = parentFolder.Path + "/" + folder.Name
	} else {
		folder.Path = "/" + folder.Name
	}

	if err := h.DB.Save(&folder).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to move folder"})
		return
	}

	// 更新所有子文件夹的路径
	h.updateChildrenPaths(oldPath, folder.Path, userID.(uint))

	c.JSON(http.StatusOK, gin.H{
		"message": "Folder moved successfully",
		"folder":  folder,
	})
}

func (h *FolderHandler) GetPath(c *gin.Context) {
	userID, _ := c.Get("user_id")
	folderID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid folder ID"})
		return
	}

	var folder models.Folder
	if err := h.DB.Where("id = ? AND user_id = ?", folderID, userID).First(&folder).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Folder not found"})
		return
	}

	breadcrumbs := h.getBreadcrumbs(uint(folderID), userID.(uint))

	c.JSON(http.StatusOK, gin.H{
		"path":        folder.Path,
		"breadcrumbs": breadcrumbs,
	})
}

// 辅助函数：更新子文件夹路径
func (h *FolderHandler) updateChildrenPaths(oldPath, newPath string, userID uint) {
	var children []models.Folder
	h.DB.Where("user_id = ? AND path LIKE ?", userID, oldPath+"/%").Find(&children)
	
	for _, child := range children {
		child.Path = strings.Replace(child.Path, oldPath, newPath, 1)
		h.DB.Save(&child)
	}
}

// 辅助函数：检查是否为后代文件夹
func (h *FolderHandler) isDescendant(targetID, folderID uint) bool {
	var folder models.Folder
	if err := h.DB.First(&folder, targetID).Error; err != nil {
		return false
	}
	
	if folder.ParentID == nil {
		return false
	}
	
	if *folder.ParentID == folderID {
		return true
	}
	
	return h.isDescendant(*folder.ParentID, folderID)
}

// 辅助函数：获取面包屑导航
func (h *FolderHandler) getBreadcrumbs(folderID, userID uint) []models.Folder {
	var breadcrumbs []models.Folder
	currentID := folderID
	
	for currentID != 0 {
		var folder models.Folder
		if err := h.DB.Where("id = ? AND user_id = ?", currentID, userID).First(&folder).Error; err != nil {
			break
		}
		
		breadcrumbs = append([]models.Folder{folder}, breadcrumbs...)
		
		if folder.ParentID == nil {
			break
		}
		currentID = *folder.ParentID
	}
	
	return breadcrumbs
}