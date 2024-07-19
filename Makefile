.PHONY: init
init:
	chmod +x .dev/initialize.sh
	.dev/initialize.sh

.PHONY: deps-up
deps-up:
	@docker-compose -f .dev/docker-compose.yaml up -d

.PHONY: deps-down
deps-down:
	@docker-compose -f .dev/docker-compose.yaml down

.PHONY: start-app
start-app:
	@air -c .dev/.app.air.toml

.PHONY: clear-postgres-data
clear-postgres-data:
	@echo "Deleting docker postgres data..."
	@rm -rf .dev/.docker/postgres-data
	@echo "Deletion complete!"

# .PHONY: test
# test:
# 	@which gotest 2>/dev/null || go get -v github.com/rakyll/gotest
# 	@gotest --race $$(go list ./... | grep -v vendor | grep -v proto | grep -v cmd)