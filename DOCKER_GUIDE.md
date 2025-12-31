# Docker 构建和部署指南

## 前提条件

1. **安装Docker Desktop**
   - 从[Docker官网](https://www.docker.com/products/docker-desktop/)下载并安装Docker Desktop
   - 启动Docker Desktop服务
   - 确保Docker服务正在运行：
     ```bash
     docker version
     ```

2. **解决连接问题**
   如果遇到 `open //./pipe/dockerDesktopLinuxEngine: The system cannot find the file specified` 错误：
   - 确保Docker Desktop正在运行
   - 重启Docker Desktop
   - 检查Windows服务中的Docker服务状态

## 构建步骤

### 方法1：使用Docker Compose（推荐）

1. **本地构建和测试**
   ```bash
   # 使用本地配置文件构建
   docker-compose -f docker-compose.local.yml up -d --build
   ```

2. **生产环境部署**
   ```bash
   # 使用生产配置文件
   docker-compose up -d --build
   ```

### 方法2：分别构建镜像

1. **构建后端镜像**
   ```bash
   cd backend
   docker build -t minigame-backend:latest .
   ```

2. **构建前端镜像**
   ```bash
   cd web
   docker build -t minigame-frontend:latest .
   ```

3. **运行容器**
   ```bash
   # 运行后端
   docker run -d -p 8888:8888 --name minigame-backend minigame-backend:latest
   
   # 运行前端
   docker run -d -p 80:80 --name minigame-frontend minigame-frontend:latest
   ```

## 验证部署

1. **检查容器状态**
   ```bash
   docker-compose ps
   # 或
   docker ps
   ```

2. **访问服务**
   - 前端应用：http://localhost
   - 后端API：http://localhost:8888

3. **查看日志**
   ```bash
   # 查看所有服务日志
   docker-compose logs -f
   
   # 查看特定服务日志
   docker-compose logs -f backend
   docker-compose logs -f frontend
   ```

## 常用命令

```bash
# 停止服务
docker-compose down

# 停止并删除数据卷
docker-compose down -v

# 重新构建并启动
docker-compose up -d --build

# 清理无用镜像
docker system prune -f
```

## 配置说明

### 环境变量

- `VITE_API_BASE`: 前端API基础地址，默认为 `http://localhost:8888`
- `TZ`: 时区设置，默认为 `Asia/Shanghai`

### 端口映射

- 后端服务：8888 → 8888
- 前端服务：80 → 80

### 数据卷

- 后端配置文件：`./backend/etc` → `/app/etc`

## 故障排除

### Docker连接问题

1. **Docker Desktop未运行**
   - 启动Docker Desktop
   - 等待服务完全启动

2. **WSL2后端问题**（Windows）
   - 重启WSL2：
     ```powershell
     wsl --shutdown
     ```
   - 重启Docker Desktop

3. **权限问题**
   - 以管理员身份运行Docker Desktop
   - 检查Windows防火墙设置

### 构建失败

1. **网络问题**
   - 检查网络连接
   - 配置Docker镜像加速器

2. **依赖下载失败**
   - 重试构建命令
   - 检查go.mod或package.json文件

### 运行时问题

1. **端口冲突**
   - 检查端口是否被占用
   - 修改docker-compose.yml中的端口映射

2. **服务启动失败**
   - 查看容器日志
   - 检查配置文件是否正确

## 推送镜像到仓库

如果需要推送到阿里云容器镜像仓库：

```bash
# 登录阿里云容器镜像服务
docker login crpi-l14b1sc286yk94go.cn-hangzhou.personal.cr.aliyuncs.com

# 推送后端镜像
docker push crpi-l14b1sc286yk94go.cn-hangzhou.personal.cr.aliyuncs.com/galgame/minigame_backend:v1.0.0

# 推送前端镜像
docker push crpi-l14b1sc286yk94go.cn-hangzhou.personal.cr.aliyuncs.com/galgame/minigame_frontend:v1.0.0
```