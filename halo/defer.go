package main

import "fmt"

func halo() {
	fmt.Println("halo")
}

func gogo() {
	fmt.Println("gogo")
}

func main() {
	// defer 语句后面跟着的函数会延迟到当前函数执行完毕后才执行

	//halo()
	//gogo()

	fmt.Println("==================================")

	//defer halo()
	//gogo()

	fmt.Println("==================================")

	// 即时求值的变量快照
	//str := "halo xxx"
	//defer fmt.Println(str)
	//str = "halo go"
	//fmt.Println(str)

	fmt.Println("==================================")

	// 延迟方法
	//b := Book{"halo go", 100}
	//defer b.showBookInfo()
	//fmt.Println("xxx")

	fmt.Println("==================================")

	// defer 栈
	// 后进先出的顺序
	//defer fmt.Println("xxx1")
	//defer fmt.Println("xxx2")
	//defer fmt.Println("xxx3")
	//fmt.Println("xxx")

	fmt.Println("==================================")

	// defer 在 return 后调用
	//plog := log()
	//fmt.Println("main str = ", str)
	//fmt.Println("main plog = ", plog)

	fmt.Println("==================================")

}

type Book struct {
	bookName  string
	bookPrice float64
}

func (b Book) showBookInfo() {
	fmt.Printf("书名是：%s，价格是：%f\n", b.bookName, b.bookPrice)
}

var str = "halo xxx"

func log() string {
	defer func() {
		str = "halo go"
	}()
	fmt.Println("log str = ", str)
	return str
}

//func fn() {
//	r := getResource()  // #0、获取资源
//
//	defer r.release()  // #1、释放资源
//	if ... {
//		...
//		return
//	}
//	...
//	if ... {
//		...
//		return
//	}
//	...
//	if ... {
//		...
//		return
//	}
// ...
//	return
//}
