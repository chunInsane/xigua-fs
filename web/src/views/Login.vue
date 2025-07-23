<template>
  <div class="login-container">
    <div class="login-hero">
      <div class="hero-content">
        <div class="logo-section">
          <div class="logo-icon">
            <svg width="64" height="64" viewBox="0 0 64 64" fill="none">
              <rect width="64" height="64" rx="16" fill="url(#gradient)"/>
              <path d="M32 16L48 28V48C48 50.2091 46.2091 52 44 52H20C17.7909 52 16 50.2091 16 48V28L32 16Z" fill="white" fill-opacity="0.9"/>
              <path d="M32 16L16 28H32V44L48 28L32 16Z" fill="white" fill-opacity="0.6"/>
              <defs>
                <linearGradient id="gradient" x1="0%" y1="0%" x2="100%" y2="100%">
                  <stop offset="0%" style="stop-color:#007AFF"/>
                  <stop offset="100%" style="stop-color:#5AC8FA"/>
                </linearGradient>
              </defs>
            </svg>
          </div>
          <h1 class="apple-text-display">🍉文件服务</h1>
          <p class="apple-text-subtitle" style="color: var(--apple-text-secondary); margin-top: 8px;">
            让文件管理变得简单高效
          </p>
        </div>
      </div>
    </div>
    
    <div class="login-form-section">
      <div class="login-card apple-card">
        <div class="card-header">
          <h2 class="apple-text-title">登录</h2>
          <p class="apple-text-body" style="color: var(--apple-text-secondary); margin-top: 8px;">
            使用您的账户信息登录
          </p>
        </div>
        
        <el-form
          ref="loginFormRef"
          :model="loginForm"
          :rules="rules"
          class="login-form"
          @submit.prevent="handleLogin"
        >
          <el-form-item prop="username">
            <div class="apple-input full-width">
              <el-input
                v-model="loginForm.username"
                placeholder="用户名"
                size="large"
                clearable
              />
            </div>
          </el-form-item>
          
          <el-form-item prop="password">
            <div class="apple-input full-width">
              <el-input
                v-model="loginForm.password"
                type="password"
                placeholder="密码"
                size="large"
                show-password
                clearable
                @keyup.enter="handleLogin"
              />
            </div>
          </el-form-item>
          
          <el-form-item>
            <el-button
              :loading="loading"
              class="apple-button-primary login-button"
              @click="handleLogin"
            >
              <span v-if="!loading">登录</span>
              <span v-else>登录中...</span>
            </el-button>
          </el-form-item>
        </el-form>
        
        <div class="login-footer">
          <div class="divider">
            <span class="divider-text apple-text-caption">或</span>
          </div>
          <div class="register-link">
            <span class="apple-text-body" style="color: var(--apple-text-secondary);">
              还没有账户？
            </span>
            <button 
              class="link-button apple-text-body"
              @click="$router.push('/register')"
            >
              立即注册
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
import { User, Lock } from '@element-plus/icons-vue'

export default {
  name: 'Login',
  setup() {
    const router = useRouter()
    const authStore = useAuthStore()
    
    const loginFormRef = ref()
    const loading = ref(false)
    
    const loginForm = reactive({
      username: '',
      password: ''
    })
    
    const rules = {
      username: [
        { required: true, message: '请输入用户名', trigger: 'blur' }
      ],
      password: [
        { required: true, message: '请输入密码', trigger: 'blur' }
      ]
    }
    
    const handleLogin = async () => {
      if (!loginFormRef.value) return
      
      await loginFormRef.value.validate(async (valid) => {
        if (!valid) return
        
        loading.value = true
        
        try {
          await authStore.login(loginForm.username, loginForm.password)
          ElMessage.success('登录成功')
          router.push('/files')
        } catch (error) {
          console.error('Login failed:', error)
        } finally {
          loading.value = false
        }
      })
    }
    
    return {
      loginFormRef,
      loginForm,
      rules,
      loading,
      handleLogin,
      User,
      Lock
    }
  }
}
</script>

