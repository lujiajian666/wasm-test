## 	1. Introduce

WebAssembly is a new type of code that can be run in modern web browsers — it is a low-level assembly-like language with a compact binary format that runs with near-native performance and provides languages such as C/C++, C# and Rust with a compilation target so that they can run on the web. It is also designed to run alongside JavaScript, allowing both to work together. [MDN WebAssembly](https://developer.mozilla.org/en-US/docs/WebAssembly)

## 2.  Workflow - js/webAssembly

![	](img/image-20221005170809780.png)
<p style="text-align: center; color: red"><a href='https://medium.com/dailyjs/understanding-v8s-bytecode-317d46c94775'>v8 js workflow(1)</a></p>


​	<img src="./img/jsflow.png" alt="js" style="zoom:200%;" />

<p style="text-align: center; color: red">v8 js workflow(2)</p>

![webassembly](img/webAssemblyflow.png)

<p style="text-align: center; color: red">wasm workflow(<a href='https://v8.dev/docs/wasm-compilation-pipeline'>[TurboFan & WASM]</a>)</p>

## 3. Advantage

- Smaller -- It takes less time to extract wasm because it is more compact than JavaScript even when compressed.
- Faster
  - wasm is closer to machine code than JavaScript and has been optimized on the server side.
  - No re-optimization is required, because wasm has types and other information built in, so the JS engine doesn't need to speculate when optimizing the way JavaScript is done.
  - <span style="color: red">Garbage collection is not required since memory is managed manually</span>
  - Execution usually takes less time, because there are fewer compiler tricks and pitfalls developers need to know to write consistent high-performance code, and wasm's instruction set is more machine-friendly.
- <span style="color: green">Step across FE and BE</span>

## 4. Usage

go version -- go1.19.3

chrome version -- 107.0.5304.87（正式版本） (x86_64)

[demo](http://127.0.0.1:8000)

## 5. Analyze

The actual running result, js performs better than wasm

+ Size-wise, code that works equally well. wasm size 1.4M, js size 705B
+ In terms of speed, the same code works. Wasm of go runs slower than js under chrome, but the opposite under safari. Wasm of c is the most faster both on safari and chrome
+ In memory, the inability to collect garbage will lead to a large memory footprint! 

## 6.  Conclusion

Wasm is not necessarily faster than javascript, but it can reuse code.

## 7. Links

[build wasm module with c -- akarin.dev](https://akarin.dev/2020/10/10/build-webassembly-module-with-c/)

[webassembly -- webassembly.org](https://webassembly.org/)

[go webassembly -- github.com](https://github.com/golang/go/wiki/WebAssembly)

[syscall/js -- golang.org](https://pkg.go.dev/syscall/js)

[https://emscripten.org/docs/porting/connecting_cpp_and_javascript/Interacting-with-code.html](https://emscripten.org/docs/porting/connecting_cpp_and_javascript/Interacting-with-code.html)