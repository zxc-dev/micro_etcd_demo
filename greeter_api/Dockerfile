FROM golang:1.16.5 AS builder
WORKDIR /app
ENV GO111MODULE=on
ENV GOPROXY https://goproxy.cn
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -o greeter_api

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /app/greeter_api .
ENV PARAMS=""
CMD ./greeter_api $PARAMS