$env:GOOS = "linux"
$env:GOARCH = "arm"
go env GOOS GOARCH
go build -o .\binary\discord-bot .\main.go