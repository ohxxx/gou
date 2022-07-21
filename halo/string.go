package main

import "fmt"

func main() {
	xxx := "halo xxx"
	fmt.Println(xxx)

	var yyy string = "halo yyy"
	fmt.Println(yyy)

	var zzz string
	zzz = `
		halo zzz
		fmt.Println(xxx)
	`
	fmt.Printf(zzz)

}
