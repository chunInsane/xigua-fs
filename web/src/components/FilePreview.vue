<template>
  <Teleport to="body">
    <div 
      v-if="visible" 
      class="apple-preview-overlay"
      @click="handleOverlayClick"
    >
      <div class="apple-preview-modal" @click.stop>
        <!-- Header -->
        <div class="preview-header">
          <div class="header-left">
            <div class="file-title">
              <div class="file-icon">
                <svg v-if="file?.is_image" width="24" height="24" viewBox="0 0 24 24" fill="none">
                  <rect x="3" y="3" width="18" height="18" rx="2" stroke="var(--apple-green)" stroke-width="2"/>
                  <circle cx="8.5" cy="8.5" r="1.5" stroke="var(--apple-green)" stroke-width="2"/>
                  <path d="M21 15L16 10L5 21" stroke="var(--apple-green)" stroke-width="2"/>
                </svg>
                <svg v-else-if="file?.is_video" width="24" height="24" viewBox="0 0 24 24" fill="none">
                  <rect x="2" y="3" width="20" height="14" rx="2" stroke="var(--apple-blue)" stroke-width="2"/>
                  <path d="M10 8L16 12L10 16V8Z" fill="var(--apple-blue)"/>
                </svg>
                <svg v-else width="24" height="24" viewBox="0 0 24 24" fill="none">
                  <path d="M14 2H6C4.9 2 4 2.9 4 4V20C4 21.1 4.9 22 6 22H18C19.1 22 20 21.1 20 20V8L14 2Z" stroke="var(--apple-gray)" stroke-width="2"/>
                  <path d="M14 2V8H20" stroke="var(--apple-gray)" stroke-width="2"/>
                </svg>
              </div>
              <h3 class="apple-text-subtitle">{{ file?.original_name || '文件预览' }}</h3>
            </div>
          </div>
          <div class="header-actions">
            <button class="action-btn" @click="downloadFile" title="下载">
              <svg width="20" height="20" viewBox="0 0 24 24" fill="none">
                <path d="M21 15V19C21 20.1 20.1 21 19 21H5C3.9 21 3 20.1 3 19V15" stroke="currentColor" stroke-width="2"/>
                <path d="M7 10L12 15L17 10" stroke="currentColor" stroke-width="2"/>
                <path d="M12 15V3" stroke="currentColor" stroke-width="2"/>
              </svg>
            </button>
            <button class="action-btn close-btn" @click="$emit('update:visible', false)" title="关闭">
              <svg width="20" height="20" viewBox="0 0 24 24" fill="none">
                <path d="M18 6L6 18M6 6L18 18" stroke="currentColor" stroke-width="2" stroke-linecap="round"/>
              </svg>
            </button>
          </div>
        </div>

        <!-- Content -->
        <div class="preview-content">
          <div v-if="file" class="preview-main">
            <!-- Image Preview -->
            <div v-if="file.is_image" class="media-preview image-preview">
              <div class="media-container">
                <img
                  :src="getPreviewUrl(file?.id)"
                  :alt="file.original_name"
                  class="preview-image"
                  @error="handleImageError"
                  @load="handleImageLoad"
                />
                <div v-if="imageLoading" class="loading-overlay">
                  <div class="loading-spinner"></div>
                </div>
              </div>
            </div>
            
            <!-- Video Preview -->
            <div v-else-if="file.is_video" class="media-preview video-preview">
              <div class="media-container">
                <video
                  :src="getPreviewUrl(file?.id)"
                  controls
                  class="preview-video"
                  playsinline
                >
                  您的浏览器不支持视频播放
                </video>
              </div>
            </div>
            
            <!-- Unsupported Preview -->
            <div v-else class="unsupported-preview">
              <div class="unsupported-content">
                <div class="unsupported-icon">
                  <svg width="64" height="64" viewBox="0 0 24 24" fill="none">
                    <path d="M14 2H6C4.9 2 4 2.9 4 4V20C4 21.1 4.9 22 6 22H18C19.1 22 20 21.1 20 20V8L14 2Z" stroke="var(--apple-gray-5)" stroke-width="1.5"/>
                    <path d="M14 2V8H20" stroke="var(--apple-gray-5)" stroke-width="1.5"/>
                    <path d="M16 13H8" stroke="var(--apple-gray-5)" stroke-width="1.5"/>
                    <path d="M16 17H8" stroke="var(--apple-gray-5)" stroke-width="1.5"/>
                    <path d="M10 9H8" stroke="var(--apple-gray-5)" stroke-width="1.5"/>
                  </svg>
                </div>
                <h4 class="apple-text-subtitle" style="color: var(--apple-text-secondary); margin: var(--apple-space) 0 var(--apple-space-xs);">
                  无法预览此文件
                </h4>
                <p class="apple-text-body" style="color: var(--apple-text-tertiary); margin-bottom: var(--apple-space-lg);">
                  此文件类型不支持在线预览
                </p>
                <button class="apple-button-primary" @click="downloadFile">
                  下载查看
                </button>
              </div>
            </div>
          </div>

          <!-- File Info Sidebar -->
          <div class="file-info-sidebar">
            <div class="info-section">
              <h4 class="apple-text-subtitle" style="margin-bottom: var(--apple-space);">文件信息</h4>
              <div class="info-grid">
                <div class="info-item">
                  <label class="info-label apple-text-caption">文件名</label>
                  <div class="info-value apple-text-body">{{ file?.original_name }}</div>
                </div>
                <div class="info-item">
                  <label class="info-label apple-text-caption">文件类型</label>
                  <div class="info-value">
                    <span class="file-type-tag" :class="getFileTypeClass(file?.file_type)">
                      {{ file?.file_type?.toUpperCase() }}
                    </span>
                  </div>
                </div>
                <div class="info-item">
                  <label class="info-label apple-text-caption">文件大小</label>
                  <div class="info-value apple-text-body">{{ formatFileSize(file?.file_size) }}</div>
                </div>
                <div class="info-item">
                  <label class="info-label apple-text-caption">上传时间</label>
                  <div class="info-value apple-text-body">{{ formatDate(file?.created_at) }}</div>
                </div>
                <div v-if="file?.width && file?.height" class="info-item">
                  <label class="info-label apple-text-caption">分辨率</label>
                  <div class="info-value apple-text-body">{{ file.width }} × {{ file.height }}</div>
                </div>
                <div v-if="file?.duration" class="info-item">
                  <label class="info-label apple-text-caption">时长</label>
                  <div class="info-value apple-text-body">{{ formatDuration(file.duration) }}</div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </Teleport>
