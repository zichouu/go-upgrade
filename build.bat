set CGO_ENABLED=0

set GOOS=windows
set GOARCH=amd64
go build -trimpath -ldflags "-s -w" -o dist/windows-amd64/upgrade.exe

set GOOS=linux
set GOARCH=arm64
go build -trimpath -ldflags "-s -w" -o dist/linux-arm64/upgrade