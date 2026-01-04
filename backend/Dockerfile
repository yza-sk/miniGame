# Stage 1: Build
FROM golang:1.24-alpine AS builder

# 设置工作目录
WORKDIR /app

# 复制go mod文件
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制源代码
COPY . .

# 构建应用
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o rank_list_api cmd/main.go

# Stage 2: Runtime
FROM alpine:latest

# 安装必要的运行时依赖
RUN apk --no-cache add ca-certificates tzdata wget

# 创建非root用户
RUN addgroup -g 1000 -S appgroup && \
    adduser -u 1000 -S appuser -G appgroup

# 设置工作目录
WORKDIR /app

# 从builder阶段复制二进制文件
COPY --from=builder /app/rank_list_api .
COPY --from=builder /app/etc ./etc

# 更改文件权限
RUN chown -R appuser:appgroup /app

# 切换到非root用户
USER appuser

# 暴露端口
EXPOSE 8888

# 启动命令：执行构建出的二进制
CMD ["./rank_list_api"]