</template>

<script>
import { ref, computed, onMounted, onUnmounted } from 'vue'
import { useAuthStore } from '@/stores/auth'
import { ElMessage } from 'element-plus'
import { Download } from '@element-plus/icons-vue'

export default {
  name: 'FilePreview',
  props: {
    visible: {
      type: Boolean,
      default: false
    },
    file: {
      type: Object,
      default: null
    }
  },
  emits: ['update:visible'],
  setup(props) {
    const authStore = useAuthStore()
    const previewUrls = ref({})
    const windowWidth = ref(window.innerWidth)
    const imageLoading = ref(true)
    
    const isMobile = computed(() => windowWidth.value <= 768)
    
    const handleResize = () => {
      windowWidth.value = window.innerWidth
    }
    
    const handleOverlayClick = () => {
      // 点击遮罩层关闭预览
      if (props.file) {
        document.dispatchEvent(new Event('update:visible', false))
      }
    }
    
    const handleImageLoad = () => {
      imageLoading.value = false
    }
    
    const getFileTypeClass = (type) => {
      const classes = {
        image: 'file-type-image',
        video: 'file-type-video', 
        audio: 'file-type-audio',
        pdf: 'file-type-pdf',
        other: 'file-type-other'
      }
      return classes[type] || 'file-type-other'
    }
    
    onMounted(() => {
      window.addEventListener('resize', handleResize)
      // 监听ESC键关闭预览
      const handleEsc = (e) => {
        if (e.key === 'Escape' && props.visible) {
          document.dispatchEvent(new Event('update:visible', false))
        }
      }
      window.addEventListener('keydown', handleEsc)
      
      return () => {
        window.removeEventListener('keydown', handleEsc)
      }
    })
    
    onUnmounted(() => {
      window.removeEventListener('resize', handleResize)
    })
    
    const downloadFile = async () => {
      if (!props.file) return
      
      try {
        const response = await fetch(`/api/files/${props.file.id}/download`, {
          headers: {
            Authorization: `Bearer ${authStore.token}`
          }
        })
        
        if (!response.ok) throw new Error('Download failed')
        
        const blob = await response.blob()
        const url = window.URL.createObjectURL(blob)
        const a = document.createElement('a')
        a.style.display = 'none'
        a.href = url
        a.download = props.file.original_name
        document.body.appendChild(a)
        a.click()
        window.URL.revokeObjectURL(url)
        document.body.removeChild(a)
      } catch (error) {
        console.error('Download error:', error)
        ElMessage.error('下载失败')
      }
    }
    
    const formatFileSize = (bytes) => {
      if (bytes === 0) return '0 B'
      const k = 1024
      const sizes = ['B', 'KB', 'MB', 'GB']
      const i = Math.floor(Math.log(bytes) / Math.log(k))
      return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i]
    }
    
    const formatDate = (dateString) => {
      return new Date(dateString).toLocaleString('zh-CN')
    }
    
    const formatDuration = (seconds) => {
      const hours = Math.floor(seconds / 3600)
      const minutes = Math.floor((seconds % 3600) / 60)
      const secs = Math.floor(seconds % 60)
      
      if (hours > 0) {
        return `${hours}:${minutes.toString().padStart(2, '0')}:${secs.toString().padStart(2, '0')}`
      } else {
        return `${minutes}:${secs.toString().padStart(2, '0')}`
      }
    }
    
    const getFileTypeColor = (type) => {
      const colors = {
        image: 'success',
        video: 'primary',
        audio: 'warning',
        pdf: 'info',
        other: ''
      }
      return colors[type] || ''
    }
    
    const getPreviewUrl = (fileId) => {
      if (!fileId) return ''
      
      if (previewUrls.value[fileId]) {
        if (previewUrls.value[fileId] === 'error') {
          return 'data:image/svg+xml;base64,PHN2ZyB3aWR0aD0iMjAwIiBoZWlnaHQ9IjEyMCIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIj48cmVjdCB3aWR0aD0iMTAwJSIgaGVpZ2h0PSIxMDAlIiBmaWxsPSIjZmVmMGYwIi8+PHRleHQgeD0iNTAlIiB5PSI1MCUiIGZpbGw9IiNmNTZjNmMiIHRleHQtYW5jaG9yPSJtaWRkbGUiIGR5PSIuM2VtIj7liqDovb3lpLHotKU8L3RleHQ+PC9zdmc+'
        }
        return previewUrls.value[fileId]
      }
      
      // 创建带认证的预览URL
      fetchPreviewBlob(fileId)
      return 'data:image/svg+xml;base64,PHN2ZyB3aWR0aD0iMjAwIiBoZWlnaHQ9IjEyMCIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIj48cmVjdCB3aWR0aD0iMTAwJSIgaGVpZ2h0PSIxMDAlIiBmaWxsPSIjZjVmN2ZhIi8+PHRleHQgeD0iNTAlIiB5PSI1MCUiIGZpbGw9IiM5MDkzOTkiIHRleHQtYW5jaG9yPSJtaWRkbGUiIGR5PSIuM2VtIj7liqDovb3kuK0uLi48L3RleHQ+PC9zdmc+'
    }
    
    const fetchPreviewBlob = async (fileId) => {
      try {
        const response = await fetch(`/api/files/${fileId}/preview`, {
          headers: {
            Authorization: `Bearer ${authStore.token}`
          }
        })
        
        if (response.ok) {
          const blob = await response.blob()
          const url = URL.createObjectURL(blob)
          previewUrls.value[fileId] = url
          
          // 清理旧的URL
          setTimeout(() => {
            URL.revokeObjectURL(url)
            delete previewUrls.value[fileId]
          }, 60000) // 1分钟后清理
        } else {
          console.error(`Preview failed with status: ${response.status}`)
          previewUrls.value[fileId] = 'error'
        }
      } catch (error) {
        console.error('Failed to fetch preview:', error)
        previewUrls.value[fileId] = 'error'
      }
    }
    
    const handleImageError = (event) => {
      console.error('Image load error:', event)
      event.target.style.display = 'none'
    }
    
    return {
      isMobile,
      imageLoading,
      handleOverlayClick,
      handleImageLoad,
      getFileTypeClass,
      downloadFile,
      formatFileSize,
      formatDate,
      formatDuration,
      getFileTypeColor,
      getPreviewUrl,
      handleImageError,
      Download
    }
  }
}
</script>

