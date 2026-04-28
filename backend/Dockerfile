# 第一阶段：构建环境
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
# 编译成名为 taskflow 的二进制文件
RUN CGO_ENABLED=0 GOOS=linux go build -o taskflow main.go

# 第二阶段：运行环境
FROM alpine:latest
WORKDIR /root/
# 只从构建阶段拷贝二进制文件
COPY --from=builder /app/taskflow .
# 拷贝配置文件（如果有的话）
COPY --from=builder /app/.env .

EXPOSE 8080
CMD ["./taskflow"]