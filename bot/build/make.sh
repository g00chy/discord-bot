#!/bin/bash
go build -o ./binary/${GOARCH}-discord-bot ./bots/bot.go
go build -o ./binary/${GOARCH}-discord-web ./web/web.go
chown ${uid}:${gid} -R ./binary
