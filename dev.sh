#!/bin/bash

echo "🍉 西瓜文件服务器 开发模式"
echo "========================="

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

# 创建tmux会话或使用现有会话
SESSION_NAME="xigua-dev"

# 检查tmux是否安装
if ! command -v tmux &> /dev/null; then
    echo "⚠️  推荐安装tmux以获得更好的开发体验"
    echo "🔧 启动后端服务器（开发模式）..."
    cd server
    go mod tidy
    go run main.go
else
    # 使用tmux分屏开发
    echo "🔧 使用tmux启动开发环境..."
    
    # 杀死现有会话（如果存在）
    tmux kill-session -t $SESSION_NAME 2>/dev/null
    
    # 创建新会话
    tmux new-session -d -s $SESSION_NAME
    
    # 分割窗口
    tmux split-window -h
    
    # 左边窗口：后端
    tmux send-keys -t $SESSION_NAME:0.0 "cd server && go mod tidy && echo '🚀 启动后端服务器...' && go run main.go" Enter
    
    # 右边窗口：前端
    tmux send-keys -t $SESSION_NAME:0.1 "cd web && npm install && echo '🌐 启动前端开发服务器...' && npm run dev" Enter
    
    # 连接到会话
    echo "✅ 开发环境启动完成！"
    echo "💡 后端服务: http://localhost:8080"
    echo "💡 前端服务: http://localhost:3000"
    echo "💡 使用 Ctrl+B 然后 D 来分离会话"
    echo "💡 使用 tmux attach -t $SESSION_NAME 重新连接"
    echo ""
    
    tmux attach -t $SESSION_NAME
fi