#!/bin/bash
go build -o ./binary/${GOARCH}-discord-bot --ldflags '-linkmode external -extldflags "-static"' ./bots/bot.go
go build -o ./binary/${GOARCH}-discord-web --ldflags '-linkmode external -extldflags "-static"' ./web/web.go
chown ${uid}:${gid} -R ./binary
