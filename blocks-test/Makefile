.PHONY: help
help: ## Print this menu
	@grep -E '^[a-zA-Z_0-9-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

run:
	env GOOS=js GOARCH=wasm go build -o blocks.wasm ./main.go
	docker build . -t blocks
	docker run -p 8081:80 -v ~/yk/git/hajimehoshi/ebiten/examples/blocks://usr/share/nginx/html blocks
