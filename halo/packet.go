package main

// error：package packet_book is not in GOROOT
// 处理：go mod init pkg-name

import (
	// 别名：book
	book "packet_book/halo/packet_book"
)

// 包的匿名导入
// 包的匿名导入是一种特殊的导入，它可以被用来导入包中的所有内容，包括子包。
import . "fmt"

//import _ "fmt"

func init() {
	println("init xxx")
}

func main() {
	bookName := "halo go"
	bookPrice := float64(100)
	bookInfo, _ := book.ShowBookInfo(bookName, bookPrice)
	Println(bookInfo)
}
