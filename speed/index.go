package main

import (
	"fmt"
	"syscall/js"
	"crypto/md5"
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
	sum := 0
	for i := 0; i < 30; i++ {
		sum += fib(max)
	}
	return sum
}
func calcMd5(this js.Value, args []js.Value) interface{} {
	ret := fmt.Sprintf("%x", md5.Sum([]byte(args[0].String())))
	return js.ValueOf(ret)
}

func main() {
	alert := js.Global().Get("alert")
	alert.Invoke("go webAssembly!")
	// 提供模块
	js.Global().Set("goFib", js.FuncOf(fibWrapper))
	js.Global().Set("goFib30Times", js.FuncOf(fib30TimesWrapper))
	js.Global().Set("CalcMd5", js.FuncOf(calcMd5))

	done := make(chan int, 0)
	<-done
}