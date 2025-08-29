FROM golang:1.25rc3-bookworm AS builder
WORKDIR /app
COPY . .
# 编译 Go 应用，-o 指定输出文件名，-ldflags "-w -s" 减小体积
RUN GOOS=linux GOARCH=amd64 go build -o main -ldflags "-w -s" .

FROM alpine:latest
WORKDIR /root/
# 从 builder 阶段拷贝编译好的二进制文件
COPY --from=builder /app/main .
# 暴露端口
EXPOSE 8080
# 运行应用
CMD ["./main"]