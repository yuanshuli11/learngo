package main

import (
	"fmt"
	"math/rand"
	"sync"
)

/*
读写锁 针对读写操作的互斥锁，
读写锁与互斥锁最大的不同就是可以分别对 读、写 进行锁定。一般用在大量读操作、少量写操作的情况：

1.同时只能有一个 goroutine 能够获得写锁定。
2.同时可以有任意多个 gorouinte 获得读锁定。
3.同时只能存在写锁定或读锁定（读和写互斥）。

 */

 var count int
 var rw sync.RWMutex
func main() {
	ch := make(chan struct{},10)

	for i:=0; i<5;i++{
		go read(i,ch)
	}
	for i:=0; i<5;i++{
		go write(i,ch)
	}
	for i := 0; i < 10; i++ {
		<-ch
	}
}

func read(n int, ch chan struct{}) {
	rw.RLock()
	fmt.Printf("goroutine %d 进入读操作...\n", n)
	v := count
	fmt.Printf("goroutine %d 读取结束，值为：%d\n", n, v)
	rw.RUnlock()
	ch <- struct{}{}
}

func write(n int, ch chan struct{}) {
	rw.Lock()
	fmt.Printf("goroutine %d 进入写操作...\n", n)
	v := rand.Intn(1000)
	count = v
	fmt.Printf("goroutine %d 写入结束，新值为：%d\n", n, v)
	rw.Unlock()
	ch <- struct{}{}
}