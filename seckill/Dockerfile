# 第一阶段：构建可执行文件
FROM golang:1.20 as builder

# 设置工作目录
WORKDIR /app

# 复制 go.mod 和 go.sum 文件
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制项目源代码
COPY . .

# 构建可执行文件
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# 第二阶段：创建轻量级镜像
FROM alpine:latest

# 安装必要的证书
RUN apk --no-cache add ca-certificates

# 设置工作目录
WORKDIR /root/

# 从第一阶段复制可执行文件
COPY --from=builder /app/main .

# 暴露服务端口，根据你的实际端口修改
EXPOSE 8080

# 启动服务
CMD ["./main"]

