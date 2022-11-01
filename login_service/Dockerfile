FROM alpine:latest

RUN apk --no-cache add ca-certificates

ENV WORKDIR /app
ENV PARAMS=""

COPY ./bin/login_service $WORKDIR/login_service

CMD $WORKDIR/login_service $PARAMS
