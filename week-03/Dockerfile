FROM golang:1.16-alpine AS builder
COPY http-server-demo /app
WORKDIR /app
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN CGO_ENABLE=0 GOOS=linux go build -o /bin/http-server main.go

FROM alpine
COPY --from=builder /bin/http-server /bin/http-server 
CMD ["/bin/http-server"]
