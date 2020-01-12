$env:GOOS = "linux"
$env:GOARCH = "arm"
go env GOOS GOARCH
go build -o .\binary\afk-bot .\afk-bot\afk-bot.go
go build -o .\binary\claim-bot .\claim-bot\claim-bot.go
go build -o .\binary\nleave-ban-bot .\nleave-ban-bot\nleave-ban-bot.go