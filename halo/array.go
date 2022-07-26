package main

import "fmt"

func main() {
	var arr1 [3]int
	arr1[0] = 1
	arr1[1] = 2
	arr1[2] = 3
	fmt.Println(arr1)

	fmt.Println("==================================")

	var arr2 = [3]int{1, 2, 3}
	fmt.Println(arr2)

	fmt.Println("==================================")

	arr3 := [3]int{1, 2, 3}
	fmt.Println(arr3)

	fmt.Println("==================================")

	arr4 := [3]int{1, 2}
	fmt.Println(arr4)

	fmt.Println("==================================")

	arr5 := [3]int{0: 000, 2: 222}
	fmt.Println(arr5)

	fmt.Println("==================================")

	arr6 := [...]int{1, 2, 3}
	fmt.Println(arr6)

	fmt.Println("==================================")

	// 数组长度是类型的一部分，所以 [3]int 和 [4]int 是不同的类型
	arr7 := [3]int{1, 2, 3}
	arr8 := [4]int{1, 2, 3, 4}
	fmt.Printf("type of arr7 is %T\n", arr7)
	fmt.Printf("type of arr8 is %T\n", arr8)

	fmt.Println("==================================")

	arr9 := [2][3]string{
		{"halo1", "xxx1", "gogogo1"},
		{"halo2", "xxx2", "gogogo2"},
	}
	fmt.Println(arr9)

	fmt.Println("==================================")

	arr10 := [...]string{"halo", "xxx"}
	fmt.Println("arr10 的长度是：", len(arr10))

	fmt.Println("==================================")

	arr11 := [...]string{"halo", "xxx"}
	for index, value := range arr11 {
		fmt.Printf("arr11[%d]=%s\n", index, value)
	}

	// _ 表示不需要使用的变量
	for _, val := range arr11 {
		fmt.Printf("%s\n", val)
	}

	fmt.Println("==================================")

	// go 中数组类型是 值类型，副本不会修改原数组中的值
	arr12 := [...]string{"hello", "xxx"}
	copyArr := arr12
	copyArr[0] = "halo"
	fmt.Println("arr12 的值是：", arr12)
	fmt.Println("copyArr 的值是：", copyArr)
}
