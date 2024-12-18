# 使用官方的轻量级 Go 镜像
FROM golang:1.22-alpine

# 设置工作目录
WORKDIR /app
# COPY . .
# COPY go.mod go.sum ./

# 复制项目文件到容器
COPY etc/config.toml etc/config.toml
COPY etc/kade-poc.json /Users/kade.chen/go-kade-project/github/mcenter/etc/kade-poc.json
# 将本地源码复制到 Docker 镜像中
# COPY vertexai@v0.13.1/genai /go/pkg/mod/cloud.google.com/go/vertexai/genai

# RUN go mod download
# 复制其他项目文件到容器
# COPY . .

# 编译 Go 程序
# RUN go build -o main .
COPY main main

# 暴露端口
EXPOSE 8010

# 设置容器启动时的运行命令
CMD ["./main", "start"]