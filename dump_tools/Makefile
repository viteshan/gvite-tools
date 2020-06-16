.PHONY: build

build:
	env GOOS=windows GOARCH=amd64 go build -o build/vite-dump-windows.exe
	env GOOS=darwin GOARCH=amd64 go build -o build/vite-dump-osx
	env GOOS=linux GOARCH=amd64 go build -o build/vite-dump-linux
