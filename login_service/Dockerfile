FROM golang:1.16.5 AS builder
WORKDIR /app
ENV GO111MODULE=on
ENV GOPROXY https://goproxy.cn
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -o login_service

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /app/login_service .
ENV PARAMS=""
ENTRYPOINT ./login_service $PARAMS