all: main.wasm ./rust-wasm/pkg/rust_wasm.js
	npx anywhere -d . --proxy ../proxy.config.js
main.wasm: wasm_exec.js index.go ../proxy.config.js
	GOOS=js GOARCH=wasm go build -ldflags="-s -w" -gcflags='-m' -o main.wasm
wasm_exec.js:
	cp $(shell go env GOROOT)/misc/wasm/wasm_exec.js ./
./rust-wasm/pkg/rust_wasm.js: ./rust-wasm/src/lib.rs
	cd ./rust-wasm && wasm-pack build --target web