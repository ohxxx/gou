package main

import (
	"fmt"
	"time"
)

func main() {
	// select 语句
	// select 语句用在多个 发送/接收 通道操作进行选择
	// 		select 语句会一直阻塞，直到 发送/接收 操作准备就绪
	// 		如果有多个通道操作准备完毕，select 语句会随机选择一个执行

	//select {
	//	case expression1:
	//		code...
	//	case expression2:
	//		code...
	//	default:
	//		code ...
	//}

	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)
	ch3 := make(chan string, 1)
	ch1 <- "halo"
	ch2 <- "xxx"
	ch3 <- "gogogo"

	select {
	case res1 := <-ch1:
		fmt.Println("ch1 = ", res1)
	case res2 := <-ch2:
		fmt.Println("ch2 = ", res2)
	case res3 := <-ch3:
		fmt.Println("ch3 = ", res3)
	default:
		fmt.Println("default")
	}
	// 以上打印随机的，只会打印一个，就看那个先发送的，进入到 case 中执行 code 然后退出

	fmt.Println("==================================")

	ch4 := make(chan string)
	ch5 := make(chan string)
	ch6 := make(chan string)
	go task1(ch4)
	go task2(ch5)
	go task3(ch6)
	select {
	case res1 := <-ch4:
		fmt.Println("ch4 = ", res1)
	case res2 := <-ch5:
		fmt.Println("ch5 = ", res2)
	case res3 := <-ch6:
		fmt.Println("ch6 = ", res3)
	}

	// 以上，没有 default，要是没有命中 select 就会阻塞造成死锁
	// 解决办法就是加上 default，如果没有命中 select 就执行 default
	// 如果是 select {} 也会造成死锁

	fmt.Println("==================================")

	// select 超时处理
	ch7 := make(chan string)
	ch8 := make(chan string)
	ch9 := make(chan string)
	timeout := make(chan bool, 1)

	go makeTimeout(timeout, 2)

	select {
	case res1 := <-ch7:
		fmt.Println("ch7 = ", res1)
	case res2 := <-ch8:
		fmt.Println("ch8 = ", res2)
	case res3 := <-ch9:
		fmt.Println("ch9 = ", res3)
	case <-timeout:
		fmt.Println("timeout")
	}

	fmt.Println("==================================")

	// 读取/写入数据
	ch10 := make(chan string, 2)
	ch10 <- "halo"
	select {
	case ch10 <- "xxx":
		fmt.Println("ch10 ", <-ch10)
		fmt.Println("ch10 ", <-ch10)
	default:
		fmt.Println("default")
	}

}

func task1(ch chan string) {
	time.Sleep(time.Second * 5)
	ch <- "halo"
}

func task2(ch chan string) {
	time.Sleep(time.Second * 2)
	ch <- "xxx"
}

func task3(ch chan string) {
	time.Sleep(time.Second * 1)
	ch <- "gogogo"
}

func makeTimeout(ch chan bool, t int) {
	time.Sleep(time.Second * time.Duration(t))
	ch <- true
}