<style scoped>
/* Apple Preview Modal */
.apple-preview-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.75);
  backdrop-filter: blur(20px);
  -webkit-backdrop-filter: blur(20px);
  z-index: 9999;
  display: flex;
  align-items: center;
  justify-content: center;
  animation: fadeIn 0.3s ease-out;
}

@keyframes fadeIn {
  from { opacity: 0; }
  to { opacity: 1; }
}

.apple-preview-modal {
  width: 90vw;
  max-width: 1200px;
  height: 85vh;
  max-height: 800px;
  background: var(--apple-background-elevated);
  border-radius: var(--apple-radius-lg);
  box-shadow: var(--apple-shadow-xl);
  border: 1px solid var(--apple-border);
  overflow: hidden;
  display: flex;
  flex-direction: column;
  animation: slideInUp 0.4s cubic-bezier(0.16, 1, 0.3, 1);
}

@keyframes slideInUp {
  from {
    opacity: 0;
    transform: translateY(60px) scale(0.95);
  }
  to {
    opacity: 1;
    transform: translateY(0) scale(1);
  }
}

/* Header */
.preview-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: var(--apple-space) var(--apple-space-lg);
  background: var(--apple-gray-1);
  border-bottom: 1px solid var(--apple-separator);
  min-height: 60px;
}

