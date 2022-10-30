run: main.wasm
	npx anywhere -d . --proxy ../proxy.config.js
main.wasm: ../static/wasm_exec.js
	GOOS=js GOARCH=wasm go build -ldflags="-s -w" -gcflags='-m' -o main.wasm
../static/wasm_exec.js:
	cp $(shell go env GOROOT)/misc/wasm/wasm_exec.js ../static