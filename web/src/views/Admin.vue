<template>
  <div class="admin-container">
    <el-header class="header">
      <div class="header-left">
        <h2>管理员面板</h2>
      </div>
      <div class="header-right">
        <el-button @click="$router.push('/files')">返回文件管理</el-button>
        <el-button @click="handleLogout">退出登录</el-button>
      </div>
    </el-header>

    <el-main class="main">
      <el-row :gutter="20" style="margin-bottom: 20px">
        <el-col :xs="12" :sm="12" :md="6" :lg="6" :xl="6">
          <el-card>
            <div class="stat-card">
              <div class="stat-icon">
                <el-icon size="24" color="#409eff"><User /></el-icon>
              </div>
              <div class="stat-content">
                <div class="stat-title">用户数量</div>
                <div class="stat-value">{{ stats.user_count }}</div>
              </div>
            </div>
          </el-card>
        </el-col>
        <el-col :xs="12" :sm="12" :md="6" :lg="6" :xl="6">
          <el-card>
            <div class="stat-card">
              <div class="stat-icon">
                <el-icon size="24" color="#67c23a"><Document /></el-icon>
              </div>
              <div class="stat-content">
                <div class="stat-title">文件数量</div>
                <div class="stat-value">{{ stats.file_count }}</div>
              </div>
            </div>
          </el-card>
        </el-col>
        <el-col :xs="12" :sm="12" :md="6" :lg="6" :xl="6">
          <el-card>
            <div class="stat-card">
              <div class="stat-icon">
                <el-icon size="24" color="#e6a23c"><Folder /></el-icon>
              </div>
              <div class="stat-content">
                <div class="stat-title">存储使用</div>
                <div class="stat-value">{{ formatFileSize(stats.total_storage) }}</div>
              </div>
            </div>
          </el-card>
        </el-col>
        <el-col :xs="12" :sm="12" :md="6" :lg="6" :xl="6">
          <el-card>
            <div class="stat-card">
              <div class="stat-icon">
                <el-icon size="24" color="#f56c6c"><TrendCharts /></el-icon>
              </div>
              <div class="stat-content">
                <div class="stat-title">平均文件大小</div>
                <div class="stat-value">
                  {{ stats.file_count > 0 ? formatFileSize(stats.total_storage / stats.file_count) : '0 B' }}
                </div>
              </div>
            </div>
          </el-card>
        </el-col>
      </el-row>

      <el-card>
        <template #header>
          <div class="card-header">
            <span>用户管理</span>
            <el-button type="primary" @click="refreshUsers">
              <el-icon><Refresh /></el-icon>
              刷新
            </el-button>
          </div>
        </template>

        <div class="table-container">
          <el-table 
            v-loading="loading" 
            :data="users" 
            style="width: 100%" 
            class="responsive-table"
          >
            <el-table-column prop="id" label="ID" width="80" class-name="mobile-hide" />
            
            <el-table-column prop="username" label="用户名" min-width="120" />
            
            <el-table-column prop="email" label="邮箱" min-width="180" class-name="mobile-hide" />
            
            <el-table-column label="角色" width="100">
              <template #default="scope">
                <el-tag :type="scope.row.is_admin ? 'danger' : 'primary'" size="small">
                  {{ scope.row.is_admin ? '管理员' : '普通用户' }}
                </el-tag>
              </template>
            </el-table-column>
            
            <el-table-column label="存储配额" width="120" class-name="mobile-hide">
              <template #default="scope">
                {{ formatFileSize(scope.row.storage_quota) }}
              </template>
            </el-table-column>
            
            <el-table-column label="已使用" width="120" class-name="mobile-hide">
              <template #default="scope">
                {{ formatFileSize(scope.row.storage_used) }}
              </template>
            </el-table-column>
            
            <el-table-column label="使用率" min-width="100">
              <template #default="scope">
                <el-progress 
                  :percentage="Math.round(scope.row.storage_used / scope.row.storage_quota * 100)"
                  :color="getProgressColor(scope.row.storage_used / scope.row.storage_quota)"
                  :show-text="false"
                  class="mobile-progress"
                />
                <div class="mobile-storage-info">
                  <span>{{ formatFileSize(scope.row.storage_used) }} / {{ formatFileSize(scope.row.storage_quota) }}</span>
                </div>
              </template>
            </el-table-column>
            
            <el-table-column prop="created_at" label="注册时间" width="150" class-name="mobile-hide">
              <template #default="scope">
                {{ formatDate(scope.row.created_at) }}
              </template>
            </el-table-column>
            
            <el-table-column label="操作" fixed="right" min-width="120">
              <template #default="scope">
                <div class="action-buttons">
                  <el-button size="small" @click="editUser(scope.row)">
                    编辑
                  </el-button>
                  <el-button
                    v-if="scope.row.id !== authStore.user?.id"
                    size="small"
                    type="danger"
                    @click="deleteUser(scope.row)"
                  >
                    删除
                  </el-button>
                </div>
              </template>
            </el-table-column>
          </el-table>
        </div>

        <el-pagination
          v-if="total > 0"
          v-model:current-page="currentPage"
          v-model:page-size="pageSize"
          :total="total"
          :page-sizes="[10, 20, 50, 100]"
          :layout="paginationLayout"
          @size-change="fetchUsers"
          @current-change="fetchUsers"
          class="pagination-responsive"
        />
      </el-card>
    </el-main>

    <el-dialog
      v-model="editDialogVisible"
      title="编辑用户"
      :width="dialogWidth"
      class="responsive-dialog"
    >
      <el-form
        ref="editFormRef"
        :model="editForm"
        :rules="editRules"
        label-width="100px"
      >
        <el-form-item label="用户名">
          <el-input v-model="editForm.username" disabled />
        </el-form-item>
        
        <el-form-item label="邮箱">
          <el-input v-model="editForm.email" disabled />
        </el-form-item>
        
        <el-form-item label="管理员" prop="is_admin">
          <el-switch v-model="editForm.is_admin" />
        </el-form-item>
        
        <el-form-item label="存储配额" prop="storage_quota">
          <el-input-number
            v-model="editForm.storage_quota"
            :min="1024 * 1024"
            :step="1024 * 1024 * 100"
            style="width: 100%"
          />
          <div class="form-help">
            {{ formatFileSize(editForm.storage_quota || 0) }}
          </div>
        </el-form-item>
      </el-form>
      
      <template #footer>
        <el-button @click="editDialogVisible = false">取消</el-button>
        <el-button type="primary" @click="saveUser">确定</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script>
