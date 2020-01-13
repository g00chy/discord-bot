#!/bin/bash
CGO_ENABLED=1 GOOS=linux GOARCH=arm go build -mod vendor -o ./binary/arm-discord-bot ./main.go