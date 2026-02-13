export CGO_ENABLED=0

export GOOS=windows
export GOARCH=amd64
go build -trimpath -ldflags "-s -w" -o dist/windows-amd64/upgrade.exe

export GOOS=linux
export GOARCH=arm64
go build -trimpath -ldflags "-s -w" -o dist/linux-arm64/upgrade