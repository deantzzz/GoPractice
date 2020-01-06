package main

import (
	"fmt"
	"gitgo/src/test/popcount"
)

func main() {
	fmt.Println(popcount.PopCount(6))
	fmt.Println(popcount.PopCountWithForLoop(6))
	fmt.Println(popcount.PopCountWithBitMove(6))
	fmt.Println(popcount.PopCountWithCleanLowest(6))
}
