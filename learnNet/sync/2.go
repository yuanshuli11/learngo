package main

import (
	"fmt"
	"sync"
)

//once do
func main() {

	var once sync.Once
	onceBody := func() {
		fmt.Println("Only once")
	}

	//done := make(chan bool)

	for i := 0; i < 10; i++ {
		go func() {
			once.Do(onceBody)
		//	done<-true
		}()
	}
	//通道的存在是为了给goroutine 足够的时间
	for i := 0; i < 10; i++ {
	//	<-done
	}
}
