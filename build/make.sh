#!/bin/bash
#CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o ./binary/amd-discord-bot ./main.go
#CGO_ENABLED=1 GOOS=linux GOARCH=arm go build -o ./binary/arm-discord-bot ./main.go
CGO_ENABLED=1 GOOS=linux GOARCH=arm go build -o ./binary/arm-discord-bot ./main.go
#CGO_ENABLED=1 GOOS=linux GOARCH=arm64 go build -o ./binary/arm64-discord-bot ./main.go