FROM golang:1.24.2-alpine3.21 AS builder

WORKDIR /build
ENV AppName="generate-id"
ENV GOPROXY="https://goproxy.cn,direct"
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${AppName}


FROM agan-harbor-registry.cn-hangzhou.cr.aliyuncs.com/saas-public/alpine:edge:3.18.4

ENV AppName="generate-id"
WORKDIR /app
COPY --from=builder /build/generate-id /app/generate-id

#ENTRYPOINT ["sh","+x","/usr/bin/docker-entrypoint.sh"]
CMD ["/app/generate-id"]
