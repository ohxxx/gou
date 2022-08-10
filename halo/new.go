package main

import "fmt"

func main() {
	// 内置函数 new 分配内存

	// 这个新的内置函数分配内存。第一个参数是一个类型。
	// 第一个参数是一个类型，而不是一个值，返回的值是一个指针，指向一个新的
	// 分配的该类型的零值的指针。
	// func new(Type) *Type

	x := new(int)
	fmt.Println(x)
	fmt.Println(*x)

	// new 为所有类型分配内存，并初始化为零值，返回指针
}
