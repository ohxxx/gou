package main

import (
	"errors"
	"fmt"
	"strconv"
)

func sum(x int, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func variableParameterFunc(args ...interface{}) {
	for k, v := range args {
		fmt.Println("args[", k, "] = ", v)
	}
}

func main() {
	// Go 是编译型语言，所以函数编写的顺序是无关紧要的

	// 定义：
	//  	func func-name(parameter-list) return-type {
	//  		function-body
	//  	}

	fmt.Println("sum(1, 2) = ", sum(1, 2))
	fmt.Println("sub(2, 1) = ", sub(2, 1))

	fmt.Println("==================================")

	// 可变参数函数
	variableParameterFunc("xxx", 666, false, []int{1, 2, 3}, map[string]string{"key1": "value1"})

	fmt.Println("==================================")

	// 解序列
	// 使用 ... 关键字，能将函数的可变参数（即切片）一个一个取出来，传递给另一个可变参数的函数
	// 而不是传递可变参数变量本身
	var slice []string
	slice = append(slice, []string{"a", "b", "c"}...)
	fmt.Println("slice = ", slice)

	fmt.Println("==================================")

	// 函数返回值
	userInfo, err := showPersonInfo("xxx", 18)
	fmt.Println("userInfo ", userInfo, " err = ", err)
	userInfo2, err := showPersonInfo2("xxx", 20)
	fmt.Println("userInfo2 ", userInfo2, " err = ", err)

	fmt.Println("==================================")

	// Go 中函数名通过首字母大小写实现控制对方法的访问权限
	//		方法的首字母为 大写 时，这个方法对于 所有包 都是 Public ，其他包可以随意调用
	//		方法的首字母为 小写 时，这个方法是 Private ，其他包是无法访问的

	fmt.Println("==================================")

	// 头等函数 （first-class function）
	// 函数可以作为参数传递给另一个函数，或者作为返回值返回到调用者

	// 赋值给变量
	haloFunc := func() {
		fmt.Println("halo xxx")
	}
	haloFunc()
	fmt.Printf("haloFunc 的类型是 %T\n", haloFunc)

	// 作为参数传递给另一个函数
	haloFunc2 := func(name string) string {
		return name
	}
	halo(haloFunc2)

	// 作为返回值返回到调用者
	haloFunc3 := halo2()
	fmt.Println("haloFunc3 = ", haloFunc3("xxx"))

	// 闭包
	x := 100
	func() {
		fmt.Println("x = ", x)
	}()
}

func showPersonInfo(name string, age int) (string, error) {
	if name == "" {
		return "", errors.New("name is empty")
	}
	if age < 0 {
		return "", errors.New("age is negative")
	}
	return "姓名：" + name + "，年龄：" + strconv.Itoa(age), nil
}

func showPersonInfo2(name string, age int) (user string, err error) {
	user = ""
	if name == "" {
		err = errors.New("name is empty")
		return
	}
	if age < 0 {
		err = errors.New("age is negative")
		return
	}
	user = "姓名：" + name + "，年龄：" + strconv.Itoa(age)
	return
}

func halo(show func(name string) string) {
	fmt.Println("halo", show("xxx"))
}

func halo2() func(name string) string {
	return func(name string) string {
		return "halo " + name
	}
}