import { ref, reactive, onMounted, computed, onUnmounted } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import { ElMessage, ElMessageBox } from 'element-plus'
import {
  User,
  Document,
  Folder,
  TrendCharts,
  Refresh
} from '@element-plus/icons-vue'
import request from '@/utils/request'

export default {
  name: 'Admin',
  setup() {
    const router = useRouter()
    const authStore = useAuthStore()
    
    const loading = ref(false)
    const isMobile = ref(false)
    const users = ref([])
    const total = ref(0)
    const currentPage = ref(1)
    const pageSize = ref(20)
    const stats = ref({
      user_count: 0,
      file_count: 0,
      total_storage: 0
    })
    
    const editDialogVisible = ref(false)
    const editFormRef = ref()
    const editForm = reactive({
      id: 0,
      username: '',
      email: '',
      is_admin: false,
      storage_quota: 0
    })
    
    const editRules = {
      storage_quota: [
        { required: true, message: '请设置存储配额', trigger: 'blur' },
        { type: 'number', min: 1024 * 1024, message: '配额不能小于1MB', trigger: 'blur' }
      ]
    }
    
    const fetchStats = async () => {
      try {
        const response = await request.get('/admin/stats')
        stats.value = response
      } catch (error) {
        console.error('Failed to fetch stats:', error)
      }
    }
    
    const fetchUsers = async () => {
      loading.value = true
      try {
        const params = {
          page: currentPage.value,
          limit: pageSize.value
        }
        
        const response = await request.get('/admin/users', { params })
        users.value = response.users
        total.value = response.total
      } catch (error) {
        console.error('Failed to fetch users:', error)
      } finally {
        loading.value = false
      }
    }
    
    const refreshUsers = () => {
      fetchStats()
      fetchUsers()
    }
    
    const editUser = (user) => {
      Object.assign(editForm, {
        id: user.id,
        username: user.username,
        email: user.email,
        is_admin: user.is_admin,
        storage_quota: user.storage_quota
      })
      editDialogVisible.value = true
    }
    
    const saveUser = async () => {
      if (!editFormRef.value) return
      
      await editFormRef.value.validate(async (valid) => {
        if (!valid) return
        
        try {
          const updateData = {
            is_admin: editForm.is_admin,
            storage_quota: editForm.storage_quota
          }
          
          await request.put(`/admin/users/${editForm.id}`, updateData)
          ElMessage.success('用户更新成功')
          editDialogVisible.value = false
          fetchUsers()
        } catch (error) {
          console.error('Failed to update user:', error)
        }
      })
    }
    
    const deleteUser = async (user) => {
      try {
        await ElMessageBox.confirm(
          `确定要删除用户 "${user.username}" 吗？这将同时删除该用户的所有文件，此操作不可恢复。`,
          '删除确认',
          {
            confirmButtonText: '确定删除',
            cancelButtonText: '取消',
            type: 'error'
          }
        )
        
        await request.delete(`/admin/users/${user.id}`)
        ElMessage.success('用户删除成功')
        fetchUsers()
        fetchStats()
      } catch (error) {
        if (error !== 'cancel') {
          console.error('Delete user error:', error)
        }
      }
    }
    
    const handleLogout = () => {
      authStore.logout()
      router.push('/login')
    }
    
    const formatFileSize = (bytes) => {
      if (bytes === 0) return '0 B'
      const k = 1024
      const sizes = ['B', 'KB', 'MB', 'GB', 'TB']
      const i = Math.floor(Math.log(bytes) / Math.log(k))
      return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i]
    }
    
    const formatDate = (dateString) => {
      return new Date(dateString).toLocaleString('zh-CN')
    }
    
    const getProgressColor = (ratio) => {
      if (ratio < 0.6) return '#67c23a'
      if (ratio < 0.8) return '#e6a23c'
      return '#f56c6c'
    }
    
    // 响应式相关
    const checkIsMobile = () => {
      isMobile.value = window.innerWidth <= 768
    }
    
    const paginationLayout = computed(() => {
      return isMobile.value ? 'prev, pager, next' : 'total, sizes, prev, pager, next, jumper'
    })
    
    const dialogWidth = computed(() => {
      return isMobile.value ? '90%' : '400px'
    })
    
    onMounted(() => {
      checkIsMobile()
      window.addEventListener('resize', checkIsMobile)
      fetchStats()
      fetchUsers()
    })
    
    onUnmounted(() => {
      window.removeEventListener('resize', checkIsMobile)
    })
    
    return {
      authStore,
      loading,
      isMobile,
      users,
      total,
      currentPage,
      pageSize,
      stats,
      editDialogVisible,
      editFormRef,
      editForm,
      editRules,
      fetchUsers,
      refreshUsers,
      editUser,
      saveUser,
      deleteUser,
      handleLogout,
      formatFileSize,
      formatDate,
      getProgressColor,
      paginationLayout,
      dialogWidth,
      User,
      Document,
      Folder,
      TrendCharts,
      Refresh
    }
  }
}
</script>

