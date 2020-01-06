package main

import (
	"fmt"
)

func main() {
	//fmt.Println(check())
	//a:=make([]string,10)
	a := []string{"a", "b", "c", "d"}
	//join := strings.Join(a, "c")
	//fmt.Println(join)
	VarAgesJoin(a, "aaa", "vvv")
	fmt.Println(a)
}

func check(vals ...int) (max, min int) {
	if len(vals) == 0 {
		return
	}
	max, min = vals[0], vals[0]
	for _, val := range vals {
		if val > max {
			max = val
		}
		if val < min {
			min = val
		}
	}
	return
}

func VarAgesJoin(a []string, seqs ...string) {
	if len(seqs) == 0 {
		return
	}
	for _, val := range a {

	}
}
