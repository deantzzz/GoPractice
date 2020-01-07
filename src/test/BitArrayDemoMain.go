package main

import (
	"fmt"
	"gitgo/src/test/geometry"
)

func main() {
	var x, y geometry.IntSet
	x.Add(0)
	x.Add(1)
	x.Add(63)
	x.Add(64)
	x.Add(65)
	x.Add(144)
	x.Add(9)
	x.AddAll(5, 6, 7)
	fmt.Println(x.String())
	fmt.Println(x.Len())
	d := x.Copy()
	fmt.Println(d)
	y.Add(9)
	y.Add(42)
	fmt.Println(y.String())

	x.UnionWith(&y)
	fmt.Println(x.String())
	fmt.Println(x.Has(0), x.Has(1), x.Has(63), x.Has(64), x.Has(65), x.Has(42), x.Has(144))

}
