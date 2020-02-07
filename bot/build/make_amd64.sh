#!/bin/bash
go mod vendor
CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o ./binary/amd-discord-bot ./bots/bot.go
CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o ./binary/amd-discord-web ./web/web.go
chown ${uid}:${gid} -R ./binary