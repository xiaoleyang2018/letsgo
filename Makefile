PKG_PATH=qinhao/letsgo
OUTPUT=app
SERVICE=app.service
TARGET=target
VERSION?=1.0.0
TARGET_DIR=app-$(VERSION)
TARGET_TAR=$(TARGET_DIR).tar.gz

APP_ROOT=/usr/local/app
APP_BIN=$(APP_ROOT)/bin
APP_ETC=$(APP_ROOT)/etc
INSTALL_ROOT=$(APP_ROOT)
SYSTEMD=/etc/systemd/system

BUILD_TIME=`date +%FT%T%z`
#LDFLAGS=-ldflags "-X main.Version=$(VERSION) -X main.BuildTime=$(BUILD_TIME)"
LDFLAGS += -X "$(PKG_PATH)/config.ReleaseVersion=$(shell git describe --tags --dirty)"
LDFLAGS += -X "$(PKG_PATH)/config.BuildTS=$(shell date -u '+%Y-%m-%d %I:%M:%S')"
LDFLAGS += -X "$(PKG_PATH)/config.GitHash=$(shell git rev-parse HEAD)"
LDFLAGS += -X "$(PKG_PATH)/config.GitBranch=$(shell git rev-parse --abbrev-ref HEAD)"


build:
	@echo Building App Server...
	@go build -ldflags '$(LDFLAGS)' -o $(OUTPUT)
	@echo Build App Server Success!

run:
	@echo "Running App Server..."
	@go run main.go

tar: build
	@echo Begin Making Tarball file...
	@mkdir -p $(TARGET)/$(TARGET_DIR)
	@cp -rf $(OUTPUT) config README* Makefile $(TARGET)/$(TARGET_DIR)
	@cd $(TARGET) && tar zcf $(TARGET_TAR) $(TARGET_DIR)
	@rm -rf $(TARGET)/$(TARGET_DIR)
	@echo Make Tarball file Success!
	@echo The target tarball can be found in $(TARGET)/$(TARGET_TAR)

install: 
	@echo "Installing App Server..."
	@sleep 1
	@if [ ! -d $(APP_ROOT) ]; then \
		mkdir -p $(APP_ETC) $(APP_BIN);\
	fi
	@cp $(OUTPUT) $(APP_BIN)
	@cp conf/app.conf $(APP_ETC)/app.conf
	@cp conf/$(SERVICE) $(SYSTEMD)/$(SERVICE)
	@chmod +x $(SYSTEMD)/$(SERVICE)
	@systemctl start $(SERVICE)
	@sleep 1
	@echo "Success to Install App Server."

vendor:

check:
	gocyclo -top 10 $(ls -d */ | grep -v vendor)
	interfacer $(go list ./... | grep -v /vendor/)
	find . -type d -not -path "./vendor/*" | xargs deadcode
	find . -type f -not -path "./vendor/*" -print0 | xargs -0 misspell

clean:
	@echo Clean up target file...
	@go clean -i ./...
	@rm -rf $(TARGET)
	@rm -rf $(OUTPUT)
	@echo Clean up target file success!
