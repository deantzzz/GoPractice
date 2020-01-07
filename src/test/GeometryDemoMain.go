package main

import (
	"fmt"
	"gitgo/src/test/geometry"
	"image/color"
)

func main() {
	q := geometry.Point{1, 2}
	p := geometry.Point{4, 6}
	fmt.Println(geometry.Distance(p, q))
	fmt.Println(p.Distance(q))
	m := geometry.Point.Distance
	fmt.Println(m(q, p))

	//三角形周长
	perim := geometry.Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	fmt.Println(perim)
	perim.ChangeV()
	fmt.Println(perim)
	fmt.Println(perim.Distance())

	//map[string][]string
	vs := geometry.Values{
		"Aa": {"A1", "A2"},
		"Bb": {"B1", "B2", "B3"},
	}
	fmt.Println(vs)
	vs.Add("Aa", "A3")
	fmt.Println(vs)
	fmt.Println(vs.GetFirst("Aa"))

	vs.ChangeMpV("Bb", []string{"c1", "c2"})
	fmt.Println(vs)
	vs.ChangeSV("Aa", "c3")
	fmt.Println(vs)

	fmt.Println(vs.GetFirst("Aa"))

	//修改vs的引用指向一个新的Values对象，当前调用方不会受任何影响
	//因为传入的是存储了内存地址的变量，你改变这个变量本身是影响不了原始的变量的
	vs.NewV()
	fmt.Println(vs)

	//6.3. 通过嵌入结构体来扩展类型
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	var p1 = geometry.ColoredPoint{&geometry.Point{1, 1}, red}
	var q1 = geometry.ColoredPoint{&geometry.Point{5, 4}, blue}

	fmt.Println(q1)
	fmt.Println(p1)
	fmt.Println(p1.Distance(*q1.Point))
	fmt.Println(q1)
	fmt.Println(p1)

	p1.ScaleBy(2)
	fmt.Println(q1)
	fmt.Println(p1)
	q1.ScaleBy(2)
	fmt.Println(q1)
	fmt.Println(p1)
	fmt.Println(p1.Distance(*q1.Point))
	fmt.Println(q1)
	fmt.Println(p1)

	q1.Point = p1.Point
	p1.ScaleBy(2)
	fmt.Println(*q1.Point)
	fmt.Println(*p1.Point)

}
