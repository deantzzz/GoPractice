package geometry

import (
	"image/color"
	"math"
	"sync"
)

//point
type Point struct {
	X, Y float64
}

//包级别函数
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

//Point类的方法
func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p Point) Add(q Point) Point {
	return Point{p.X + q.X, p.Y + q.Y}
}

func (p Point) Sub(q Point) Point {
	return Point{p.X - q.X, p.Y - q.Y}
}

// path
type Path []Point

func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

func (path Path) TranslateBy(offset Point, add bool) {
	var op func(p, q Point) Point
	if add {
		op = Point.Add
	} else {
		op = Point.Sub
	}
	for i := range path {
		op(path[i], offset)
	}
}

func (path Path) ChangeV() {
	for i := range path {
		path[i] = Point{5, 5}
	}
}

//map
// Values maps a string key to a list of values.
//type Values map[string][]string
type Values map[string][]string

func (v Values) NewV() {
	v = Values{
		"Dd":   {"D1", "D2"},
		"Dddd": {"D1", "D2", "D3"},
	}
}

// Get returns the first value associated with the given key,
// or "" if there are none.
func (v Values) GetFirst(key string) string {
	if vs := v[key]; len(vs) > 0 {
		return vs[0]
	}
	return ""
}

// Add adds the value to key.
// It appends to any existing values associated with key.
func (v Values) Add(key, value string) {
	v[key] = append(v[key], value)
}

func (v Values) ChangeMpV(key string, value []string) {
	v[key] = value
}

func (v Values) ChangeSV(key, value string) {
	v[key][0] = value
}

type ColoredPoint struct {
	*Point
	Color color.RGBA
}

//cache lock
var (
	mu      sync.Mutex
	mapping = make(map[string]string)
)

func LockUp(key string) string {
	mu.Lock()
	v := mapping[key]
	mu.Unlock()
	return v
}

var cache = struct {
	sync.Mutex
	mapping map[string]string
}{
	mapping: make(map[string]string),
}

func LockUp1(key string) string {
	cache.Lock()
	v := cache.mapping[key]
	cache.Unlock()
	return v
}