<style scoped>
.login-container {
  min-height: 100vh;
  background: var(--apple-background-secondary);
  display: flex;
  flex-direction: column;
}

.login-hero {
  flex: 1;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: var(--apple-space-2xl) var(--apple-space-lg);
  background: linear-gradient(135deg, 
    var(--apple-blue) 0%, 
    var(--apple-blue-light) 50%, 
    var(--apple-purple) 100%);
  position: relative;
  overflow: hidden;
}

.login-hero::before {
  content: '';
  position: absolute;
  top: -50%;
  left: -50%;
  width: 200%;
  height: 200%;
  background: radial-gradient(circle, rgba(255,255,255,0.1) 0%, transparent 70%);
  animation: float 6s ease-in-out infinite;
}

@keyframes float {
  0%, 100% { transform: translate(0, 0) rotate(0deg); }
  33% { transform: translate(30px, -30px) rotate(120deg); }
  66% { transform: translate(-20px, 20px) rotate(240deg); }
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
  animation: bounce 2s ease-in-out infinite;
}

@keyframes bounce {
  0%, 20%, 50%, 80%, 100% { transform: translateY(0); }
  40% { transform: translateY(-10px); }
  60% { transform: translateY(-5px); }
}

.logo-section h1 {
  color: white;
  margin: 0;
  margin-bottom: var(--apple-space-sm);
}

.login-form-section {
  flex: 1;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: var(--apple-space-2xl) var(--apple-space-lg);
}

.login-card {
  width: 100%;
  max-width: 400px;
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

.login-form {
  margin-bottom: var(--apple-space-lg);
}

.login-form .el-form-item {
  margin-bottom: var(--apple-space);
}

.full-width {
  width: 100%;
}

.full-width .el-input {
  width: 100%;
}

.login-button {
  width: 100%;
  height: 52px;
  font-size: var(--apple-font-size-lg);
  font-weight: var(--apple-font-weight-semibold);
}

.login-footer {
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

.register-link {
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
  .login-hero {
    padding: var(--apple-space-xl) var(--apple-space);
    min-height: 40vh;
  }
  
  .login-form-section {
    padding: var(--apple-space-xl) var(--apple-space);
    min-height: 60vh;
  }
  
  .logo-section h1 {
    font-size: var(--apple-font-size-3xl);
  }
  
  .login-card {
    max-width: 100%;
  }
  
  .login-button {
    height: 48px;
    font-size: var(--apple-font-size);
  }
}

@media (max-width: 480px) {
  .login-container {
    flex-direction: column;
    min-height: 100vh;
  }
  
  .login-hero {
    flex: 0.4;
    padding: var(--apple-space-lg) var(--apple-space);
    min-height: 35vh;
  }
  
  .login-form-section {
    flex: 0.6;
    padding: var(--apple-space-lg) var(--apple-space);
    display: flex;
    align-items: flex-start;
    justify-content: center;
  }
  
  .logo-section h1 {
    font-size: var(--apple-font-size-2xl);
    margin-bottom: var(--apple-space-xs);
  }
  
  .logo-icon {
    margin-bottom: var(--apple-space-sm);
  }
  
  .logo-icon svg {
    width: 48px;
    height: 48px;
  }
  
  .login-button {
    height: 44px;
    font-size: var(--apple-font-size-sm);
  }
  
  .card-header h2 {
    font-size: var(--apple-font-size-lg);
  }
  
  .login-form .el-form-item {
    margin-bottom: var(--apple-space-sm);
  }
  
  /* 确保输入框在移动端不会被缩放 */
  .apple-input .el-input__inner {
    font-size: 16px !important;
    min-height: 44px !important;
  }
}

/* 微交互 */
.apple-input .el-input__inner:focus {
  transform: translateY(-1px);
}

.login-button:not(:disabled):hover {
  transform: translateY(-2px);
}

.login-button:not(:disabled):active {
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