package main

import (
	"fmt"
	"unicode"
)

func main() {
	//slice reverse
	a := []int{1, 2, 3, 4, 5}
	reverseSlice(a)
	fmt.Printf("slice reverse: %v\n", a) // [5 4 3 2 1]

	//array reverse
	b := [5]int{1, 2, 3, 4, 5}
	reverseArray(&b)
	fmt.Printf("array reverse: %v\n\n", b)

	//向左旋转n个元素
	n := 2
	s := []int{0, 1, 2, 3, 4, 5}
	reverseSlice(s[:n])
	reverseSlice(s[n:])
	reverseSlice(s)
	fmt.Printf("rotate 2: %v\n\n", s) // [2 3 4 5 0 1]

	//原地消除相邻重复str
	str := []string{"aa", "bb", "bb", "bb", "aa", "aa", "cc"}
	str = removeDupStr(str)
	fmt.Printf("remove adjacent duplicate string: %v\n\n", str)

	//相邻空格全部合并为一个空格
	byt := []byte("asd 2123      hahaasdsad              x  1    2")
	fmt.Printf("before merge Space: %s\n", byt)
	byt = mergeSpace(byt)
	fmt.Printf("after merge Space: %s\n\n", byt)
}

//反转 reverse
func reverseSlice(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

//重写reverse函数，使用数组指针代替slice。
func reverseArray(s *[5]int) {
	for i, j := 0, len(*s)-1; i < j; i, j = i+1, j-1 {
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	}
}

//旋转
func rotate(s []int, n int) {

}

//原地消除[]string中相邻重复字符串
func removeDupStr(s []string) []string {
	t := s[:0]
	var oldVal string
	for _, val := range s {
		if val != oldVal {
			t = append(t, val)
		}
		oldVal = val
	}
	return t
}

//编写一个函数，
//原地将一个UTF-8编码的[]byte类型的slice中相邻的空格替换成一个空格返回
//（参考unicode.IsSpace)
func mergeSpace(s []byte) []byte {
	t := s[:0]
	var countSpace int
	for _, val := range s {
		if unicode.IsSpace(rune(val)) {
			countSpace++
			if countSpace == 1 {
				t = append(t, val)
			}
			continue
		}
		t = append(t, val)
		countSpace = 0
	}
	return t
}
