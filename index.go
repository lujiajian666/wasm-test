package main

import (
	"syscall/js"
	// "fmt"
)

func fib(this js.Value, args []js.Value) interface{} {
	max := args[0].Int();
	slice := []int{1, 2}
	if max < 3 {
		return slice[max - 1]
	}
	for i := 2; i < max; i++ {
		slice[0], slice[1] = slice[1], slice[0]+ slice[1]
	}

	return js.ValueOf(slice[1])
}

func main() {
	alert := js.Global().Get("alert")
	alert.Invoke("go webAssembly!")
	// 提供模块
	done := make(chan int, 0)
	js.Global().Set("goFib", js.FuncOf(fib))
	<-done
}