package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	defer printStack()
	f1(3)
}

func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
}

func f1(x int) {
	fmt.Printf("f(%d)\n", x+0/x)
	defer fmt.Printf("defer %d\n", x)
	f1(x - 1)
}
