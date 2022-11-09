all: main.wasm ./emsdk/indexc.wasm
	npx anywhere -d . --proxy ../proxy.config.js
main.wasm: wasm_exec.js index.go ../proxy.config.js
	GOOS=js GOARCH=wasm go build -ldflags="-s -w" -gcflags='-m' -o main.wasm
wasm_exec.js:
	cp $(shell go env GOROOT)/misc/wasm/wasm_exec.js ./
./emsdk/indexc.wasm: ./emsdk/index.c
	cd ./emsdk && source ./emsdk_env.sh && emcc index.c -O3 -s SIDE_MODULE=1 -o indexc.wasm