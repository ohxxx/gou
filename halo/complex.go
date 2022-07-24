package main

import "fmt"

func main() {
	var xxx complex64 = complex(1, 2)
	var yyy complex128 = complex(3, 4)
	var zzz complex128 = complex(5, 6)
	fmt.Println("xxx = ", xxx)
	fmt.Println("yyy = ", yyy)
	fmt.Println("zzz = ", zzz)

	xxx1 := 1 + 2i
	yyy1 := 3 + 4i
	zzz1 := 5 + 6i
	fmt.Println("xxx1 = ", xxx1)
	fmt.Println("yyy1 = ", yyy1)
	fmt.Println("zzz1 = ", zzz1)

	// real 返回复数的实部
	fmt.Println("real(xxx) = ", real(xxx))
	// imag 返回复数的虚部
	fmt.Println("imag(xxx) = ", imag(xxx))
	fmt.Println("yyy * zzz = ", yyy*zzz)
}
