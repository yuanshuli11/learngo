package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Cond实现了一个条件变量，一个线程集合地，供线程等待或者宣布某事件的发生。
每个Cond实例都有一个相关的锁（一般是*Mutex或*RWMutex类型的值），
它必须在改变条件时或者调用Wait方法时保持锁定。Cond可以创建为其他结构体的字段，
Cond在开始使用后不能被拷贝。

 */

var count int = 4

func main() {
	ch := make(chan struct{}, 5)

	//新建 cond
	var l sync.Mutex
	cond := sync.NewCond(&l)

	for i := 0; i < 5; i++ {
		go func(i int) {
			cond.L.Lock()
			defer func() {
				cond.L.Unlock()
				ch <- struct{}{}
			}()

			//条件是否达成
			for count >i {
				cond.Wait()
				fmt.Printf("收到一个通知 goroutine%d\n",i)
			}
			fmt.Printf("goroutine%d 执行结束\n",i)
		}(i)

	}
	time.Sleep(time.Microsecond*200)
	fmt.Println("broadcast...")
	cond.L.Lock()
	count -= 1
	//广播发出去，然后符合条件的协程就会继续处理
	cond.Broadcast()
	cond.L.Unlock()

	time.Sleep(time.Second)
	fmt.Println("signal...")
	cond.L.Lock()
	count -=2
	cond.Signal()
	cond.L.Unlock()

	time.Sleep(time.Second)
	fmt.Println("broadcast...")
	cond.L.Lock()
	count -=1
	cond.Broadcast()
	cond.L.Unlock()

	for i:=0;i<5;i++{
		<-ch
	}

}
