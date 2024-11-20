# 使用官方的轻量级 Go 镜像
FROM golang:1.22-alpine

# 设置工作目录
WORKDIR /app

# 复制项目文件到容器
COPY go.mod go.sum ./
RUN go mod download
COPY etc/config.toml etc/config.toml

COPY . .

# 编译 Go 程序
RUN go build -o main .

# 暴露端口
EXPOSE 8010

# 设置运行命令
CMD ["./main", "start"]
