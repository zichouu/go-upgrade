set CGO_ENABLED=0

set GOOS=linux
set GOARCH=amd64
go build -trimpath -ldflags "-s -w" -o build/linux-amd64/ ./...

set GOOS=linux
set GOARCH=arm64
go build -trimpath -ldflags "-s -w" -o build/linux-arm64/ ./...

set GOOS=android
set GOARCH=arm64
go build -trimpath -ldflags "-s -w" -o build/android-arm64/ ./...

set GOOS=windows
set GOARCH=amd64
go build -trimpath -ldflags "-s -w" -o build/windows-amd64/ ./...
