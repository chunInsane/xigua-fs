<template>
  <el-config-provider :locale="zhCn">
    <div id="app">
      <router-view />
    </div>
  </el-config-provider>
</template>

<script>
import zhCn from 'element-plus/es/locale/lang/zh-cn'

export default {
  name: 'App',
  data() {
    return {
      zhCn
    }
  }
}
</script>

<style>
/* Apple 设计系统 - CSS 变量 */
:root {
  /* Apple 色彩系统 */
  --apple-blue: #007AFF;
  --apple-blue-light: #5AC8FA;
  --apple-blue-dark: #0051D5;
  --apple-green: #34C759;
  --apple-red: #FF3B30;
  --apple-orange: #FF9500;
  --apple-yellow: #FFCC00;
  --apple-purple: #AF52DE;
  --apple-pink: #FF2D92;
  --apple-gray: #8E8E93;
  
  /* Apple 中性色 */
  --apple-gray-1: #F2F2F7;
  --apple-gray-2: #E5E5EA;
  --apple-gray-3: #D1D1D6;
  --apple-gray-4: #C7C7CC;
  --apple-gray-5: #AEAEB2;
  --apple-gray-6: #8E8E93;
  --apple-text-primary: #000000;
  --apple-text-secondary: #6D6D70;
  --apple-text-tertiary: #8E8E93;
  --apple-background: #FFFFFF;
  --apple-background-secondary: #F2F2F7;
  --apple-background-tertiary: #FFFFFF;
  
  /* 深色模式支持 */
  --apple-background-elevated: #FFFFFF;
  --apple-border: rgba(0, 0, 0, 0.04);
  --apple-separator: rgba(0, 0, 0, 0.08);
  
  /* Apple 阴影 */
  --apple-shadow-sm: 0 1px 3px rgba(0, 0, 0, 0.04);
  --apple-shadow: 0 4px 6px rgba(0, 0, 0, 0.04);
  --apple-shadow-md: 0 10px 15px rgba(0, 0, 0, 0.04);
  --apple-shadow-lg: 0 20px 25px rgba(0, 0, 0, 0.06);
  --apple-shadow-xl: 0 25px 50px rgba(0, 0, 0, 0.08);
  
  /* Apple 圆角 */
  --apple-radius-sm: 6px;
  --apple-radius: 12px;
  --apple-radius-md: 16px;
  --apple-radius-lg: 20px;
  --apple-radius-xl: 24px;
  --apple-radius-full: 50%;
  
  /* Apple 间距 */
  --apple-space-xs: 4px;
  --apple-space-sm: 8px;
  --apple-space: 16px;
  --apple-space-md: 24px;
  --apple-space-lg: 32px;
  --apple-space-xl: 48px;
  --apple-space-2xl: 64px;
  
  /* Apple 字体 */
  --apple-font-size-xs: 12px;
  --apple-font-size-sm: 14px;
  --apple-font-size: 16px;
  --apple-font-size-lg: 18px;
  --apple-font-size-xl: 20px;
  --apple-font-size-2xl: 24px;
  --apple-font-size-3xl: 30px;
  --apple-font-size-4xl: 36px;
  --apple-font-size-5xl: 48px;
  
  /* Apple 字重 */
  --apple-font-weight-light: 300;
  --apple-font-weight-normal: 400;
  --apple-font-weight-medium: 500;
  --apple-font-weight-semibold: 600;
  --apple-font-weight-bold: 700;
}

/* 深色模式 */
@media (prefers-color-scheme: dark) {
  :root {
    --apple-text-primary: #FFFFFF;
    --apple-text-secondary: #EBEBF5;
    --apple-text-tertiary: #EBEBF5;
    --apple-background: #000000;
    --apple-background-secondary: #1C1C1E;
    --apple-background-tertiary: #2C2C2E;
    --apple-background-elevated: #1C1C1E;
    --apple-gray-1: #1C1C1E;
    --apple-gray-2: #2C2C2E;
    --apple-gray-3: #3A3A3C;
    --apple-gray-4: #48484A;
    --apple-gray-5: #636366;
    --apple-gray-6: #8E8E93;
    --apple-border: rgba(255, 255, 255, 0.08);
    --apple-separator: rgba(255, 255, 255, 0.1);
  }
}

