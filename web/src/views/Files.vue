<template>
  <div class="files-container">
    <!-- Apple风格Header -->
    <div class="apple-header">
      <div class="header-content">
        <div class="header-left">
          <div class="app-title">
            <div class="app-icon">
              <svg width="32" height="32" viewBox="0 0 64 64" fill="none">
                <rect width="64" height="64" rx="12" fill="url(#headerGradient)"/>
                <path d="M32 16L48 28V48C48 50.2091 46.2091 52 44 52H20C17.7909 52 16 50.2091 16 48V28L32 16Z" fill="white" fill-opacity="0.9"/>
                <path d="M32 16L16 28H32V44L48 28L32 16Z" fill="white" fill-opacity="0.6"/>
                <defs>
                  <linearGradient id="headerGradient" x1="0%" y1="0%" x2="100%" y2="100%">
                    <stop offset="0%" style="stop-color:#007AFF"/>
                    <stop offset="100%" style="stop-color:#5AC8FA"/>
                  </linearGradient>
                </defs>
              </svg>
            </div>
            <div class="title-info">
              <h2 class="apple-text-title">🍉文件服务</h2>
              <span v-if="authStore.user" class="user-info apple-text-caption">
                欢迎，{{ authStore.user.username }}
              </span>
            </div>
          </div>
        </div>
        <div class="header-right">
          <button
            v-if="authStore.isAdmin"
            class="apple-button-secondary header-btn"
            title="管理面板"
            @click="$router.push('/admin')"
          >
            <svg width="16" height="16" viewBox="0 0 24 24" fill="none">
              <circle cx="12" cy="12" r="3" stroke="currentColor" stroke-width="2"/>
              <path d="M19.4 15C19.2 14.2 19 13.1 19 12s.2-2.2.4-3C18.8 8.4 18.4 8 17.8 7.6L16.2 6.8C15.6 6.4 14.8 6.6 14.4 7.2L14 8C13.6 8.8 13 9.2 12 9.2s-1.6-.4-2-.8l-.4-.8C9.2 6.6 8.4 6.4 7.8 6.8L6.2 7.6C5.6 8 5.2 8.4 4.6 9C4.8 9.8 5 10.9 5 12s-.2 2.2-.4 3C5.2 15.6 5.6 16 6.2 16.4L7.8 17.2C8.4 17.6 9.2 17.4 9.6 16.8L10 16C10.4 15.2 11 14.8 12 14.8s1.6.4 2 .8l.4.8c.4.6 1.2.8 1.8.4L17.8 16.4C18.4 16 18.8 15.6 19.4 15z" stroke="currentColor" stroke-width="2"/>
            </svg>
            <span>管理面板</span>
          </button>
          <button 
            class="apple-button-secondary header-btn" 
            title="退出登录"
            @click="handleLogout"
          >
            <svg width="16" height="16" viewBox="0 0 24 24" fill="none">
              <path d="M9 21H5C4.46957 21 3.96086 20.7893 3.58579 20.4142C3.21071 20.0391 3 19.5304 3 19V5C3 4.46957 3.21071 3.96086 3.58579 3.58579C3.96086 3.21071 4.46957 3 5 3H9" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>
              <path d="M16 17L21 12L16 7" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>
              <path d="M21 12H9" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>
            </svg>
            <span>退出登录</span>
          </button>
        </div>
      </div>
    </div>

    <div class="main-content">
      <!-- Apple风格面包屑导航 -->
      <div class="macos-breadcrumb">
        <div class="breadcrumb-path">
          <div class="breadcrumb-item root" @click="navigateToFolder(null)">
            <el-icon class="breadcrumb-icon"><HomeFilled /></el-icon>
            <span>根目录</span>
          </div>
          <template v-for="crumb in breadcrumbs" :key="crumb.id">
            <el-icon class="breadcrumb-separator"><ArrowRight /></el-icon>
            <div class="breadcrumb-item" @click="navigateToFolder(crumb.id)">
              <span>{{ crumb.name }}</span>
            </div>
          </template>
        </div>
      </div>

      <!-- Apple风格工具栏 -->
      <div class="apple-toolbar">
        <div class="toolbar-section">
          <button class="apple-button-primary toolbar-btn" @click="showCreateFolderDialog">
            <svg width="16" height="16" viewBox="0 0 24 24" fill="none">
              <path d="M10 4H4C3.46957 4 2.96086 4.21071 2.58579 4.58579C2.21071 4.96086 2 5.46957 2 6V18C2 18.5304 2.21071 19.0391 2.58579 19.4142C2.96086 19.7893 3.46957 20 4 20H20C20.5304 20 21.0391 19.7893 21.4142 19.4142C21.7893 19.0391 22 18.5304 22 18V8C22 7.46957 21.7893 6.96086 21.4142 6.58579C21.0391 6.21071 20.5304 6 20 6H12L10 4Z" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>
              <path d="M12 10V14M10 12H14" stroke="currentColor" stroke-width="2" stroke-linecap="round"/>
            </svg>
            <span>新建</span>
          </button>
          
          <el-upload
            ref="uploadRef"
            :action="`/api/files/upload`"
            :headers="{ Authorization: `Bearer ${authStore.token}` }"
            :data="{ folder_id: currentFolderId }"
            :on-success="handleUploadSuccess"
            :on-error="handleUploadError"
            :before-upload="beforeUpload"
            :show-file-list="false"
            multiple
            class="upload-wrapper"
          >
            <button class="apple-button-secondary toolbar-btn">
              <svg width="16" height="16" viewBox="0 0 24 24" fill="none">
                <path d="M21 15V19C21 20.1 20.1 21 19 21H5C3.9 21 3 20.1 3 19V15" stroke="currentColor" stroke-width="2"/>
                <path d="M17 8L12 3L7 8" stroke="currentColor" stroke-width="2"/>
                <path d="M12 3V15" stroke="currentColor" stroke-width="2"/>
              </svg>
              <span>上传</span>
            </button>
          </el-upload>
        </div>

        <div class="toolbar-section">
          <div class="file-type-filter">
            <select v-model="selectedFileType" @change="fetchData" class="apple-select">
              <option value="">所有类型</option>
              <option value="image">图片</option>
              <option value="video">视频</option>
              <option value="audio">音频</option>
              <option value="pdf">文档</option>
              <option value="other">其他</option>
            </select>
          </div>
          
          <div class="view-toggle">
            <button 
              :class="['toggle-button', { active: viewMode === 'grid' }]"
              @click="viewMode = 'grid'"
              title="网格视图"
            >
              <svg width="16" height="16" viewBox="0 0 24 24" fill="none">
                <rect x="3" y="3" width="7" height="7" rx="1" stroke="currentColor" stroke-width="2"/>
                <rect x="14" y="3" width="7" height="7" rx="1" stroke="currentColor" stroke-width="2"/>
                <rect x="14" y="14" width="7" height="7" rx="1" stroke="currentColor" stroke-width="2"/>
                <rect x="3" y="14" width="7" height="7" rx="1" stroke="currentColor" stroke-width="2"/>
              </svg>
            </button>
            <button 
              :class="['toggle-button', { active: viewMode === 'list' }]"
              @click="viewMode = 'list'"
              title="列表视图"
            >
              <svg width="16" height="16" viewBox="0 0 24 24" fill="none">
                <path d="M8 6H21" stroke="currentColor" stroke-width="2" stroke-linecap="round"/>
                <path d="M8 12H21" stroke="currentColor" stroke-width="2" stroke-linecap="round"/>
                <path d="M8 18H21" stroke="currentColor" stroke-width="2" stroke-linecap="round"/>
                <path d="M3 6H3.01" stroke="currentColor" stroke-width="2" stroke-linecap="round"/>
                <path d="M3 12H3.01" stroke="currentColor" stroke-width="2" stroke-linecap="round"/>
                <path d="M3 18H3.01" stroke="currentColor" stroke-width="2" stroke-linecap="round"/>
              </svg>
            </button>
          </div>
        </div>
      </div>

      <!-- 文件内容区域 -->
      <div v-loading="loading" class="files-content">
        <div v-if="viewMode === 'grid'" class="macos-files-grid">
          <!-- 文件夹 -->
          <div
            v-for="folder in folders"
            :key="`folder-${folder.id}`"
            class="macos-file-card folder-card"
            @click="navigateToFolder(folder.id)"
          >
            <div class="card-preview">
              <div class="folder-icon-wrapper">
                <div class="folder-icon">
                  <el-icon><Folder /></el-icon>
                </div>
              </div>
            </div>
            <div class="card-info">
              <div class="card-title" :title="folder.name">
                {{ folder.name }}
              </div>
              <div class="card-meta">
                文件夹 • {{ formatDate(folder.created_at) }}
              </div>
            </div>
            <div class="card-actions" @click.stop>
              <el-dropdown trigger="click" placement="bottom-end">
                <button class="action-button">
                  <el-icon><MoreFilled /></el-icon>
                </button>
                <template #dropdown>
                  <el-dropdown-menu class="macos-dropdown">
                    <el-dropdown-item @click="showRenameFolderDialog(folder)">
                      <el-icon><Edit /></el-icon>
                      重命名
                    </el-dropdown-item>
                    <el-dropdown-item @click="showMoveFolderDialog(folder)">
                      <el-icon><Position /></el-icon>
                      移动
                    </el-dropdown-item>
                    <el-dropdown-item @click="deleteFolder(folder)" divided class="danger">
                      <el-icon><Delete /></el-icon>
                      删除
                    </el-dropdown-item>
                  </el-dropdown-menu>
                </template>
              </el-dropdown>
            </div>
          </div>
          
          <!-- 文件 -->
          <div
            v-for="file in files"
            :key="`file-${file.id}`"
            class="macos-file-card file-card"
            @click="handleFileClick(file)"
          >
            <div class="card-preview">
              <div class="preview-wrapper">
                <img
                  v-if="file.is_image"
                  :src="getImagePreviewUrl(file.id)"
                  :alt="file.original_name"
                  class="preview-image"
                  @load="handleImageLoad"
                  @error="handleImageError"
                />
                <video
                  v-else-if="file.is_video"
                  :src="getVideoPreviewUrl(file.id)"
                  class="preview-video"
                  preload="metadata"
                />
                <div v-else class="file-icon-wrapper">
                  <div class="file-icon" :class="getFileIconClass(file.file_type)">
                    <el-icon><Document /></el-icon>
                  </div>
                </div>
              </div>
            </div>
            <div class="card-info">
              <div class="card-title" :title="file.original_name">
                {{ file.original_name }}
              </div>
              <div class="card-meta">
                {{ formatFileSize(file.file_size) }} • {{ formatDate(file.created_at) }}
              </div>
            </div>
            <div class="card-actions" @click.stop>
              <el-dropdown trigger="click" placement="bottom-end">
                <button class="action-button">
                  <el-icon><MoreFilled /></el-icon>
                </button>
                <template #dropdown>
                  <el-dropdown-menu class="macos-dropdown">
                    <el-dropdown-item @click="downloadFile(file)">
                      <el-icon><Download /></el-icon>
                      下载
                    </el-dropdown-item>
                    <el-dropdown-item @click="showMoveFileDialog(file)">
                      <el-icon><Position /></el-icon>
                      移动
                    </el-dropdown-item>
                    <el-dropdown-item @click="deleteFile(file)" divided class="danger">
                      <el-icon><Delete /></el-icon>
                      删除
                    </el-dropdown-item>
                  </el-dropdown-menu>
                </template>
              </el-dropdown>
            </div>
          </div>
        </div>

        <el-table v-else :data="[...folders, ...files]" style="width: 100%">
          <el-table-column prop="name" label="名称" min-width="200">
            <template #default="scope">
              <div class="file-name-cell">
                <el-icon v-if="scope.row.name">
                  <Folder />
                </el-icon>
                <el-icon v-else-if="scope.row.is_image"><Picture /></el-icon>
                <el-icon v-else-if="scope.row.is_video"><VideoCamera /></el-icon>
                <el-icon v-else><Document /></el-icon>
                <span 
                  @click="scope.row.name ? navigateToFolder(scope.row.id) : handleFileClick(scope.row)" 
                  class="file-link"
                >
                  {{ scope.row.name || scope.row.original_name }}
                </span>
              </div>
            </template>
          </el-table-column>
          
          <el-table-column prop="file_type" label="类型" width="100">
            <template #default="scope">
              <el-tag v-if="scope.row.name" type="warning" size="small">
                文件夹
              </el-tag>
              <el-tag v-else size="small" :type="getFileTypeColor(scope.row.file_type)">
                {{ scope.row.file_type }}
              </el-tag>
            </template>
          </el-table-column>
          
          <el-table-column prop="file_size" label="大小" width="120">
            <template #default="scope">
              <span v-if="!scope.row.name">
                {{ formatFileSize(scope.row.file_size) }}
              </span>
              <span v-else>-</span>
            </template>
          </el-table-column>
          
          <el-table-column prop="created_at" label="创建时间" width="180">
            <template #default="scope">
              {{ formatDate(scope.row.created_at) }}
            </template>
          </el-table-column>
          
          <el-table-column label="操作" width="120" fixed="right">
            <template #default="scope">
              <el-dropdown trigger="click">
                <el-button size="small" :icon="MoreFilled" />
                <template #dropdown>
                  <el-dropdown-menu>
                    <template v-if="scope.row.name">
                      <el-dropdown-item @click="showRenameFolderDialog(scope.row)">
                        重命名
                      </el-dropdown-item>
                      <el-dropdown-item @click="showMoveFolderDialog(scope.row)">
                        移动
                      </el-dropdown-item>
                      <el-dropdown-item @click="deleteFolder(scope.row)" divided>
                        删除
                      </el-dropdown-item>
                    </template>
                    <template v-else>
                      <el-dropdown-item @click="downloadFile(scope.row)">
                        下载
                      </el-dropdown-item>
                      <el-dropdown-item @click="showMoveFileDialog(scope.row)">
                        移动
                      </el-dropdown-item>
                      <el-dropdown-item @click="deleteFile(scope.row)" divided>
                        删除
                      </el-dropdown-item>
                    </template>
                  </el-dropdown-menu>
                </template>
              </el-dropdown>
            </template>
          </el-table-column>
        </el-table>

        <!-- 空状态 -->
        <div v-if="!loading && files.length === 0 && folders.length === 0" class="macos-empty-state">
          <div class="empty-icon">
            <el-icon><Folder /></el-icon>
          </div>
          <div class="empty-title">暂无文件</div>
          <div class="empty-description">点击上传按钮开始使用</div>
        </div>
      </div>

      <!-- Apple风格分页 -->
      <div v-if="total > 0" class="macos-pagination">
        <el-pagination
          v-model:current-page="currentPage"
          v-model:page-size="pageSize"
          :total="total"
          :page-sizes="[10, 20, 50, 100]"
          layout="total, sizes, prev, pager, next, jumper"
          @size-change="fetchData"
          @current-change="fetchData"
          background
        />
      </div>
    </div>

    <!-- 文件预览 -->
    <FilePreview
      v-model:visible="previewVisible"
      :file="currentFile"
    />

    <!-- 创建文件夹对话框 -->
    <el-dialog v-model="createFolderVisible" title="新建文件夹" width="400px">
      <el-form :model="folderForm" :rules="folderRules" ref="folderFormRef">
        <el-form-item label="文件夹名称" prop="name">
          <el-input v-model="folderForm.name" placeholder="请输入文件夹名称" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="createFolderVisible = false">取消</el-button>
        <el-button type="primary" @click="createFolder" :loading="createFolderLoading">
          确定
        </el-button>
      </template>
    </el-dialog>

    <!-- 重命名文件夹对话框 -->
    <el-dialog v-model="renameFolderVisible" title="重命名文件夹" width="400px">
      <el-form :model="renameForm" :rules="folderRules" ref="renameFormRef">
        <el-form-item label="文件夹名称" prop="name">
          <el-input v-model="renameForm.name" placeholder="请输入新的文件夹名称" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="renameFolderVisible = false">取消</el-button>
        <el-button type="primary" @click="renameFolder" :loading="renameFolderLoading">
          确定
        </el-button>
      </template>
    </el-dialog>

    <!-- 移动文件夹对话框 -->
    <el-dialog v-model="moveFolderVisible" title="移动文件夹" width="400px">
      <div class="move-dialog-content">
        <p>移动到：</p>
        <el-tree
          :data="folderTree"
          :props="{ children: 'children', label: 'name' }"
          node-key="id"
          :default-expand-all="true"
          :expand-on-click-node="false"
          @node-click="selectMoveTarget"
          ref="folderTreeRef"
        >
          <template #default="{ node, data }">
            <span :class="{ 'selected-folder': selectedMoveTarget?.id === data.id }">
              <el-icon><Folder /></el-icon>
              {{ data.name }}
            </span>
          </template>
        </el-tree>
        <div style="margin-top: 10px;">
          <el-button @click="selectMoveTarget({ id: null, name: '根目录' })" size="small">
            移动到根目录
          </el-button>
        </div>
      </div>
      <template #footer>
        <el-button @click="moveFolderVisible = false">取消</el-button>
        <el-button type="primary" @click="moveFolder" :loading="moveFolderLoading">
          确定
        </el-button>
      </template>
    </el-dialog>

    <!-- 移动文件对话框 -->
    <el-dialog v-model="moveFileVisible" title="移动文件" width="400px">
      <div class="move-dialog-content">
        <p>移动到：</p>
        <el-tree
          :data="folderTree"
          :props="{ children: 'children', label: 'name' }"
          node-key="id"
          :default-expand-all="true"
          :expand-on-click-node="false"
          @node-click="selectMoveTarget"
          ref="fileTreeRef"
        >
          <template #default="{ node, data }">
            <span :class="{ 'selected-folder': selectedMoveTarget?.id === data.id }">
              <el-icon><Folder /></el-icon>
              {{ data.name }}
            </span>
          </template>
        </el-tree>
        <div style="margin-top: 10px;">
          <el-button @click="selectMoveTarget({ id: null, name: '根目录' })" size="small">
            移动到根目录
          </el-button>
        </div>
      </div>
      <template #footer>
        <el-button @click="moveFileVisible = false">取消</el-button>
        <el-button type="primary" @click="moveFile" :loading="moveFileLoading">
          确定
        </el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script>
