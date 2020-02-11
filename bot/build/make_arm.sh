#!/bin/bash
CC=arm-linux-gnueabi-gcc CGO_ENABLED=1 GOOS=linux GOARCH=arm go build -o ./binary/arm-discord-bot ./bots/bot.go
CC=arm-linux-gnueabi-gcc CGO_ENABLED=1 GOOS=linux GOARCH=arm go build -o ./binary/arm-discord-web ./web/web.go
chown ${uid}:${gid} -R ./binary
chown ${uid}:${gid} -R ./vendor