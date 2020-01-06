//练习 4.8： 修改charcount程序，使用unicode.IsLetter等相关的函数，统计字母、数字等 Unicode中不同的字符类别。
//练习 4.9： 编写一个程序wordfreq程序，报告输入文本中每个单词出现的频率。在第一次调用 Scan前先调用input.Split(bufio.ScanWords)函数，这样可以按单词而不是按行输入。
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"unicode"
)

func main() {
	wordfreq()
}

//练习 4.8： 修改charcount程序，使用unicode.IsLetter等相关的函数，统计字母、数字等 Unicode中不同的字符类别。
func charcount() {
	counts := make(map[string]int, 2)
	invalid := 0

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(os.Stderr, "%v\n", err)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		if 48 <= r && r <= 97 {
			counts["number"]++
		} else if r > 97 {
			counts["string"]++
		}
		fmt.Printf("%v\n", counts)
	}

	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")

	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}

//练习 4.9： 编写一个程序wordfreq程序，报告输入文本中每个单词出现的频率。在第一次调用 Scan前先调用input.Split(bufio.ScanWords)函数，这样可以按单词而不是按行输入。
func wordfreq() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		word := input.Text()
		counts[word]++
	}
	for k, v := range counts {
		fmt.Printf("%s: %s\n", k, v)
	}
}
