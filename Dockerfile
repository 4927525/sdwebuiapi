FROM golang:1.20

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /app
COPY . .
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o main

# 指定运行时环境变量
ENV GIN_MODE=release
EXPOSE 8009
ENTRYPOINT ["./main"]