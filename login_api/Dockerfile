FROM alpine:latest

RUN apk --no-cache add ca-certificates

ENV WORKDIR /app
ENV PARAMS=""

COPY ./bin/login_api $WORKDIR/login_api

CMD $WORKDIR/login_api $PARAMS
