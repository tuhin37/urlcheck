# Define paths
BINS_DIR := ./bin
INSTALL_DIR_UNIX := /usr/local/bin
INSTALL_DIR_WINDOWS := C:/Program Files

# Determine OS and architecture
OS := $(shell uname -s | tr '[:upper:]' '[:lower:]')
ARCH := $(shell uname -m)
ifeq ($(OS), darwin)
    ifeq ($(ARCH), x86_64)
        BIN := urlcheck-darwin-amd64
    else ifeq ($(ARCH), arm64)
        BIN := urlcheck-darwin-arm64
    endif
else ifeq ($(OS), linux)
    ifeq ($(ARCH), x86_64)
        BIN := urlcheck-linux-amd64
    else ifeq ($(ARCH), arm64)
        BIN := urlcheck-linux-arm64
    endif
else
    $(error Unsupported OS or Architecture)
endif

# Handle Windows separately
ifeq ($(OS), mingw32)
    BIN := urlcheck-windows-amd64.exe
endif

# Install command
install:
ifeq ($(OS), mingw32)
	@echo "Installing for Windows..."
	copy $(BINS_DIR)/$(BIN) $(INSTALL_DIR_WINDOWS)/urlcheck.exe
else
	@echo "Installing for Unix-like OS..."
	cp $(BINS_DIR)/$(BIN) $(INSTALL_DIR_UNIX)/urlcheck
	chmod +x $(INSTALL_DIR_UNIX)/urlcheck
endif

# Uninstall command
uninstall:
ifeq ($(OS), mingw32)
	@echo "Uninstalling from Windows..."
	del $(INSTALL_DIR_WINDOWS)/urlcheck.exe
else
	@echo "Uninstalling from Unix-like OS..."
	rm -f $(INSTALL_DIR_UNIX)/urlcheck
endif

.PHONY: install uninstall






# Docker commands
docker-build:
	docker build -t urlcheck:dev -f Dockerfile .

docker-run:
	docker run -it urlcheck:dev 

# builds binaries from go
gobuild:
	./build/build.sh