package main

import "fmt"

func main(){

	ch1 := make(chan int, 1)
	ch1 <- 1
	a := <-ch1
	ch1 <- 2
	fmt.Println(a)



}
