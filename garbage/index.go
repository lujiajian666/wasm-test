package main

import (
	"syscall/js"
)

func generateGarbage(this js.Value, args []js.Value) any {
	list := make([]interface{}, 1000000);
	// js.Global().Set("list", js.ValueOf(list))
	return list
}

func main() {
	alert := js.Global().Get("alert")
	alert.Invoke("go webAssembly!")
	// 提供模块
	js.Global().Set("goGenGar", js.FuncOf(generateGarbage))
	
	done := make(chan int, 0)
	<-done
}