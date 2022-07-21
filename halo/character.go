package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// 字符串中每一个元素叫做"字符"，定义字符时使用单引号
	// Go 语言的字符有两种
	// 	 1、byte 类型、1个字节数、表示 UTF-8 字符串的单个字节的值，表示的是 ASCII 码表中的一个字符， uint8 的别名类型
	// 	 2、rune 类型、4个字节数、表示单个 unicode 字符，int32 的别名类型

	var x byte = 66
	var y rune = 66
	// var z byte = '棒' // error
	// 中文的话，就得使用 rune 类型

	fmt.Printf("x = %c, y = %c\n", x, y) // x = B, y = B
	fmt.Printf("x 占用 %d 个字节，y 占用 %d 个字节\n", unsafe.Sizeof(x), unsafe.Sizeof(y))
}
