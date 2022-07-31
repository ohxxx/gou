package main

import "fmt"

func main() {
	//  switch 表达式 {
	// 	case 条件1:
	// 		语句1
	// 	case 条件2:
	// 		语句2
	// 	case ...:
	// 		语句...
	// 	default:
	// 		语句3

	language := "JS"
	switch language {
	case "JS":
		fmt.Println("JS 很神奇的语言")
	case "TS":
		fmt.Println("TS 很有趣的语言")
	case "Go":
		fmt.Println("Go 很好玩的语言")
	default:
		fmt.Println("不是有手就行？")
	}

	fmt.Println("==================================")

	switch month := 6; month {
	case 1, 3, 5, 7, 8, 10, 12:
		fmt.Println("31 天")
	case 4, 6, 9, 11:
		fmt.Println("30 天")
	case 2:
		fmt.Println("闰年29 天，非闰年 28 天")
	default:
		fmt.Println("Error")
	}

	fmt.Println("==================================")

	chinese := 100
	math := 98
	switch getResult(chinese, math) {
	case true:
		fmt.Println("Pass")
	case false:
		fmt.Println("Fail")
	}

	fmt.Println("==================================")

	score := 96.6
	switch {
	case 90 <= score && score <= 100:
		fmt.Println("等级为 A")
	case 80 <= score && score < 90:
		fmt.Println("等级为 B")
	case 70 <= score && score < 80:
		fmt.Println("等级为 C")
	case 60 <= score && score < 70:
		fmt.Println("等级为 D")
	case 60 > score:
		fmt.Println("等级为 E")
	}

	fmt.Println("==================================")

	lang := "TS"
	switch lang {
	case "JS":
		fmt.Println("JS 很神奇的语言")

	case "TS":
		fmt.Println("TS 很有趣的语言")
		fallthrough // >>> !
	case "Go":
		fmt.Println("Go 很好玩的语言")
	default:
		fmt.Println("不是有手就行？")
	}
	// log：TS 很有趣的语言
	// log：Go 很好玩的语言

}

func getResult(args ...int) bool {
	for _, v := range args {
		if v < 60 {
			return false
		}
	}
	return true
}
