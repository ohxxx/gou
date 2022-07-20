package main

import (
	"fmt"
	"math"
	"unsafe"
)

func main() {
	var num1 float32 = math.MaxFloat32
	var num2 float64 = math.MaxFloat64
	fmt.Printf("num1 的值是 %g， num1 的类型是 %T，num1 的大小是 %d\n", num1, num1, unsafe.Sizeof(num1))
	fmt.Printf("num2 的值是 %g， num2 的类型是 %T，num2 的大小是 %d\n", num2, num2, unsafe.Sizeof(num2))
}
