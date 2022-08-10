package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 反射
	// Go 提供了一种机制，可以在运行时通过反射机制读取程序的信息
	// 反射机制可以让我们在运行时分析一个对象的类型，从而对对象进行操作

	// reflect.Type
	// 表示 interface{} 的具体类型
	// reflect.TypeOf() 方法返回 reflect.Type
	reflectType(666)
	reflectType("halo xxx")

	fmt.Println("==================================")

	// reflect.Value
	// 表示 interface{} 的具体值
	// reflect.ValueOf() 方法返回 reflect.Value
	reflectValue(666)
	reflectValue("halo xxx")

	fmt.Println("==================================")

	// reflect.Kind
	// 表示的是种类，当需要区分一个大品种的类型的时候就会用到 Kind（种类）
	// Kind 是一个枚举类型，有以下几种类型：
	//
	// Kind 表示 Type 表示的特定类型的类型。
	// 零种类不是有效种类。
	// type Kind uint
	//const (
	//	Invalid Kind = iota
	//	Bool
	//	Int
	//	Int8
	//	Int16
	//	Int32
	//	Int64
	//	Uint
	//	Uint8
	//	Uint16
	//	Uint32
	//	Uint64
	//	Uintptr
	//	Float32
	//	Float64
	//	Complex64
	//	Complex128
	//	Array
	//	Chan
	//	Func
	//	Interface
	//	Map
	//	Pointer
	//	Slice
	//	String
	//	Struct
	//	UnsafePointer
	//)
	var a user
	reflectType2(a)

	fmt.Println("==================================")

	// reflect.NumField
	// 返回 struct 类型的字段数量
	var b user2
	reflectNumField(b)

	fmt.Println("==================================")

	// reflect.Field(i int)
	// 返回 i 的 reflect.Value
	var c = user2{name: "xxx", age: 18}
	reflectNumField2(c)

	fmt.Println("==================================")

	// 反射的三大定律
	// 1、反射可以将“接口类型变量”转换为“反射类型对象”。
	// 2、反射可以将“反射类型对象”转换为“接口类型变量”
	// 3、如果要修改“反射类型对象”其值必须是“可写的”

	// 第一定律
	// 反射可以将“接口类型变量”转换为“反射类型对象”。
	// 反射类型：reflect.Type 和 reflect.Value
	var d interface{} = "halo xxx"
	fmt.Println("d = ", d)
	dt := reflect.TypeOf(d)
	dv := reflect.ValueOf(d)
	fmt.Printf("从接口类型变量到反射类型对象：Type 对象类型为 %T\n", dt)
	fmt.Printf("从接口类型变量到反射类型对象：Value 对象类型为 %T\n", dv)

	fmt.Println("==================================")

	// 第二定律
	// 反射可以将“反射类型对象”转换为“接口类型变量”
	fdv := dv.Interface()
	fmt.Printf("从反射类型对象到接口类型变量：Type 对象类型为 %T\n", fdv)
	// 使用类型断言进行转换
	fdv2 := dv.Interface().(string)
	fmt.Printf("fdv2 类型 %T，值为 %v\n", fdv2, fdv2)

	fmt.Println("==================================")

	// 第三定律
	// 如果要修改“反射类型对象”其值必须是“可写的”
	var e float64 = 3.14
	fe := reflect.ValueOf(e)
	//fe.SetFloat(41.2) // error：panic: reflect: reflect.Value.SetFloat using unaddressable value
	fmt.Println("可写行：", fe.CanSet())
	fe2 := reflect.ValueOf(&e).Elem()
	fmt.Println("可写行：", fe2.CanSet())
	fe2.SetFloat(41.2)
	fmt.Println("fe2 = ", fe2)
}

func reflectType(x interface{}) {
	t := reflect.TypeOf(x)
	fmt.Println(t)
}

func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	fmt.Println(v)
}

type user struct{}

func reflectType2(x interface{}) {
	typex := reflect.TypeOf(x)
	fmt.Println("typex = ", typex)
	fmt.Println("typex.Kind() = ", typex.Kind())
}

type user2 struct {
	name string
	age  int
}

func reflectNumField(x interface{}) {
	if reflect.ValueOf(x).Kind() == reflect.Struct { // 检查是不是 struct 类型
		v := reflect.ValueOf(x)
		fmt.Println("v.NumField() = ", v.NumField())
	}
}

func reflectNumField2(x interface{}) {
	if reflect.ValueOf(x).Kind() == reflect.Struct {
		v := reflect.ValueOf(x)
		fmt.Println("v.NumField() = ", v.NumField())
		for i := 0; i < v.NumField(); i++ {
			fmt.Printf("Field = %d type = %T value = %v\n", i, v.Field(i), v.Field(i))
		}
	}
}
