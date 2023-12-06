package main

import (
	"fmt"
	"time"
)

// 死锁1，无缓冲只读不写
func deadLock1() {
	// fatal error: all goroutines are asleep - deadlock!
	ch := make(chan int)
	<-ch
}

// 死锁2，无缓冲只写不读
func deadLock2() {
	// fatal error: all goroutines are asleep - deadlock!
	ch := make(chan int)
	//go func() {
	//	<-ch
	//}()
	ch <- 1
}

// 死锁3，无缓冲区读在写后面
func deadLock3() {
	// fatal error: all goroutines are asleep - deadlock!
	//ch := make(chan int)
	//<-ch
	//ch <- 1

	ch := make(chan int)
	ch <- 100 //  这里会发生一直阻塞的情况，执行不到下面一句
	go func() {
		num := <-ch
		fmt.Println("num=", num)
	}()
	time.Sleep(time.Second)
}

// 死锁4：空读
func deadLock4() {
	ch := make(chan int)
	fmt.Println(<-ch)
}

func mutex() {
	//a := sync.Mutex{}
}

func main() {
	//deadLock1()
	//deadLock2()
	//deadLock3()
}
