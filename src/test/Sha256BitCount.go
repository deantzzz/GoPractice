//编写一个函数，
//计算两个SHA256哈希码中不同bit的数目。
//（参考2.6.2节的 PopCount函数。)
package main

import (
	"crypto/sha256"
	"fmt"
	"gitgo/src/test/popcount"
)

func main() {
	//sha256
	sha1 := sha256.Sum256([]byte("X"))
	fmt.Println(sha1)
	fmt.Printf("%x\n %b\n %T", sha1, sha1, sha1)

	sha2 := sha256.Sum256([]byte("Y"))

	var counts = make(map[int]int)

	for _, val := range sha1 {
		counts[popcount.PopCount(uint64(val))]++
	}
	fmt.Println(counts)

	for _, val := range sha2 {
		counts[popcount.PopCount(uint64(val))]--
	}

	fmt.Println(counts)

}
