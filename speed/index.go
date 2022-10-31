package main

import (
	"syscall/js"
)

func fib(max int) int {
	if max < 2 {
		return 1
	}
	return fib(max - 1) + fib(max - 2)
}
func fibWrapper(this js.Value, args []js.Value) any {
	max := args[0].Int();
	return fib(max)
}
func fib30TimesWrapper(this js.Value, args []js.Value) any {
	max := args[0].Int();
	for i := 0; i < 30; i++ {
		fib(max)
	}
	return "end"
}

func main() {
	alert := js.Global().Get("alert")
	alert.Invoke("go webAssembly!")
	// 提供模块
	js.Global().Set("goFib", js.FuncOf(fibWrapper))
	js.Global().Set("goFib30Times", js.FuncOf(fib30TimesWrapper))
	
	done := make(chan int, 0)
	<-done
}