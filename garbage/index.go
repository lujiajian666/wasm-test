package main

import (
	"syscall/js"
)

func fib(this js.Value, args []js.Value) any {
	max := args[0].Int();
	slice := []int{1, 1}
	if max < 3 {
		return slice[max - 1]
	}
	for i := 2; i < max; i++ {
		slice[0], slice[1] = slice[1], slice[0]+ slice[1]
	}

	return slice[1]
}
func fib2(max int) int {
	if max < 2 {
		return 1
	}
	return fib2(max - 1) + fib2(max - 2)
}
func fib2wrapper(this js.Value, args []js.Value) any {
	max := args[0].Int();
	return fib2(max)
}
func fib3wrapper(this js.Value, args []js.Value) any {
	max := args[0].Int();
	for i := 0; i < 30; i++ {
		fib2(max)
	}
	return "end"
}

func generateGarbage(this js.Value, args []js.Value) any {
	time := args[0].Int();
	len := args[1].Int();
	var list []interface{}
	for i := 0; i < time; i++ {
		list = make([]interface{}, len)
	}
	return list
}

func main() {
	alert := js.Global().Get("alert")
	alert.Invoke("go webAssembly!")
	// 提供模块
	js.Global().Set("goFib", js.FuncOf(fib))
	js.Global().Set("goFib2", js.FuncOf(fib2wrapper))
	js.Global().Set("goFib3", js.FuncOf(fib3wrapper))
	js.Global().Set("goGenGar", js.FuncOf(generateGarbage))
	
	done := make(chan int, 0)
	<-done
}