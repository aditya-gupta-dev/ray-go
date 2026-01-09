package main

import (
	"fmt"
	"ray-go/window"
	"unsafe"
)

type LifeCycle struct{}

func main() {

	var window = window.Window{}

	fmt.Println(unsafe.Sizeof(window))
}
