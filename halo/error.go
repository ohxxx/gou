package main

import (
	"errors"
	"fmt"
)

func main() {
	// Go 中错误使用内置的 error 类型表示，error 类型是个接口类型
	//type error interface {
	//	Error() string
	//}

	// 手动报错操作 - 读取不存在的文件
	//file, err := os.Open("./halo/xxx.ts")
	//if err != nil {
	//	fmt.Println("err = ", err)
	//	return
	//}
	//fmt.Println("file = ", file)

	fmt.Println("==================================")

	// 自定义错误
	val, err := add(1, 2)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	fmt.Println("add result = ", val)

}

func add(a int, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, errors.New("a or b < 0")
	}
	return a + b, nil
}
