#!/bin/bash
go mod download
CC=aarch64-linux-gnu-gcc CGO_ENABLED=1 GOOS=linux GOARCH=arm64 go build -o ./binary/arm64-discord-bot ./bots/bot.go
CC=aarch64-linux-gnu-gcc CGO_ENABLED=1 GOOS=linux GOARCH=arm64 go build -o ./binary/arm64-discord-web ./web/web.go