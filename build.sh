export CGO_ENABLED=0

export GOOS=linux
export GOARCH=amd64
go build -trimpath -ldflags "-s -w" -o build/linux-amd64/ ./...

export GOOS=linux
export GOARCH=arm64
go build -trimpath -ldflags "-s -w" -o build/linux-arm64/ ./...

export GOOS=android
export GOARCH=arm64
go build -trimpath -ldflags "-s -w" -o build/android-arm64/ ./...

export GOOS=windows
export GOARCH=amd64
go build -trimpath -ldflags "-s -w" -o build/windows-amd64/ ./...
