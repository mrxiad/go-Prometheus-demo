# 使用 Golang 官方镜像作为构建环境
FROM golang:1.20-alpine AS builder

# 设置工作目录
WORKDIR /app

# 将当前目录内容复制到工作目录中
COPY . .

# 删除现有的 go.mod 和 go.sum 文件
RUN rm -f go.mod go.sum

# 初始化 Go 模块
RUN go mod init my-go-app

# 添加依赖
RUN go get github.com/prometheus/client_golang

# 运行 go mod tidy 以下载所有依赖
RUN go mod tidy

# 编译 Go 应用程序
RUN go build -o my-go-app .

# 使用一个较小的基础镜像
FROM alpine:latest

# 设置工作目录
WORKDIR /app

# 从构建阶段复制编译后的二进制文件
COPY --from=builder /app/my-go-app .

# 设置时区信息
ENV ZONEINFO /usr/local/go/lib/time/zoneinfo.zip

# 暴露应用程序的端口
EXPOSE 8080

# 运行二进制文件
CMD ["./my-go-app"]
