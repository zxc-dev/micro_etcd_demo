FROM alpine:latest

RUN apk --no-cache add ca-certificates

ENV WORKDIR /app
ENV PARAMS=""

COPY ./bin/login_service $WORKDIR/login_service
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

CMD $WORKDIR/login_service $PARAMS
