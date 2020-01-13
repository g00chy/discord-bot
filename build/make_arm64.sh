#!/bin/bash
CGO_ENABLED=1 GOOS=linux GOARCH=arm64 go build -mod vendor -o ./binary/arm64-discord-bot ./main.go