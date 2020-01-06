package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	//bigSlowOperation被调时，trace会返回一个函数值，该函数值会在bigSlowOperation退出时被调用
	//bigSlowOperation()
	fmt.Println(double1(2))
}

func bigSlowOperation() {
	defer trace("bigSlowOperation")()
	time.Sleep(10 * time.Second)
}

func trace(s string) func() {
	start := time.Now()
	log.Printf("enter %s", s)
	return func() {
		log.Printf("exit %s (%s)", s, time.Since(start))
	}
}

func double(x int) int {
	return x + x
}

func double1(x int) (result int) {
	defer func() {
		fmt.Println(x, result)
		result += 1
	}()
	return double(x)
}
