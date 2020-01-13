$env:GOOS = "linux"
$env:GOARCH = "arm"
go build -o .\binary\arm-discord-bot .\main.go
$env:GOOS = "linux"
$env:GOARCH = "arm64"
go build -o .\binary\arm64-discord-bot .\main.go
$env:GOOS = "linux"
$env:GOARCH = "amd64"
go build -o .\binary\amd64-discord-bot .\main.go