.header-left {
  display: flex;
  align-items: center;
  gap: var(--apple-space);
}

.close-btn {
  background: var(--apple-red) !important;
  color: white !important;
}

.close-btn:hover {
  background: #e6332a !important;
}

.file-title {
  display: flex;
  align-items: center;
  gap: var(--apple-space-sm);
}

.file-icon {
  display: flex;
  align-items: center;
  justify-content: center;
}

.header-actions {
  display: flex;
  gap: var(--apple-space-sm);
}

.action-btn {
  width: 36px;
  height: 36px;
  border-radius: var(--apple-radius-sm);
  border: none;
  background: var(--apple-gray-2);
  color: var(--apple-text-secondary);
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all 0.2s ease;
}

.action-btn:hover {
  background: var(--apple-gray-3);
  color: var(--apple-text-primary);
  transform: translateY(-1px);
}

/* Content */
.preview-content {
  flex: 1;
  display: flex;
  overflow: hidden;
}

.preview-main {
  flex: 1;
  display: flex;
  align-items: center;
  justify-content: center;
  background: var(--apple-gray-1);
  position: relative;
}

/* Media Preview */
.media-preview {
  width: 100%;
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: var(--apple-space-lg);
}

.media-container {
  position: relative;
  max-width: 100%;
  max-height: 100%;
  border-radius: var(--apple-radius);
  overflow: hidden;
  box-shadow: var(--apple-shadow-lg);
}

.preview-image {
  max-width: 100%;
  max-height: 100%;
  object-fit: contain;
  display: block;
}

.preview-video {
  max-width: 100%;
  max-height: 100%;
  border-radius: var(--apple-radius);
}

.loading-overlay {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: var(--apple-gray-1);
  display: flex;
  align-items: center;
  justify-content: center;
}

.loading-spinner {
  width: 32px;
  height: 32px;
  border: 3px solid var(--apple-gray-3);
  border-top: 3px solid var(--apple-blue);
  border-radius: 50%;
  animation: spin 1s linear infinite;
}

@keyframes spin {
  0% { transform: rotate(0deg); }
  100% { transform: rotate(360deg); }
}

/* Unsupported Preview */
.unsupported-preview {
  width: 100%;
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
}

.unsupported-content {
  text-align: center;
  max-width: 400px;
  padding: var(--apple-space-xl);
}

.unsupported-icon {
  margin-bottom: var(--apple-space-lg);
  opacity: 0.6;
}

/* File Info Sidebar */
.file-info-sidebar {
  width: 300px;
  background: var(--apple-background-elevated);
  border-left: 1px solid var(--apple-separator);
  padding: var(--apple-space-lg);
  overflow-y: auto;
}

.info-section h4 {
  margin: 0;
  color: var(--apple-text-primary);
}

.info-grid {
  display: flex;
  flex-direction: column;
  gap: var(--apple-space);
}

.info-item {
  display: flex;
  flex-direction: column;
  gap: var(--apple-space-xs);
}

.info-label {
  color: var(--apple-text-tertiary);
  text-transform: uppercase;
  font-weight: var(--apple-font-weight-semibold);
  letter-spacing: 0.5px;
}

.info-value {
  color: var(--apple-text-primary);
  word-break: break-all;
}

/* File Type Tags */
.file-type-tag {
  display: inline-block;
  padding: var(--apple-space-xs) var(--apple-space-sm);
  border-radius: var(--apple-radius-sm);
  font-size: var(--apple-font-size-xs);
  font-weight: var(--apple-font-weight-semibold);
  text-transform: uppercase;
  letter-spacing: 0.5px;
}

.file-type-image {
  background: rgba(52, 199, 89, 0.1);
  color: var(--apple-green);
}

.file-type-video {
  background: rgba(0, 122, 255, 0.1);
  color: var(--apple-blue);
}

.file-type-audio {
  background: rgba(255, 149, 0, 0.1);
  color: var(--apple-orange);
}

