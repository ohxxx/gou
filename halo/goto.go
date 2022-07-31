package main

import "fmt"

func main() {
	// goto 语句
	// 可以跳转到指定的标签

	fmt.Println("halo xxx")
	goto label
	// goto 语句下面的语句不会执行
	// goto 语句与标签之间不能有变量声明，否则编译报错
	fmt.Println("halo go") // 不会执行
	//var name = "xxx" // error: name declared and not used

label:
	fmt.Println("halo gogo")

}
