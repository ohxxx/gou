package main

import "fmt"

func main() {
	// 常量（只读属性）
	const length int = 10
	fmt.Println("length 的值是：", length)

	fmt.Println("==================================")

	// const 来定义枚举类型
	const (
		BEIJING   = "北京"
		SHANGHAI  = "上海"
		GUANGZHOU = "广州"
		ShenZhen  = "深圳"
	)
	fmt.Println("BEIJING = ", BEIJING)
	fmt.Println("SHANGHAI = ", SHANGHAI)
	fmt.Println("GUANGZHOU = ", GUANGZHOU)
	fmt.Println("ShenZhen = ", ShenZhen)

	fmt.Println("==================================")

	// iota 是 Go 语言的常量计数器，只能在常量的表达式中使用
	// iota 在 const 关键字出现时将被重置为 0 (const 内部的第一行之前)
	// const 中每新增一行常量声明将使 iota 计数一次 (iota 可理解为 const 语句块中的行索引)
	// 使用 iota 能简化定义
	const (
		aaa = 0
		bbb
		ccc = 10 * iota
		ddd
	)
	fmt.Println("aaa = ", aaa)
	fmt.Println("bbb = ", bbb)
	fmt.Println("ccc = ", ccc)
	fmt.Println("ddd = ", ddd)

}
