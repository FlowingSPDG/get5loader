# Go
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOINSTALL=$(GOCMD) install
BINARY_NAME=get5
DIST_DIR=build
SERVER_DIR=bacjebd
WEB_DIR=web
OS_Linux=linux
ARCH_AMD64=amd64

.DEFAULT_GOAL := build-backend

test:
	cd ./backend && $(GOTEST) -v ./...
clean:
	@$(GOCLEAN)
	-@$(RM) $(DIST_DIR)/*
deps: deps-web
deps-web:
	@cd ./front && yarn
up:
	docker compose up -d --build
down:
	docker compose down
generate:
	@cd ./backend && \
	$(GOINSTALL) golang.org/x/tools/cmd/stringer@v0.11.1 && \
	$(GOINSTALL) github.com/sqlc-dev/sqlc/cmd/sqlc@latest && \
	go generate ./...
# compile for go
build-backend: build-prepare
	cd ./backend/cmd && \
	gox \
	-os="$(OS_Linux)" \
	-arch="$(ARCH_AMD64)" \
	--output "../$(DIST_DIR)/$(BINARY_NAME)_{{.OS}}_{{.Arch}}/$(BINARY_NAME)"
build-prepare: clean generate
	@$(GOINSTALL) github.com/mitchellh/gox@latest
	-@$(RM) ./$(DIST_DIR)/*/static
build-front: deps-web
	@cd ./front && yarn build
coverage:
	cd backend && \
	go test -v -count=1 -covermode=count -coverprofile=coverage.out ./... | tee test_output.txt && \
	go tool cover -func=coverage.out | awk '/total:/ {print "| **" $$1 "** | **" $$3 "** |"}' | tee coverage.txt && \
	cat test_output.txt | grep 'ok.*coverage' | awk '{sub("github.com/FlowingSPDG/get5loader/", "", $$2); print "| " $$2 " | " $$5 " |"}' | tee -a coverage.txt && \
	echo "## Test Coverage Report" > coverage_with_header.txt && \
	echo "| Package           | Coverage |" >> coverage_with_header.txt && \
	echo "|-------------------|----------|" >> coverage_with_header.txt && \
	cat coverage.txt >> coverage_with_header.txt && \
	mv coverage_with_header.txt coverage.txt && \
	cat coverage.txt