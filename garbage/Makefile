all: main.wasm ./rust-wasm/pkg/rust_wasm.js ./emsdk/indexc.wasm
	npx anywhere -d . --proxy ../proxy.config.js
main.wasm: wasm_exec.js index.go ../proxy.config.js
	GOOS=js GOARCH=wasm go build -ldflags="-s -w" -gcflags='-m' -o main.wasm
wasm_exec.js:
	cp $(shell go env GOROOT)/misc/wasm/wasm_exec.js ./
./rust-wasm/pkg/rust_wasm.js: ./rust-wasm/src/lib.rs
	cd ./rust-wasm && wasm-pack build --target web --features wee_alloc
./emsdk/indexc.wasm: ./emsdk/index/index.c
	cd ./emsdk && source ./emsdk_env.sh && emcc index.c -O3 -s SIDE_MODULE=1 -o indexc.wasm