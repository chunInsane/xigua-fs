# 🍉文件服务

一个功能完整的多媒体文件管理和预览服务器系统，支持多用户、权限控制和丰富的文件操作功能。

## ✨ 功能特性

- 🔐 **多账户支持** - 完全的数据隔离，每个用户独立的文件空间
- 📁 **文件管理** - 上传、下载、删除、列表显示等完整文件操作
- 🖼️ **媒体预览** - 支持图片、视频在线预览，无需下载
- 📱 **响应式界面** - 移动端友好的Web界面，支持手机访问
- 👑 **管理员功能** - 用户管理、存储配额控制、系统监控
- 🔒 **隐私保护** - 仅登录用户可访问，文件权限完全隔离
- 💾 **存储控制** - 灵活的存储配额管理
- 🎨 **现代UI** - 基于Element Plus的美观界面

## 🛠 技术栈

- **后端**: Go + Gin Framework + GORM + SQLite + JWT
- **前端**: Vue 3 + Element Plus + Vite + Pinia
- **存储**: 本地文件系统 + SQLite数据库
- **认证**: JWT Token + bcrypt密码加密

## 📁 项目结构

```
xigua/
├── server/                 # Go后端服务
│   ├── main.go            # 服务入口
│   ├── config/            # 配置管理
│   ├── models/            # 数据模型 (User, FileInfo)
│   ├── handlers/          # API处理器 (auth, file, admin)
│   ├── middleware/        # 中间件 (认证, 权限)
│   ├── utils/             # 工具函数 (JWT, 文件处理)
│   └── storage/           # 文件存储目录
│       ├── xigua.db       # SQLite数据库
│       └── files/         # 用户文件存储
├── web/                   # Vue前端应用
│   ├── src/
│   │   ├── views/         # 页面组件
│   │   ├── components/    # 通用组件
│   │   ├── stores/        # Pinia状态管理
│   │   ├── utils/         # 工具函数
│   │   └── router/        # 路由配置
│   └── dist/              # 构建产物
├── start.sh               # 生产环境启动脚本
├── dev.sh                 # 开发环境启动脚本
└── README.md
```

## 🚀 快速开始

### 环境要求

- Node.js 16+ 
- Go 1.21+
- 现代浏览器

### 方式一：一键启动（推荐）

```bash
# 克隆项目（如果从仓库获取）
git clone <repository-url>
cd xigua

# 生产环境启动
./start.sh
```

### 方式二：开发模式

```bash
# 开发模式（推荐安装tmux获得更好体验）
./dev.sh
```

### 方式三：手动启动

#### 1. 启动后端服务

```bash
cd server
go mod tidy
go run main.go
```

#### 2. 启动前端开发服务（另开终端）

```bash
cd web
npm install
npm run dev
```

#### 3. 生产环境构建

```bash
# 构建前端
cd web
npm run build

# 启动后端（会自动服务前端静态文件）
cd ../server
go run main.go
```

## 📝 使用说明

### 首次使用

1. 访问 `http://localhost:8080`（生产环境）或 `http://localhost:3000`（开发环境）
2. 点击"立即注册"创建账户
3. **第一个注册的用户自动成为管理员**
4. 登录后即可开始上传和管理文件

### 主要功能

#### 普通用户功能
- **文件上传**: 支持拖拽上传，最大100MB
- **文件管理**: 网格/列表视图切换，按类型筛选
- **文件预览**: 图片、视频在线预览
- **文件下载**: 一键下载到本地
- **存储监控**: 查看存储使用情况

#### 管理员功能
- **用户管理**: 查看所有用户，编辑用户权限
- **存储控制**: 设置用户存储配额
- **系统监控**: 查看系统统计信息
- **用户删除**: 删除用户及其所有文件

### 支持的文件类型

- **图片**: JPG, PNG, GIF, BMP, WebP
- **视频**: MP4, AVI, MOV, WMV, FLV, WebM, MKV  
- **音频**: MP3, WAV, FLAC, AAC, OGG
- **文档**: PDF, TXT, MD
- **其他**: 所有文件类型均可上传存储

## 🔧 配置说明

### 环境变量

后端服务支持以下环境变量配置：

```bash
PORT=8080                    # 服务端口
DATABASE_PATH=./storage/xigua.db  # 数据库文件路径
STORAGE_PATH=./storage/files      # 文件存储路径  
JWT_SECRET=your-secret-key        # JWT密钥
```

### 默认配置

- 服务端口: `8080`
- 默认存储配额: `1GB`
- 最大文件大小: `100MB`
- JWT Token有效期: `24小时`

## 🔒 安全特性

- **密码加密**: 使用bcrypt加密存储
- **JWT认证**: 无状态的安全认证
- **文件隔离**: 用户文件完全隔离访问
- **权限控制**: 严格的API权限验证
- **SQL注入防护**: 使用GORM防止SQL注入

## 📊 API接口

### 认证接口
- `POST /api/auth/register` - 用户注册
- `POST /api/auth/login` - 用户登录  
- `GET /api/auth/me` - 获取用户信息

### 文件接口
- `POST /api/files/upload` - 文件上传
- `GET /api/files` - 获取文件列表
- `GET /api/files/:id/download` - 文件下载
- `GET /api/files/:id/preview` - 文件预览
- `DELETE /api/files/:id` - 删除文件

### 管理员接口
- `GET /api/admin/users` - 获取用户列表
- `PUT /api/admin/users/:id` - 更新用户信息
- `DELETE /api/admin/users/:id` - 删除用户
- `GET /api/admin/stats` - 系统统计

## 🤝 贡献指南

欢迎提交Issue和Pull Request来改进项目！

## 📄 许可证

本项目采用MIT许可证 - 查看 [LICENSE](LICENSE) 文件了解详情。

## 🎯 Roadmap

- [ ] 文件夹功能
- [ ] 文件分享链接  
- [ ] 批量操作
- [ ] 文件搜索
- [ ] 缩略图生成
- [ ] 视频转码
- [ ] Docker部署
- [ ] 对象存储支持

---

**🍉文件服务** - 让文件管理变得简单高效