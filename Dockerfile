# 第一阶段：构建包含 goctls 的 Go 环境
FROM golang:1.22.5-alpine AS builder
# 安装必要的依赖
RUN apk update && apk add --no-cache git
# 设置工作目录
WORKDIR /go/src/app

# 安装 goctls
RUN go install github.com/suyuan32/goctls@latest

# 第二阶段：创建 Alpine 镜像并设置环境变量
FROM alpine:3.19

# Define the project name | 定义项目名称
ARG PROJECT=onion
# Define the config file name | 定义配置文件名
ARG CONFIG_FILE=onion.yaml
# Define the author | 定义作者
ARG AUTHOR="lei.wang@example.com"

# 复制 go 环境从第一阶段
COPY --from=builder /usr/local/go /usr/local/go

# 复制 goctls 工具从第一阶段
COPY --from=builder /go/bin/goctls /usr/local/bin/goctls

# 设置环境变量
ENV PATH="/usr/local/go/bin:/usr/local/bin:${PATH}"
#ENV PATH="/usr/local/bin:${PATH}"

# 安装必要的依赖
RUN apk update && apk add --no-cache tzdata protobuf

# 安装 protoc-gen-go
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
# 设置时区
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

# 复制应用代码和配置文件
COPY ./app /app/
COPY ./etc/onion.yaml /app/etc/config.yaml

# 设置工作目录
workdir /app/

# 暴露端口
EXPOSE 2165

# 设置入口点
ENTRYPOINT ["./app"]