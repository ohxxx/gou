package main

import "fmt"

func main() {
	// 声明一个变量, 默认的值是0
	var a int
	fmt.Println("a 的值是：", a)
	fmt.Printf("a 的类型：%T\n", a)

	fmt.Println("==================================")

	// 声明一个变量，并初始化值
	var b int = 6
	fmt.Println("b 的值是：", b)
	fmt.Printf("b 的类型：%T\n", b)

	fmt.Println("==================================")

	// 声明一个变量，省去数据类型
	var c = "666"
	fmt.Println("c 的值是：", c)
	fmt.Printf("c 的类型：%T\n", c)

	fmt.Println("==================================")

	// 短声明，使用 := 定义变量
	// 短声明是在函数或方法内部使用，不支持全局变量声明
	d := 888
	fmt.Println("d 的值是：", d)
	fmt.Printf("d 的类型：%T\n", d)

	fmt.Println("==================================")

	// 多变量声明
	var e, f, g int = 1, 2, 3
	var h, i, j = "halo", "xxx", "gogogo"
	var (
		k = 666
		l = "xxx"
	)
	fmt.Println(e, f, g, h, i, j, k, l)
}
