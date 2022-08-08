package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

func main() {
	// 互斥锁（Mutex）和 读写锁（RWMutex）处理竞争条件

	a := &Deposits{amount: 0}
	a.Add(100)
	a.Add(200)
	a.Add(300)
	fmt.Println("a = ", a.Total())

	fmt.Println("==================================")

	// 临界区
	// 当程序并发运行时，多个 Go 协程不应该同时访问哪些修改共享资源的代码，这些修改共享资源的代码称为临界区
	var wg sync.WaitGroup
	b := &Deposits{amount: 0}

	n := 100
	wg.Add(n)
	for i := 1; i <= n; i++ {
		go func() {
			b.Add(100)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("b = ", b.Total()) // 9800、9900、10000...
	// 上面的代码没啥问题，但是要是有多个协程并发时，就会发送错误，这种情况被称为数据竞争（data race）

	fmt.Println("==================================")

	// 互斥锁
	// Mutex 用于提供一种加锁机制，以确保在某个时刻只有一个协程在临界区运行，以防止出现竞争。
	// 也是为了来保护一个资源不会被并发操作而引起冲突导致数据不准确

	// Mutex 中有两个方法：Lock 和 Unlock
	// Lock 方法用于获取锁，Unlock 方法用于释放锁
	var wg2 sync.WaitGroup
	c := &DepositsV2{}
	n2 := 100
	wg2.Add(n2)

	for i := 1; i <= n2; i++ {
		go func() {
			c.Add(100)
			wg2.Done()
		}()
	}
	wg2.Wait()
	fmt.Println("c = ", c.Total()) // 10000

	fmt.Println("==================================")

	// 读写锁 RWMutex，适用于读多写少的场景
	// 读锁：RLock 开启锁， RUnlock 释放锁
	// 写锁：Lock 开启锁，Unlock 释放锁
	var wg3 sync.WaitGroup
	d := &DepositsV3{}
	n3 := 100
	wg3.Add(n3)

	for i := 1; i <= n3; i++ {
		go func() {
			d.Add(100)
			wg3.Done()
		}()
	}
	wg3.Wait()
	fmt.Println("d = ", d.Total()) // 10000

	fmt.Println("==================================")

	// 条件变量（Cond）
	// 条件变量用于等待和通知
	//type Cond struct {
	//	...
	//	L Locker
	//	...
	//}
	//// 创建一个带锁的条件变量，Locker 通常是一个 *Mutex 或 *RWMutex
	//func NewCond(l Locker) *Cond
	//// 唤醒所有因等待条件变量 c 阻塞的 goroutine
	//func (c *Cond) Broadcast()
	//// 唤醒一个因等待条件变量 c 阻塞的 goroutine
	//func (c *Cond) Signal()
	//// 等待 c.L 解锁并挂起 goroutine，在稍后恢复执行后，Wait 返回前锁定 c.L，
	//// 只有当被 Broadcast 和 Signal 唤醒，Wait 才能返回。
	//func (c *Cond) Wait()

	s1 := []string{"xxx"}
	s2 := []string{"yyy"}
	s3 := []string{"zzz"}
	var m sync.Mutex
	cond := sync.NewCond(&m)

	// listener 1
	go listen("666", s1, cond)
	// listener 2
	go listen("777", s2, cond)
	// listener 3
	go listen("888", s3, cond)

	// broadcast
	go broadcast("贵宾", cond)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
}

type Deposits struct {
	amount int
}

func (d *Deposits) Add(a int) {
	d.amount += a
}

func (d *Deposits) Total() int {
	return d.amount
}

type DepositsV2 struct {
	amount int
	mu     sync.Mutex
}

func (d *DepositsV2) Add(a int) {
	d.mu.Lock()
	d.amount += a
	d.mu.Unlock()
}

func (d *DepositsV2) Total() int {
	return d.amount
}

type DepositsV3 struct {
	amount  int
	rwMutex sync.RWMutex
}

func (d *DepositsV3) Add(a int) {
	d.rwMutex.Lock()
	d.amount += a
	d.rwMutex.Unlock()
}

func (d *DepositsV3) Total() (dd int) {
	d.rwMutex.RLock()
	dd = d.amount
	d.rwMutex.RUnlock()
	return
}

func listen(name string, s []string, c *sync.Cond) {
	c.L.Lock()
	c.Wait()
	fmt.Println(name, "接待：", s)
	c.L.Unlock()
}

func broadcast(event string, c *sync.Cond) {
	time.Sleep(time.Second)
	c.L.Lock()
	fmt.Println("欢迎光临", event)
	c.Broadcast()
	c.L.Unlock()
}
