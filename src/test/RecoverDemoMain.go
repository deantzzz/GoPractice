//**练习5.19：** 使用panic和recover编写一个不包含return语句但能返回一个非零值的函数。
package main

import "fmt"

func main() {
	defer func() {
		if p := recover(); p != nil {
			fmt.Print(p)
		}
	}()
	recv(10)
}

func recv(x int) {
	panic(x)

}