<style scoped>
.admin-container {
  height: 100vh;
  display: flex;
  flex-direction: column;
}

.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 0 20px;
  border-bottom: 1px solid #ebeef5;
  background: white;
  min-height: 60px;
}

.header-left h2 {
  margin: 0;
  color: #303133;
  font-size: 18px;
}

.header-right {
  display: flex;
  gap: 8px;
}

.main {
  flex: 1;
  padding: 20px;
  overflow: auto;
}

.stat-card {
  display: flex;
  align-items: center;
  gap: 12px;
}

.stat-icon {
  width: 48px;
  height: 48px;
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
  background: #f5f7fa;
}

.stat-content {
  flex: 1;
}

.stat-title {
  font-size: 14px;
  color: #909399;
  margin-bottom: 4px;
}

.stat-value {
  font-size: 24px;
  font-weight: 600;
  color: #303133;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.form-help {
  font-size: 12px;
  color: #909399;
  margin-top: 4px;
}

/* 表格容器 */
.table-container {
  overflow-x: auto;
}

/* 响应式表格 */
.responsive-table {
  min-width: 600px;
}

/* 移动端进度条 */
.mobile-progress {
  margin-bottom: 4px;
}

.mobile-storage-info {
  font-size: 12px;
  color: #909399;
  text-align: center;
}

