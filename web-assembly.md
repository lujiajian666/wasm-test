## 介绍

WebAssembly 是一种新的编码方式，可以在现代的网络浏览器中运行 － 它是一种低级的类汇编语言，具有紧凑的二进制格式，可以接近原生的性能运行，并为诸如 C / C ++等语言提供一个编译目标，以便它们可以在 Web 上运行。它也被设计为可以与 JavaScript 共存，允许两者一起工作。 [MDN WebAssembly](https://developer.mozilla.org/zh-CN/docs/WebAssembly)

## js/webAssembly工作方式

![js](img/image-20221005170809780.png)

​	<img src="./img/jsflow.png" alt="js" style="zoom:200%;" />

<p style="text-align: center; color: red">js 即时编译流程</p>

![webassembly](img/webAssemblyflow.png)

<p style="text-align: center; color: red">web Assembly 流程</p>

## 优点

- 提取WebAssembly所需的时间更少，因为即使压缩后，它也比JavaScript更紧凑。
- 解码WebAssembly所需的时间少于解析JavaScript所需的时间。
- 编译和优化所需的时间更少
  - WebAssembly比JavaScript更接近机器代码，并且已经在服务器端进行了优化。
  - 不需要进行重新优化，因为WebAssembly内置了类型和其他信息，因此JS引擎在优化JavaScript方式时无需推测。
- 执行通常会花费较少的时间，因为开发人员在编写一致的高性能代码时需要知道的编译器技巧和陷阱就更少了，另外WebAssembly的指令集更适合机器。
- 由于内存是手动管理的，因此不需要垃圾回收。

## 实际运行情况

[demo](./static/index.html)

## 相关链接

[go webassembly 文档 -- github.com](https://github.com/golang/go/wiki/WebAssembly)

[syscall/js 官方文档 -- golang.org](https://pkg.go.dev/syscall/js)