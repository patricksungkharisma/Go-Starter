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

.PHONY: test
test:
	@go test --race $$(go list ./... | grep -v vendor | grep -v proto | grep -v cmd)

.PHONY: clear-postgres-data
clear-postgres-data:
	@echo "Deleting docker postgres data..."
	@rm -rf .dev/.docker/postgres-data
	@echo "Deletion complete!"

