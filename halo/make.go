package main

import "fmt"

func main() {

	// make 只分配和初始化类型为 slice、map 或 chan 的对象
	// 第一个参数：类型，与 new 不同，make 的返回类型与其参数类型相同，而不是指针
	// 第二个参数：长度，如果是 slice 或 map，则是容量，如果是 chan，则是缓冲区大小
	// 第三个参数：容量，如果是 slice 或 map，则是长度，如果是 chan，则是缓冲区大小

	a := make([]int, 6, 10)
	fmt.Println(a)

	b := make(map[string]int, 10)
	fmt.Println(b)

	c := make(chan int, 10)
	fmt.Println(c)
}
