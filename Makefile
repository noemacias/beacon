APP_NAME := beacon
VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo 0.1.0)

DIST_DIR := dist
BUILD_FLAGS := -ldflags="-s -w"

.PHONY: clean build release

clean:
	@echo "Cleaning build directories"
	rm -rf bin/ $(DIST_DIR)/

build: clean
	@echo "Go mod tidy"
	go mod tidy

	@echo "Building local binary"
	mkdir -p bin
	CGO_ENABLED=0 go build $(BUILD_FLAGS) -o bin/$(APP_NAME) main.go

release: clean
	@echo "Go mod tidy"
	go mod tidy

	@echo "Creating release artifacts"
	mkdir -p $(DIST_DIR)

	# Linux AMD64
	mkdir -p $(DIST_DIR)/$(APP_NAME)_$(VERSION)_linux_amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
		go build $(BUILD_FLAGS) \
		-o $(DIST_DIR)/$(APP_NAME)_$(VERSION)_linux_amd64/$(APP_NAME) main.go
	cp README.md LICENSE $(DIST_DIR)/$(APP_NAME)_$(VERSION)_linux_amd64/
	tar -czf $(DIST_DIR)/$(APP_NAME)_$(VERSION)_linux_amd64.tar.gz \
		-C $(DIST_DIR) $(APP_NAME)_$(VERSION)_linux_amd64
	rm -rf $(DIST_DIR)/$(APP_NAME)_$(VERSION)_linux_amd64

	# Linux ARM64
	mkdir -p $(DIST_DIR)/$(APP_NAME)_$(VERSION)_linux_arm64
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 \
		go build $(BUILD_FLAGS) \
		-o $(DIST_DIR)/$(APP_NAME)_$(VERSION)_linux_arm64/$(APP_NAME) main.go
	cp README.md LICENSE $(DIST_DIR)/$(APP_NAME)_$(VERSION)_linux_arm64/
	tar -czf $(DIST_DIR)/$(APP_NAME)_$(VERSION)_linux_arm64.tar.gz \
		-C $(DIST_DIR) $(APP_NAME)_$(VERSION)_linux_arm64
	rm -rf $(DIST_DIR)/$(APP_NAME)_$(VERSION)_linux_arm64

	# macOS AMD64
	mkdir -p $(DIST_DIR)/$(APP_NAME)_$(VERSION)_darwin_amd64
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 \
		go build $(BUILD_FLAGS) \
		-o $(DIST_DIR)/$(APP_NAME)_$(VERSION)_darwin_amd64/$(APP_NAME) main.go
	cp README.md LICENSE $(DIST_DIR)/$(APP_NAME)_$(VERSION)_darwin_amd64/
	tar -czf $(DIST_DIR)/$(APP_NAME)_$(VERSION)_darwin_amd64.tar.gz \
		-C $(DIST_DIR) $(APP_NAME)_$(VERSION)_darwin_amd64
	rm -rf $(DIST_DIR)/$(APP_NAME)_$(VERSION)_darwin_amd64

	# macOS ARM64 (Apple Silicon)
	mkdir -p $(DIST_DIR)/$(APP_NAME)_$(VERSION)_darwin_arm64
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 \
		go build $(BUILD_FLAGS) \
		-o $(DIST_DIR)/$(APP_NAME)_$(VERSION)_darwin_arm64/$(APP_NAME) main.go
	cp README.md LICENSE $(DIST_DIR)/$(APP_NAME)_$(VERSION)_darwin_arm64/
	tar -czf $(DIST_DIR)/$(APP_NAME)_$(VERSION)_darwin_arm64.tar.gz \
		-C $(DIST_DIR) $(APP_NAME)_$(VERSION)_darwin_arm64
	rm -rf $(DIST_DIR)/$(APP_NAME)_$(VERSION)_darwin_arm64

	# Windows AMD64
	mkdir -p $(DIST_DIR)/$(APP_NAME)_$(VERSION)_windows_amd64
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 \
		go build $(BUILD_FLAGS) \
		-o $(DIST_DIR)/$(APP_NAME)_$(VERSION)_windows_amd64/$(APP_NAME).exe main.go
	cp README.md LICENSE $(DIST_DIR)/$(APP_NAME)_$(VERSION)_windows_amd64/
	cd $(DIST_DIR) && zip -rq $(APP_NAME)_$(VERSION)_windows_amd64.zip \
		$(APP_NAME)_$(VERSION)_windows_amd64
	rm -rf $(DIST_DIR)/$(APP_NAME)_$(VERSION)_windows_amd64

	# Windows ARM64
	mkdir -p $(DIST_DIR)/$(APP_NAME)_$(VERSION)_windows_arm64
	CGO_ENABLED=0 GOOS=windows GOARCH=arm64 \
		go build $(BUILD_FLAGS) \
		-o $(DIST_DIR)/$(APP_NAME)_$(VERSION)_windows_arm64/$(APP_NAME).exe main.go
	cp README.md LICENSE $(DIST_DIR)/$(APP_NAME)_$(VERSION)_windows_arm64/
	cd $(DIST_DIR) && zip -rq $(APP_NAME)_$(VERSION)_windows_arm64.zip \
		$(APP_NAME)_$(VERSION)_windows_arm64
	rm -rf $(DIST_DIR)/$(APP_NAME)_$(VERSION)_windows_arm64

	@echo "Generating checksums"
	cd $(DIST_DIR) && sha256sum *.tar.gz *.zip > checksums.txt

	@echo
	@echo "Release artifacts created in $(DIST_DIR)/"