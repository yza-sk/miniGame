# Go 后端 Dockerfile
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
# 构建可执行文件
RUN go mod tidy && go build -o server main.go
WORKDIR /app
COPY . .
# 构建启用 CGO 的可执行文件
RUN go mod tidy && go build -o server main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/server .
COPY etc ./etc
EXPOSE 8888
CMD ["./server"]
