<template>
  <div class="register-container">
    <div class="register-hero">
      <div class="hero-content">
        <div class="logo-section">
          <div class="logo-icon">
            <svg width="64" height="64" viewBox="0 0 64 64" fill="none">
              <rect width="64" height="64" rx="16" fill="url(#gradient)"/>
              <path d="M32 16L48 28V48C48 50.2091 46.2091 52 44 52H20C17.7909 52 16 50.2091 16 48V28L32 16Z" fill="white" fill-opacity="0.9"/>
              <path d="M32 16L16 28H32V44L48 28L32 16Z" fill="white" fill-opacity="0.6"/>
              <defs>
                <linearGradient id="gradient" x1="0%" y1="0%" x2="100%" y2="100%">
                  <stop offset="0%" style="stop-color:#34C759"/>
                  <stop offset="100%" style="stop-color:#007AFF"/>
                </linearGradient>
              </defs>
            </svg>
          </div>
          <h1 class="apple-text-display">开始使用</h1>
          <p class="apple-text-subtitle" style="color: rgba(255,255,255,0.8); margin-top: 8px;">
            创建您的西瓜文件服务器账户
          </p>
        </div>
      </div>
    </div>
    
    <div class="register-form-section">
      <div class="register-card apple-card">
        <div class="card-header">
          <h2 class="apple-text-title">注册</h2>
          <p class="apple-text-body" style="color: var(--apple-text-secondary); margin-top: 8px;">
            填写以下信息来创建您的账户
          </p>
        </div>
        
        <el-form
          ref="registerFormRef"
          :model="registerForm"
          :rules="rules"
          class="register-form"
          @submit.prevent="handleRegister"
        >
          <el-form-item prop="username">
            <div class="apple-input full-width">
              <el-input
                v-model="registerForm.username"
                placeholder="用户名"
                size="large"
                clearable
              />
            </div>
          </el-form-item>
          
          <el-form-item prop="email">
            <div class="apple-input full-width">
              <el-input
                v-model="registerForm.email"
                placeholder="邮箱地址"
                size="large"
                clearable
              />
            </div>
          </el-form-item>
          
          <el-form-item prop="password">
            <div class="apple-input full-width">
              <el-input
                v-model="registerForm.password"
                type="password"
                placeholder="密码"
                size="large"
                show-password
                clearable
              />
            </div>
          </el-form-item>
          
          <el-form-item prop="confirmPassword">
            <div class="apple-input full-width">
              <el-input
                v-model="registerForm.confirmPassword"
                type="password"
                placeholder="确认密码"
                size="large"
                show-password
                clearable
                @keyup.enter="handleRegister"
              />
            </div>
          </el-form-item>
          
          <el-form-item>
            <el-button
              :loading="loading"
              class="apple-button-primary register-button"
              @click="handleRegister"
            >
              <span v-if="!loading">创建账户</span>
              <span v-else>创建中...</span>
            </el-button>
          </el-form-item>
        </el-form>
        
        <div class="register-footer">
          <div class="divider">
            <span class="divider-text apple-text-caption">或</span>
          </div>
          <div class="login-link">
            <span class="apple-text-body" style="color: var(--apple-text-secondary);">
              已有账户？
            </span>
            <button 
              class="link-button apple-text-body"
              @click="$router.push('/login')"
            >
              立即登录
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { ref, reactive } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import { ElMessage } from 'element-plus'
import { User, Lock, Message } from '@element-plus/icons-vue'

export default {
  name: 'Register',
  setup() {
    const router = useRouter()
    const authStore = useAuthStore()
    
    const registerFormRef = ref()
    const loading = ref(false)
    
    const registerForm = reactive({
      username: '',
      email: '',
      password: '',
      confirmPassword: ''
    })
    
    const validateConfirmPassword = (rule, value, callback) => {
      if (value !== registerForm.password) {
        callback(new Error('两次输入的密码不一致'))
      } else {
        callback()
      }
    }
    
    const rules = {
      username: [
        { required: true, message: '请输入用户名', trigger: 'blur' },
        { min: 3, max: 20, message: '用户名长度应在3-20个字符', trigger: 'blur' }
      ],
      email: [
        { required: true, message: '请输入邮箱地址', trigger: 'blur' },
        { type: 'email', message: '请输入正确的邮箱地址', trigger: 'blur' }
      ],
      password: [
        { required: true, message: '请输入密码', trigger: 'blur' },
        { min: 6, message: '密码长度不能少于6个字符', trigger: 'blur' }
      ],
      confirmPassword: [
        { required: true, message: '请确认密码', trigger: 'blur' },
        { validator: validateConfirmPassword, trigger: 'blur' }
      ]
    }
    
    const handleRegister = async () => {
      if (!registerFormRef.value) return
      
      await registerFormRef.value.validate(async (valid) => {
        if (!valid) return
        
        loading.value = true
        
        try {
          await authStore.register(
            registerForm.username,
            registerForm.password,
            registerForm.email
          )
          ElMessage.success('注册成功')
          router.push('/files')
        } catch (error) {
          console.error('Register failed:', error)
        } finally {
          loading.value = false
        }
      })
    }
    
    return {
      registerFormRef,
      registerForm,
      rules,
      loading,
      handleRegister,
      User,
      Lock,
      Message
    }
  }
}
</script>

