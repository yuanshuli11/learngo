package main

import (
	"fmt"
	"sync"
)

//互斥锁
func main(){

	ch := make(chan struct{},10)

	var l sync.Mutex

	for i:=0;i<10;i++{
		go func(index int) {
			l.Lock()
			defer l.Unlock()
			fmt.Printf("goroutine %d: 我抢到了，我会锁定大概 1s \n",index)
			//time.Sleep(time.Second*1)
			fmt.Printf("goroutine %d: 我解锁了，你们去抢吧 \n",index)
			ch <- struct{}{}
		}(i)
	}

	for i:=0;i<10 ;i++  {
		<-ch
	}

}