.file-type-pdf {
  background: rgba(255, 59, 48, 0.1);
  color: var(--apple-red);
}

.file-type-other {
  background: rgba(142, 142, 147, 0.1);
  color: var(--apple-gray);
}

/* 移动端适配 */
@media (max-width: 768px) {
  .apple-preview-modal {
    width: 95vw;
    height: 90vh;
  }
  
  .preview-content {
    flex-direction: column;
  }
  
  .file-info-sidebar {
    width: 100%;
    max-height: 30%;
    border-left: none;
    border-top: 1px solid var(--apple-separator);
  }
  
  .preview-header {
    padding: var(--apple-space-sm) var(--apple-space-sm);
    min-height: 60px;
    flex-wrap: nowrap;
    justify-content: space-between;
    align-items: center;
    overflow: visible; /* 允许看到按钮 */
    position: relative;
  }
  
  .header-left {
    flex: 1;
    min-width: 0;
    margin-right: var(--apple-space-sm);
    overflow: hidden;
  }
  
  .file-title {
    gap: var(--apple-space-sm);
    min-width: 0;
    overflow: hidden;
    align-items: center;
  }
  
  .file-title h3 {
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
    max-width: calc(100vw - 180px); /* 为按钮预留空间 */
    margin: 0;
    font-size: var(--apple-font-size);
  }
  
  .header-actions {
    flex-shrink: 0;
    min-width: 96px; /* 确保两个按钮的最小宽度 */
    gap: var(--apple-space-xs);
    display: flex;
    align-items: center;
    justify-content: flex-end;
  }
  
  .action-btn {
    width: 44px;
    height: 44px;
    flex-shrink: 0;
    position: relative;
  }
}

@media (max-width: 480px) {
  .apple-preview-modal {
    width: 100vw;
    height: 100vh;
    border-radius: 0;
  }
  
  .preview-header {
    padding: var(--apple-space-sm) var(--apple-space-sm);
    min-height: 56px;
    flex-direction: row;
    justify-content: space-between;
    align-items: center;
    flex-wrap: nowrap;
    overflow: hidden;
  }
  
  .header-left {
    flex: 1;
    min-width: 0;
    margin-right: var(--apple-space-sm);
    overflow: hidden;
  }
  
  .file-title {
    gap: var(--apple-space-xs);
    align-items: center;
    min-width: 0;
    overflow: hidden;
  }
  
  .file-title h3 {
    font-size: var(--apple-font-size-sm);
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
    margin: 0;
    max-width: calc(100vw - 120px); /* 为小屏幕预留更多按钮空间 */
    flex: 1;
    min-width: 0;
  }
  
  .file-icon {
    flex-shrink: 0;
    margin-right: var(--apple-space-xs);
  }
  
  .file-icon svg {
    width: 20px;
    height: 20px;
  }
  
  .header-actions {
    flex-shrink: 0;
    min-width: 88px; /* 确保小屏幕上两个按钮的空间 */
    gap: var(--apple-space-xs);
    display: flex;
    align-items: center;
    justify-content: flex-end;
  }
  
  .action-btn {
    width: 40px;
    height: 40px;
    flex-shrink: 0;
    border-radius: var(--apple-radius);
  }
  
  .action-btn svg {
    width: 18px;
    height: 18px;
  }
  
  .file-info-sidebar {
    padding: var(--apple-space);
  }
  
  .media-preview {
    padding: var(--apple-space);
  }
  
  .unsupported-content {
    padding: var(--apple-space-lg);
  }
}

/* 极小屏幕优化 */
@media (max-width: 360px) {
  .preview-header {
    padding: var(--apple-space-xs) var(--apple-space-sm);
    min-height: 52px;
  }
  
  .header-left {
    margin-right: var(--apple-space-xs);
  }
  
  .file-title h3 {
    font-size: var(--apple-font-size-xs);
    max-width: calc(100vw - 100px); /* 极小屏幕预留按钮空间 */
  }
  
  .file-icon svg {
    width: 18px;
    height: 18px;
  }
  
  .action-btn {
    width: 36px;
    height: 36px;
    flex-shrink: 0;
  }
  
  .action-btn svg {
    width: 16px;
    height: 16px;
  }
}

/* 深色模式优化 */
@media (prefers-color-scheme: dark) {
  .apple-preview-overlay {
    background: rgba(0, 0, 0, 0.85);
  }
  
  .preview-header {
    background: var(--apple-gray-2);
  }
  
  .preview-main {
    background: var(--apple-gray-2);
  }
}
</style>