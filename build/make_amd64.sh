#!/bin/bash
CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go -mod vendor build -o ./binary/amd-discord-bot ./main.go