package main

import "fmt"

/*
占位符		说明
%s		直接输出字符串或者[]byte
%q		该值对应的双引号括起来的go语法字符串字面值，必要时会采用安全的转义表示
%x		每个字节用两字符十六进制数表示（使用a-f
%X		每个字节用两字符十六进制数表示（使用A-F）
*/

func main() {
	s := "小王子"
	fmt.Printf("%s\n", s)
	fmt.Printf("%q\n", s)
	fmt.Printf("%x\n", s)
	fmt.Printf("%X\n", s)
}
