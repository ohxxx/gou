package main

// error：package packet_book is not in GOROOT
// 处理：go mod init pkg-name

//import _ "fmt"

func init() {
	println("init xxx")
}

func main() {
	//bookName := "halo go"
	//bookPrice := float64(100)
	//bookInfo, _ := book.ShowBookInfo(bookName, bookPrice)
	//Println(bookInfo)
}
