.PHONY: generate

generate:
	@go generate ./...
	@echo "[OK] Files added to embed box!"

security:
	@gosec ./...
	@echo "[OK] Go security check was completed!"

build: generate security
	@go build -o ./build/server ./cmd/app/*.go
	@echo "[OK] App binary was created!"

run:
	@./build/server
