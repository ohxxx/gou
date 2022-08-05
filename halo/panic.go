package main

import "fmt"

func main() {
	//panic("halo xxx")

	//defer fmt.Println("main：halo xxx")
	//test()

	// recover()
	// 可以恢复 panic 引发的异常
	arr(8)                    // panic: runtime error: index out of range
	fmt.Println("gogogogogo") // 这里会执行，被捕获了
}

func test() {
	defer fmt.Println("func：halo xxx")
	panic("halo go")
}

func arr(a int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("arr：", err)
		}
	}()
	var arr [6]int
	arr[a] = 666
}
