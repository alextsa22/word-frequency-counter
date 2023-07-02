FROM golang:alpine

# setup preferred platform
ENV GOOS=linux
ENV GOARCH=amd64

COPY . /github.com/alextsa22/word-frequency-counter
WORKDIR /github.com/alextsa22/word-frequency-counter

# install bash
RUN apk update && apk add bash

RUN go build -o main ./cmd/main.go

# use time command to show the execution time
CMD echo '[AIM]' && time ./scripts/aim_linux_64bit.sh ./assets/mobydick.txt && echo '' \
    && echo '[MAIN]' && time ./main ./assets/mobydick.txt
