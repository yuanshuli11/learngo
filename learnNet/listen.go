package main

import (
	"fmt"
	"net"
)


//监听tcp连接
func main() {

	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		// handle error
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
			continue
		}
		go handleConnection(conn)
	}

}

func handleConnection(c net.Conn){
	fmt.Println("RemoteAddr:",c.RemoteAddr())
	fmt.Println("LocalAddr:",c.LocalAddr())

}