/* 操作按钮 */
.action-buttons {
  display: flex;
  gap: 4px;
  flex-wrap: wrap;
}

/* 分页响应式 */
.pagination-responsive {
  margin-top: 20px;
  text-align: center;
}

/* 移动端适配 */
@media (max-width: 768px) {
  .header {
    padding: 0 12px;
    flex-wrap: wrap;
    gap: 8px;
  }
  
  .header-left h2 {
    font-size: 16px;
  }
  
  .header-right {
    gap: 4px;
  }
  
  .header-right .el-button {
    padding: 8px 12px;
    font-size: 12px;
  }
  
  .main {
    padding: 12px;
  }
  
  /* 统计卡片移动端优化 */
  .stat-card {
    gap: 8px;
  }
  
  .stat-icon {
    width: 40px;
    height: 40px;
  }
  
  .stat-title {
    font-size: 13px;
  }
  
  .stat-value {
    font-size: 20px;
  }
  
  /* 表格移动端优化 */
  .table-container {
    overflow-x: auto;
    -webkit-overflow-scrolling: touch;
  }
  
  .responsive-table {
    min-width: 480px;
  }
  
  /* 隐藏移动端不重要的列 */
  :deep(.mobile-hide) {
    display: none !important;
  }
  
  /* 移动端操作按钮 */
  .action-buttons {
    flex-direction: column;
    gap: 2px;
  }
  
  .action-buttons .el-button {
    padding: 4px 8px;
    font-size: 12px;
    width: 100%;
  }
  
  /* 移动端分页 */
  .pagination-responsive {
    margin-top: 16px;
  }
  
  :deep(.el-pagination) {
    justify-content: center;
  }
  
  :deep(.el-pager li) {
    min-width: 32px;
    height: 32px;
    line-height: 32px;
    margin: 0 2px;
  }
  
  /* 移动端对话框 */
  .responsive-dialog {
    margin: 0 !important;
  }
  
  :deep(.el-dialog) {
    margin-top: 5vh !important;
    margin-bottom: 5vh !important;
  }
  
  :deep(.el-form-item__label) {
    font-size: 14px;
  }
  
  /* 移动端卡片间距 */
  :deep(.el-row) {
    margin-left: -8px !important;
    margin-right: -8px !important;
  }
  
  :deep(.el-col) {
    padding-left: 8px !important;
    padding-right: 8px !important;
    margin-bottom: 12px;
  }
}

/* 平板适配 */
@media (min-width: 769px) and (max-width: 1024px) {
  .main {
    padding: 16px;
  }
  
  .stat-value {
    font-size: 22px;
  }
  
  .responsive-table {
    min-width: 700px;
  }
}

/* 小屏幕移动端 */
@media (max-width: 480px) {
  .header {
    padding: 0 8px;
    min-height: 50px;
  }
  
  .header-left h2 {
    font-size: 14px;
  }
  
  .header-right .el-button {
    padding: 6px 8px;
    font-size: 11px;
  }
  
  .main {
    padding: 8px;
  }
  
  .stat-icon {
    width: 36px;
    height: 36px;
  }
  
  .stat-title {
    font-size: 12px;
  }
  
  .stat-value {
    font-size: 18px;
  }
  
  .responsive-table {
    min-width: 400px;
  }
  
  :deep(.el-table th) {
    padding: 8px 4px !important;
    font-size: 12px;
  }
  
  :deep(.el-table td) {
    padding: 8px 4px !important;
    font-size: 12px;
  }
  
  .action-buttons .el-button {
    padding: 2px 6px;
    font-size: 11px;
  }
}

/* 进度条在移动端的优化 */
@media (max-width: 768px) {
  .mobile-storage-info {
    display: block;
  }
}

@media (min-width: 769px) {
  .mobile-storage-info {
    display: none;
  }
}
</style>