/* 基础样式重置 */
#app {
  font-family: -apple-system, BlinkMacSystemFont, 'SF Pro Display', 'SF Pro Text', system-ui, sans-serif;
  height: 100vh;
  margin: 0;
  padding: 0;
  background: var(--apple-background);
  color: var(--apple-text-primary);
  font-size: var(--apple-font-size);
  line-height: 1.47059;
  font-weight: var(--apple-font-weight-normal);
  letter-spacing: -0.022em;
}

body {
  margin: 0;
  padding: 0;
  background: var(--apple-background);
  /* Apple 优化 */
  -webkit-tap-highlight-color: transparent;
  -webkit-touch-callout: none;
  -webkit-text-size-adjust: none;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-rendering: optimizeLegibility;
}

* {
  box-sizing: border-box;
}

/* Apple 按钮样式覆盖 */
.apple-button-primary {
  background: var(--apple-blue) !important;
  border: none !important;
  border-radius: var(--apple-radius) !important;
  color: white !important;
  font-weight: var(--apple-font-weight-semibold) !important;
  font-size: var(--apple-font-size) !important;
  padding: 12px 24px !important;
  transition: all 0.2s ease !important;
  box-shadow: var(--apple-shadow) !important;
  min-height: unset !important;
  line-height: 1.2 !important;
}

.apple-button-primary:hover {
  background: var(--apple-blue-dark) !important;
  box-shadow: var(--apple-shadow-md) !important;
  transform: translateY(-1px) !important;
}

.apple-button-secondary {
  background: var(--apple-gray-1) !important;
  border: 1px solid var(--apple-border) !important;
  border-radius: var(--apple-radius) !important;
  color: var(--apple-text-primary) !important;
  font-weight: var(--apple-font-weight-medium) !important;
  font-size: var(--apple-font-size) !important;
  padding: 12px 24px !important;
  transition: all 0.2s ease !important;
  min-height: unset !important;
  line-height: 1.2 !important;
}

.apple-button-secondary:hover {
  background: var(--apple-gray-2) !important;
  transform: translateY(-1px) !important;
}

/* Apple 输入框样式 */
.apple-input .el-input__inner {
  background: var(--apple-background-elevated) !important;
  border: 1px solid var(--apple-border) !important;
  border-radius: var(--apple-radius) !important;
  color: var(--apple-text-primary) !important;
  font-size: var(--apple-font-size) !important;
  padding: 14px 16px !important;
  transition: all 0.2s ease !important;
  box-shadow: var(--apple-shadow-sm) !important;
}

.apple-input .el-input__inner:focus {
  border-color: var(--apple-blue) !important;
  box-shadow: 0 0 0 3px rgba(0, 122, 255, 0.1) !important;
}

.apple-input .el-input__inner::placeholder {
  color: var(--apple-text-tertiary) !important;
  font-weight: var(--apple-font-weight-normal) !important;
}

/* Apple 卡片样式 */
.apple-card {
  background: var(--apple-background-elevated) !important;
  border: 1px solid var(--apple-border) !important;
  border-radius: var(--apple-radius-lg) !important;
  box-shadow: var(--apple-shadow) !important;
  padding: var(--apple-space-lg) !important;
  transition: all 0.3s ease !important;
}

.apple-card:hover {
  box-shadow: var(--apple-shadow-md) !important;
  transform: translateY(-2px) !important;
}

/* Apple 对话框样式 */
.apple-dialog .el-dialog {
  background: var(--apple-background-elevated) !important;
  border-radius: var(--apple-radius-lg) !important;
  box-shadow: var(--apple-shadow-xl) !important;
  border: 1px solid var(--apple-border) !important;
}

.apple-dialog .el-dialog__header {
  padding: var(--apple-space-lg) var(--apple-space-lg) var(--apple-space) !important;
  border-bottom: 1px solid var(--apple-separator) !important;
}

.apple-dialog .el-dialog__title {
  color: var(--apple-text-primary) !important;
  font-size: var(--apple-font-size-xl) !important;
  font-weight: var(--apple-font-weight-semibold) !important;
  line-height: 1.3 !important;
}

