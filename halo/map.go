package main

import "fmt"

func main() {
	// map 是散列表（哈希表）的引用。它是一个拥有键值对元素的`无序集合`

	// 创建一个 map：
	// 		1、make(map[key-type]value-type)
	// 		2、map[key-type]value-type{key1:value1, key2:value2, ...}
	map1 := make(map[string]int)
	fmt.Println("map1 = ", map1)
	map1["key1"] = 1
	map1["key2"] = 2
	fmt.Println("map1 = ", map1)

	map2 := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}
	fmt.Println("map2 = ", map2)

	// map 操作
	map3 := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}
	fmt.Println("初始化的 map3 = ", map3)
	// 添加
	map3["key3"] = "value3"
	map3["key4"] = "value4"
	fmt.Println("添加后的 map3 = ", map3)
	// 删除
	delete(map3, "key4")
	fmt.Println("删除后的 map3 = ", map3)
	// 更新
	map3["key3"] = "666"
	fmt.Println("更新后的 map3 = ", map3)
	// 查询
	fmt.Println("map3[\"key3\"] = ", map3["key3"])
	// 判断是否存在
	// value, ok := map[key]
	v, ok := map3["key4"]
	fmt.Println("map3[\"key3\"] = ", v, " ok = ", ok)
	// map 的长度
	fmt.Println("map3 len = ", len(map3))
	// map 的遍历
	for k, v := range map3 {
		fmt.Println("map3[\"", k, "\"] = ", v)
	}

	// map 是引用类型
	map4 := map3
	map4["key5"] = "value5"
	fmt.Println("map3 = ", map3)
	fmt.Println("map4 = ", map4)
}
