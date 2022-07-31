package main

import "fmt"

func main() {
	// if 条件1 {
	// 	语句1
	// } else if 条件2 {
	// 	语句2
	// } else if ... {
	// 	语句...
	// } else {
	// 	语句3
	// }

	str1 := "halo xxx"
	if str1 == "halo xxx" {
		fmt.Println("gogogo")
	}

	if true {
		fmt.Println("gogogo")
	}

	if !false {
		fmt.Println("gogogo")
	}

	for i := 0; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}

}
