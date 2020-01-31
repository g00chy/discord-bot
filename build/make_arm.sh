#!/bin/bash
go mod download
CC=arm-linux-gnueabi-gcc CGO_ENABLED=1 GOOS=linux GOARCH=arm go build -o ./binary/arm-discord-bot ./bot.go