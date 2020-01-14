#!/bin/bash
CC=aarch64-linux-gnu-gcc CGO_ENABLED=1 GOOS=linux GOARCH=arm64 go build -o ./binary/arm64-discord-bot ./main.go