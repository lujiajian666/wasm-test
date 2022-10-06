run: static/main.wasm
	npx anywhere -d static --proxy proxy.config.js
static/main.wasm: static/wasm_exec.js
	GOOS=js GOARCH=wasm go build -ldflags="-s -w" -o static/main.wasm
static/wasm_exec.js:
	cp $(shell go env GOROOT)/misc/wasm/wasm_exec.js static