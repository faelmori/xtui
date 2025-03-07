ROOT_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))
BINARY_NAME := $(ROOT_DIR)xtui
CMD_DIR := $(ROOT_DIR)cmd
INSTALL_SCRIPT=$(ROOT_DIR)scripts/install.sh
ARGS :=

# Alvo para build
build:
	@go build -ldflags "-s -w -X main.version=$(git describe --tags) -X main.commit=$(git rev-parse HEAD) -X main.date=$(date +%Y-%m-%d)" -trimpath -o $(BINARY_NAME) ${CMD_DIR} &&\
    upx $(BINARY_NAME) --force-overwrite --lzma --no-progress --no-color
	@echo "Built $(BINARY_NAME)"

build-dev:
	@if [ -f $(BINARY_NAME) ]; then rm $(BINARY_NAME); fi
	@go build -ldflags "-s -w -X main.version=$(git describe --tags) -X main.commit=$(git rev-parse HEAD) -X main.date=$(date +%Y-%m-%d)" -trimpath -o $(BINARY_NAME) ${CMD_DIR}
	@echo "Built $(BINARY_NAME)"

install:
	@sh $(INSTALL_SCRIPT) install $(ARGS)
	@echo "Installed $(BINARY_NAME)"

# Limpar o binário gerado
clean:
	@if [ -f $(BINARY_NAME) ]; then rm $(BINARY_NAME); fi
	@echo "Cleaned up build artifacts"

# Alvo de ajuda
help:
	@echo "Available targets:"
	@echo "  make build      - Build the binary using install script"
	@echo "  make build-dev  - Build the binary without using install script"
	@echo "  make install    - Install the binary and configure environment"
	@echo "  make clean      - Clean up build artifacts"
	@echo "  make help       - Display this help message"
	@echo ""
	@echo "Usage with arguments:"
	@echo "  make install ARGS='--custom-arg value' - Pass custom arguments to the install script"