<style scoped>
.register-container {
  min-height: 100vh;
  background: var(--apple-background-secondary);
  display: flex;
  flex-direction: column;
}

.register-hero {
  flex: 1;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: var(--apple-space-2xl) var(--apple-space-lg);
  background: linear-gradient(135deg, 
    var(--apple-green) 0%, 
    var(--apple-blue) 50%, 
    var(--apple-purple) 100%);
  position: relative;
  overflow: hidden;
}

.register-hero::before {
  content: '';
  position: absolute;
  top: -50%;
  right: -50%;
  width: 200%;
  height: 200%;
  background: radial-gradient(circle, rgba(255,255,255,0.1) 0%, transparent 70%);
  animation: floatReverse 8s ease-in-out infinite;
}

@keyframes floatReverse {
  0%, 100% { transform: translate(0, 0) rotate(0deg); }
  33% { transform: translate(-30px, 30px) rotate(-120deg); }
  66% { transform: translate(20px, -20px) rotate(-240deg); }
}

.hero-content {
  text-align: center;
  z-index: 1;
  position: relative;
}

.logo-section {
  animation: fadeInUp 0.8s ease-out;
}

@keyframes fadeInUp {
  from {
    opacity: 0;
    transform: translateY(30px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.logo-icon {
  margin-bottom: var(--apple-space);
  display: inline-block;
  animation: pulse 3s ease-in-out infinite;
}

@keyframes pulse {
  0%, 100% { transform: scale(1); }
  50% { transform: scale(1.05); }
}

.logo-section h1 {
  color: white;
  margin: 0;
  margin-bottom: var(--apple-space-sm);
}

.register-form-section {
  flex: 1;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: var(--apple-space-2xl) var(--apple-space-lg);
}

.register-card {
  width: 100%;
  max-width: 420px;
  animation: slideInUp 0.6s ease-out 0.3s both;
}

@keyframes slideInUp {
  from {
    opacity: 0;
    transform: translateY(50px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.card-header {
  text-align: center;
  margin-bottom: var(--apple-space-lg);
}

.card-header h2 {
  margin: 0;
  color: var(--apple-text-primary);
}

.register-form {
  margin-bottom: var(--apple-space-lg);
}

.register-form .el-form-item {
  margin-bottom: var(--apple-space);
}

.full-width {
  width: 100%;
}

.full-width .el-input {
  width: 100%;
}

.register-button {
  width: 100%;
  height: 52px;
  font-size: var(--apple-font-size-lg);
  font-weight: var(--apple-font-weight-semibold);
}

.register-footer {
  margin-top: var(--apple-space-lg);
}

.divider {
  position: relative;
  text-align: center;
  margin: var(--apple-space-lg) 0;
}

.divider::before {
  content: '';
  position: absolute;
  top: 50%;
  left: 0;
  right: 0;
  height: 1px;
  background: var(--apple-separator);
}

.divider-text {
  background: var(--apple-background-elevated);
  padding: 0 var(--apple-space);
  color: var(--apple-text-tertiary);
  position: relative;
  z-index: 1;
}

.login-link {
  text-align: center;
  margin-top: var(--apple-space);
}

.link-button {
  background: none;
  border: none;
  color: var(--apple-blue);
  cursor: pointer;
  text-decoration: none;
  font-weight: var(--apple-font-weight-medium);
  margin-left: var(--apple-space-xs);
  transition: all 0.2s ease;
  padding: 0;
}

.link-button:hover {
  color: var(--apple-blue-dark);
  text-decoration: underline;
}

.link-button:focus {
  outline: 2px solid var(--apple-blue);
  outline-offset: 2px;
  border-radius: var(--apple-radius-sm);
}

/* 移动端适配 */
@media (max-width: 768px) {
  .register-hero {
    padding: var(--apple-space-xl) var(--apple-space);
  }
  
  .register-form-section {
    padding: var(--apple-space-xl) var(--apple-space);
  }
  
  .logo-section h1 {
    font-size: var(--apple-font-size-3xl);
  }
  
  .register-card {
    max-width: 100%;
  }
}

@media (max-width: 480px) {
  .register-container {
    flex-direction: column;
  }
  
  .register-hero {
    flex: 0.5;
    padding: var(--apple-space-lg) var(--apple-space);
  }
  
  .register-form-section {
    flex: 1.2;
    padding: var(--apple-space-lg) var(--apple-space);
  }
  
  .logo-section h1 {
    font-size: var(--apple-font-size-2xl);
  }
  
  .register-button {
    height: 48px;
  }
}

/* 微交互 */
.apple-input .el-input__inner:focus {
  transform: translateY(-1px);
}

.register-button:not(:disabled):hover {
  transform: translateY(-2px);
}

.register-button:not(:disabled):active {
  transform: translateY(0);
}

/* 表单验证错误样式优化 */
.el-form-item.is-error .apple-input .el-input__inner {
  border-color: var(--apple-red);
  box-shadow: 0 0 0 3px rgba(255, 59, 48, 0.1);
}

.el-form-item__error {
  color: var(--apple-red);
  font-size: var(--apple-font-size-sm);
  font-weight: var(--apple-font-weight-medium);
  margin-top: var(--apple-space-xs);
}
</style>