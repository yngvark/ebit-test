.PHONY: help
help: ## Print this menu
	@grep -E '^[a-zA-Z_0-9-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

build: build/app.wasm ## -

build/app.wasm: ## -
	env GOOS=js GOARCH=wasm go build -o build/app.wasm ./main.go

.PHONY: run-go
run-go: ## Run natively
	export PKG_CONFIG_PATH=/usr/lib/x86_64-linux-gnu/pkgconfig
	go run .

.PHONY: run
run: ## Run in Docker, and with Webassembly
	docker build . -t game
	(docker stop game && docker rm game) || true
	docker run --name game -d -p 8081:8080 game
	@echo Now open: http://localhost:8081
