package main

import "fmt"

func main() {
	// 方法就是一个函数，在 func 关键字和方法名中间加入了一个特殊的接收器类型
	// func (接收器类型) 方法名(参数列表) (返回值列表) {}
	// func (t Type) MethodName(parameterList) (returnList) {}

	// Go 中：相同名字的方法可以定义在不同类型上，而相同名字的函数是不被允许的

	user := Person{name: "halo", age: 18}
	user.personInfo()

	fmt.Println("==================================")

	user2 := Person{name: "xxx", age: 19}
	personInfo2(user2)

	fmt.Println("==================================")

	user3 := Author{name: "xxx"}
	user3.personInfo()

	fmt.Println("==================================")

	// 指针接收器 与 值接收器
	// 指针接收器：func (p *Person) MethodName(parameterList) (returnList) {}
	// 值接收器：	 func (p Person) MethodName(parameterList) (returnList) {}

	// 指针接收器可以修改接收器类型的值
	// 值接收器不能修改接收器类型的值
	user4 := Author{name: "xxx"}
	user4.authorInfo("halo")
	fmt.Println("user4 = ", user4)
	user4.authorInfo2("hello")
	fmt.Println("user4 = ", user4)

	fmt.Println("==================================")

	// 在方法中使用值接收器 与 在函数中使用值参数
	user5 := Author{name: "xxx"}
	authorInfo3(user5)
	user5.authorInfo3()

	user6 := &user5
	//authorInfo3(user6) // error: cannot use user6 (type *Author) as type Author in argument to authorInfo3
	authorInfo3(*user6)
	user6.authorInfo3()

	fmt.Println("==================================")

	// 在非结构体上的方法
	a := xxxInt(1)
	b := xxxInt(3)
	fmt.Println(a.add(b))

}

type Person struct {
	name string
	age  int
}

func (p Person) personInfo() {
	fmt.Println("name", p.name)
	fmt.Println("age", p.age)
}

func personInfo2(p Person) {
	fmt.Println("name", p.name)
	fmt.Println("age", p.age)
}

type Author struct {
	name string
}

func (a Author) personInfo() {
	fmt.Println("name", a.name)
}

func (a Author) authorInfo(name string) {
	a.name = name
}

func (a *Author) authorInfo2(name string) {
	a.name = name
}

func (a Author) authorInfo3() {
	fmt.Println(a.name)
}

func authorInfo3(a Author) {
	fmt.Println(a.name)
}

type xxxInt int

func (a xxxInt) add(b xxxInt) xxxInt {
	return a + b
}