import { ref, reactive, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import { ElMessage, ElMessageBox } from 'element-plus'
import {
  Upload,
  Download,
  Delete,
  Document,
  Picture,
  VideoCamera,
  Folder,
  FolderAdd,
  HomeFilled,
  MoreFilled,
  ArrowRight,
  Grid,
  List,
  Edit,
  Position,
  Setting,
  SwitchButton
} from '@element-plus/icons-vue'
import request from '@/utils/request'
import FilePreview from '@/components/FilePreview.vue'

export default {
  name: 'Files',
  components: {
    FilePreview
  },
  setup() {
    const router = useRouter()
    const authStore = useAuthStore()
    
    const loading = ref(false)
    const files = ref([])
    const folders = ref([])
    const folderTree = ref([])
    const total = ref(0)
    const currentPage = ref(1)
    const pageSize = ref(20)
    const selectedFileType = ref('')
    const viewMode = ref('grid')
    const previewVisible = ref(false)
    const currentFile = ref(null)
    const currentFolderId = ref(null)
    const breadcrumbs = ref([])
    
    // 文件夹操作对话框
    const createFolderVisible = ref(false)
    const createFolderLoading = ref(false)
    const renameFolderVisible = ref(false)
    const renameFolderLoading = ref(false)
    const moveFolderVisible = ref(false)
    const moveFolderLoading = ref(false)
    const moveFileVisible = ref(false)
    const moveFileLoading = ref(false)
    
    // 表单数据
    const folderForm = reactive({ name: '' })
    const renameForm = reactive({ name: '', id: null })
    const currentMoveItem = ref(null)
    const selectedMoveTarget = ref(null)
    
    // 表单验证规则
    const folderRules = {
      name: [
        { required: true, message: '请输入文件夹名称', trigger: 'blur' },
        { min: 1, max: 255, message: '文件夹名称长度在 1 到 255 个字符', trigger: 'blur' }
      ]
    }
    
    // 获取文件和文件夹数据
    const fetchData = async () => {
      await Promise.all([fetchFiles(), fetchFolders(), fetchFolderTree()])
    }
    
    const fetchFiles = async () => {
      try {
        const params = {
          page: currentPage.value,
          limit: pageSize.value
        }
        if (selectedFileType.value) {
          params.type = selectedFileType.value
        }
        if (currentFolderId.value) {
          params.folder_id = currentFolderId.value
        }
        
        const response = await request.get('/files', { params })
        files.value = response.files
        total.value = response.total
      } catch (error) {
        console.error('Failed to fetch files:', error)
      }
    }
    
    const fetchFolders = async () => {
      try {
        const params = {}
        if (currentFolderId.value) {
          params.parent_id = currentFolderId.value
        }
        
        const response = await request.get('/folders', { params })
        folders.value = response.folders
      } catch (error) {
        console.error('Failed to fetch folders:', error)
      }
    }
    
    const fetchFolderTree = async () => {
      try {
        const response = await request.get('/folders')
        folderTree.value = buildFolderTree(response.folders)
      } catch (error) {
        console.error('Failed to fetch folder tree:', error)
      }
    }
    
    const buildFolderTree = (folders) => {
      const map = {}
      const roots = []
      
      // 首先创建所有节点的映射
      folders.forEach(folder => {
        map[folder.id] = { ...folder, children: [] }
      })
      
      // 然后构建树结构
      folders.forEach(folder => {
        if (folder.parent_id) {
          if (map[folder.parent_id]) {
            map[folder.parent_id].children.push(map[folder.id])
          }
        } else {
          roots.push(map[folder.id])
        }
      })
      
      return roots
    }
    
    const handleUploadSuccess = (response) => {
      ElMessage.success('文件上传成功')
      fetchData()
    }
    
    // 导航到文件夹
    const navigateToFolder = async (folderId) => {
      currentFolderId.value = folderId
      currentPage.value = 1
      
      if (folderId) {
        try {
          const response = await request.get(`/folders/${folderId}/path`)
          breadcrumbs.value = response.breadcrumbs
        } catch (error) {
          console.error('Failed to fetch folder path:', error)
        }
      } else {
        breadcrumbs.value = []
      }
      
      loading.value = true
      try {
        await fetchData()
      } finally {
        loading.value = false
      }
    }
    
    // 创建文件夹
    const showCreateFolderDialog = () => {
      folderForm.name = ''
      createFolderVisible.value = true
    }
    
    const createFolder = async () => {
      createFolderLoading.value = true
      try {
        const data = { name: folderForm.name }
        if (currentFolderId.value) {
          data.parent_id = currentFolderId.value
        }
        
        await request.post('/folders', data)
        ElMessage.success('文件夹创建成功')
        createFolderVisible.value = false
        fetchData()
      } catch (error) {
        console.error('Failed to create folder:', error)
        ElMessage.error('文件夹创建失败')
      } finally {
        createFolderLoading.value = false
      }
    }
    
    // 重命名文件夹
    const showRenameFolderDialog = (folder) => {
      renameForm.name = folder.name
      renameForm.id = folder.id
      renameFolderVisible.value = true
    }
    
    const renameFolder = async () => {
      renameFolderLoading.value = true
      try {
        await request.put(`/folders/${renameForm.id}`, { name: renameForm.name })
        ElMessage.success('文件夹重命名成功')
        renameFolderVisible.value = false
        fetchData()
      } catch (error) {
        console.error('Failed to rename folder:', error)
        ElMessage.error('文件夹重命名失败')
      } finally {
        renameFolderLoading.value = false
      }
    }
    
    // 删除文件夹
    const deleteFolder = async (folder) => {
      try {
        await ElMessageBox.confirm(
          `确定要删除文件夹 "${folder.name}" 吗？`,
          '删除确认',
          {
            confirmButtonText: '确定',
            cancelButtonText: '取消',
            type: 'warning'
          }
        )
        
        await request.delete(`/folders/${folder.id}`)
        ElMessage.success('文件夹删除成功')
        fetchData()
      } catch (error) {
        if (error !== 'cancel') {
          console.error('Delete folder error:', error)
          ElMessage.error('文件夹删除失败')
        }
      }
    }
    
    // 移动文件夹
    const showMoveFolderDialog = (folder) => {
      currentMoveItem.value = folder
      selectedMoveTarget.value = null
      moveFolderVisible.value = true
    }
    
    const moveFolder = async () => {
      moveFolderLoading.value = true
      try {
        const data = { parent_id: selectedMoveTarget.value?.id || null }
        await request.put(`/folders/${currentMoveItem.value.id}/move`, data)
        ElMessage.success('文件夹移动成功')
        moveFolderVisible.value = false
        fetchData()
      } catch (error) {
        console.error('Failed to move folder:', error)
        ElMessage.error('文件夹移动失败')
      } finally {
        moveFolderLoading.value = false
      }
    }
    
    // 移动文件
    const showMoveFileDialog = (file) => {
      currentMoveItem.value = file
      selectedMoveTarget.value = null
      moveFileVisible.value = true
    }
    
    const moveFile = async () => {
      moveFileLoading.value = true
      try {
        const data = { folder_id: selectedMoveTarget.value?.id || null }
        await request.put(`/files/${currentMoveItem.value.id}/move`, data)
        ElMessage.success('文件移动成功')
        moveFileVisible.value = false
        fetchData()
      } catch (error) {
        console.error('Failed to move file:', error)
        ElMessage.error('文件移动失败')
      } finally {
        moveFileLoading.value = false
      }
    }
    
    // 选择移动目标
    const selectMoveTarget = (target) => {
      selectedMoveTarget.value = target
    }
    
    const handleUploadError = (error) => {
      console.error('Upload error:', error)
      ElMessage.error('文件上传失败')
    }
    
    const beforeUpload = (file) => {
      const maxSize = 100 * 1024 * 1024 // 100MB
      if (file.size > maxSize) {
        ElMessage.error('文件大小不能超过100MB')
        return false
      }
      return true
    }
    
    const downloadFile = async (file) => {
      try {
        const response = await fetch(`/api/files/${file.id}/download`, {
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
        a.download = file.original_name
        document.body.appendChild(a)
        a.click()
        window.URL.revokeObjectURL(url)
        document.body.removeChild(a)
      } catch (error) {
        console.error('Download error:', error)
        ElMessage.error('下载失败')
      }
    }
    
    const deleteFile = async (file) => {
      try {
        await ElMessageBox.confirm(
          `确定要删除文件 "${file.original_name}" 吗？`,
          '删除确认',
          {
            confirmButtonText: '确定',
            cancelButtonText: '取消',
            type: 'warning'
          }
        )
        
        await request.delete(`/files/${file.id}`)
        ElMessage.success('文件删除成功')
        fetchData()
      } catch (error) {
        if (error !== 'cancel') {
          console.error('Delete error:', error)
        }
      }
    }
    
    const handleFileClick = (file) => {
      if (file.is_image || file.is_video) {
        currentFile.value = file
        previewVisible.value = true
      } else {
        downloadFile(file)
      }
    }
    
    const handleLogout = () => {
      authStore.logout()
      router.push('/login')
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
    
    const getFileIconClass = (type) => {
      const classes = {
        image: 'file-icon-image',
        video: 'file-icon-video',
        audio: 'file-icon-audio',
        pdf: 'file-icon-pdf',
        other: 'file-icon-other'
      }
      return classes[type] || 'file-icon-other'
    }
    
    const previewCache = ref(new Map())
    
    const getImagePreviewUrl = (fileId) => {
      // 对于图片，必须使用带认证的blob URL
      if (!previewCache.value.has(fileId)) {
        createAuthenticatedPreview(fileId)
        // 返回占位符直到blob URL准备好
        return 'data:image/svg+xml;base64,PHN2ZyB3aWR0aD0iMjAwIiBoZWlnaHQ9IjEyMCIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIj48cmVjdCB3aWR0aD0iMTAwJSIgaGVpZ2h0PSIxMDAlIiBmaWxsPSIjZjVmN2ZhIi8+PHRleHQgeD0iNTAlIiB5PSI1MCUiIGZpbGw9IiM5MDkzOTkiIHRleHQtYW5jaG9yPSJtaWRkbGUiIGR5PSIuM2VtIj7liqDovb3kuK0uLi48L3RleHQ+PC9zdmc+'
      }
      
      const cachedUrl = previewCache.value.get(fileId)
      if (cachedUrl === 'error') {
        // 返回错误图标
        return 'data:image/svg+xml;base64,PHN2ZyB3aWR0aD0iMjAwIiBoZWlnaHQ9IjEyMCIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIj48cmVjdCB3aWR0aD0iMTAwJSIgaGVpZ2h0PSIxMDAlIiBmaWxsPSIjZmVmMGYwIi8+PHRleHQgeD0iNTAlIiB5PSI1MCUiIGZpbGw9IiNmNTZjNmMiIHRleHQtYW5jaG9yPSJtaWRkbGUiIGR5PSIuM2VtIj7liqDovb3lpLHotKU8L3RleHQ+PC9zdmc+'
      }
      
      return cachedUrl || 'data:image/svg+xml;base64,PHN2ZyB3aWR0aD0iMjAwIiBoZWlnaHQ9IjEyMCIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIj48cmVjdCB3aWR0aD0iMTAwJSIgaGVpZ2h0PSIxMDAlIiBmaWxsPSIjZjVmN2ZhIi8+PHRleHQgeD0iNTAlIiB5PSI1MCUiIGZpbGw9IiM5MDkzOTkiIHRleHQtYW5jaG9yPSJtaWRkbGUiIGR5PSIuM2VtIj7liqDovb3kuK0uLi48L3RleHQ+PC9zdmc+'
    }
    
    const getVideoPreviewUrl = (fileId) => {
      // 视频直接使用API接口，浏览器会自动处理认证
      return `/api/files/${fileId}/preview`
    }
    
    const createAuthenticatedPreview = async (fileId) => {
      try {
        const response = await fetch(`/api/files/${fileId}/preview`, {
          headers: {
            'Authorization': `Bearer ${authStore.token}`
          }
        })
        
        if (response.ok) {
          const blob = await response.blob()
          const url = URL.createObjectURL(blob)
          previewCache.value.set(fileId, url)
          
          // 触发重新渲染以显示新的blob URL
          files.value = [...files.value]
          
          // 60秒后清理缓存
          setTimeout(() => {
            URL.revokeObjectURL(url)
            previewCache.value.delete(fileId)
          }, 60000)
        } else {
          console.error(`Preview failed with status: ${response.status}`)
          // 认证失败时显示错误图标
          previewCache.value.set(fileId, 'error')
        }
      } catch (error) {
        console.error('Failed to create preview:', error)
        previewCache.value.set(fileId, 'error')
      }
    }
    
    const handleImageLoad = (event) => {
      // 图片加载成功
      console.log('Image loaded successfully')
    }
    
    const handleImageError = (event) => {
      console.error('Image load failed, trying fallback')
      // 如果图片加载失败，显示默认图标
      const img = event.target
      img.style.display = 'none'
      
      // 可以在这里添加重试逻辑或显示默认图标
      const parent = img.parentElement
      if (parent && !parent.querySelector('.fallback-icon')) {
        const fallbackDiv = document.createElement('div')
        fallbackDiv.className = 'fallback-icon'
        fallbackDiv.innerHTML = '<i class="el-icon-picture" style="font-size: 48px; color: #ccc;"></i>'
        parent.appendChild(fallbackDiv)
      }
    }
    
    onMounted(() => {
      loading.value = true
      fetchData().finally(() => {
        loading.value = false
      })
    })
    
    return {
      authStore,
      loading,
      files,
      folders,
      folderTree,
      total,
      currentPage,
      pageSize,
      selectedFileType,
      viewMode,
      previewVisible,
      currentFile,
      currentFolderId,
      breadcrumbs,
      
      // 对话框状态
      createFolderVisible,
      createFolderLoading,
      renameFolderVisible,
      renameFolderLoading,
      moveFolderVisible,
      moveFolderLoading,
      moveFileVisible,
      moveFileLoading,
      
      // 表单数据
      folderForm,
      renameForm,
      folderRules,
      selectedMoveTarget,
      
      // 方法
      fetchData,
      navigateToFolder,
      showCreateFolderDialog,
      createFolder,
      showRenameFolderDialog,
      renameFolder,
      deleteFolder,
      showMoveFolderDialog,
      moveFolder,
      showMoveFileDialog,
      moveFile,
      selectMoveTarget,
      handleUploadSuccess,
      handleUploadError,
      beforeUpload,
      downloadFile,
      deleteFile,
      handleFileClick,
      handleLogout,
      formatFileSize,
      formatDate,
      getFileTypeColor,
      getImagePreviewUrl,
      getVideoPreviewUrl,
      handleImageLoad,
      handleImageError,
      getFileIconClass,
      
      // 图标
      Upload,
      Download,
      Delete,
      Document,
      Picture,
      VideoCamera,
      Folder,
      FolderAdd,
      HomeFilled,
      MoreFilled,
      ArrowRight,
      Grid,
      List,
      Edit,
      Position,
      Setting,
      SwitchButton
    }
  }
}
</script>

<style scoped>
/* Apple风格主题配置 */
:root {
  --apple-primary: #007AFF;
  --apple-secondary: #34C759;
  --apple-danger: #FF3B30;
  --apple-warning: #FF9500;
  --apple-gray-1: #F2F2F7;
  --apple-gray-2: #E5E5EA;
  --apple-gray-3: #C7C7CC;
  --apple-gray-4: #8E8E93;
  --apple-gray-5: #636366;
  --apple-gray-6: #48484A;
  --apple-background: #FFFFFF;
  --apple-surface: #FBFBFD;
  --apple-border: rgba(0, 0, 0, 0.08);
  --apple-shadow: rgba(0, 0, 0, 0.1);
  --apple-shadow-light: rgba(0, 0, 0, 0.05);
  --apple-text-primary: #1C1C1E;
  --apple-text-secondary: #8E8E93;
}

.files-container {
  height: 100vh;
  display: flex;
  flex-direction: column;
  background: var(--apple-surface);
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', 'Roboto', sans-serif;
}

/* Apple风格Header */
.apple-header {
  position: relative;
  background: var(--apple-background-elevated);
  border-bottom: 1px solid var(--apple-separator);
  backdrop-filter: blur(20px);
  -webkit-backdrop-filter: blur(20px);
}

.header-content {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: var(--apple-space) var(--apple-space-lg);
  min-height: 64px;
}

.header-left {
  display: flex;
  align-items: center;
}

.app-title {
  display: flex;
  align-items: center;
  gap: var(--apple-space);
}

.app-icon {
  display: flex;
  align-items: center;
  justify-content: center;
}

.title-info {
  display: flex;
  flex-direction: column;
  gap: var(--apple-space-xs);
}

.title-info h2 {
  margin: 0;
  color: var(--apple-text-primary);
}

.user-info {
  color: var(--apple-text-tertiary);
}

.header-right {
  display: flex;
  gap: var(--apple-space-sm);
  align-items: center;
}

.header-btn {
  display: flex;
  align-items: center;
  gap: var(--apple-space-xs);
  padding: var(--apple-space-sm) var(--apple-space);
  font-size: var(--apple-font-size-sm);
  height: 36px;
}

.header-btn svg {
  opacity: 0.7;
  transition: opacity 0.2s ease;
}

.header-btn:hover svg {
  opacity: 1;
}

.macos-button.primary {
  background: var(--apple-primary);
  color: white;
}

.macos-button.primary:hover {
  background: #0051D2;
}

.macos-button.admin-button {
  background: var(--apple-warning);
  color: white;
}

.macos-button.admin-button:hover {
  background: #E6820A;
}

.macos-button.logout-button {
  background: var(--apple-danger);
  color: white;
}

.macos-button.logout-button:hover {
  background: #D70015;
}

/* 主内容区域 */
.main-content {
  flex: 1;
  padding: 20px;
  overflow: auto;
  background: var(--apple-surface);
}

/* Apple风格面包屑导航 */
.macos-breadcrumb {
  margin-bottom: 20px;
  padding: 12px 16px;
  background: var(--apple-background);
  border-radius: 12px;
  border: 1px solid var(--apple-border);
}

.breadcrumb-path {
  display: flex;
  align-items: center;
  gap: 8px;
}

.breadcrumb-item {
  display: flex;
  align-items: center;
  gap: 6px;
  padding: 4px 8px;
  border-radius: 6px;
  cursor: pointer;
  transition: all 0.2s;
  font-size: 14px;
  font-weight: 500;
  color: var(--apple-text-primary);
}

.breadcrumb-item:hover {
  background: var(--apple-gray-1);
  color: var(--apple-primary);
}

.breadcrumb-item.root {
  color: var(--apple-primary);
}

.breadcrumb-icon {
  font-size: 14px;
}

.breadcrumb-separator {
  color: var(--apple-text-secondary);
  font-size: 12px;
}

/* Apple风格工具栏 */
.apple-toolbar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: var(--apple-space-lg);
  padding: var(--apple-space) var(--apple-space-lg);
  background: var(--apple-background-elevated);
  border-radius: var(--apple-radius-lg);
  border: 1px solid var(--apple-border);
  box-shadow: var(--apple-shadow-sm);
}

.toolbar-section {
  display: flex;
  align-items: center;
  gap: var(--apple-space);
}

.toolbar-btn {
  display: flex;
  align-items: center;
  gap: var(--apple-space-xs);
  padding: var(--apple-space-sm) var(--apple-space);
  font-size: var(--apple-font-size-sm);
  height: 36px;
}

.toolbar-btn svg {
  opacity: 0.8;
  transition: opacity 0.2s ease;
}

.toolbar-btn:hover svg {
  opacity: 1;
}

.upload-wrapper {
  display: inline-block;
}

/* Apple风格选择器 */
.apple-select {
  padding: var(--apple-space-sm) var(--apple-space);
  border: 1px solid var(--apple-border);
  border-radius: var(--apple-radius);
  background: var(--apple-background-elevated);
  color: var(--apple-text-primary);
  font-size: var(--apple-font-size-sm);
  font-weight: var(--apple-font-weight-medium);
  cursor: pointer;
  transition: all 0.2s ease;
  outline: none;
  min-width: 120px;
}

.apple-select:hover {
  background: var(--apple-gray-1);
  border-color: var(--apple-gray-3);
}

.apple-select:focus {
  border-color: var(--apple-blue);
  box-shadow: 0 0 0 3px rgba(0, 122, 255, 0.1);
}

/* 视图切换按钮 */
.view-toggle {
  display: flex;
  background: var(--apple-gray-1);
  border-radius: var(--apple-radius);
  padding: 2px;
  border: 1px solid var(--apple-border);
}

.toggle-button {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 36px;
  height: 32px;
  border: none;
  border-radius: var(--apple-radius-sm);
  background: transparent;
  color: var(--apple-text-secondary);
  cursor: pointer;
  transition: all 0.2s ease;
}

.toggle-button.active {
  background: var(--apple-background-elevated);
  color: var(--apple-blue);
  box-shadow: var(--apple-shadow-sm);
}

.toggle-button:hover:not(.active) {
  background: var(--apple-gray-2);
  color: var(--apple-text-primary);
}

.toggle-button svg {
  transition: all 0.2s ease;
}

/* Apple风格文件网格 */
.macos-files-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(180px, 1fr));
  gap: 16px;
}

/* Apple风格文件卡片 */
.macos-file-card {
  background: var(--apple-background);
  border: 1px solid var(--apple-border);
  border-radius: 12px;
  padding: 16px;
  cursor: pointer;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  position: relative;
  overflow: hidden;
}

.macos-file-card:hover {
  transform: translateY(-2px);
  box-shadow: 0 8px 25px var(--apple-shadow);
  border-color: var(--apple-primary);
}

.macos-file-card:active {
  transform: translateY(-1px);
}

/* 文件夹卡片特殊样式 */
.macos-file-card.folder-card {
  background: linear-gradient(135deg, #c1c4d3ff 0%, #7298baff 100%);
  border: none;
  color: white;
}

.macos-file-card.folder-card:hover {
  box-shadow: 0 8px 25px rgba(102, 126, 234, 0.3);
  transform: translateY(-3px);
}

.macos-file-card.folder-card .card-title,
.macos-file-card.folder-card .card-meta {
  color: white;
}

.macos-file-card.folder-card .folder-icon {
  color: rgba(255, 255, 255, 0.9);
}

/* 卡片预览区域 */
.card-preview {
  height: 120px;
  display: flex;
  align-items: center;
  justify-content: center;
  margin-bottom: 12px;
  border-radius: 8px;
  overflow: hidden;
  background: var(--apple-gray-1);
}

.folder-card .card-preview {
  background: rgba(255, 255, 255, 0.1);
}

.folder-icon-wrapper {
  display: flex;
  align-items: center;
  justify-content: center;
}

.folder-icon {
  font-size: 48px;
  color: rgba(255, 255, 255, 0.9);
}

.preview-wrapper {
  width: 100%;
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
}

.preview-image {
  max-width: 100%;
  max-height: 100%;
  object-fit: cover;
  border-radius: 6px;
}

.preview-video {
  max-width: 100%;
  max-height: 100%;
  border-radius: 6px;
}

.file-icon-wrapper {
  display: flex;
  align-items: center;
  justify-content: center;
}

.file-icon {
  font-size: 48px;
  transition: all 0.2s;
}

.file-icon-image {
  color: var(--apple-secondary);
}

.file-icon-video {
  color: var(--apple-primary);
}

.file-icon-audio {
  color: var(--apple-warning);
}

.file-icon-pdf {
  color: var(--apple-danger);
}

.file-icon-other {
  color: var(--apple-text-secondary);
}

/* 卡片信息区域 */
.card-info {
  margin-bottom: 8px;
}

.card-title {
  font-weight: 600;
  color: var(--apple-text-primary);
  font-size: 14px;
  line-height: 1.3;
  margin-bottom: 4px;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.card-meta {
  font-size: 12px;
  color: var(--apple-text-secondary);
  line-height: 1.2;
}

/* 卡片操作区域 */
.card-actions {
  position: absolute;
  top: 8px;
  right: 8px;
  opacity: 0;
  transition: all 0.2s;
}

.macos-file-card:hover .card-actions {
  opacity: 1;
}

.action-button {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 28px;
  height: 28px;
  border: none;
  border-radius: 6px;
  background: rgba(255, 255, 255, 0.9);
  color: var(--apple-text-secondary);
  cursor: pointer;
  transition: all 0.2s;
}

.action-button:hover {
  background: white;
  color: var(--apple-text-primary);
  box-shadow: 0 2px 8px var(--apple-shadow-light);
}

.folder-card .action-button {
  background: rgba(255, 255, 255, 0.2);
  color: rgba(255, 255, 255, 0.8);
}

.folder-card .action-button:hover {
  background: rgba(255, 255, 255, 0.3);
  color: white;
}

/* Apple风格下拉菜单 */
.macos-dropdown .el-dropdown-menu__item {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 8px 12px;
  font-size: 13px;
  color: var(--apple-text-primary);
}

.macos-dropdown .el-dropdown-menu__item.danger {
  color: var(--apple-danger);
}

/* 空状态 */
.macos-empty-state {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 200px;
  padding: 40px;
  text-align: center;
}

.empty-icon {
  font-size: 64px;
  color: var(--apple-text-secondary);
  margin-bottom: 16px;
}

.empty-title {
  font-size: 18px;
  font-weight: 600;
  color: var(--apple-text-primary);
  margin-bottom: 8px;
}

.empty-description {
  font-size: 14px;
  color: var(--apple-text-secondary);
}

/* Apple风格分页 */
.macos-pagination {
  display: flex;
  justify-content: center;
  margin-top: 32px;
  padding: 16px;
}

/* 列表视图样式保持原有Element Plus样式 */
.file-name-cell {
  display: flex;
  align-items: center;
  gap: 8px;
}

.file-link {
  cursor: pointer;
  color: var(--apple-primary);
  transition: all 0.2s;
}

.file-link:hover {
  text-decoration: underline;
}

/* 对话框相关样式 */
.move-dialog-content {
  max-height: 300px;
  overflow-y: auto;
}

.selected-folder {
  background-color: var(--apple-primary);
  color: white;
  padding: 2px 8px;
  border-radius: 6px;
  font-weight: 500;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .apple-header .header-content {
    padding: var(--apple-space) var(--apple-space);
    min-height: auto;
    flex-direction: column;
    align-items: flex-start;
    gap: var(--apple-space);
  }
  
  .header-left {
    width: 100%;
  }
  
  .app-title {
    justify-content: space-between;
    width: 100%;
  }
  
  .title-info h2 {
    font-size: var(--apple-font-size-lg);
  }
  
  .header-right {
    width: 100%;
    justify-content: flex-end;
    flex-wrap: wrap;
    gap: var(--apple-space-xs);
  }
  
  .header-btn {
    height: 44px; /* 更好的触摸目标 */
    padding: var(--apple-space-sm) var(--apple-space);
    font-size: var(--apple-font-size-sm);
    min-width: 120px;
    justify-content: center;
  }
  
  .main-content {
    padding: 16px;
  }
  
  .apple-toolbar {
    flex-direction: column;
    align-items: stretch;
    gap: var(--apple-space);
    padding: var(--apple-space);
  }
  
  .toolbar-section {
    flex-wrap: nowrap;
    justify-content: flex-start;
    gap: var(--apple-space-sm);
    overflow-x: auto; /* 如果内容过多可以滚动 */
  }
  
  .toolbar-section:first-child {
    justify-content: flex-start;
  }
  
  .toolbar-section:last-child {
    justify-content: flex-end;
  }
  
  .toolbar-btn {
    height: 44px; /* 更好的触摸目标 */
    min-width: 110px; /* 减小宽度以适应同一行 */
    flex-shrink: 0;
    justify-content: center;
  }
  
  .apple-select {
    min-width: 110px;
    height: 44px;
    text-align: center;
    flex-shrink: 0;
  }
  
  .view-toggle {
    height: 44px;
    flex-shrink: 0;
  }
  
  .toggle-button {
    width: 44px;
    height: 40px;
  }
  
  .macos-breadcrumb {
    padding: var(--apple-space-sm) var(--apple-space);
    margin-bottom: var(--apple-space);
  }
  
  .breadcrumb-path {
    flex-wrap: wrap;
    gap: var(--apple-space-xs);
  }
  
  .breadcrumb-item {
    padding: var(--apple-space-xs) var(--apple-space-sm);
    font-size: var(--apple-font-size-sm);
    min-height: 36px;
    display: flex;
    align-items: center;
  }
  
  .macos-files-grid {
    grid-template-columns: repeat(auto-fill, minmax(160px, 1fr));
    gap: 12px;
  }
  
  .macos-file-card {
    padding: 12px;
  }
  
  .card-preview {
    height: 100px;
  }
  
  /* 改善表格在移动端的显示 */
  .el-table {
    font-size: var(--apple-font-size-sm);
  }
  
  .el-table th,
  .el-table td {
    padding: 8px 4px;
  }
}

@media (max-width: 480px) {
  .apple-header {
    position: sticky;
    top: 0;
    z-index: 1000;
  }
  
  .apple-header .header-content {
    padding: var(--apple-space-sm) var(--apple-space);
    flex-direction: column;
    gap: var(--apple-space-sm);
  }
  
  .app-title {
    flex-direction: row;
    align-items: center;
    gap: var(--apple-space-sm);
    width: 100%;
    justify-content: flex-start;
  }
  
  .app-icon {
    flex-shrink: 0;
  }
  
  .app-icon svg {
    width: 28px;
    height: 28px;
  }
  
  .title-info h2 {
    font-size: var(--apple-font-size);
    margin: 0;
  }
  
  .user-info {
    font-size: var(--apple-font-size-xs);
  }
  
  .header-right {
    width: 100%;
    justify-content: space-between;
    gap: var(--apple-space-xs);
  }
  
  .header-btn {
    flex: 1;
    height: 40px;
    padding: var(--apple-space-xs) var(--apple-space-sm);
    font-size: var(--apple-font-size-xs);
    min-width: 0;
  }
  
  .header-btn span {
    display: none; /* 隐藏文字，只显示图标 */
  }
  
  .header-btn svg {
    margin: 0;
  }
  
  /* 为管理员按钮和退出按钮添加工具提示替代 */
  .header-btn[title]:before {
    content: attr(title);
    position: absolute;
    bottom: -30px;
    left: 50%;
    transform: translateX(-50%);
    background: var(--apple-text-primary);
    color: var(--apple-background);
    padding: 4px 8px;
    border-radius: 4px;
    font-size: 12px;
    white-space: nowrap;
    opacity: 0;
    pointer-events: none;
    transition: opacity 0.2s;
    z-index: 1001;
  }
  
  .header-btn:hover:before {
    opacity: 1;
  }
  
  .main-content {
    padding: var(--apple-space-sm);
  }
  
  .macos-breadcrumb {
    padding: var(--apple-space-xs) var(--apple-space-sm);
    margin-bottom: var(--apple-space-sm);
  }
  
  .breadcrumb-item {
    padding: var(--apple-space-xs);
    font-size: var(--apple-font-size-xs);
    min-height: 32px;
  }
  
  .breadcrumb-item span {
    max-width: 80px;
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
  }
  
  .apple-toolbar {
    padding: var(--apple-space-sm);
    gap: var(--apple-space-sm);
  }
  
  .toolbar-section {
    gap: var(--apple-space-xs);
  }
  
  .toolbar-btn {
    height: 40px;
    min-width: 100px;
    font-size: var(--apple-font-size-xs);
    padding: var(--apple-space-xs) var(--apple-space-sm);
  }
  
  .apple-select {
    min-width: 100px;
    height: 40px;
    font-size: var(--apple-font-size-xs);
  }
  
  .view-toggle {
    height: 40px;
  }
  
  .toggle-button {
    width: 40px;
    height: 36px;
  }
  
  .macos-files-grid {
    grid-template-columns: repeat(auto-fill, minmax(140px, 1fr));
    gap: var(--apple-space-xs);
  }
  
  .macos-file-card {
    padding: var(--apple-space-xs);
    border-radius: var(--apple-radius);
  }
  
  .card-preview {
    height: 90px;
    margin-bottom: var(--apple-space-xs);
  }
  
  .folder-icon,
  .file-icon {
    font-size: 36px;
  }
  
  .card-title {
    font-size: var(--apple-font-size-xs);
    line-height: 1.2;
  }
  
  .card-meta {
    font-size: 10px;
    margin-top: 2px;
  }
  
  .action-button {
    width: 24px;
    height: 24px;
  }
  
  /* 表格在小屏幕上的优化 */
  .el-table .el-table__header-wrapper {
    display: none;
  }
  
  .el-table .el-table__body-wrapper {
    overflow-x: auto;
  }
  
  .el-table td {
    padding: 4px 2px;
    font-size: var(--apple-font-size-xs);
  }
  
  /* 分页组件移动端优化 */
  .macos-pagination {
    margin-top: var(--apple-space);
    padding: var(--apple-space-sm);
  }
  
  .macos-pagination .el-pagination {
    flex-wrap: wrap;
    justify-content: center;
    gap: var(--apple-space-xs);
  }
  
  .macos-pagination .el-pagination .el-pager li {
    min-width: 32px;
    height: 32px;
    line-height: 32px;
    font-size: var(--apple-font-size-xs);
  }
  
  .macos-pagination .el-pagination .btn-prev,
  .macos-pagination .el-pagination .btn-next {
    min-width: 32px;
    height: 32px;
  }
  
  /* 对话框移动端适配 */
  :deep(.el-dialog) {
    width: 95% !important;
    margin: 5vh auto !important;
    max-height: 90vh;
    border-radius: var(--apple-radius) !important;
  }
  
  :deep(.el-dialog__body) {
    max-height: 60vh;
    overflow-y: auto;
    padding: var(--apple-space) !important;
  }
  
  :deep(.el-dialog__header) {
    padding: var(--apple-space) var(--apple-space) var(--apple-space-sm) !important;
  }
  
  :deep(.el-dialog__footer) {
    padding: var(--apple-space-sm) var(--apple-space) var(--apple-space) !important;
  }
  
  /* 上传组件移动端优化 */
  .upload-wrapper {
    width: 100%;
  }
  
  .upload-wrapper .toolbar-btn {
    width: 100%;
  }
  
  /* 空状态移动端优化 */
  .macos-empty-state {
    height: 150px;
    padding: var(--apple-space);
  }
  
  .empty-icon {
    font-size: 48px;
    margin-bottom: var(--apple-space-sm);
  }
  
  .empty-title {
    font-size: var(--apple-font-size);
    margin-bottom: var(--apple-space-xs);
  }
  
  .empty-description {
    font-size: var(--apple-font-size-xs);
  }
}

/* 深色模式支持 */
@media (prefers-color-scheme: dark) {
  :root {
    --apple-background: #1C1C1E;
    --apple-surface: #000000;
    --apple-border: rgba(255, 255, 255, 0.1);
    --apple-shadow: rgba(0, 0, 0, 0.3);
    --apple-shadow-light: rgba(0, 0, 0, 0.15);
    --apple-text-primary: #FFFFFF;
    --apple-text-secondary: #8E8E93;
    --apple-gray-1: #2C2C2E;
    --apple-gray-2: #3A3A3C;
    --apple-gray-3: #48484A;
  }
}
</style>