package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func main() {
	// 结构体（struct）是一种聚合的数据类型

	// 定义：
	// 		type struct-name struct {
	//			attribute-name1 attribute-type1
	//			attribute-name2 attribute-type2
	//			...
	//		}

	// Person 称为 命名的结构体(Named Structure)
	type Person struct {
		name string
		age  int
	}

	// 创建一个 Person 对象
	s1 := Person{name: "xxx", age: 18}
	s2 := Person{"yyy", 19}
	fmt.Println("s1 = ", s1)
	fmt.Println("s2 = ", s2)

	fmt.Println("==================================")

	// 匿名结构体
	// 匿名结构体是一种结构体，它的名字是编译器自动生成的，不需要我们自己定义
	// 创建一个匿名结构体
	s3 := struct {
		name string
		age  int
	}{
		name: "xxx",
		age:  18,
	}
	fmt.Println("s3 = ", s3)

	fmt.Println("==================================")

	// 结构体的零值
	// 定义好的结构体没有被显式初始化，结构体的字段将会赋为相应类型的零值
	s4 := Person{}
	fmt.Println("s4 = ", s4)

	fmt.Println("==================================")

	// 初始化结构体
	// 可以使用结构体的字段名来初始化结构体
	s5 := Person{name: "xxx"}
	fmt.Println("s5 = ", s5)

	fmt.Println("==================================")

	// 访问结构体的字段
	// 可以使用结构体的字段名来访问结构体的字段
	s6 := Person{name: "xxx", age: 20}
	fmt.Println("s6.name = ", s6.name)
	fmt.Println("s6.age = ", s6.age)
	s6.age = 18
	fmt.Println("s6.age = ", s6.age)

	fmt.Println("==================================")

	// 指向结构体的指针
	s7 := &Person{name: "xxx", age: 18}
	fmt.Println("s7 = ", (*s7).name)
	fmt.Println("s7 = ", s7)

	fmt.Println("==================================")

	// 匿名字段
	type Person2 struct {
		string
		int
	}
	s8 := Person2{"xxx", 18}
	fmt.Println("s8 = ", s8)
	fmt.Println("s8.string = ", s8.string)
	fmt.Println("s8.int = ", s8.int)

	fmt.Println("==================================")

	// 嵌套结构体
	type Address struct {
		city     string
		district string
	}
	type Person3 struct {
		name string
		age  int
		addr Address
	}
	s9 := Person3{
		name: "xxx",
		age:  18,
		addr: Address{
			city:     "上海",
			district: "徐汇",
		},
	}
	fmt.Println("s9 = ", s9)
	s9.addr = Address{
		city:     "北京",
		district: "朝阳",
	}
	fmt.Println("s9 = ", s9)

	fmt.Println("==================================")

	// 提升结构体的字段
	// 可以使用结构体的字段名来提升结构体的字段
	type Address2 struct {
		city string
		name string
	}
	type Person4 struct {
		name string
		age  int
		Address2
	}
	s10 := Person4{
		name: "xxx",
		age:  18,
	}
	s10.Address2 = Address2{
		city: "深圳",
		name: "张三",
	}
	fmt.Println("s10 = ", s10)

	fmt.Println("==================================")

	// 结构体比较
	// 结构体比较是比较两个结构体的字段的值，不会比较两个结构体的指针
	s11 := Person4{
		name: "xxx",
		age:  18,
	}
	s12 := Person4{
		name: "xxx",
		age:  18,
	}
	fmt.Println("s11 == s12 ? ", s11 == s12)
	fmt.Println("s11 == s12 ? ", reflect.DeepEqual(s11, s12))
	fmt.Println("s11.name == s12.name ? ", s11.name == s12.name)

	fmt.Println("==================================")

	// 给结构体定义方法
	s13 := Person5{name: "xxx", age: 18}
	s13.showPersonInfo()

	fmt.Println("==================================")

	// 方法的参数传递方法
	s14 := Person5{name: "xxx", age: 19}
	s14.showPersonInfo()
	s14.addAge(2)
	s14.showPersonInfo()

	fmt.Println("==================================")

	// Tag 的使用
	type User struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
		City string `json:"city,omitempty"` // omitempty 忽略空值
		// omitempty 会忽略 false、0、空指针、空接口、空数组、空切片、空映射、空字符串
	}
	u1 := User{Name: "xxx", Age: 19}
	// 结构体转为 JSON
	data, err := json.Marshal(u1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("data = %s\n", data)

	u2 := User{
		Name: "xxx",
		Age:  19,
		City: "shanghai",
	}
	// 结构体转为 JSON
	data2, err := json.Marshal(u2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("data2 = %s\n", data2)

	fmt.Println("==================================")

	// 获取字段
	field1, _ := reflect.TypeOf(u2).FieldByName("Name")
	field2 := reflect.ValueOf(u2).Type().Field(1) // Filed(i) i 第几个字段
	field3 := reflect.ValueOf(&u2).Elem().Type().Field(2)
	fmt.Println("field1 = ", field1)
	fmt.Println("field2 = ", field2)
	fmt.Println("field3 = ", field3)

	// 获取Tag
	tag1 := field1.Tag
	tag2 := field2.Tag
	tag3 := field3.Tag
	fmt.Println("tag1 = ", tag1, "tag1 json = ", tag1.Get("json"))
	fmt.Println("tag2 = ", tag2, "tag2 json = ", tag2.Get("json"))
	fmt.Println("tag3 = ", tag3, "tag3 json = ", tag3.Get("json"))

	// 获取键值对
	labelValue := tag1.Get("json")
	labelValue2, ok := tag2.Lookup("json")
	fmt.Println("labelValue = ", labelValue)
	fmt.Println("labelValue2 = ", labelValue2, "ok = ", ok)
}

type Person5 struct {
	name string
	age  int
}

// showPersonInfo 方法名
// p Person5 >>> 表示将此方法与 Person5 的实例绑定
func (p Person5) showPersonInfo() {
	fmt.Println("name = ", p.name)
	fmt.Println("age = ", p.age)
}

func (p *Person5) addAge(a int) {
	p.age = p.age + a
}
