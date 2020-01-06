//编写一个程序，
//默认情况下打印标准输入的SHA256编码，
//并支持通过命令行flag 定制，输出SHA384或SHA512哈希算法。
package main

import (
	"crypto/sha256"
	sha5122 "crypto/sha512"
	"flag"
	"fmt"
)

//第一个是的命令行标志参数的名字“xxx”，
//然后是该标志参数的默认值（这里是false），
//最后是该标志参数对 应的描述信息。
var sha384 = flag.Bool("sha384", false, "sha384")
var sha512 = flag.Bool("sha512", false, "sha512")

func main() {
	flag.Parse()

	args := flag.Args()
	if *sha384 {
		fmt.Printf("sha384: %x\n", sha5122.Sum384([]byte(args[0])))
	} else if *sha512 {
		fmt.Printf("sha512: %x\n", sha5122.Sum512([]byte(args[0])))
	} else {
		fmt.Printf("sha256: %x\n", sha256.Sum256([]byte(args[0])))
	}
}
