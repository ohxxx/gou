package main

import "fmt"

func main() {
	// 通道（channel）就是一个管道，Go 协程之间通信的管道，它是一种队列式的数据结构，遵循先进先出的原则。

	// 声明一个通道
	// 		1、var chanel-name chan channel-type
	//    	   chanel-name = make(chan channel-type)
	//  	2、chanel-name := make(chan channel-type)

	var ch chan string
	// 上面声明了一个类型为 string 的通道，名字为 ch，值为 nil
	fmt.Println("ch = ", ch)
	ch = make(chan string)

	fmt.Println("==================================")

	ch2 := make(chan string)
	fmt.Println("ch2 = ", ch2)

	fmt.Println("==================================")

	// 使用通道发送和接收数据
	// 		1、发送数据：channel-name <- data
	//        （即把 data 数据写入到 channel-name 通道中）
	// 		2、接收数据：data := <- channel-name
	//        （即从 channel-name 通道中读取数据到 data）

	// 创建通道
	ch3 := make(chan string)
	fmt.Println("开始")
	// 开启协程
	go func() {
		ch3 <- "xxx" // 发送数据
	}()

	res := <-ch3 // 接收数据
	fmt.Println("halo：", res)
	fmt.Println("结束")

	// TIPS：通道的发送和接收默认是阻塞的
	// 如果通道没有数据，那么发送或接收将会一直阻塞，直到有数据可以发送或接收。

	fmt.Println("==================================")

	// 通道的关闭
	// 	close(channel-name)
	// 创建通道
	ch4 := make(chan string)
	fmt.Println("开始")
	// 开启协程
	go func() {
		ch4 <- "xxx" // 发送数据
		close(ch4)   // 关闭通道
	}()

	value, ok := <-ch4 // 接收数据
	fmt.Println("halo：", value, ok)
	fmt.Println("结束")

	fmt.Println("==================================")

}
