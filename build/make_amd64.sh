#!/bin/bash
go mod download
CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o ./binary/amd-discord-bot ./main.go