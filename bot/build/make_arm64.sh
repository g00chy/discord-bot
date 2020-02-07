#!/bin/bash
go mod vendor
CC=aarch64-linux-gnu-gcc CGO_ENABLED=1 GOOS=linux GOARCH=arm64 go build -o ./binary/arm64-discord-bot ./bots/bot.go
CC=aarch64-linux-gnu-gcc CGO_ENABLED=1 GOOS=linux GOARCH=arm64 go build -o ./binary/arm64-discord-web ./web/web.go
chown ${uid}:${gid} -R ./binary
chown ${uid}:${gid} -R ./vendor