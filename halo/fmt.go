package main

import "fmt"

func main() {
	// fmt 格式输出
	fmt.Println("halo xxx")
	// %%
	fmt.Printf("%%\n")
	// %b
	fmt.Printf("%b\n", 10)
	// %c
	fmt.Printf("%c\n", 'x')
	// %d
	fmt.Printf("%d\n", 10)
	// %f
	fmt.Printf("%f\n", 1+2i)
	// %o
	fmt.Printf("%o\n", 10)
	// %q
	fmt.Printf("%q\n", 'x')
	// %s
	fmt.Printf("%s\n", "halo xxx")
	// %t
	fmt.Printf("%t\n", true)
	// %T
	fmt.Printf("%T\n", "halo xxx")
	// %x
	fmt.Printf("%x\n", 'x')
	// %X
	fmt.Printf("%X\n", 'X')
}
