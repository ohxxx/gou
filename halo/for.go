package main

import "fmt"

func main() {
	// for init; condition; post {
	// 	code...
	// }

	// for condition {
	// 	code...
	// }

	// for range_expression {
	//	code...
	// }

	//for {
	//	code
	//}

	for i := 0; i < 6; i++ {
		fmt.Println(i)
	}

	fmt.Println("==================================")

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Println(i, j)
		}
	}

	fmt.Println("==================================")

	str := "halo xxx"
	for i, v := range str {
		fmt.Printf("i = %d, v = %c\n", i, v)
	}

	fmt.Println("==================================")

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	fmt.Println("==================================")

	i := 0
	for i < 6 {
		fmt.Println(i)
		i++
	}

	fmt.Println("==================================")

	//i1 := 0
	//for {
	//	fmt.Println(i1)
	//	i1++
	//}

	fmt.Println("==================================")

	i2 := 0
	for {
		fmt.Println(i2)
		if i2 >= 10 {
			break
		}
		i2++
	}

	fmt.Println("==================================")

	i3 := 0
	for {
		i3++
		if i3%2 == 0 {
			continue
		}
		fmt.Println(i3)
		if i3 >= 50 {
			break
		}
	}

	fmt.Println("==================================")

	for i := 250; i <= 340; i++ {
		fmt.Println(i, " - ", string(rune(i)), " - ", []byte(string(rune(i))))
	}
	foo := "a"
	fmt.Println(foo)
	fmt.Printf("%T \n", foo)

	fmt.Println("==================================")

	for i := 50; i <= 140; i++ {
		fmt.Printf("%v - %v - %v \n", i, string(rune(i)), []byte(string(rune(i))))
	}
}
