mock-generate: ## Generate mocks
	go mod vendor
	docker run --rm -v "$(PWD):/app" -w /app -t vektra/mockery --all --dir ./internal/user/domain --case underscore
	docker run --rm -v "$(PWD):/app" -w /app -t vektra/mockery --all --dir ./internal/chat --case underscore
