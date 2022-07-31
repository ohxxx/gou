package main

import "fmt"

func main() {
	// 切片定义：
	// 	1、[]Type
	// 	2、make([]Type, size, cap)
	var slice1 []int
	var slice2 = []int{}
	fmt.Printf("slice1 的类型：%T，slice1 的长度是：%d，slice1 的值是：%d\n", slice1, len(slice1), slice1)
	fmt.Printf("slice2 的类型：%T，slice2 的长度是：%d，slice2 的值是：%d\n", slice2, len(slice2), slice2)

	slice3 := make([]int, 3, 5)
	fmt.Printf("slice3 的类型：%T，slice3 的长度是：%d，slice3 的值是：%d\n", slice3, len(slice3), slice3)

	slice4 := [3]string{"hello", "xxx", "gogogo"}
	slice4[0] = "halo"
	slice5 := slice4[1:2] // 截取切片 [start:end]
	fmt.Printf("slice4 的类型：%T，slice4 的长度是：%d，slice4 的值是：%s\n", slice4, len(slice4), slice4)
	fmt.Printf("slice5 的类型：%T，slice5 的长度是：%d，slice5 的值是：%s\n", slice5, len(slice5), slice5)

	// 切片由三部分组成：指针、长度、容量
	slice6 := make([]string, 3, 5)
	fmt.Println("slice6 的长度是：", len(slice6))
	fmt.Println("slice6 的容量是：", cap(slice6))

	//slice7 := make([]int, 2, 5)
	//fmt.Println("slice7：", slice7[10]) // error：panic: runtime error: index out of range [10] with length 2

	// TIPS：slice 是引用类型，所以不对其进行赋值的话，它的默认值就是 nil，切片之间不能比较
	var slice8 []int
	fmt.Println(slice8 == nil) // true

	// 切片元素的修改
	// 切片自己不拥有任何数据，它只是底层数组的一种表示。对切片的任何操作都会对底层数组进行操作。
	slice9 := []int{1, 2, 3, 4, 5}
	slice10 := slice9[:] // >>> [:] 默认就是 0 到 len(slice9)
	fmt.Println("slice9 的值是：", slice9)
	fmt.Println("slice10 的值是：", slice10)
	slice10[0] = 666
	fmt.Println("slice9 的值是：", slice9)
	fmt.Println("slice10 的值是：", slice10)

	// 追加切片元素
	slice11 := []int{1, 2, 3, 4, 5}
	fmt.Println("slice11 的值是：", slice11)
	fmt.Println("slice11 的容量是：", cap(slice11))
	slice11 = append(slice11, 6)
	fmt.Println("slice11 的值是：", slice11)
	fmt.Println("slice11 的容量是：", cap(slice11))
	slice11 = append(slice11, []int{7, 8, 9, 10, 11}...)
	fmt.Println("slice11 的值是：", slice11)
	fmt.Println("slice11 的容量是：", cap(slice11))
	// 当容量不足时，会创建一个新的数组，将现有数组复制到这个新的数组中，并返回新的引用。切片的容量是久切片的 2 倍

	// 切片的删除
	slice12 := []int{1, 2, 3, 4, 5}
	slice12 = append(slice12, 6)
	fmt.Println("slice12 的值是：", slice12)
	fmt.Println("slice12 的容量是：", cap(slice12))
	slice12 = slice12[:len(slice12)-1]
	fmt.Println("slice12 的值是：", slice12)
	fmt.Println("slice12 的容量是：", cap(slice12))

	// 切片的遍历
	slice13 := []int{1, 2, 3, 4, 5}
	for i, v := range slice13 {
		fmt.Printf("slice13 的第 %d 个元素是：%d\n", i, v)
	}

	// 切片的复制
	slice14 := []int{1, 2, 3, 4, 5}
	slice15 := slice14[:]
	fmt.Println("slice14 的值是：", slice14)
	fmt.Println("slice15 的值是：", slice15)
	slice14[0] = 666
	fmt.Println("slice14 的值是：", slice14)
	fmt.Println("slice15 的值是：", slice15)
	slice15[0] = 777
	fmt.Println("slice14 的值是：", slice14)
	fmt.Println("slice15 的值是：", slice15)

	// 多维切片
	slice16 := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println("slice16 的值是：", slice16)
}