.apple-dialog .el-dialog__body {
  padding: var(--apple-space-lg) !important;
  color: var(--apple-text-primary) !important;
}

/* Apple 表格样式 */
.apple-table .el-table {
  background: var(--apple-background-elevated) !important;
  border: 1px solid var(--apple-border) !important;
  border-radius: var(--apple-radius-lg) !important;
  box-shadow: var(--apple-shadow) !important;
}

.apple-table .el-table th {
  background: var(--apple-gray-1) !important;
  color: var(--apple-text-secondary) !important;
  font-weight: var(--apple-font-weight-semibold) !important;
  border-bottom: 1px solid var(--apple-separator) !important;
}

.apple-table .el-table td {
  border-bottom: 1px solid var(--apple-separator) !important;
  color: var(--apple-text-primary) !important;
}

/* 移动端优化 */
@media (max-width: 768px) {
  :root {
    --apple-font-size: 17px;
    --apple-space: 16px;
    --apple-space-lg: 24px;
    --apple-radius: 12px;
    --apple-radius-lg: 16px;
  }
  
  .el-input__inner {
    font-size: 17px !important; /* 防止移动端Safari缩放 */
    min-height: 44px !important; /* 更好的触摸目标 */
  }
  
  .el-button {
    min-height: 44px !important; /* 移动端最小触摸目标 */
    font-size: 16px !important;
  }
  
  .el-dialog {
    width: 95% !important;
    margin: 5vh auto !important;
    max-height: 90vh;
    border-radius: var(--apple-radius-lg) !important;
  }
  
  .el-dialog__body {
    max-height: 70vh;
    overflow-y: auto;
  }
  
  /* 移动端触摸优化 */
  .el-dropdown-menu__item {
    min-height: 44px !important;
    padding: 12px 16px !important;
  }
  
  .el-table td {
    min-height: 44px !important;
  }
  
  /* 移动端滚动优化 */
  .el-scrollbar__wrap {
    -webkit-overflow-scrolling: touch;
  }
}

/* 小屏幕优化 */
@media (max-width: 480px) {
  :root {
    --apple-font-size: 16px;
    --apple-space: 12px;
    --apple-space-lg: 20px;
    --apple-radius: 10px;
    --apple-radius-lg: 14px;
  }
  
  body {
    font-size: 16px; /* 防止Safari放大 */
  }
  
  .el-input__inner {
    font-size: 16px !important;
    min-height: 40px !important;
  }
  
  .el-button {
    min-height: 40px !important;
    font-size: 14px !important;
    padding: 8px 16px !important;
  }
  
  /* 防止水平滚动 */
  * {
    max-width: 100%;
    word-wrap: break-word;
  }
  
  /* 优化触摸滚动 */
  body, html {
    touch-action: manipulation;
    -webkit-text-size-adjust: 100%;
  }
}

/* 通用工具类 */
.apple-text-display {
  font-size: var(--apple-font-size-4xl) !important;
  font-weight: var(--apple-font-weight-bold) !important;
  line-height: 1.07143 !important;
  letter-spacing: -0.022em !important;
}

.apple-text-title {
  font-size: var(--apple-font-size-2xl) !important;
  font-weight: var(--apple-font-weight-semibold) !important;
  line-height: 1.16667 !important;
  letter-spacing: -0.022em !important;
}

.apple-text-subtitle {
  font-size: var(--apple-font-size-lg) !important;
  font-weight: var(--apple-font-weight-medium) !important;
  line-height: 1.22222 !important;
}

.apple-text-body {
  font-size: var(--apple-font-size) !important;
  font-weight: var(--apple-font-weight-normal) !important;
  line-height: 1.42857 !important;
}

.apple-text-caption {
  font-size: var(--apple-font-size-sm) !important;
  font-weight: var(--apple-font-weight-medium) !important;
  color: var(--apple-text-secondary) !important;
  line-height: 1.33333 !important;
}

/* 滚动条美化 */
::-webkit-scrollbar {
  width: 6px;
  height: 6px;
}

::-webkit-scrollbar-track {
  background: transparent;
}

::-webkit-scrollbar-thumb {
  background: var(--apple-gray-4);
  border-radius: 3px;
}

::-webkit-scrollbar-thumb:hover {
  background: var(--apple-gray-5);
}
</style>