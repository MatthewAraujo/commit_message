BIN_DIR=bin

LINUX_BINARY=$(BIN_DIR)/commit_message-linux-amd64
DARWIN_BINARY=$(BIN_DIR)/commit_message-darwin-amd64
WINDOWS_BINARY=$(BIN_DIR)/commit_message.exe

all: build-linux build-darwin build-windows

build-linux:
	@mkdir -p $(BIN_DIR)
	GOOS=linux GOARCH=amd64 go build -o $(LINUX_BINARY)
	@echo "Built Linux binary: $(LINUX_BINARY)"

build-darwin:
	@mkdir -p $(BIN_DIR)
	GOOS=darwin GOARCH=amd64 go build -o $(DARWIN_BINARY)
	@echo "Built Darwin binary: $(DARWIN_BINARY)"

build-windows:
	@mkdir -p $(BIN_DIR)
	GOOS=windows GOARCH=amd64 go build -o $(WINDOWS_BINARY)
	@echo "Built Windows binary: $(WINDOWS_BINARY)"

build:
	go build -o commit_message
	sudo mv commit_message /usr/local/bin/
	@echo "Built and installed commit_message to /usr/local/bin/"

clean:
	rm -rf $(BIN_DIR)
	@echo "Cleaned binaries in $(BIN_DIR)"

