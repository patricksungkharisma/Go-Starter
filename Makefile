# Init Project
.PHONY: init
init:
	@echo "=== Initialize Requirements ==="
	@chmod +x .dev/initialize.sh
	@.dev/initialize.sh

# Setup Dependencies
.PHONY: download
download:
	@echo "=== Downloading Dependencies ==="
	@go mod download
	@sleep 1
	@echo "=== Removing Vendor Folder ==="
	@rm -rf vendor
	@echo "=== Setup Vendor in Go Pkg Cache ==="
	@go mod tidy
	@sleep 1
	@echo "=== Dependencies Setup All OK ==="

# Start App
.PHONY: start-app
start-app:
	@echo "=== Running Program ==="
	@sudo chown -R $$(id -u):$$(id -g) .dev
	@echo "=== DOCKER UP ==="
	@docker-compose -p app -f .dev/docker-compose.yaml up --build ${SERVICE}

# Stop App
.PHONY: stop-app
stop-app:
	@docker-compose -p app -f .dev/docker-compose.yaml down

# Unit Test
.PHONY: Test
test:
	@which gotest 2>/dev/null || go get -v github.com/rakyll/gotest
	@gotest --race $$(go list ./... | grep -v vendor | grep -v proto | grep -v cmd)