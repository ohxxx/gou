package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// 通道（channel）就是一个管道，Go 协程之间通信的管道，它是一种队列式的数据结构，遵循先进先出的原则。

	// 声明一个通道
	// 		1、var chanel-name chan channel-type
	//    	   chanel-name = make(chan channel-type)
	//  	2、chanel-name := make(chan channel-type)
	//  	3、chanel-name := make(chan channel-type, channel-capacity)

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

	// 通道的容量和长度
	// 	cap(channel-name) >>> 容量
	//  	当容量为 0 时，说明通道中不能存放数据，在发送数据时，必须要求立马有人接收，否则会报错。此时通道称为无缓冲通道
	//		当容量为 1 时，说明通道中只能存放一个数据，再发送数据会造成程序阻塞。这种情况可以用来做锁
	// 		当容量大于 1 时，说明通道中可以存放多个数据，可以用于多个协程之间的通信管道，共享资源
	// 	len(channel-name) >>> 长度
	// 创建通道
	ch5 := make(chan int, 3)
	fmt.Println("开始")
	fmt.Println("ch5 的 len = ", len(ch5))
	fmt.Println("ch5 的 cap = ", cap(ch5))
	ch5 <- 6
	ch5 <- 7
	fmt.Println("添加后：ch5 的 len = ", len(ch5))
	fmt.Println("添加后：ch5 的 cap = ", cap(ch5))
	<-ch5
	fmt.Println("删除后：ch5 的 len = ", len(ch5))
	fmt.Println("删除后：ch5 的 cap = ", cap(ch5))
	fmt.Println("结束")

	fmt.Println("==================================")

	// 缓冲通道和无缓冲通道

	// 缓冲通道：允许存储多个数据，设置缓冲区后，发送端和接收端可以处于异步的状态
	ch6 := make(chan int, 6)
	fmt.Println("ch6 的 cap = ", cap(ch6))

	// 无缓存通道：通道里无法存储数据，接收端必须先于发送端准备好，以保证发送完数据后，有人立马接收数据，否则发送端就会阻塞
	// 原因就是：通道中无法存储数据，也就是说发送端和接收端是同步运行的
	ch7 := make(chan int)
	ch8 := make(chan int, 0)
	fmt.Println("ch7 的 cap = ", cap(ch7))
	fmt.Println("ch8 的 cap = ", cap(ch8))

	fmt.Println("==================================")

	// 双向通道
	// 创建一个通道
	ch9 := make(chan int)
	// 发送数据
	go func() {
		ch9 <- 1
	}()
	// 接收数据
	go func() {
		r := <-ch9
		fmt.Println("xxx：", r)
	}()
	// 主协程休眠
	time.Sleep(time.Millisecond)

	fmt.Println("==================================")

	// 单向通道
	// 单项通道只能发送或接收数据
	// 		可以细分为：只写通道和只读通道

	// 只读通道（<-chan）
	ch10 := make(chan string)
	type Receiver = <-chan string
	var receiver Receiver = ch10
	fmt.Println("receiver = ", receiver)

	// 只写通道（chan<-）
	ch11 := make(chan string)
	type Sender = chan<- string
	var sender Sender = ch11
	fmt.Println("sender = ", sender)

	var ch12 = make(chan string)
	// 开启一个协程
	go func() {
		// 只写通道
		var sender Sender = ch12
		fmt.Println("写入 ...")
		sender <- "xxx"
	}()
	// 开启一个协程
	go func() {
		// 只读通道
		var receiver Receiver = ch12
		msg := <-receiver
		fmt.Println("读取：", msg)
	}()
	time.Sleep(time.Millisecond)

	fmt.Println("==================================")

	// 遍历通道
	// 使用 for range 循环进行遍历
	// 但是在遍历时要确保通道是处于关闭状态，否则循环会被阻塞

	ch13 := make(chan int, 6)
	go logChannel(ch13)
	for v := range ch13 {
		fmt.Println("v = ", v)
	}

	fmt.Println("==================================")

	// 用通道做锁（通道容量为 1 的情况下）
	ch14 := make(chan bool, 1)
	var x int
	for i := 0; i < 1000; i++ {
		go increment(ch14, &x)
	}
	time.Sleep(time.Millisecond)
	fmt.Println("x = ", x)

	fmt.Println("==================================")

	// 死锁
	//ch15 := make(chan string)
	//ch15 <- "xxx" // error：fatal error: all goroutines are asleep - deadlock!

	//ch15 := make(chan string)
	//ch15 <- "xxx"
	//fmt.Println("ch15 = ", <-ch15)  // error：fatal error: all goroutines are asleep - deadlock!

	ch15 := make(chan string)
	go fnRecieve(ch15) // 将接收者代码放在协程中
	ch15 <- "halo xxx"
	time.Sleep(time.Millisecond)

	//ch15 := make(chan string, 1)
	//ch15 <- "halo xxx"
	//fmt.Println("ch15 = ", <-ch15)

	fmt.Println("==================================")

	// waitGroup 用于等待一组协程执行完毕
	// 	 Add：增加一个等待协程，初始值为 0
	// 	 Done：当某个子协程执行完毕后，调用 Done 函数，等待协程数减 1，即子协程数减 1，通常使用 defer 来调用
	// 	 Wait：阻塞当前协程，直到实例里的等待协程数为 0

	//isDone := make(chan bool)
	//go func() {
	//	for i := 0; i < 3; i++ {
	//		fmt.Println("i = ", i)
	//	}
	//	isDone <- true
	//}()
	//<-isDone

	//var wg sync.WaitGroup
	// 实例化 sync.WaitGroup
	var wg sync.WaitGroup
	// 传入子协程的数量
	wg.Add(3)
	// 开启子协程 1 及其实例 waitGroup
	go task(1, &wg)
	// 开启子协程 2 及其实例 waitGroup
	go task(2, &wg)
	// 开启子协程 3 及其实例 waitGroup
	go task(3, &wg)
	// 实例 waitGroup 阻塞当前协程，直到子协程数为 0
	wg.Wait()
}

func logChannel(c chan int) {
	for i := 0; i < 6; i++ {
		c <- i
	}
	close(c) // >>> 重点：关闭通道，否则会阻塞
}

func increment(c chan bool, x *int) {
	c <- true
	*x++ // x++ 不是原子操作，所以应避免多个协程对 x 进行操作
	<-c
}

func fnRecieve(c chan string) {
	fmt.Println(<-c)
}

func task(taskNum int, wg *sync.WaitGroup) {
	// 延迟调用，执行完子协程计数器减一
	defer wg.Done()
	for i := 0; i < 3; i++ {
		fmt.Println("taskNum = ", taskNum, " i = ", i)
	}
}
