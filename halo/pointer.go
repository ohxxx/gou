package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// 指针也会是一种类型，也可以创建变量，称之为指针变量
	// 指针变量的类型是 *Type，该指针指向一个 Type 类型的变量
	// 指针变量最大的特点就是：存储的某个实际变量的内存地址，通过记录某个变量的地址，从而间接的操作该变量。
	// ----------------------------------------------------------------------------------------------
	//                                             Memory Address                       Memory
	//  [variable]   var x int = 100   ------->       0x0201                             100
	//                                                0x0202
	//                                                0x0203
	//                                                0x0204
	//                                                0x0205
	//  [pointer]     var p *int = &x   ------->      0x0206                            0x0201
	//                                                0x0207
	//                                                0x0208
	// ----------------------------------------------------------------------------------------------

	// 创建指针变量
	// 		1、&variable
	// 		2、new(Type)
	// 		3、*pointer

	// [先定义普通变量，再通过获取该普通变量的地址创建指针（&variable）]
	// 定义普通变量
	pointer1 := "halo xxx"
	// 取普通变量 pointer1 的地址指向 pointer2
	pointer2 := &pointer1
	fmt.Println("pointer2 = ", pointer2)
	fmt.Println("*pointer2 = ", *pointer2)

	fmt.Println("==================================")

	// [先创建指针并分配好内存，再给指针指向的内存地址写入对应的值（new(string)）]
	// 创建指针
	pointer3 := new(string)
	// 给指针指向的地址写入对应的值
	*pointer3 = "halo xxx"
	fmt.Println("pointer3 = ", pointer3)
	fmt.Println("*pointer3 = ", *pointer3)

	fmt.Println("==================================")

	// [先声明一个指针变量，再从其他变量获取内存地址给指针变量（*pointer）]
	// 定义变量
	pointer4 := "halo xxx"
	// 声明指针变量
	var pointer5 *string
	// 将 pointer4 的地址赋值给 pointer5
	pointer5 = &pointer4
	fmt.Println("pointer5 = ", pointer5)
	fmt.Println("*pointer5 = ", *pointer5)

	fmt.Println("==================================")

	// TIPS：
	// 		& 操作符可以从一个变量中获取其内存地址
	// 		* 操作符如果在赋值操作符值的
	//		 	左侧，指该指针指向的变量
	//          右侧，指从一个指针变量中取得变量值，又称指针的解引用
	pointer6 := "halo xxx"
	pointer7 := &pointer6
	fmt.Println("pointer6 = ", pointer6)
	fmt.Println("*pointer7 = ", *pointer7)
	fmt.Println("&pointer6 = ", &pointer6)
	fmt.Println("pointer7 = ", pointer7)

	fmt.Println("==================================")

	// 指针的类型
	// *(指向变量值的数据类型) 就是对应的指针类型
	pStr := "halo xxx"
	pInt := 1
	pBool := false
	pFloat := 1.1
	fmt.Printf("type of &pStr is :%T\n", &pStr)
	fmt.Printf("type of &pInt is :%T\n ", &pInt)
	fmt.Printf("type of &pBool is :%T\n ", &pBool)
	fmt.Printf("type of &pFloat is :%T\n ", &pFloat)

	fmt.Println("==================================")

	// 指针的零值
	str := "halo xxx"
	var pZero *string
	fmt.Println("pZero = ", pZero) // <nil>
	pZero = &str
	fmt.Println("pZero = ", pZero)

	fmt.Println("==================================")

	// 函数传递指针参数
	pNum := 001
	pFunc := &pNum
	fmt.Println("*pFunc = ", *pFunc)
	changeByPointer(pFunc)
	fmt.Println("*pFunc = ", *pFunc)

	fmt.Println("==================================")

	// 指针与切片都是引用类型
	// 如果想通过一个函数改变一个数组的值
	//   1、将该数组的切片当作参数传给函数（推荐）
	//   2、将这个数组的指针当作参数传给函数
	pNum2 := [3]int{1, 2, 3}
	fmt.Println("pNum2 = ", pNum2)
	changeSlice(pNum2[:])
	fmt.Println("pNum2 = ", pNum2)
	changeArray(&pNum2)
	fmt.Println("pNum2 = ", pNum2)

	fmt.Println("==================================")

	// Go 中不支持直接指针运算，但可以使用 unsafe 实现
	//xxx := [...]int{20, 30, 40}
	//pp := &xxx
	//pp++ // error

	var a []int
	alloc(a, 5)

}

func changeByPointer(value *int) {
	*value = 666
}

func changeSlice(value []int) {
	value[0] = 666
}

func changeArray(value *[3]int) {
	(*value)[1] = 666
}

func alloc(a []int, n int) {
	a = make([]int, n)
	for i := 0; i < len(a); i++ {
		a[i] = i*i + 1
	}

	var ptr *int = &a[0]

	for i := 0; i < n; i++ {
		fmt.Printf("i=%d, *p=%d;\n", i, *ptr)
		p := uintptr(unsafe.Pointer(ptr))
		p = p + unsafe.Sizeof(int(0)) // 指针移动
		ptr = (*int)(unsafe.Pointer(p))
	}
}
