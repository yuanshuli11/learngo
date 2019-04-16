package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	counter int64
	wg      sync.WaitGroup
)
//WaitGroup用于等待一组线程的结束。父线程调用Add方法来设定应等待的线程的数量。
// 每个被等待的线程在结束时应调用Done方法。同时，主线程里可以调用Wait方法阻塞至所有线程结束。
func main() {

	wg.Add(2)
	go incCounter(1)
	go incCounter(2)
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

func incCounter(id int) {
	defer wg.Done()
	//if id >1{
	//	time.Sleep(2*time.Second)
	//}
	for count := 0; count < 2; count++ {
		atomic.AddInt64(&counter, 1)

		fmt.Printf("go%d:%d \n",id,counter)
	}
	runtime.Gosched()
}
