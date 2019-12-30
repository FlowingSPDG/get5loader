# Go
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=get5
DIST_DIR=build
SERVER_DIR=server
WEB_DIR=web
GAME_DIR=game_plugin

test:
	$(GOTEST) -v ./...
clean:
	$(GOCLEAN)
	rm -rf $(DIST_DIR)/*
deps:
	@$(GOGET) -v \
	github.com/FlowingSPDG/get5-web-go/server \
	github.com/mitchellh/gox \
	github.com/go-ini/ini \
	github.com/gorilla/mux \
	github.com/hydrogen18/stalecucumber \
	github.com/solovev/steam_go \
	github.com/go-sql-driver/mysql \
	github.com/jinzhu/gorm \
	github.com/kataras/go-sessions \
	github.com/Acidic9/steam \
	github.com/kidoman/go-steam
	@yarn
# Cross compile for go
build-all: build-prepare build-web clean
	@cd ./server && gox \
	-os="windows darwin linux" \
	-arch="386 amd64" \
	--output "../$(DIST_DIR)/$(BINARY_NAME)_{{.OS}}_{{.Arch}}/$(BINARY_NAME)"
	@cp -R ./web/build/ ./$(DIST_DIR)/$(BINARY_NAME)_linux_386/static
	@cp -R ./web/build/ ./$(DIST_DIR)/$(BINARY_NAME)_linux_amd64/static
	@cp -R ./web/build/ ./$(DIST_DIR)/$(BINARY_NAME)_windows_386/static
	@cp -R ./web/build/ ./$(DIST_DIR)/$(BINARY_NAME)_windows_amd64/static
	@cp -R ./web/build/ ./$(DIST_DIR)/$(BINARY_NAME)_darwin_386/static
	@cp -R ./web/build/ ./$(DIST_DIR)/$(BINARY_NAME)_darwin_amd64/static

	@cp ./server/config.ini.template ./$(DIST_DIR)/$(BINARY_NAME)_linux_386/config.ini.template
	@cp ./server/config.ini.template ./$(DIST_DIR)/$(BINARY_NAME)_linux_amd64/config.ini.template
	@cp ./server/config.ini.template ./$(DIST_DIR)/$(BINARY_NAME)_windows_386/config.ini.template
	@cp ./server/config.ini.template ./$(DIST_DIR)/$(BINARY_NAME)_windows_amd64/config.ini.template
	@cp ./server/config.ini.template ./$(DIST_DIR)/$(BINARY_NAME)_darwin_386/config.ini.template
	@cp ./server/config.ini.template ./$(DIST_DIR)/$(BINARY_NAME)_darwin_amd64/config.ini.template
	
build-prepare:
	@cd ./server && $(GOGET) github.com/mitchellh/gox
	@cd ./server && $(GOGET) github.com/konsorten/go-windows-terminal-sequences
	@cd ./server && $(GOGET) github.com/FlowingSPDG/get5-web-go/server
build-linux: build-prepare build-web
	@cd ./server && gox \
	-os="linux" \
	-arch="386 amd64" \
	--output "../$(DIST_DIR)/$(BINARY_NAME)_{{.OS}}_{{.Arch}}/$(BINARY_NAME)"
	@cp -R ./web/build/ ./$(DIST_DIR)/$(BINARY_NAME)_linux_386/static
	@cp -R ./web/build/ ./$(DIST_DIR)/$(BINARY_NAME)_linux_amd64/static
	@cp ./server/config.ini.template ./$(DIST_DIR)/$(BINARY_NAME)_linux_386/config.ini.template
	@cp ./server/config.ini.template ./$(DIST_DIR)/$(BINARY_NAME)_linux_amd64/config.ini.template
build-linux-server-only: build-prepare 
	@cd ./server && gox \
	-os="linux" \
	-arch="386 amd64" \
	--output "../$(DIST_DIR)/$(BINARY_NAME)_{{.OS}}_{{.Arch}}/$(BINARY_NAME)"
	@cp ./server/config.ini.template ./$(DIST_DIR)/$(BINARY_NAME)_linux_386/config.ini.template
	@cp ./server/config.ini.template ./$(DIST_DIR)/$(BINARY_NAME)_linux_amd64/config.ini.template
build-windows: build-prepare build-web
	@cd ./server && gox \
	-os="windows" \
	-arch="386 amd64" \
	--output "../$(DIST_DIR)/$(BINARY_NAME)_{{.OS}}_{{.Arch}}/$(BINARY_NAME)"
	@cp -R ./web/build/ ./$(DIST_DIR)/$(BINARY_NAME)_windows_386/static
	@cp -R ./web/build/ ./$(DIST_DIR)/$(BINARY_NAME)_windows_amd64/static
	@cp ./server/config.ini.template ./$(DIST_DIR)/$(BINARY_NAME)_windows_386/config.ini.template
	@cp ./server/config.ini.template ./$(DIST_DIR)/$(BINARY_NAME)_windows_amd64/config.ini.template
build-windows-server-only: build-prepare
	@cd ./server && gox \
	-os="windows" \
	-arch="386 amd64" \
	--output "../$(DIST_DIR)/$(BINARY_NAME)_{{.OS}}_{{.Arch}}/$(BINARY_NAME)"
	@cp ./server/config.ini.template ./$(DIST_DIR)/$(BINARY_NAME)_windows_386/config.ini.template
	@cp ./server/config.ini.template ./$(DIST_DIR)/$(BINARY_NAME)_windows_amd64/config.ini.template
build-mac: build-prepare build-web
	@cd ./server && gox \
	-os="darwin" \
	-arch="386 amd64" \
	--output "../$(DIST_DIR)/$(BINARY_NAME)_{{.OS}}_{{.Arch}}/$(BINARY_NAME)"
	@cp -R ./web/build/ ./$(DIST_DIR)/$(BINARY_NAME)_darwin_386/static
	@cp -R ./web/build/ ./$(DIST_DIR)/$(BINARY_NAME)_darwin_amd64/static
	@cp ./server/config.ini.template ./$(DIST_DIR)/$(BINARY_NAME)_darwin_386/config.ini.template
	@cp ./server/config.ini.template ./$(DIST_DIR)/$(BINARY_NAME)_darwin_amd64/config.ini.template
build-mac-server-only: build-prepare
	@cd ./server && gox \
	-os="darwin" \
	-arch="386 amd64" \
	--output "../$(DIST_DIR)/$(BINARY_NAME)_{{.OS}}_{{.Arch}}/$(BINARY_NAME)"
	@cp ./server/config.ini.template ./$(DIST_DIR)/$(BINARY_NAME)_darwin_386/config.ini.template
	@cp ./server/config.ini.template ./$(DIST_DIR)/$(BINARY_NAME)_darwin_amd64/config.ini.template
build-web:
	@cd ./web && yarn run build

# Source Mod compile
# TODO
#build-game:
	#@cd ./game_plugin/scripting