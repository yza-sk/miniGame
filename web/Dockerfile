# Stage 1: Build
FROM node:18-alpine AS build
WORKDIR /app
# 只复制包清单，避免锁文件不存在时COPY失败
COPY package.json .
# 优先使用npm ci（若无锁或失败则回退npm install）
RUN npm set fund false && (npm ci || npm install)
# 复制源码
COPY . .
# You can pass VITE_API_BASE at build time
ARG VITE_API_BASE
ENV VITE_API_BASE=${VITE_API_BASE}
RUN npm run build

# Stage 2: Serve with Nginx
FROM nginx:latest AS runtime
WORKDIR /usr/share/nginx/html
COPY --from=build /app/dist .
COPY nginx.conf /etc/nginx/conf.d/default.conf
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
