FROM golang:1.25rc3-bookworm AS builder
WORKDIR /app
COPY . .
RUN GOOS=linux GOARCH=arm64 go build -o main -ldflags "-w -s" .

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]