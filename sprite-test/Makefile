.PHONY: help
help: ## Print this menu
	@grep -E '^[a-zA-Z_0-9-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: run-go
run-go: ## Run natively
	export PKG_CONFIG_PATH=/usr/lib/x86_64-linux-gnu/pkgconfig
	go run .

.PHONY: run
run: ## Run in Webassembly
	docker build . -t game
	docker run -d -p 8081:80 -v $(pwd)://usr/share/nginx/html game
	@echo Now open: http://localhost:8081
