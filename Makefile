build: static/wasm_exec.js
	GOOS=js GOARCH=wasm go build -o static/main.wasm
static/wasm_exec.js:
	cp $(shell go env GOROOT)/misc/wasm/wasm_exec.js static