all: linux_arm64 linux_amd64 windows_amd64

linux_amd64:
	@go env -w GOOS=linux
	@go env -w GOARCH=amd64
	@go build -o consulapi_amd64 main.go

linux_arm64:
	@go env -w GOOS=linux
	@go env -w GOARCH=arm64
	@go build -o consulapi_arm64 main.go

windows_amd64:
	@go env -w GOOS=windows
	@go env -w GOARCH=amd64
	@go build -o consulapi.exe main.go