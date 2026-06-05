clean:
	@echo "Cleaning bin directory"
	rm -rf bin/

build: clean
	@echo "Go mod tidy"
	go mod tidy

	@echo
	@echo "Building binaries"
	CGO_ENABLED=0 GOOS=linux  GOARCH=amd64 go build -o bin/beacon main.go
	