package main

import (
	"fmt"
	"math"
	"unsafe"
)

// 整型分两大类
// 	 - 有符号整型：int, int8, int16, int32, int64
// 	 - 无符号整型：uint, uint8, uint16, uint32, uint64

func integer() {
	var num int = math.MaxInt
	var num8 int8 = 11
	var num16 int16 = 12
	var num32 int32 = math.MaxInt32
	var num64 int64 = math.MaxInt64
	fmt.Printf("num 的值是 %d， num 的类型是 %T，num 的大小是 %d\n", num, num, unsafe.Sizeof(num))
	fmt.Printf("num8 的值是 %d， num8 的类型是 %T，num8 的大小是 %d\n", num8, num8, unsafe.Sizeof(num8))
	fmt.Printf("num16 的值是 %d， num16 的类型是 %T，num16 的大小是 %d\n", num16, num16, unsafe.Sizeof(num16))
	fmt.Printf("num32 的值是 %d， num32 的类型是 %T，num32 的大小是 %d\n", num32, num32, unsafe.Sizeof(num32))
	fmt.Printf("num64 的值是 %d， num64 的类型是 %T，num64 的大小是 %d\n", num64, num64, unsafe.Sizeof(num64))
}

func unsignedInteger() {
	var num uint = math.MaxUint
	var num8 uint8 = 11
	var num16 uint16 = 12
	var num32 uint32 = math.MaxUint32
	var num64 uint64 = math.MaxUint64
	fmt.Printf("num 的值是 %d， num 的类型是 %T，num 的大小是 %d\n", num, num, unsafe.Sizeof(num))
	fmt.Printf("num8 的值是 %d， num8 的类型是 %T，num8 的大小是 %d\n", num8, num8, unsafe.Sizeof(num8))
	fmt.Printf("num16 的值是 %d， num16 的类型是 %T，num16 的大小是 %d\n", num16, num16, unsafe.Sizeof(num16))
	fmt.Printf("num32 的值是 %d， num32 的类型是 %T，num32 的大小是 %d\n", num32, num32, unsafe.Sizeof(num32))
	fmt.Printf("num64 的值是 %d， num64 的类型是 %T，num64 的大小是 %d\n", num64, num64, unsafe.Sizeof(num64))
}

func main() {
	// 有符号整型
	integer()

	fmt.Println("==================================")

	// 无符号整型
	unsignedInteger()
}
