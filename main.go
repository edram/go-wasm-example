package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	jsonFunc := js.FuncOf(func(this js.Value, args []js.Value) any {
		return "213"
	})

	js.Global().Set("formatJSON", jsonFunc)

	fmt.Println("123")
}
