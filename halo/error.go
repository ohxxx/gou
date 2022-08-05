package main

import (
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
	//val, err := add(-1, 2)
	//if err != nil {
	//	fmt.Println("err = ", err)
	//	return
	//}
	//fmt.Println("add result = ", val)

	fmt.Println("==================================")

	//val2, err2 := add(-1, 2)
	//if err2 != nil {
	//	if err3, ok := err2.(*addError); ok {
	//		fmt.Println("err3 = ", err3.msg, err3.length)
	//		return
	//	}
	//	fmt.Println(err2)
	//	return
	//}
	//fmt.Println("add result = ", val2)

	fmt.Println("==================================")

	val3, err4 := superAdd(-1, 2)
	if err4 != nil {
		if err5, ok := err4.(*addError2); ok {
			if err5.aLessThanZero() {
				fmt.Println("Yeeha a < 0", err5.a)
			}

			if err5.bLessThanZero() {
				fmt.Println("Yeeha b < 0", err5.b)
			}

			return
		}
	}

	fmt.Println("add result = ", val3)

}

func add(a int, b int) (int, error) {
	if a < 0 || b < 0 {
		//return 0, errors.New("a or b < 0")
		//return 0, fmt.Errorf("(a = %d or b = %d) < 0", a, b)
		return 0, &addError{"a or b < 0", a + b}
	}
	return a + b, nil
}

type addError struct {
	msg    string // 错误信息
	length int    // 错误长度
}

func (e *addError) Error() string {
	return fmt.Sprintf("%s, length = %d", e.msg, e.length)
}

type addError2 struct {
	msg string
	a   int
	b   int
}

func (e *addError2) Error() string {
	return e.msg
}

func (e *addError2) aLessThanZero() bool {
	return e.a < 0
}

func (e *addError2) bLessThanZero() bool {
	return e.b < 0
}

func superAdd(a, b int) (int, error) {
	err := ""

	if a < 0 {
		err += "a < 0"
	}

	if b < 0 {
		err += "b < 0"
	}

	if err != "" {
		return 0, &addError2{err, a, b}
	}

	return a + b, nil
}
