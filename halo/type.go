package main

import "fmt"

func main() {
	// 静态类型
	var a int    // int 是静态类型
	var b string // string 是静态类型

	fmt.Println(a)
	fmt.Println(b)

	// 动态类型 - 程序运行时系统才能看到的类型
	var c interface{} // c 的静态类型是 interface{}
	c = 100           // c 的静态类型是 interface{} 动态类型是 int
	c = "halo xxx"    // c 的静态类型是 interface{} 动态类型是 string
	fmt.Println(c)

	// 接口组成
	d := (int)(666)
	fmt.Printf("d type = %T，value = %v\n", d, d)
}
