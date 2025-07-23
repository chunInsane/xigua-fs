#!/bin/bash

echo "🍉 西瓜文件服务器 启动脚本"
echo "=========================="

# 检查Node.js
if ! command -v node &> /dev/null; then
    echo "❌ Node.js 未安装，请先安装 Node.js"
    exit 1
fi

# 检查Go
if ! command -v go &> /dev/null; then
    echo "❌ Go 未安装，请先安装 Go"
    exit 1
fi

# 安装前端依赖
echo "📦 安装前端依赖..."
cd web
npm install
if [ $? -ne 0 ]; then
    echo "❌ 前端依赖安装失败"
    exit 1
fi

# 构建前端
echo "🔨 构建前端..."
npm run build
if [ $? -ne 0 ]; then
    echo "❌ 前端构建失败"
    exit 1
fi

# 返回根目录
cd ..

# 安装后端依赖
echo "📦 安装后端依赖..."
cd server
go mod tidy
if [ $? -ne 0 ]; then
    echo "❌ 后端依赖安装失败"
    exit 1
fi

# 启动服务器
echo "🚀 启动服务器..."
echo "服务将在 http://localhost:8080 启动"
echo "按 Ctrl+C 停止服务"
echo ""

go run main.go